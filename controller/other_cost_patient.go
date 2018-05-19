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

	var otherCostPatientValues []string
	otherCostPatientSets := []string{
		"clinic_triage_patient_id",
		"clinic_other_cost_id",
		"order_sn",
		"soft_sn",
		"amount",
		"operation_id",
		"illustration",
	}

	orderSn := FormatPayOrderSn(clinicTriagePatientID, "6")

	for index, v := range results {
		clinicOtherCostID := v["clinic_other_cost_id"]
		times := v["amount"]
		illustration := v["illustration"]
		fmt.Println("clinicOtherCostID====", clinicOtherCostID)
		var sl []string
		var sm []string
		clinicOtherCostSQL := `select coc.id as clinic_other_cost_id,coc.price,coc.is_discount,oc.name,oc.unit_name from clinic_other_cost coc
		left join other_cost oc on oc.id = coc.other_cost_id
		where coc.id=$1`
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
		price := clinicOtherCost["price"].(int64)
		name := clinicOtherCost["name"].(string)
		unitName := clinicOtherCost["unit_name"].(string)
		amount, _ := strconv.Atoi(times)
		total := int(price) * amount

		sl = append(sl, clinicTriagePatientID, clinicOtherCostID, "'"+orderSn+"'", strconv.Itoa(index), times, personnelID)
		sm = append(sm, clinicTriagePatientID, "6", clinicOtherCostID, "'"+orderSn+"'", strconv.Itoa(index), "'"+name+"'", strconv.FormatInt(price, 10), strconv.Itoa(amount), "'"+unitName+"'", strconv.Itoa(total), strconv.Itoa(total), personnelID)

		if illustration == "" {
			sl = append(sl, `null`)
		} else {
			sl = append(sl, "'"+illustration+"'")
		}

		tstr := "(" + strings.Join(sl, ",") + ")"
		otherCostPatientValues = append(otherCostPatientValues, tstr)
		mstr := "(" + strings.Join(sm, ",") + ")"
		mzUnpaidOrdersValues = append(mzUnpaidOrdersValues, mstr)
	}
	tSetStr := strings.Join(otherCostPatientSets, ",")
	tValueStr := strings.Join(otherCostPatientValues, ",")

	mSetStr := strings.Join(mzUnpaidOrdersSets, ",")
	mvValueStr := strings.Join(mzUnpaidOrdersValues, ",")

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}
	_, errdlp := tx.Exec("delete from other_cost_patient where clinic_triage_patient_id=$1", clinicTriagePatientID)
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

	inserttSQL := "insert into other_cost_patient (" + tSetStr + ") values " + tValueStr
	fmt.Println("inserttSQL===", inserttSQL)

	_, errt := tx.Exec(inserttSQL)
	if errt != nil {
		fmt.Println("errt ===", errt)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errt.Error()})
		return
	}

	insertmSQL := "insert into mz_unpaid_orders (" + mSetStr + ") values " + mvValueStr
	fmt.Println("insertmSQL===", insertmSQL)
	_, errm := tx.Exec(insertmSQL)
	if errm != nil {
		fmt.Println("errm ===", errm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": "请其它费用是否漏填"})
		return
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

	rows, err := model.DB.Queryx(`select ocp.*, oc.name, oc.unit_name from other_cost_patient ocp 
		left join clinic_other_cost coc on ocp.clinic_other_cost_id = coc.id 
		left join other_cost oc on coc.other_cost_id = oc.id
		where ocp.clinic_triage_patient_id = $1`, clinicTriagePatientID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "--1", "msg": err.Error()})
	}

	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}
