package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

//PrescriptionWesternPatientCreate 开西/成药处方
func PrescriptionWesternPatientCreate(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	personnelID := ctx.PostValue("personnel_id")
	items := ctx.PostValue("items")

	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}
	if items == "" {
		ctx.JSON(iris.Map{"code": "200", "data": nil})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("results===", results)
	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}
	row := model.DB.QueryRowx(`select id,status from clinic_triage_patient where id=$1 limit 1`, clinicTriagePatientID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存西/成药处方失败,分诊记录错误"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "保存西/成药处方失败,操作员错误"})
		return
	}
	clinicTriagePatient := FormatSQLRowToMap(row)
	personnel := FormatSQLRowToMap(prow)

	_, ok := clinicTriagePatient["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "1", "msg": "分诊记录不存在"})
		return
	}
	status := clinicTriagePatient["status"]
	if status.(int64) != 30 {
		ctx.JSON(iris.Map{"code": "1", "msg": "分诊记录当前状态错误"})
		return
	}
	_, pok := personnel["id"]
	if !pok {
		ctx.JSON(iris.Map{"code": "1", "msg": "操作员错误"})
		return
	}

	var mzUnpaidOrdersValues []string
	mzUnpaidOrdersSets := []string{
		"clinic_triage_patient_id",
		"charge_project_type_id",
		"charge_project_id",
		"order_sn",
		"soft_sn",
		"name",
		"price",
		"amount",
		"unit",
		"total",
		"fee",
		"operation_id",
	}

	var prescriptionWesternPatientValues []string
	prescriptionWesternPatientSets := []string{
		"clinic_triage_patient_id",
		"drug_stock_id",
		"order_sn",
		"soft_sn",
		"once_dose",
		"once_dose_unit_id",
		"route_administration_id",
		"frequency_id",
		"amount",
		"fetch_address",
		"operation_id",
		"eff_day",
		"illustration",
	}
	orderSn := FormatPayOrderSn(clinicTriagePatientID, "1")

	for index, v := range results {
		drugStockID := v["drug_stock_id"]
		onceDose := v["once_dose"]
		onceDoseUnitID := v["once_dose_unit_id"]
		routeAdministrationID := v["route_administration_id"]
		frequencyID := v["frequency_id"]
		times := v["amount"]
		illustration := v["illustration"]
		fetchAddress := v["fetch_address"]
		effDay := v["eff_day"]
		fmt.Println("drugStockID====", drugStockID)
		var sl []string
		var sm []string
		laboratorySQL := `select ds.id,d.name,ds.ret_price,du.name as packing_unit_name from drug_stock ds
			left join drug d on d.id = ds.drug_id
			left join dose_unit du on du.id = ds.packing_unit_id
			where ds.id=$1`
		trow := model.DB.QueryRowx(laboratorySQL, drugStockID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "西/成药处方项错误"})
			return
		}
		drugStock := FormatSQLRowToMap(trow)
		fmt.Println("====", drugStock)
		_, ok := drugStock["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "选择的西/成药处方药错误"})
			return
		}

		price := drugStock["ret_price"].(int64)
		name := drugStock["name"].(string)
		unitName := drugStock["packing_unit_name"].(string)
		amount, _ := strconv.Atoi(times)
		total := int(price) * amount

		sl = append(sl, clinicTriagePatientID, drugStockID, "'"+orderSn+"'", strconv.Itoa(index), onceDose, onceDoseUnitID, routeAdministrationID, frequencyID, times, fetchAddress, personnelID)
		sm = append(sm, clinicTriagePatientID, "1", drugStockID, "'"+orderSn+"'", strconv.Itoa(index), "'"+name+"'", strconv.FormatInt(price, 10), strconv.Itoa(amount), "'"+unitName+"'", strconv.Itoa(total), strconv.Itoa(total), personnelID)

		if illustration == "" {
			sl = append(sl, `null`)
		} else {
			sl = append(sl, "'"+illustration+"'")
		}

		if effDay == "" {
			sl = append(sl, `null`)
		} else {
			sl = append(sl, effDay)
		}

		tstr := "(" + strings.Join(sl, ",") + ")"
		prescriptionWesternPatientValues = append(prescriptionWesternPatientValues, tstr)
		mstr := "(" + strings.Join(sm, ",") + ")"
		mzUnpaidOrdersValues = append(mzUnpaidOrdersValues, mstr)
	}
	tSetStr := strings.Join(prescriptionWesternPatientSets, ",")
	tValueStr := strings.Join(prescriptionWesternPatientValues, ",")

	mSetStr := strings.Join(mzUnpaidOrdersSets, ",")
	mvValueStr := strings.Join(mzUnpaidOrdersValues, ",")

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}
	_, errdlp := tx.Exec("delete from prescription_western_patient where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if errdlp != nil {
		fmt.Println("errdlp ===", errdlp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdlp.Error()})
		return
	}
	_, errdm := tx.Exec("delete from mz_unpaid_orders where clinic_triage_patient_id=$1 and charge_project_type_id=1", clinicTriagePatientID)
	if errdm != nil {
		fmt.Println("errdm ===", errdm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdm.Error()})
		return
	}

	inserttSQL := "insert into prescription_western_patient (" + tSetStr + ") values " + tValueStr
	fmt.Println("inserttSQL===", inserttSQL)

	_, errt := tx.Exec(inserttSQL)
	if errt != nil {
		fmt.Println("errt ===", errt)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errt.Error()})
		return
	}

	insertmSQL := "insert into mz_unpaid_orders (" + mSetStr + ") values " + mvValueStr
	fmt.Println("insertmSQL===", insertmSQL)

	_, errm := tx.Exec(insertmSQL)
	if errm != nil {
		fmt.Println("errm ===", errm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": "请检查是否漏填"})
		return
	}
	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//PrescriptionChinesePatientCreate 开中药处方
func PrescriptionChinesePatientCreate(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	dayDose := ctx.PostValue("day_dose")
	doseUnitID := ctx.PostValue("dose_unit_id")
	routeAdministrationID := ctx.PostValue("route_administration_id")
	frequencyID := ctx.PostValue("frequency_id")
	amount := ctx.PostValue("amount")
	medicineIllustration := ctx.PostValue("medicine_illustration")
	fetchAddress := ctx.PostValue("fetch_address")
	effDay := ctx.PostValue("eff_day")
	personnelID := ctx.PostValue("personnel_id")
	items := ctx.PostValue("items")

	if clinicTriagePatientID == "" || routeAdministrationID == "" || frequencyID == "" || amount == "" || fetchAddress == "" || doseUnitID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	if items == "" {
		ctx.JSON(iris.Map{"code": "200", "data": nil})
		return
	}

	var results []map[string]string
	errj := json.Unmarshal([]byte(items), &results)
	fmt.Println("results===", results)
	if errj != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": errj.Error()})
		return
	}
	row := model.DB.QueryRowx(`select id,status from clinic_triage_patient where id=$1 limit 1`, clinicTriagePatientID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存中药处方失败,分诊记录错误"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存中药处方失败,操作员错误"})
		return
	}
	clinicTriagePatient := FormatSQLRowToMap(row)
	personnel := FormatSQLRowToMap(prow)

	_, ok := clinicTriagePatient["id"]
	if !ok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "分诊记录不存在"})
		return
	}
	status := clinicTriagePatient["status"]
	if status.(int64) != 30 {
		ctx.JSON(iris.Map{"code": "-1", "msg": "分诊记录当前状态错误"})
		return
	}
	_, pok := personnel["id"]
	if !pok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "操作员错误"})
		return
	}

	var mzUnpaidOrdersValues []string
	mzUnpaidOrdersSets := []string{
		"clinic_triage_patient_id",
		"charge_project_type_id",
		"charge_project_id",
		"order_sn",
		"soft_sn",
		"name",
		"price",
		"amount",
		"unit",
		"total",
		"fee",
		"operation_id",
	}

	var prescriptionChinesePatientValues []string
	prescriptionChinesePatientSets := []string{
		"clinic_triage_patient_id",
		"order_sn",
		"dose_unit_id",
		"route_administration_id",
		"frequency_id",
		"amount",
		"fetch_address",
		"operation_id",
		"day_dose",
		"eff_day",
		"medicine_illustration",
	}

	var prescriptionChineseItemValues []string
	prescriptionChineseItemSets := []string{
		"drug_stock_id",
		"order_sn",
		"soft_sn",
		"once_dose",
		"once_dose_unit_id",
		"amount",
		"prescription_chinese_patient_id",
		"special_illustration",
	}
	orderSn := FormatPayOrderSn(clinicTriagePatientID, "2")
	prescriptionChinesePatientValues = append(prescriptionChinesePatientValues, clinicTriagePatientID, "'"+orderSn+"'", doseUnitID, routeAdministrationID, frequencyID, amount, fetchAddress, personnelID)
	if dayDose == "" {
		prescriptionChinesePatientValues = append(prescriptionChinesePatientValues, `null`)
	} else {
		prescriptionChinesePatientValues = append(prescriptionChinesePatientValues, dayDose)
	}
	if effDay == "" {
		prescriptionChinesePatientValues = append(prescriptionChinesePatientValues, `null`)
	} else {
		prescriptionChinesePatientValues = append(prescriptionChinesePatientValues, effDay)
	}
	if medicineIllustration == "" {
		prescriptionChinesePatientValues = append(prescriptionChinesePatientValues, `null`)
	} else {
		prescriptionChinesePatientValues = append(prescriptionChinesePatientValues, "'"+medicineIllustration+"'")
	}

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}
	_, errdlp := tx.Exec("delete from prescription_chinese_patient where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if errdlp != nil {
		fmt.Println("errdlp ===", errdlp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errdlp.Error()})
		return
	}

	_, errdm := tx.Exec("delete from mz_unpaid_orders where clinic_triage_patient_id=$1 and charge_project_type_id=2", clinicTriagePatientID)
	if errdm != nil {
		fmt.Println("errdm ===", errdm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errdm.Error()})
		return
	}
	tSetStr := strings.Join(prescriptionChinesePatientSets, ",")
	tValueStr := strings.Join(prescriptionChinesePatientValues, ",")

	insertpwpSQL := "insert into prescription_chinese_patient (" + tSetStr + ") values (" + tValueStr + ") RETURNING id"
	fmt.Println("insertpwpSQL===", insertpwpSQL)

	pcprow := tx.QueryRow(insertpwpSQL)
	var prescriptionChinesePatientID string
	errp := pcprow.Scan(&prescriptionChinesePatientID)
	if errp != nil {
		fmt.Println("errp ===", errp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存中药处方错误"})
		return
	}
	fmt.Println("====", prescriptionChinesePatientID)
	for index, v := range results {
		drugStockID := v["drug_stock_id"]
		onceDose := v["once_dose"]
		onceDoseUnitID := v["once_dose_unit_id"]
		times := v["amount"]
		illustration := v["special_illustration"]
		fmt.Println("drugStockID====", drugStockID)
		var sl []string
		var sm []string
		laboratorySQL := `select ds.id,d.name,ds.ret_price,du.name as packing_unit_name from drug_stock ds
			left join drug d on d.id = ds.drug_id
			left join dose_unit du on du.id = ds.packing_unit_id
			where ds.id=$1`
		trow := model.DB.QueryRowx(laboratorySQL, drugStockID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "中药处方项错误"})
			return
		}
		drugStock := FormatSQLRowToMap(trow)
		fmt.Println("====", drugStock)
		_, ok := drugStock["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "选择的中药处方药错误"})
			return
		}

		price := drugStock["ret_price"].(int64)
		name := drugStock["name"].(string)
		unitName := drugStock["packing_unit_name"].(string)
		drugAmount, _ := strconv.Atoi(times)
		total := int(price) * drugAmount

		sl = append(sl, drugStockID, "'"+orderSn+"'", strconv.Itoa(index), onceDose, onceDoseUnitID, times, prescriptionChinesePatientID)
		sm = append(sm, clinicTriagePatientID, "2", drugStockID, "'"+orderSn+"'", strconv.Itoa(index), "'"+name+"'", strconv.FormatInt(price, 10), strconv.Itoa(drugAmount), "'"+unitName+"'", strconv.Itoa(total), strconv.Itoa(total), personnelID)

		if illustration == "" {
			sl = append(sl, `null`)
		} else {
			sl = append(sl, "'"+illustration+"'")
		}

		pcstr := "(" + strings.Join(sl, ",") + ")"
		prescriptionChineseItemValues = append(prescriptionChineseItemValues, pcstr)
		mstr := "(" + strings.Join(sm, ",") + ")"
		mzUnpaidOrdersValues = append(mzUnpaidOrdersValues, mstr)
	}

	mSetStr := strings.Join(mzUnpaidOrdersSets, ",")
	mvValueStr := strings.Join(mzUnpaidOrdersValues, ",")

	pcSetStr := strings.Join(prescriptionChineseItemSets, ",")
	pcValueStr := strings.Join(prescriptionChineseItemValues, ",")

	insertpSQL := "insert into prescription_chinese_item (" + pcSetStr + ") values " + pcValueStr
	fmt.Println("insertpSQL===", insertpSQL)
	_, errpci := tx.Exec(insertpSQL)
	if errpci != nil {
		fmt.Println("errpci ===", errpci)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": "请检查是否漏填"})
		return
	}

	insertmSQL := "insert into mz_unpaid_orders (" + mSetStr + ") values " + mvValueStr
	fmt.Println("insertmSQL===", insertmSQL)
	_, errm := tx.Exec(insertmSQL)
	if errm != nil {
		fmt.Println("errm ===", errm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": "请检查是否漏填"})
		return
	}
	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errc.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//PrescriptionWesternPatientGet 获取西药处方
func PrescriptionWesternPatientGet(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	rows, err := model.DB.Queryx(`select pwp.*, d.name as drug_name,d.specification,ds.stock_amount from prescription_western_patient pwp 
		left join drug_stock ds on pwp.drug_stock_id = ds.id 
		left join drug d on ds.drug_id = d.id
		where pwp.clinic_triage_patient_id = $1`, clinicTriagePatientID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
	}

	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}

//PrescriptionChinesePatientGet 获取中药处方
func PrescriptionChinesePatientGet(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	pcprow := model.DB.QueryRowx(`select * from prescription_chinese_patient where clinic_triage_patient_id = $1`, clinicTriagePatientID)
	prescriptionChinesePatient := FormatSQLRowToMap(pcprow)
	prescriptionChinesePatientID := prescriptionChinesePatient["id"]

	rows, err := model.DB.Queryx(`select pci.*, d.name as drug_name,d.specification,ds.stock_amount from prescription_chinese_item pci 
		left join drug_stock ds on pci.drug_stock_id = ds.id 
		left join drug d on ds.drug_id = d.id
		where pci.prescription_chinese_patient_id = $1`, prescriptionChinesePatientID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
	}

	result := FormatSQLRowsToMapArray(rows)
	prescriptionChinesePatient["items"] = result
	ctx.JSON(iris.Map{"code": "200", "data": prescriptionChinesePatient})
}
