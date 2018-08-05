package controller

import (
	"clinicSystemGo/model"

	"github.com/kataras/iris"
)

//PlatformTotalAmountDay 挂账还款
func PlatformTotalAmount(ctx iris.Context) {
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")

	typeCode := ctx.PostValue("typeCode")

	if startDate == "" || endDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	groupSQL := `to_char(cd.created_time, 'YYYY-MM-DD')`
	if typeCode == "2" {
		groupSQL = `to_char(cd.created_time, 'YYYY-MM')`
	}

	SQL := `select c.id  as clinic_id,c.name,` + groupSQL + ` as created_time,sum(cd.balance_money) as balance_money  from charge_detail cd 
	left join personnel p on cd.operation_id = p.id 
	left join clinic c on c.id = p.clinic_id where cd.created_time BETWEEN $1 and $2 group by ` + groupSQL + `,c.id,c.name order by created_time ASC,clinic_id ASC   
	`
	rows, _ := model.DB.Queryx(SQL, startDate, endDate)
	rowsMap := FormatSQLRowsToMapArray(rows)
	totalSQL := `select sum(balance_money) as total_money from charge_detail where created_time BETWEEN $1 and $2 `

	clinicSQL := `SELECT c.id,c.name FROM charge_detail cd 
	left join personnel p on cd.operation_id = p.id 
	left join clinic c on c.id = p.clinic_id where cd.created_time BETWEEN $1 and $2 group by c.id,c.name order by c.id ASC
	`
	row := model.DB.QueryRowx(totalSQL, startDate, endDate)
	rowMap := FormatSQLRowToMap(row)

	crows, _ := model.DB.Queryx(clinicSQL, startDate, endDate)
	crowMap := FormatSQLRowsToMapArray(crows)

	ctx.JSON(iris.Map{"code": "200", "data": rowsMap, "total": rowMap, "clinic": crowMap})
	return

}
