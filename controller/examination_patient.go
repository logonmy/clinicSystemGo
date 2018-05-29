package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/kataras/iris"
)

//ExaminationPatientCreate 开检查
func ExaminationPatientCreate(ctx iris.Context) {
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
		ctx.JSON(iris.Map{"code": "1", "msg": "保存检查失败,分诊记录错误"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "1", "msg": "保存检查失败,操作员错误"})
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

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errb})
		return
	}

	orderSn := FormatPayOrderSn(clinicTriagePatientID, "4")

	_, errdlp := tx.Exec("delete from examination_patient where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if errdlp != nil {
		fmt.Println("errdlp ===", errdlp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdlp.Error()})
		return
	}
	_, errdm := tx.Exec("delete from mz_unpaid_orders where clinic_triage_patient_id=$1 and charge_project_type_id=4", clinicTriagePatientID)
	if errdm != nil {
		fmt.Println("errdm ===", errdm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errdm.Error()})
		return
	}

	inserttSQL := `insert into examination_patient (
		clinic_triage_patient_id,
		clinic_examination_id,
		order_sn,
		soft_sn,
		times,
		organ,
		operation_id,
		illustration
	) values ($1, $2, $3, $4, $5, $6, $7, $8)`

	insertmSQL := `insert into mz_unpaid_orders (clinic_triage_patient_id,
		charge_project_type_id,
		charge_project_id,
		order_sn,
		soft_sn,
		name,
		price,
		amount,
		unit,
		total,
		discount,
		fee,
		operation_id) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	for index, v := range results {
		clinicExaminationID := v["clinic_examination_id"]
		times := v["times"]
		illustration := v["illustration"]
		organ := v["organ"]
		clinicExaminationSQL := `select * from clinic_examination	where id=$1`
		trow := model.DB.QueryRowx(clinicExaminationSQL, clinicExaminationID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "1", "msg": "检查项错误"})
			return
		}
		clinicExamination := FormatSQLRowToMap(trow)
		fmt.Println("====", clinicExamination)
		_, ok := clinicExamination["id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "1", "msg": "选择的检查项错误"})
			return
		}

		_, errt := tx.Exec(inserttSQL, clinicTriagePatientID, clinicExaminationID, orderSn, index, times, ToNullString(organ), personnelID, ToNullString(illustration))
		if errt != nil {
			fmt.Println("errt ===", errt)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "1", "msg": errt.Error()})
			return
		}

		price := int(clinicExamination["price"].(int64))
		discountPrice := int(clinicExamination["discount_price"].(int64))
		isDiscount := clinicExamination["is_discount"].(bool)
		name := clinicExamination["name"].(string)
		unitName := clinicExamination["unit_name"].(string)
		amount, _ := strconv.Atoi(times)
		total := price * amount
		discount := 0
		fee := total
		if isDiscount {
			discount = int(discountPrice) * amount
			fee = total - discount
		}
		chargeProjectTypeID := 4

		_, errm := tx.Exec(insertmSQL,
			clinicTriagePatientID,
			chargeProjectTypeID,
			clinicExaminationID,
			orderSn,
			index,
			name,
			price,
			amount,
			unitName,
			total,
			fee,
			discount,
			personnelID)
		if errm != nil {
			fmt.Println("errm ===", errm)
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": errm.Error()})
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

//ExaminationPatientGet 获取检查
func ExaminationPatientGet(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	rows, err := model.DB.Queryx(`select ep.*, ce.name from examination_patient ep 
		left join clinic_examination ce on ep.clinic_examination_id = ce.id 
		where ep.clinic_triage_patient_id = $1`, clinicTriagePatientID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
	}

	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}
