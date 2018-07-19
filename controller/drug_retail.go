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

// CreateDrugRetailOrder 创建药品零售订单
func CreateDrugRetailOrder(ctx iris.Context) {
	items := ctx.PostValue("items")
	if items == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}
	rand.Seed(time.Now().UnixNano())

	tradeNo := time.Now().Format("20060102150405") + strconv.Itoa((rand.Intn(8999) + 1000))

	sql := "INSERT INTO drug_retail_temp VALUES "

	var values []string
	for _, v := range results {
		var s []string

		s = append(s, tradeNo)
		s = append(s, v["clinic_drug_id"])
		s = append(s, v["amount"])
		s = append(s, v["total_fee"])

		str := strings.Join(s, ",")
		str = "(" + str + ")"
		values = append(values, str)
	}

	valueStr := strings.Join(values, ",")

	sql += valueStr

	_, erre := model.DB.Exec(sql)
	if erre != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": erre.Error()})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "msg": "ok", "data": tradeNo})

}

// CreateDrugRetailPaymentOrder 创建支付订单
func CreateDrugRetailPaymentOrder(ctx iris.Context) {
	outTradeNo := ctx.PostValue("out_trade_no")
	payMethod := ctx.PostValue("pay_method")
	authCode := ctx.PostValue("auth_code")
	totalMoney := ctx.PostValue("total_money")
	discountMoney := ctx.PostValue("discount_money")
	medicalMoney := ctx.PostValue("medical_money")
	balanceMoney := ctx.PostValue("balance_money")
	operationID := ctx.PostValue("operation_id")

	if outTradeNo == "" || payMethod == "" || discountMoney == "" || medicalMoney == "" || totalMoney == "" || balanceMoney == "" || operationID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if authCode == "" && (payMethod == "alipay" || payMethod == "wechat") {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少认证吗"})
		return
	}

	if payMethod == "alipay" {

	}

}
