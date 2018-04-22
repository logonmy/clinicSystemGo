package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris"
)

// ChargeUnPayCreate 创建收费列表
func ChargeUnPayCreate(ctx iris.Context) {

	items := ctx.PostValue("items")
	if items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	var results []map[string]interface{}
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	timestamp := time.Now().Unix()
	random := strconv.FormatFloat(rand.Float64(), 'E', 10, 64)[2:5]
	orderSn := strconv.FormatInt(timestamp, 10) + random

	var sets []string

	for index, v := range results {

		var set []string

		if v["registration_id"] == nil || v["charge_project_type_id"] == nil || v["charge_project_id"] == nil || v["amount"] == nil || v["discount"] == nil || v["operation_id"] == nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
			return
		}

		registrationID := strconv.FormatFloat(v["registration_id"].(float64), 'f', 0, 64)
		chargeProjecttypeid := strconv.FormatFloat(v["charge_project_type_id"].(float64), 'f', 0, 64)
		chargeprojectid := strconv.FormatFloat(v["charge_project_id"].(float64), 'f', 0, 64)
		amount := int(v["amount"].(float64))
		discount := int(v["discount"].(float64))
		operation := strconv.FormatFloat(v["operation_id"].(float64), 'f', 0, 64)

		var (
			name  string
			unit  string
			price int
			total int
			fee   int
		)

		switch chargeProjecttypeid {
		case "8":
			rows := model.DB.QueryRowx("select * from charge_project_treatment where project_type_id=" + chargeProjecttypeid + " AND id=" + chargeprojectid)
			result := FormatSQLRowToMap(rows)
			_, exsist := result["id"]
			if !exsist {
				ctx.JSON(iris.Map{"code": "-1", "msg": "未找到指定收费项 " + chargeprojectid})
				return
			}
			name = result["name"].(string)
			unit = "次"
			price = int(result["fee"].(int64))
			total = price * amount
			fee = total - discount
			break
		default:
			ctx.JSON(iris.Map{"code": "-1", "msg": "charge_project_type_id 无效"})
			return
		}

		set = append(set, registrationID)
		set = append(set, chargeProjecttypeid)
		set = append(set, chargeprojectid)
		set = append(set, operation)
		set = append(set, "'"+orderSn+"'")
		set = append(set, strconv.Itoa(index+1))
		set = append(set, "'"+name+"'")
		set = append(set, "'"+unit+"'")
		set = append(set, strconv.Itoa(price))
		set = append(set, strconv.Itoa(amount))
		set = append(set, strconv.Itoa(total))
		set = append(set, strconv.Itoa(discount))
		set = append(set, strconv.Itoa(fee))
		setn := "(" + strings.Join(set, ",") + ")"

		sets = append(sets, setn)
	}

	setst := strings.Join(sets, ",")

	sql := "INSERT INTO unpaid_orders (registration_id, charge_project_type_id, charge_project_id, operation_id, order_sn, soft_sn, name, unit, price, amount, total, discount, fee ) VALUES " + setst

	_, err1 := model.DB.Query(sql)

	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})

}

// ChargeUnPayDelete 删除缴费项目
func ChargeUnPayDelete(ctx iris.Context) {
	id := ctx.PostValue("id")
	if id == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	_, err := model.DB.Query("DELETE FROM unpaid_orders id=" + id)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})

}

// ChargeUnPayList 根据预约编码查询待缴费列表
func ChargeUnPayList(ctx iris.Context) {
	registrationid := ctx.PostValue("registration_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if registrationid == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if offset == "" {
		offset = "0"
	}

	if limit == "" {
		limit = "10"
	}

	_, err := strconv.Atoi(offset)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "offset 必须为数字"})
		return
	}
	_, err = strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "limit 必须为数字"})
		return
	}

	total := model.DB.QueryRowx(`select count(id) as total from unpaid_orders where registration_id=$1`, registrationid)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rowSQL := `select * from unpaid_orders where registration_id=$1 offset $2 limit $3`

	rows, err1 := model.DB.Queryx(rowSQL, registrationid, offset, limit)

	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}

	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})

}

