package controller

import (
	"bytes"
	"clinicSystemGo/lib/hcpay"
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/kataras/iris"
)

var partnerID = "GZHXDMD"
var partnerKey = "C1F013D2B90E45E9995E16A3411A6910"
var hcMerchantID = map[string]string{"wx": "1260939901", "ali": "2017081008129270"}
var hcOrdertype = map[string]string{"normal": "hcpay.trade.unifiedorder", "public": "hcpay.trade.registerorder"}

// var url = "https://pay.med.gzhc365.com/api/hcpay/gateway"
var url = "https://upay.med.gzhc365.com/api/hcpay/gateway"

//CreateHcOrder 创建订单
func CreateHcOrder(ctx iris.Context) {
	payMode := ctx.PostValue("pay_mode")
	subjectType := ctx.PostValue("subject_type")
	subjectID := ctx.PostValue("subject_id")
	subjectName := ctx.PostValue("subject_name")
	openID := ctx.PostValue("openid")
	logonAccount := ctx.PostValue("logon_account")
	totalFee := ctx.PostValue("total_fee")
	body := ctx.PostValue("body")
	goodsDetail := ctx.PostValue("goods_detail")
	deviceInfo := ctx.PostValue("device_info")
	merchantID := ctx.PostValue("merchant_id")
	orderType := ctx.PostValue("order_type")

	if payMode == "" || totalFee == "" || body == "" || merchantID == "" || orderType == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	outTradeNo := hcpay.CreateTradeNo(20)

	requestIP := "47.93.206.157"
	// requestIP := ctx.Host()
	// requestIP = requestIP[0:strings.LastIndex(requestIP, ":")]
	noncestr := hcpay.GenerateNonceString(32)

	var m map[string]string
	m = make(map[string]string, 0)
	m["service_code"] = hcOrdertype[orderType]
	m["merchant_id"] = hcMerchantID[merchantID]
	m["partner_id"] = partnerID
	m["nonce_str"] = noncestr
	m["sign_type"] = "MD5"
	m["pay_mode"] = payMode
	m["out_trade_no"] = outTradeNo
	m["openid"] = openID
	m["logon_account"] = logonAccount
	m["time_start"] = time.Now().Format("20060102150405")
	m["time_expire"] = time.Now().Add(5 * time.Minute).Format("20060102150405")
	m["body"] = body
	m["notify_url"] = "notify_url"
	m["callback_url"] = "callback_url"
	m["create_ip"] = requestIP
	m["total_fee"] = totalFee
	m["subject_type"] = subjectType
	m["subject_id"] = subjectID
	m["subject_name"] = subjectName
	m["goods_detail"] = goodsDetail
	m["device_info"] = deviceInfo
	sign := hcpay.GetSign(m, partnerKey)
	m["sign"] = sign

	resData := request("POST", m)

	if resData["code"] == "error" {
		ctx.JSON(iris.Map{"code": "-1", "msg": resData["msg"]})
	} else {
		ctx.JSON(iris.Map{"code": "200", "data": resData})
	}
}

