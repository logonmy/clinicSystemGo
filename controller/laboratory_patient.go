package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

//LaboratoryPatientCreate 开检验
func LaboratoryPatientCreate(ctx iris.Context) {
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存检验失败,分诊记录错误"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存检验失败,操作员错误"})
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

	clinicLaboratorySets := []string{
		"clinic_triage_patient_id",
		"clinic_laboratory_id",
		"order_sn",
		"soft_sn",
		"times",
		"operation_id",
		"illustration",
	}
	tSetStr := strings.Join(clinicLaboratorySets, ",")
	mSetStr := strings.Join(mzUnpaidOrdersSets, ",")

	orderSn := FormatPayOrderSn(clinicTriagePatientID, "3")

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}
	_, errdlp := tx.Exec("delete from laboratory_patient where clinic_triage_patient_id=$1", clinicTriagePatientID)
	if errdlp != nil {
		fmt.Println("errdlp ===", errdlp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errdlp.Error()})
		return
	}
	_, errdm := tx.Exec("delete from mz_unpaid_orders where clinic_triage_patient_id=$1 and charge_project_type_id=3", clinicTriagePatientID)
	if errdm != nil {
		fmt.Println("errdm ===", errdm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errdm.Error()})
		return
	}
	inserttSQL := "insert into laboratory_patient (" + tSetStr + ") values ($1,$2,$3,$4,$5,$6,$7)"
	insertmSQL := "insert into mz_unpaid_orders (" + mSetStr + ") values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)"

	for index, v := range results {
		clinicLaboratoryID := v["clinic_laboratory_id"]
		times := v["times"]
		illustration := v["illustration"]
		fmt.Println("clinicLaboratoryID====", clinicLaboratoryID)
		clinicLaboratorySQL := `select id as clinic_laboratory_id,price,is_discount,name,unit_name,discount_price from clinic_laboratory where id=$1`
		trow := model.DB.QueryRowx(clinicLaboratorySQL, clinicLaboratoryID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "检验项错误"})
			return
		}
		clinicLaboratory := FormatSQLRowToMap(trow)
		fmt.Println("====", clinicLaboratory)
		_, ok := clinicLaboratory["clinic_laboratory_id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "选择的检验项错误"})
			return
		}
		isDiscount := clinicLaboratory["is_discount"].(bool)
		price := clinicLaboratory["price"].(int64)
		discountPrice := clinicLaboratory["discount_price"].(int64)
		name := clinicLaboratory["name"].(string)
		unitName := clinicLaboratory["unit_name"].(string)
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
			ToNullInt64(clinicLaboratoryID),
			ToNullString(orderSn),
			index,
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
			3,
			ToNullInt64(clinicLaboratoryID),
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

//LaboratoryPatientGet 获取检验
func LaboratoryPatientGet(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	rows, err := model.DB.Queryx(`select lp.*, cl.name as laboratory_name, cl.price from laboratory_patient lp 
		left join clinic_laboratory cl on lp.clinic_laboratory_id = cl.id 
		where lp.clinic_triage_patient_id = $1`, clinicTriagePatientID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
	}

	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}