// ChargePay 缴费
func ChargePay(ctx iris.Context) {

	orderSn := ctx.PostValue("order_sn")
	softSn := ctx.PostValue("soft_sn")
	confrimID := ctx.PostValue("confrim_id")
	payTypeCode := ctx.PostValue("pay_type_code")
	payMethodCode := ctx.PostValue("pay_method_code")
	balanceMoney := ctx.PostValue("balance_money")
	totalMoney := ctx.PostValue("total_money")

	discountRate := ctx.PostValue("discount_rate")

	derateMoney, _ := strconv.Atoi(ctx.PostValue("derate_money"))
	medicalMoney, _ := strconv.Atoi(ctx.PostValue("medical_money"))
	onCreditMoney, _ := strconv.Atoi(ctx.PostValue("on_credit_money"))
	voucherMoney, _ := strconv.Atoi(ctx.PostValue("voucher_money"))
	bonusPointsMoney, _ := strconv.Atoi(ctx.PostValue("bonus_points_money"))

	if orderSn == "" || softSn == "" || balanceMoney == "" || totalMoney == "" || confrimID == "" || payTypeCode == "" || payMethodCode == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	balanceMoneyInt, _ := strconv.Atoi(balanceMoney)
	totalMoneyInt, _ := strconv.Atoi(totalMoney)
	if discountRate == "" {
		discountRate = "100"
	}

	discountRateInt, _ := strconv.Atoi(discountRate)

	sBalance := totalMoneyInt*(discountRateInt/100) - (derateMoney + medicalMoney + onCreditMoney + voucherMoney + bonusPointsMoney)

	if sBalance != balanceMoneyInt {
		cmap := map[string]interface{}{
			"balance":      sBalance,
			"balanceMoney": balanceMoneyInt,
		}
		ctx.JSON(iris.Map{"code": "-1", "data": cmap, "msg": "实收款结算错误"})
		return
	}

	rows := model.DB.QueryRowx("select SUM(fee) from unpaid_orders where order_sn='" + orderSn + "' AND soft_sn in (" + softSn + ")")
	result := FormatSQLRowToMap(rows)
	var total int
	sum, exsist := result["sum"]
	if !exsist {
		total = 0
	} else {
		total = int(sum.(int64))
	}

	if total != totalMoneyInt {
		tmap := map[string]interface{}{
			"total":      total,
			"totalMoney": totalMoneyInt,
		}
		ctx.JSON(iris.Map{"code": "-1", "data": tmap, "msg": "应收收款结算错误"})
		return
	}

	tx, err := model.DB.Beginx()

	if err != nil {
		ctx.JSON(iris.Map{"code": "1", "msg": err.Error()})
		return
	}

	sql1 := "insert into paid_orders (id,registration_id,charge_project_type_id,charge_project_id,order_sn,soft_sn,name,price,amount,unit,total,discount,fee,operation_id,confrim_id)" +
		" select id,registration_id,charge_project_type_id,charge_project_id,order_sn,soft_sn,name,price,amount,unit,total,discount,fee,operation_id," + confrimID + " from unpaid_orders where order_sn='" + orderSn + "' AND soft_sn in (" + softSn + ")"

	_, errtx := tx.Query(sql1)
	if errtx != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errtx.Error()})
		return
	}

	sql2 := "insert into paid_record (soft_sns,order_sn,confrim_id,pay_type_code,pay_method_code,discount_rate,derate_money,medical_money,on_credit_money,voucher_money,bonus_points_money,total_money,balance_money) " +
		"values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) RETURNING id"
	var ID int
	errtx2 := tx.QueryRow(sql2, softSn, orderSn, confrimID, payTypeCode, payMethodCode, discountRateInt, derateMoney, medicalMoney, onCreditMoney, voucherMoney, bonusPointsMoney, totalMoneyInt, balanceMoneyInt).Scan(&ID)
	if errtx2 != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errtx2.Error()})
		return
	}

	sql3 := "DELETE from unpaid_orders where order_sn='" + orderSn + "' AND soft_sn in (" + softSn + ")"
	_, errtx3 := tx.Query(sql3)
	if errtx != nil {
		tx.Rollback()
		ctx.JSON(iris.Map{"code": "1", "msg": errtx3.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": ID})

}

// ChargePaidList 根据预约编码查询已缴费缴费列表
func ChargePaidList(ctx iris.Context) {
	registrationid := ctx.PostValue("registration_id")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")

	if registrationid == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if offset == "" {
		offset = "0"
	}

	if limit == "" {
		limit = "10"
	}

	_, err := strconv.Atoi(offset)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "offset 必须为数字"})
		return
	}
	_, err = strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "limit 必须为数字"})
		return
	}

	total := model.DB.QueryRowx(`select count(id) as total from paid_orders where registration_id=$1`, registrationid)

	pageInfo := FormatSQLRowToMap(total)
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rowSQL := `select * from paid_orders where registration_id=$1 offset $2 limit $3`

	rows, err1 := model.DB.Queryx(rowSQL, registrationid, offset, limit)

	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1})
		return
	}

	result := FormatSQLRowsToMapArray(rows)
	ctx.JSON(iris.Map{"code": "200", "data": result, "page_info": pageInfo})

}