//QueryHcOrder 订单查询
func QueryHcOrder(ctx iris.Context) {
	outTradeNo := ctx.PostValue("out_trade_no")
	transactionNo := ctx.PostValue("trade_no")
	merchantID := ctx.PostValue("merchant_id")

	if outTradeNo == "" || merchantID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	noncestr := hcpay.GenerateNonceString(32)
	var tradeNo string
	var openID string
	var payTime string
	payState := false
	var m map[string]string
	m = make(map[string]string, 0)
	m["service_code"] = "hcpay.trade.orderquery"
	m["merchant_id"] = hcMerchantID[merchantID]
	m["partner_id"] = partnerID
	m["nonce_str"] = noncestr
	m["sign_type"] = "MD5"
	m["transaction_id"] = transactionNo
	m["out_trade_no"] = outTradeNo
	sign := hcpay.GetSign(m, partnerKey)
	m["sign"] = sign

	updateSQL := `update pay_order set 
		trade_no=$2,pay_state=$3,openid=$4,pay_time=$5 where out_trade_no=$1`

	resData := request("POST", m)

	if resData["code"] == "error" {
		ctx.JSON(iris.Map{"code": "-1", "msg": resData["msg"]})
	} else {
		tradeNo = resData["transaction_id"].(string)
		openID = resData["openid"].(string)
		payTime = resData["payed_time"].(string)
		if resData["trade_status"] == "SUCCESS" {
			payState = true
		}
		_, err := model.DB.Exec(updateSQL, outTradeNo, tradeNo, payState, openID, payTime)
		if err != nil {
			fmt.Println("err1 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
		ctx.JSON(iris.Map{"code": "200", "data": resData})
	}
}

//HcRefund 退款
func HcRefund(ctx iris.Context) {
	outTradeNo := ctx.PostValue("out_trade_no")
	transactionNo := ctx.PostValue("trade_no")
	refundFeeStr := ctx.PostValue("refund_fee")
	refundReason := ctx.PostValue("refund_reason")
	merchantID := ctx.PostValue("merchant_id")

	if outTradeNo == "" || merchantID == "" || refundFeeStr == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	refundFee, _ := strconv.ParseInt(refundFeeStr, 10, 64)

	crow := model.DB.QueryRowx("select out_trade_no,total_fee,refund_fee_total from pay_order where out_trade_no=$1 limit 1", outTradeNo)
	if crow == nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": "退款失败"})
		return
	}
	payOrder := FormatSQLRowToMap(crow)
	_, cok := payOrder["out_trade_no"]
	if !cok {
		ctx.JSON(iris.Map{"code": "-1", "msg": "支付订单不存在"})
		return
	}
	var refundState string
	totalFee := payOrder["total_fee"].(int64)
	refundFeeTotal := payOrder["refund_fee_total"].(int64)
	if refundFee < 0 {
		ctx.JSON(iris.Map{"code": "-1", "msg": "退款金额小于0"})
		return
	}
	if totalFee < refundFee+refundFeeTotal {
		ctx.JSON(iris.Map{"code": "-1", "msg": "总退款金额大于支付金额"})
		return
	}
	if totalFee > refundFee+refundFeeTotal {
		refundState = "20"
	}
	if totalFee == refundFee+refundFeeTotal {
		refundState = "30"
	}

	noncestr := hcpay.GenerateNonceString(32)
	outRefundNo := hcpay.CreateTradeNo(20)
	var refundTradeNo string

	var m map[string]string
	m = make(map[string]string, 0)
	m["service_code"] = "hcpay.trade.refund"
	m["merchant_id"] = hcMerchantID[merchantID]
	m["partner_id"] = partnerID
	m["nonce_str"] = noncestr
	m["sign_type"] = "MD5"
	m["transaction_id"] = transactionNo
	m["out_trade_no"] = outTradeNo
	m["out_refund_no"] = outRefundNo
	m["refund_fee"] = refundFeeStr
	sign := hcpay.GetSign(m, partnerKey)
	m["sign"] = sign

	updateSQL := `update pay_order set 
	refund_state=$2,refund_fee_total=refund_fee_total+$3 where out_trade_no=$1`

	insertSQL := `INSERT INTO refund_order (
		out_trade_no,refund_fee,refund_reason,out_refund_no,refund_trade_no,refund_result) 
		VALUES 
		($1,$2,$3,$4,$5,$6)`

	resData := request("POST", m)

	if resData["code"] == "error" {
		ctx.JSON(iris.Map{"code": "-1", "msg": resData["msg"]})
	} else {
		mjson, _ := json.Marshal(m)
		refundResult := string(mjson)
		refundTradeNo = resData["refund_id"].(string)
		tx, txErr := model.DB.Beginx()
		if txErr != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": txErr.Error()})
			return
		}
		_, err1 := tx.Exec(updateSQL, outTradeNo, refundState, refundFee)
		if err1 != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": err1.Error()})
			return
		}

		_, erru := tx.Exec(insertSQL, outTradeNo, refundFee, refundReason, outRefundNo, refundTradeNo, refundResult)
		if erru != nil {
			tx.Rollback()
			ctx.JSON(iris.Map{"code": "-1", "msg": erru.Error()})
			return
		}

		erre := tx.Commit()
		if erre != nil {
			ctx.JSON(iris.Map{"code": "-1", "msg": erre.Error()})
			return
		}

		ctx.JSON(iris.Map{"code": "200", "data": resData})
	}
}

