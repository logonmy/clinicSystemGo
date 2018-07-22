package controller

import (
	"clinicSystemGo/model"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/kataras/iris"
)

// DrugRetailList 获取药品零售表
func DrugRetailList(ctx iris.Context) {

	refundStatus := ctx.PostValue("refundStatus")
	offset := ctx.PostValue("offset")
	limit := ctx.PostValue("limit")
	startDate := ctx.PostValue("start_date")
	endDate := ctx.PostValue("end_date")

	if refundStatus == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "6"
	}

	if startDate == "" || endDate == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "请输入正确的时间范围"})
		return
	}

	queryMap := map[string]interface{}{
		"offset":    ToNullInt64(offset),
		"limit":     ToNullInt64(limit),
		"startDate": ToNullString(startDate),
		"endDate":   ToNullString(endDate),
	}

	sql := `from drug_retail_pay_record pr 
	left join drug_retail_refund_record rr on pr.out_trade_no = rr.out_trade_no 
	left join personnel p on p.id = pr.operation_id 
	where pr.pay_time between :startDate AND :endDate and pr.status = 2 
	group by pr.out_trade_no,pr.pay_method,p.name,pr.balance_money,pr.pay_time `

	querySQL := `SELECT * FROM (SELECT pr.out_trade_no,pr.pay_method,p.name,pr.balance_money,pr.pay_time,sum(rr.refund_money) as  refund_money ` + sql + `) AS u `
	countSQL := `SELECT COUNT(*) AS total from (select pr.out_trade_no, sum(rr.refund_money) as refund_money  ` + sql + `) as u `

	if refundStatus == "2" {
		querySQL += `where u.refund_money > 0 `
		countSQL += `where u.refund_money > 0 `
	}

	total, err2 := model.DB.NamedQuery(countSQL, queryMap)

	if err2 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err2.Error()})
		return
	}

	pageInfo := FormatSQLRowsToMapArray(total)[0]
	pageInfo["offset"] = offset
	pageInfo["limit"] = limit

	rows, err1 := model.DB.NamedQuery(querySQL+` order BY pay_time DESC offset :offset limit :limit`, queryMap)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}
	results := FormatSQLRowsToMapArray(rows)

	ctx.JSON(iris.Map{"code": "200", "data": results, "page_info": pageInfo})

}

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

	queryMap := map[string]interface{}{
		"outTradeNo":    ToNullString(outTradeNo),
		"payMethod":     ToNullString(payMethod),
		"authCode":      ToNullString(authCode),
		"status":        -1,
		"totalMoney":    ToNullInt64(totalMoney),
		"discountMoney": ToNullInt64(discountMoney),
		"medicalMoney":  ToNullInt64(medicalMoney),
		"balanceMoney":  ToNullInt64(balanceMoney),
		"operationID":   ToNullInt64(operationID),
	}

	row := model.DB.QueryRowx("select out_trade_no, sum(total_fee) as fee from drug_retail_temp where out_trade_no = $1 GROUP by out_trade_no ", outTradeNo)
	rowMap := FormatSQLRowToMap(row)

	if strconv.FormatInt(rowMap["fee"].(int64), 10) != totalMoney {
		ctx.JSON(iris.Map{"code": "-1", "msg": "金额不一致"})
		return
	}

	requestIP := ctx.Host()
	requestIP = requestIP[0:strings.LastIndex(requestIP, ":")]

	model.DB.NamedExec("DELETE from drug_retail_pay_record where out_trade_no = :outTradeNo", queryMap)

	_, err1 := model.DB.NamedExec("INSERT INTO drug_retail_pay_record (out_trade_no,pay_method,auth_code,total_money,discount_money,medical_money,balance_money,operation_id) VALUES (:outTradeNo,:payMethod,:authCode,:totalMoney,:discountMoney,:medicalMoney,:balanceMoney,:operationID)", queryMap)
	if err1 != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
		return
	}

	if payMethod == "alipay" || payMethod == "wechat" {

		merID := "ali"
		payModel := "alipay_f2f"
		if payMethod == "wechat" {
			merID = "wx"
			payModel = "weixin_f2f"
		}

		fmt.Println(outTradeNo, authCode, merID, payModel, balanceMoney, requestIP)

		result := FaceToFace(outTradeNo, authCode, merID, payModel, "ls", balanceMoney, "药品零售", "127.0.0.1", "", "")
		if result["code"].(string) != "200" {
			if result["code"].(string) == "2" {
				_, err := model.DB.Exec("update drug_retail_pay_record set status = 1 where out_trade_no = $1", outTradeNo)
				if err != nil {
					ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
					return
				}
				ctx.JSON(iris.Map{"code": "300", "msg": result["msg"]})
			} else {
				model.DB.Exec("update drug_retail_pay_record set status = 3 where out_trade_no = $1", outTradeNo)
				ctx.JSON(iris.Map{"code": result["code"], "msg": result["msg"]})
			}
		} else {
			err := paySuccessNotice(outTradeNo)
			if err != nil {
				ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
				return
			}
			ctx.JSON(iris.Map{"code": "200", "data": 1})
		}
	} else if payMethod == "cash" {
		err := paySuccessNotice(outTradeNo)
		if err != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": err.Error()})
			return
		}
		ctx.JSON(iris.Map{"code": "200", "data": 1})
	} else {
		ctx.JSON(iris.Map{"code": "-1", "msg": "不支持的支付方式"})
	}
}

