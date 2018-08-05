package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

//MaterialPatientCreate 开材料费
func MaterialPatientCreate(ctx iris.Context) {
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
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存材料费失败,分诊记录错误"})
		return
	}

	prow := model.DB.QueryRowx("select id from personnel where id=$1 limit 1", personnelID)
	if prow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "保存材料费失败,操作员错误"})
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

	clinicMaterialSets := []string{
		"clinic_triage_patient_id",
		"clinic_material_id",
		"order_sn",
		"soft_sn",
		"amount",
		"operation_id",
		"illustration",
	}

	orderSn := FormatPayOrderSn(clinicTriagePatientID, "5")

	tSetStr := strings.Join(clinicMaterialSets, ",")
	mSetStr := strings.Join(mzUnpaidOrdersSets, ",")

	tx, errb := model.DB.Begin()
	if errb != nil {
		fmt.Println("errb ===", errb)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errb})
		return
	}
	_, errdlp := tx.Exec("delete from material_patient where clinic_triage_patient_id=$1 and order_sn in (select order_sn from mz_unpaid_orders where clinic_triage_patient_id = $1 and charge_project_type_id=5)", clinicTriagePatientID)
	if errdlp != nil {
		fmt.Println("errdlp ===", errdlp)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errdlp.Error()})
		return
	}
	_, errdm := tx.Exec("delete from mz_unpaid_orders where clinic_triage_patient_id=$1 and charge_project_type_id=5", clinicTriagePatientID)
	if errdm != nil {
		fmt.Println("errdm ===", errdm)
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "-1", "msg": errdm.Error()})
		return
	}
	inserttSQL := "insert into material_patient (" + tSetStr + ") values ($1,$2,$3,$4,$5,$6,$7)"
	insertmSQL := "insert into mz_unpaid_orders (" + mSetStr + ") values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)"

	for index, v := range results {
		clinicMaterialID := v["clinic_material_id"]
		times := v["amount"]
		illustration := v["illustration"]
		fmt.Println("clinicMaterialID====", clinicMaterialID)
		if times == "" {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": "请填写数量"})
			return
		}

		clinicMaterialSQL := `select id as clinic_material_id,ret_price,discount_price,is_discount,name,unit_name from clinic_material where id=$1`
		trow := model.DB.QueryRowx(clinicMaterialSQL, clinicMaterialID)
		if trow == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "材料费项错误"})
			return
		}
		clinicMaterial := FormatSQLRowToMap(trow)
		fmt.Println("====", clinicMaterial)
		_, ok := clinicMaterial["clinic_material_id"]
		if !ok {
			ctx.JSON(iris.Map{"code": "-1", "msg": "选择的材料费项错误"})
			return
		}
		isDiscount := clinicMaterial["is_discount"].(bool)
		price := clinicMaterial["ret_price"].(int64)
		discountPrice := clinicMaterial["discount_price"].(int64)
		name := clinicMaterial["name"].(string)
		unitName := clinicMaterial["unit_name"].(string)
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
			ToNullInt64(clinicMaterialID),
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
			5,
			ToNullInt64(clinicMaterialID),
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

//MaterialPatientGet 获取材料费
func MaterialPatientGet(ctx iris.Context) {
	clinicTriagePatientID := ctx.PostValue("clinic_triage_patient_id")
	if clinicTriagePatientID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	rows, err := model.DB.Queryx(`select mp.*, cm.name, cm.specification, cm.unit_name, cm.ret_price as price, sum( ms.stock_amount ) as stock_amount, 
	case when mpo.id is not null then true else false end as paid_status 
	 from material_patient mp 
	left join clinic_material cm on mp.clinic_material_id = cm.id 
	left join material_stock ms on cm.id = ms.clinic_material_id 
	left join mz_paid_orders mpo on mpo.clinic_triage_patient_id = mp.clinic_triage_patient_id and mp.order_sn=mpo.order_sn and mp.soft_sn=mpo.soft_sn
	where mp.clinic_triage_patient_id = $1
	group by (mpo.id, mp.id, mp.clinic_triage_patient_id, mp.clinic_material_id, mp.order_sn, mp.soft_sn, mp.amount, mp.illustration, mp.operation_id, mp.created_time, cm.name, cm.specification, cm.unit_name,cm.ret_price)`, clinicTriagePatientID)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
	}

	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result})
}
