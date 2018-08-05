package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

//OtherCostPatientCreate 开其它费用
func OtherCostPatientCreate(ctx iris.Context) {
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
		ctx.JSON(iris.Map{"code": "--1", "msg": errj.Error()})
		return
	}
	row := model.DB.QueryRowx(`select id,status from clinic_triage_patient where id=$1 limit 1`, clinicTriagePatientID)
	if row == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存其它费用失败,分诊记录错误"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存其它费用失败,操作员错误"})
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
		"cost",
		"price",
		"amount",
		"unit",
		"total",
		"discount",
		"fee",
		"operation_id",
	}

	otherCostPatientSets := []string{
		"clinic_triage_patient_id",
		"clinic_other_cost_id",
		"order_sn",
		"soft_sn",
		"amount",
		"operation_id",
		"illustration",
	}
	tSetStr := strings.Join(otherCostPatientSets, ",")
	mSetStr := strings.Join(mzUnpaidOrdersSets, ",")
	orderSn := FormatPayOrderSn(clinicTriagePatientID, "6")

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}
	_, errdlp := tx.Exec("delete from other_cost_patient where clinic_triage_patient_id=$1 and order_sn in (select order_sn from mz_unpaid_orders where clinic_triage_patient_id = $1 and charge_project_type_id=6)", clinicTriagePatientID)
	if errdlp != nil {
		fmt.Println("errdlp ===", errdlp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errdlp.Error()})
		return
	}
	_, errdm := tx.Exec("delete from mz_unpaid_orders where clinic_triage_patient_id=$1 and charge_project_type_id=6", clinicTriagePatientID)
	if errdm != nil {
		fmt.Println("errdm ===", errdm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errdm.Error()})
		return
	}
	clinicOtherCostSQL := `select id as clinic_other_cost_id,price,is_discount,name,unit_name,discount_price,COALESCE(cost, 0) as cost_price from clinic_other_cost where id=$1`
	inserttSQL := "insert into other_cost_patient (" + tSetStr + ") values ($1,$2,$3,$4,$5,$6,$7)"
	insertmSQL := "insert into mz_unpaid_orders (" + mSetStr + ") values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)"

	for index, v := range results {
		clinicOtherCostID := v["clinic_other_cost_id"]
		times := v["amount"]
		illustration := v["illustration"]
		fmt.Println("clinicOtherCostID====", clinicOtherCostID)
		if times == "" {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": "请填写次数"})
			return
		}

		trow := model.DB.QueryRowx(clinicOtherCostSQL, clinicOtherCostID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "其它费用项错误"})
			return
		}
		clinicOtherCost := FormatSQLRowToMap(trow)
		fmt.Println("====", clinicOtherCost)
		_, ok := clinicOtherCost["clinic_other_cost_id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "选择的其它费用项错误"})
			return
		}
		isDiscount := clinicOtherCost["is_discount"].(bool)
		costPrice := clinicOtherCost["cost_price"].(int64)
		price := clinicOtherCost["price"].(int64)
		discountPrice := clinicOtherCost["discount_price"].(int64)
		name := clinicOtherCost["name"].(string)
		unitName := clinicOtherCost["unit_name"].(string)
		amount, _ := strconv.Atoi(times)
		total := int(price) * amount
		cost := int(costPrice) * amount
		discount := 0
		fee := total
		if isDiscount {
			discount = int(discountPrice) * amount
			fee = total - discount
		}

		_, errt := tx.Exec(inserttSQL,
			ToNullInt64(clinicTriagePatientID),
			ToNullInt64(clinicOtherCostID),
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
			6,
			ToNullInt64(clinicOtherCostID),
			ToNullString(orderSn),
			index,
			ToNullString(name),
			cost,
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
			ctx.JSON(iris.Map{"code": "-1", "msg": "请其它费用是否漏填"})
			return
		}
	}

	errc := tx.Commit()
	if errc != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "--1", "msg": errc.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})
}

//OtherCostPatientGet 获取其它费
func OtherCostPatientGet(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	rows, err := model.DB.Queryx(`select ocp.*, coc.name, coc.unit_name, coc.price from other_cost_patient ocp 
		left join clinic_other_cost coc on ocp.clinic_other_cost_id = coc.id 
		where ocp.clinic_triage_patient_id = $1`, clinicTriagePatientID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "--1", "msg": err.Error()})
	}

	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}