// 支付成功后通知
func paySuccessNotice(outTradeNo string) error {
	fmt.Println(outTradeNo)
	payRecordRow := model.DB.QueryRowx("SELECT * FROM drug_retail_pay_record WHERE out_trade_no = $1", outTradeNo)
	payRecordMap := FormatSQLRowToMap(payRecordRow)
	if payRecordMap["out_trade_no"] == nil {
		return errors.New("未找到指定的缴费记录")
	}

	if payRecordMap["status"].(int64) > 1 {
		return errors.New("订单已处理过")
	}

	_, err := model.DB.Exec("update drug_retail_pay_record set status = 2,pay_time = LOCALTIMESTAMP where out_trade_no = $1", outTradeNo)
	if err != nil {
		return err
	}

	rows, _ := model.DB.Queryx("SELECT * FROM drug_retail_temp WHERE out_trade_no = $1", outTradeNo)
	rowsMap := FormatSQLRowsToMapArray(rows)

	tx, txErr := model.DB.Beginx()
	if txErr != nil {
		return txErr
	}

	for _, item := range rowsMap {

		clinicDrugID := item["clinic_drug_id"]
		amount := item["amount"]
		price := item["total_fee"].(int64) / item["amount"].(int64)
		uperr := updateDrugStock(tx, clinicDrugID.(int64), amount.(int64), outTradeNo, price)

		if uperr != nil {
			return uperr
		}

	}

	crr := tx.Commit()
	if crr != nil {
		return err
	}

	return nil
}

func updateDrugStock(tx *sqlx.Tx, clinicDrugID int64, amount int64, outTradeNo string, price int64) error {
	if amount < 0 {
		return errors.New("库存数量有误")
	}
	if amount == 0 {
		return nil
	}

	timeNow := time.Now().Format("2006-01-02")
	fmt.Println("clinicDrugID", clinicDrugID)
	row := model.DB.QueryRowx("select * from drug_stock where clinic_drug_id = $1 and stock_amount > 0 and eff_date > $2 ORDER by created_time asc limit 1", clinicDrugID, timeNow)
	rowMap := FormatSQLRowToMap(row)
	if rowMap["stock_amount"] == nil {
		return errors.New("库存不足")
	}

	stockAmount := rowMap["stock_amount"].(int64)

	if stockAmount >= amount {
		_, err := tx.Exec("update drug_stock set stock_amount = $1 where id = $2", stockAmount-amount, rowMap["id"])
		if err != nil {
			tx.Rollback()
			return err
		}
		_, errIn := tx.Exec("insert into drug_retail (out_trade_no,clinic_drug_id,drug_stock_id,amount,total_fee) VALUES ($1,$2,$3,$4,$5)", outTradeNo, clinicDrugID, rowMap["id"], amount, price*amount)
		if errIn != nil {
			tx.Rollback()
			return errIn
		}

		return nil
	} else {
		_, err := tx.Exec("update drug_stock set 0 where id = $1", rowMap["id"])
		if err != nil {
			tx.Rollback()
			return err
		}

		_, errIn := tx.Exec("insert into drug_retail (out_trade_no,clinic_drug_id,drug_stock_id,amount,total_fee) VALUES ($1,$2,$3,$4,$5)", outTradeNo, clinicDrugID, rowMap["id"], stockAmount, price*stockAmount)
		if errIn != nil {
			tx.Rollback()
			return errIn
		}

		return updateDrugStock(tx, clinicDrugID, amount-stockAmount, outTradeNo, price)
	}
}
