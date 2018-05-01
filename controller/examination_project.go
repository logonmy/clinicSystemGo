package controller

import (
	"clinicSystemGo/model"
	"time"

	"github.com/kataras/iris"
)

// ExaminationProjectCreate 创建检查缴费项目
func ExaminationProjectCreate(ctx iris.Context) {

	name := ctx.PostValue("name")
	price := ctx.PostValue("price")

	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	idcCode := ctx.PostValue("idc_code")
	unit := ctx.PostValue("unit")
	cost := ctx.PostValue("cost")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")
	organ := ctx.PostValue("organ")
	remark := ctx.PostValue("remark")

	if name == "" || price == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	sql := `INSERT INTO  examination_project ( name, en_name, py_code, idc_code, unit, cost, price, status, is_discount, organ, remark ) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id`

	var id int
	err := model.DB.QueryRow(sql, name, enName, pyCode, idcCode, unit, cost, price, status, isDiscount, organ, remark).Scan(&id)
	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": id})

}

// ExaminationProjectUpdate 更新检查缴费项目
func ExaminationProjectUpdate(ctx iris.Context) {

	ID := ctx.PostValue("id")

	name := ctx.PostValue("name")
	price := ctx.PostValue("price")
	enName := ctx.PostValue("en_name")
	pyCode := ctx.PostValue("py_code")
	idcCode := ctx.PostValue("idc_code")
	unit := ctx.PostValue("unit")
	cost := ctx.PostValue("cost")
	status := ctx.PostValue("status")
	isDiscount := ctx.PostValue("is_discount")
	organ := ctx.PostValue("organ")
	remark := ctx.PostValue("remark")

	if ID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	sql := `UPDATE examination_project SET name=$1, en_name=$2, py_code=$3, idc_code=$4, unit=$5, cost=$6, price=$7, status=$8, is_discount=$9, organ=$10, remark=$11, updated_time=$12 where id=$13`

	_, err := model.DB.Exec(sql, name, enName, pyCode, idcCode, unit, cost, price, status, isDiscount, organ, remark, time.Now(), ID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": nil})

}

// ExaminationProjectOnOff 启用和停用
func ExaminationProjectOnOff(ctx iris.Context) {
	ID := ctx.PostValue("id")
	status := ctx.PostValue("status")
	if ID == "" {
		ctx.JSON(iris.Map{"code": "1", "msg": "缺少参数"})
		return
	}

	sql := `UPDATE examination_project SET status=$1,updated_time=$2 where id=$3`

	_, err := model.DB.Exec(sql, status, time.Now(), ID)
	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": nil})

}

// ExaminationProjectList 查询
func ExaminationProjectList(ctx iris.Context) {

	keyword := ctx.PostValue("keyword")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}

	countSQL := `select count(id) as total from examination_project where name ~$1 or en_name ~$1 or py_code ~$1 or idc_code ~$1`

	selectSQL := `select * from examination_project where name ~$1 or en_name ~$1 or py_code ~$1 or idc_code ~$1 ORDER BY created_time DESC offset $2 limit $3`

	total := model.DB.QueryRowx(countSQL, keyword)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err1 := model.DB.Queryx(selectSQL, keyword, offset, limit)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}
	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})

}