//QueryHcRefund 退款查询
func QueryHcRefund(ctx iris.Context) {
	outTradeNo := ctx.PostValue("out_trade_no")
	transactionNo := ctx.PostValue("trade_no")
	outRefundNo := ctx.PostValue("out_refund_no")
	refundID := ctx.PostValue("refund_trade_no")
	merchantID := ctx.PostValue("merchant_id")

	if merchantID == "" || (outTradeNo == "" && transactionNo == "" && outRefundNo == "" && refundID == "") {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	noncestr := hcpay.GenerateNonceString(32)

	var m map[string]string
	m = make(map[string]string, 0)
	m["service_code"] = "hcpay.trade.refundquery"
	m["merchant_id"] = hcMerchantID[merchantID]
	m["partner_id"] = partnerID
	m["nonce_str"] = noncestr
	m["sign_type"] = "MD5"
	m["transaction_id"] = transactionNo
	m["out_trade_no"] = outTradeNo
	m["out_refund_no"] = outRefundNo
	m["refund_id"] = refundID
	sign := hcpay.GetSign(m, partnerKey)
	m["sign"] = sign

	resData := request("POST", m)

	if resData["code"] == "error" {
		ctx.JSON(iris.Map{"code": "-1", "msg": resData["msg"]})
	} else {
		ctx.JSON(iris.Map{"code": "200", "data": resData})
	}
}

//HcOrderClose 订单关闭
func HcOrderClose(ctx iris.Context) {
	outTradeNo := ctx.PostValue("out_trade_no")
	merchantID := ctx.PostValue("merchant_id")

	if merchantID == "" || outTradeNo == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	noncestr := hcpay.GenerateNonceString(32)

	var m map[string]string
	m = make(map[string]string, 0)
	m["service_code"] = "hcpay.trade.close"
	m["merchant_id"] = hcMerchantID[merchantID]
	m["partner_id"] = partnerID
	m["nonce_str"] = noncestr
	m["sign_type"] = "MD5"
	m["out_trade_no"] = outTradeNo
	sign := hcpay.GetSign(m, partnerKey)
	m["sign"] = sign

	resData := request("POST", m)

	if resData["code"] == "error" {
		ctx.JSON(iris.Map{"code": "-1", "msg": resData["msg"]})
	} else {
		ctx.JSON(iris.Map{"code": "200", "data": resData})
	}
}

//FaceToFace 当面付
func FaceToFace(ctx iris.Context) {
	payMode := ctx.PostValue("pay_mode")
	totalFee := ctx.PostValue("total_fee")
	body := ctx.PostValue("body")
	goodsDetail := ctx.PostValue("goods_detail")
	deviceInfo := ctx.PostValue("device_info")
	authCode := ctx.PostValue("auth_code")
	merchantID := ctx.PostValue("merchant_id")

	if payMode == "" || totalFee == "" || body == "" || merchantID == "" || authCode == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	var tradeNo string
	var openID string
	var payTime string
	var originalData string
	outTradeNo := hcpay.CreateTradeNo(20)

	requestIP := "47.93.206.157"
	// requestIP := ctx.Host()
	// requestIP = requestIP[0:strings.LastIndex(requestIP, ":")]
	noncestr := hcpay.GenerateNonceString(32)

	var m map[string]string
	m = make(map[string]string, 0)
	m["service_code"] = "hcpay.trade.f2f"
	m["merchant_id"] = hcMerchantID[merchantID]
	m["partner_id"] = partnerID
	m["auth_code"] = authCode
	m["nonce_str"] = noncestr
	m["sign_type"] = "MD5"
	m["pay_mode"] = payMode
	m["out_trade_no"] = outTradeNo
	m["body"] = body
	m["create_ip"] = requestIP
	m["total_fee"] = totalFee
	m["goods_detail"] = goodsDetail
	m["device_info"] = deviceInfo
	sign := hcpay.GetSign(m, partnerKey)
	m["sign"] = sign

	mjson, _ := json.Marshal(m)
	originalData = string(mjson)

	resData := request("POST", m)
	insertSQL := `INSERT INTO pay_order (
		trade_no,out_trade_no,total_fee,body,pay_state,original_data,openid,trade_type,pay_time) 
		VALUES 
		($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	if resData["code"] == "error" {
		_, err := model.DB.Exec(insertSQL, tradeNo, outTradeNo, totalFee, body, false, originalData, openID, payMode, payTime)
		if err != nil {
			fmt.Println("err1 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
		ctx.JSON(iris.Map{"code": "-1", "msg": resData["msg"]})
	} else {
		tradeNo = resData["transaction_id"].(string)
		openID = resData["openid"].(string)
		payTime = resData["payed_time"].(string)
		_, err := model.DB.Exec(insertSQL, tradeNo, outTradeNo, totalFee, body, true, originalData, openID, payMode, payTime)
		if err != nil {
			fmt.Println("err1 ===", err)
			ctx.JSON(iris.Map{"code": "-1", "msg": err})
			return
		}
		ctx.JSON(iris.Map{"code": "200", "data": resData})
	}
}

//FaceToFaceCancel 当面付支付撤销
func FaceToFaceCancel(ctx iris.Context) {
	outTradeNo := ctx.PostValue("out_trade_no")
	merchantID := ctx.PostValue("merchant_id")

	if merchantID == "" || outTradeNo == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	noncestr := hcpay.GenerateNonceString(32)

	var m map[string]string
	m = make(map[string]string, 0)
	m["service_code"] = "hcpay.trade.reverse"
	m["merchant_id"] = hcMerchantID[merchantID]
	m["partner_id"] = partnerID
	m["nonce_str"] = noncestr
	m["sign_type"] = "MD5"
	m["out_trade_no"] = outTradeNo
	sign := hcpay.GetSign(m, partnerKey)
	m["sign"] = sign

	resData := request("POST", m)

	if resData["code"] == "error" {
		ctx.JSON(iris.Map{"code": "-1", "msg": resData["msg"]})
	} else {
		ctx.JSON(iris.Map{"code": "200", "data": resData})
	}
}

//DownloadBill 下载对账单
func DownloadBill(ctx iris.Context) {
	merchantID := ctx.PostValue("merchant_id")
	payChannel := ctx.PostValue("pay_channel")
	billDate := ctx.PostValue("bill_date")
	numPerPage := ctx.PostValue("num_per_page")
	pageNo := ctx.PostValue("page_no")

	if merchantID == "" || payChannel == "" || billDate == "" || numPerPage == "" || pageNo == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	noncestr := hcpay.GenerateNonceString(32)

	var m map[string]string
	m = make(map[string]string, 0)
	m["service_code"] = "hcpay.bills.download"
	m["merchant_id"] = hcMerchantID[merchantID]
	m["partner_id"] = partnerID
	m["nonce_str"] = noncestr
	m["sign_type"] = "MD5"
	m["pay_channel"] = payChannel
	m["bill_date"] = payChannel
	m["num_per_page"] = numPerPage
	m["page_no"] = pageNo
	sign := hcpay.GetSign(m, partnerKey)
	m["sign"] = sign

	resData := request("POST", m)

	if resData["code"] == "error" {
		ctx.JSON(iris.Map{"code": "-1", "msg": resData["msg"]})
	} else {
		ctx.JSON(iris.Map{"code": "200", "data": resData})
	}
}

//request 请求方法
func request(method string, m map[string]string) map[string]interface{} {
	post, er := json.Marshal(m)
	if er != nil {
		fmt.Println("er=========", er)
		return nil
	}

	var jsonStr = []byte(post)
	fmt.Println("new_str", bytes.NewBuffer(jsonStr))

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	resData, _ := ioutil.ReadAll(resp.Body)

	var results map[string]interface{}
	errb := json.Unmarshal(resData, &results)

	if errb != nil {
		fmt.Println("errb=========", errb)
		return nil
	}
	fmt.Println("results=========", results)

	if results["return_code"] == "0" {
		if results["result_code"] == "0" {
			return results
		}
		return map[string]interface{}{"code": "error", "msg": results["result_msg"]}
	}
	return map[string]interface{}{"code": "error", "msg": results["return_msg"]}
}