package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
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

	var results []map[string]string
	err := json.Unmarshal([]byte(items), &results)

	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
		return
	}

	timestamp := time.Now().Unix()
	random := strconv.FormatFloat(rand.Float64(), 'E', 10, 64)[2:5]
	orderSn := strconv.FormatInt(timestamp, 10) + random

	var sets []string

	for _, v := range results {

		var set []string

		registrationID := v["registration_id"]
		chargeProjecttypeid := v["charge_project_type_id"]
		chargeprojectid := v["charge_project_id"]
		name := v["name"]
		unit := v["unit"]
		price, _ := strconv.Atoi(v["price"])
		amount, _ := strconv.Atoi(v["amount"])
		total := price * amount
		discount, _ := strconv.Atoi(v["discount"])
		fee := total - discount
		operation := v["operation_id"]

		switch v["charge_project_type_id"] {
		case "8":
			rows := model.DB.QueryRowx("select * from charge_project_treatment where project_type_id=" + chargeProjecttypeid + " AND id=" + chargeprojectid)
			result := FormatSQLRowToMap(rows)
			name = result["name"].(string)
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

	sql := "INSERT INTO unpaid_orders (registration_id, charge_project_type_id, charge_project_id, operation_id, order_sn, name, unit, price, amount, total, discount, fee ) VALUES " + setst

	fmt.Println(sql)
	_, err1 := model.DB.Query(sql)

	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	ctx.JSON(iris.Map{"code": "200", "data": nil})

}
