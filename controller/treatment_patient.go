package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

//TreatmentPatientCreate 开治疗
func TreatmentPatientCreate(ctx iris.Context) {
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
		ctx.JSON(iris.Map{"code": "1", "msg": "保存治疗失败,分诊记录错误"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "保存治疗失败,操作员错误"})
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

	var treatmentPatientValues []string
	treatmentPatientSets := []string{
		"clinic_triage_patient_id",
		"clinic_treatment_id",
		"order_sn",
		"soft_sn",
		"times",
		"operation_id",
		"illustration",
	}
	orderSn := FormatPayOrderSn(clinicTriagePatientID, "7")

	for index, v := range results {
		clinicTreatmentID := v["clinic_treatment_id"]
		times := v["times"]
		illustration := v["illustration"]
		fmt.Println("clinicTreatmentID====", clinicTreatmentID)
		var st []string
		var sm []string
		treatmentSQL := `select ct.id as clinic_treatment_id,ct.price,ct.is_discount,t.name,t.unit_name from clinic_treatment ct
		left join treatment t on t.id = ct.treatment_id
			where ct.id=$1`
		trow := model.DB.QueryRowx(treatmentSQL, clinicTreatmentID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "治疗项错误"})
			return
		}
		treatment := FormatSQLRowToMap(trow)
		fmt.Println("====", treatment)
		_, ok := treatment["clinic_treatment_id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "选择的治疗项错误"})
			return
		}
		price := treatment["price"].(int64)
		name := treatment["name"].(string)
		unitName := treatment["unit_name"].(string)
		amount, _ := strconv.Atoi(times)
		total := int(price) * amount

		st = append(st, clinicTriagePatientID, clinicTreatmentID, "'"+orderSn+"'", strconv.Itoa(index), times, personnelID)
		sm = append(sm, clinicTriagePatientID, "7", clinicTreatmentID, "'"+orderSn+"'", strconv.Itoa(index), "'"+name+"'", strconv.FormatInt(price, 10), strconv.Itoa(amount), "'"+unitName+"'", strconv.Itoa(total), strconv.Itoa(total), personnelID)

		if illustration == "" {
			st = append(st, `null`)
		} else {
			st = append(st, "'"+illustration+"'")
		}

		tstr := "(" + strings.Join(st, ",") + ")"
		treatmentPatientValues = append(treatmentPatientValues, tstr)
		mstr := "(" + strings.Join(sm, ",") + ")"
		mzUnpaidOrdersValues = append(mzUnpaidOrdersValues, mstr)
	}
	tSetStr := strings.Join(treatmentPatientSets, ",")
	tValueStr := strings.Join(treatmentPatientValues, ",")

	mSetStr := strings.Join(mzUnpaidOrdersSets, ",")
	mvValueStr := strings.Join(mzUnpaidOrdersValues, ",")

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}

	_, errdtp := tx.Exec("delete from treatment_patient where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if errdtp != nil {
		fmt.Println("errdtp ===", errdtp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdtp.Error()})
		return
	}
	_, errdm := tx.Exec("delete from mz_unpaid_orders where clinic_triage_patient_id=$1 and charge_project_type_id=7", clinicTriagePatientID)
	if errdm != nil {
		fmt.Println("errdm ===", errdm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdm.Error()})
		return
	}

	inserttSQL := "insert into treatment_patient (" + tSetStr + ") values " + tValueStr
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

//TreatmentPatientGet 查询治疗
func TreatmentPatientGet(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	rows, err := model.DB.Queryx(`select tp.*, t.name as treatment_name, t.unit_name from treatment_patient tp left join clinic_treatment ct on tp.clinic_treatment_id = ct.id 
		left join treatment t on ct.treatment_id = t.id
		where tp.clinic_triage_patient_id = $1`, clinicTriagePatientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
	}

	result := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": result})
}
