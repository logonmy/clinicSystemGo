package controller

import (
	"time"

	"clinicSystemGo/model"

	"github.com/kataras/iris"
)

//ChargeDayReportByPayWay 收费日报表按支付方式
func ChargeDayReportByPayWay(ctx iris.Context) {
	startDateStr := ctx.PostValue("start_date")
	endDateStr := ctx.PostValue("end_date")
	clinicID := ctx.PostValue("clinic_id")

	if startDateStr == "" || endDateStr == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}

	_, errs := time.Parse("2006-01-02", startDateStr)
	if errs != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "start_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	_, erre := time.Parse("2006-01-02", endDateStr)
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "end_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	querySQL := `select 
  c.id as clinic_id,
  c.name as clinic_name,
  sum(cd.total_money) as total_money,
  sum(cd.discount_money) as discount_money,
  sum(cd.balance_money) as balance_money,
  sum(cd.cash) as cash,
  sum(cd.bank) as bank,
  sum(cd.wechat) as wechat,
  sum(cd.alipay) as alipay,
  sum(cd.discount_money) as discount_money,
  sum(cd.medical_money) as medical_money,
  sum(cd.voucher_money) as voucher_money,
  sum(cd.bonus_points_money) as bonus_points_money,
  sum(cd.derate_money) as derate_money,
  sum(cd.on_credit_money) as on_credit_money
 from charge_detail cd 
left join personnel p on p.id = cd.operation_id
left join clinic c on p.clinic_id = c.id 
where cd.created_time between '2018-01-02' and '2018-08-05'`

	if clinicID != "" {
		querySQL += " and clinic_id = :clinic_id"
	}

	var queryOptions = map[string]interface{}{
		"clinic_id":  ToNullInt64(clinicID),
		"start_date": ToNullString(startDateStr),
		"end_date":   ToNullInt64(endDateStr),
	}

	rows, err := model.DB.NamedQuery(querySQL+" group by (c.id, c.name);", queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results})
}

//ChargeDayReportByBusiness 收费日报表按业务类型
func ChargeDayReportByBusiness(ctx iris.Context) {
	startDateStr := ctx.PostValue("start_date")
	endDateStr := ctx.PostValue("end_date")
	clinicID := ctx.PostValue("clinic_id")

	if startDateStr == "" || endDateStr == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "参数错误"})
		return
	}

	_, errs := time.Parse("2006-01-02", startDateStr)
	if errs != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "start_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	_, erre := time.Parse("2006-01-02", endDateStr)
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "end_date 必须为 YYYY-MM-DD 的 有效日期格式"})
		return
	}

	querySQL := `select 
  c.id as clinic_id,
  c.name as clinic_name,
  sum(cd.total_money) as total_money,
  sum(cd.discount_money) as discount_money,
  sum(cd.diagnosis_treatment_fee) as diagnosis_treatment_fee,
  sum(cd.traditional_medical_fee) as traditional_medical_fee,
  sum(cd.western_medicine_fee) as western_medicine_fee,
  sum(cd.examination_fee) as examination_fee,
  sum(cd.labortory_fee) as labortory_fee,
  sum(cd.treatment_fee) as treatment_fee,
  sum(cd.material_fee) as material_fee,
  sum(cd.other_fee) as other_fee,
  sum(cd.retail_fee) as retail_fee
 from charge_detail cd 
left join personnel p on p.id = cd.operation_id
left join clinic c on p.clinic_id = c.id 
where cd.created_time between '2018-01-02' and '2018-08-05'`

	if clinicID != "" {
		querySQL += " and clinic_id = :clinic_id"
	}

	var queryOptions = map[string]interface{}{
		"clinic_id":  ToNullInt64(clinicID),
		"start_date": ToNullString(startDateStr),
		"end_date":   ToNullInt64(endDateStr),
	}

	rows, err := model.DB.NamedQuery(querySQL+" group by (c.id, c.name);", queryOptions)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": results})
}