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
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存治疗失败,分诊记录错误"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存治疗失败,操作员错误"})
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
		"discount",
		"fee",
		"operation_id",
	}
	mSetStr := strings.Join(mzUnpaidOrdersSets, ",")

	treatmentPatientSets := []string{
		"clinic_triage_patient_id",
		"clinic_treatment_id",
		"order_sn",
		"soft_sn",
		"times",
		"left_times",
		"operation_id",
		"illustration",
	}
	tSetStr := strings.Join(treatmentPatientSets, ",")

	orderSn := FormatPayOrderSn(clinicTriagePatientID, "7")
	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}

	_, errdtp := tx.Exec("delete from treatment_patient where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if errdtp != nil {
		fmt.Println("errdtp ===", errdtp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errdtp.Error()})
		return
	}
	_, errdm := tx.Exec("delete from mz_unpaid_orders where clinic_triage_patient_id=$1 and charge_project_type_id=7", clinicTriagePatientID)
	if errdm != nil {
		fmt.Println("errdm ===", errdm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errdm.Error()})
		return
	}
	inserttSQL := "insert into treatment_patient (" + tSetStr + ") values ($1,$2,$3,$4,$5,$6,$7,$8)"
	insertmSQL := "insert into mz_unpaid_orders (" + mSetStr + ") values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)"

	for index, v := range results {
		clinicTreatmentID := v["clinic_treatment_id"]
		times := v["times"]
		illustration := v["illustration"]
		fmt.Println("clinicTreatmentID====", clinicTreatmentID)

		treatmentSQL := `select id as clinic_treatment_id,price,discount_price,is_discount,name,unit_name from clinic_treatment where id=$1`
		trow := model.DB.QueryRowx(treatmentSQL, clinicTreatmentID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "治疗项错误"})
			return
		}
		clinicTreatment := FormatSQLRowToMap(trow)
		fmt.Println("====", clinicTreatment)
		_, ok := clinicTreatment["clinic_treatment_id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "选择的治疗项错误"})
			return
		}
		isDiscount := clinicTreatment["is_discount"].(bool)
		price := clinicTreatment["price"].(int64)
		discountPrice := clinicTreatment["discount_price"].(int64)
		name := clinicTreatment["name"].(string)
		unitName := clinicTreatment["unit_name"].(string)
		amount, _ := strconv.Atoi(times)
		total := int(price) * amount
		discount := 0
		fee := total
		if isDiscount {
			discount = int(discountPrice) * amount
			fee = total - discount
		}

		_, errt := tx.Exec(inserttSQL,
			ToNullInt64(clinicTriagePatientID),
			ToNullInt64(clinicTreatmentID),
			ToNullString(orderSn),
			index,
			ToNullInt64(times),
			ToNullInt64(times),
			ToNullInt64(personnelID),
			ToNullString(illustration),
		)
		if errt != nil {
			fmt.Println("errt ===", errt)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": errt.Error()})
			return
		}
		_, errm := tx.Exec(insertmSQL,
			ToNullInt64(clinicTriagePatientID),
			7,
			ToNullInt64(clinicTreatmentID),
			ToNullString(orderSn),
			index,
			ToNullString(name),
			price,
			amount,
			ToNullString(unitName),
			total,
			discount,
			fee,
			personnelID,
		)
		if errm != nil {
			fmt.Println("errm ===", errm)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": "请检查是否漏填"})
			return
		}
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

	rows, err := model.DB.Queryx(`select tp.*, ct.name as treatment_name, ct.unit_name, ct.price from treatment_patient tp 
		left join clinic_treatment ct on tp.clinic_treatment_id = ct.id 
		where tp.clinic_triage_patient_id = $1`, clinicTriagePatientID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
	}

	result := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": result})
}
