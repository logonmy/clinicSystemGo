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
)

var partnerID = "GZHXDMD"
var partnerKey = "C1F013D2B90E45E9995E16A3411A6910"

var hcMerchantID = map[string]string{"wx": "1260939901", "ali": "2017081008129270"}
var hcOrdertype = map[string]string{"normal": "hcpay.trade.unifiedorder", "public": "hcpay.trade.registerorder"}

// var url = "https://pay.med.gzhc365.com/api/hcpay/gateway"

//测试地址
var url = "https://upay.med.gzhc365.com/api/hcpay/gateway"

//FaceToFace 当面付
/**
 * @param outTradeNo 系统交易号 必须
 * @param authCode 付款码 必须
 * @param merchantID 支付类型 wx 微信 ali 支付宝 必须
 * @param payMode 支付模式 weixin_f2f –微信当面付 alipay_f2f – 支付宝当面付 mybank_weixin_f2f 网商微信当面付 mybank_alipay_f2f 网商支付宝当面付 必须
 * @param businessType 业务类型 必须
 * @param totalFee 费用 必须
 * @param body 商品描述 必须
 * @param requestIP 请求地址 必须
 * @param goodsDetail 商品详情 非必须
 * @param deviceInfo 设备信息 非必须
 */
func FaceToFace(outTradeNo string, authCode string, merchantID string, payMode string, businessType string, totalFee string, body string, requestIP string, goodsDetail string, deviceInfo string) map[string]interface{} {
	if outTradeNo == "" || businessType == "" || payMode == "" || totalFee == "" || body == "" || merchantID == "" || authCode == "" || requestIP == "" {
		return map[string]interface{}{"code": "-1", "msg": "缺少参数"}
	}

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
	originalData := string(mjson)

	insertSQL := `INSERT INTO pay_order (
		out_trade_no,total_fee,body,order_status,original_data,trade_type,business_type,merchant_id) 
		VALUES 
		($1,$2,$3,$4,$5,$6,$7,$8)`

	updateSQL := `update pay_order set 
	trade_no=$2,order_status=$3,openid=$4,pay_time=$5 where out_trade_no=$1`

	_, err := model.DB.Exec(insertSQL, outTradeNo, totalFee, body, "NOTPAY", originalData, payMode, businessType, merchantID)
	if err != nil {
		fmt.Println("err ===", err)
		return map[string]interface{}{"code": "-1", "msg": err.Error()}
	}

	resData := request("POST", m)

	if resData["code"] == "error" {
		if resData["data"] != nil && resData["data"].(string) == "2" {
			_, err := model.DB.Exec("update pay_order set order_status=$2 where out_trade_no=$1", outTradeNo, "USERPAYING")
			if err != nil {
				fmt.Println("err ===", err)
				return map[string]interface{}{"code": "-1", "msg": err.Error()}
			}
			return map[string]interface{}{"code": "2", "msg": resData["msg"]}
		}

		_, errf := model.DB.Exec("update pay_order set order_status=$2 where out_trade_no=$1", outTradeNo, "FAIL")
		if errf != nil {
			fmt.Println("errf ===", errf)
			return map[string]interface{}{"code": "-1", "msg": errf.Error()}
		}

		return map[string]interface{}{"code": "-1", "msg": resData["msg"]}
	}
	tradeNo := resData["transaction_id"]
	openID := resData["openid"]
	payTime := resData["payed_time"]

	_, erre := model.DB.Exec(updateSQL, outTradeNo, tradeNo, "SUCCESS", openID, payTime)
	if erre != nil {
		fmt.Println("erre ===", erre)
		return map[string]interface{}{"code": "-1", "msg": erre.Error()}
	}
	return map[string]interface{}{"code": "200", "data": resData}
}

//FaceToFaceCancel 当面付支付撤销
/**
 * @param outTradeNo 系统交易号 必须
 */
func FaceToFaceCancel(outTradeNo string) map[string]interface{} {

	if outTradeNo == "" {
		return map[string]interface{}{"code": "-1", "msg": "缺少参数"}
	}

	selectSQL := `select out_trade_no,merchant_id from pay_order where out_trade_no=$1`

	row := model.DB.QueryRowx(selectSQL, outTradeNo)
	if row == nil {
		return map[string]interface{}{"code": "-1", "msg": "支付订单查询失败"}
	}

	payOrder := FormatSQLRowToMap(row)
	_, ok := payOrder["out_trade_no"]
	if !ok {
		return map[string]interface{}{"code": "-1", "msg": "支付订单不存在"}
	}

	noncestr := hcpay.GenerateNonceString(32)

	var m map[string]string
	m = make(map[string]string, 0)
	m["service_code"] = "hcpay.trade.reverse"
	m["merchant_id"] = hcMerchantID[payOrder["merchant_id"].(string)]
	m["partner_id"] = partnerID
	m["nonce_str"] = noncestr
	m["sign_type"] = "MD5"
	m["out_trade_no"] = outTradeNo
	sign := hcpay.GetSign(m, partnerKey)
	m["sign"] = sign

	updateSQL := `update pay_order set order_status=$2,updated_time=LOCALTIMESTAMP where out_trade_no=$1`

	resData := request("POST", m)

	if resData["code"] == "error" {
		if resData["data"] != nil && resData["data"].(string) == "2" {
			_, err := model.DB.Exec("update pay_order set order_status=$2,updated_time=LOCALTIMESTAMP where out_trade_no=$1", outTradeNo, "UNKNOW")
			if err != nil {
				fmt.Println("err ===", err)
				return map[string]interface{}{"code": "-1", "msg": err.Error()}
			}
			return map[string]interface{}{"code": "2", "msg": resData["msg"]}
		}
		return map[string]interface{}{"code": resData["code"], "msg": resData["msg"]}
	}
	_, err := model.DB.Exec(updateSQL, outTradeNo, "CLOSE")
	if err != nil {
		fmt.Println("err ===", err)
		return map[string]interface{}{"code": "-1", "msg": err.Error()}
	}
	return map[string]interface{}{"code": "200", "data": resData}
}

//HcRefund 退款
/**
 * @param outTradeNo 系统交易号 必须
 * @param refundFeeStr 退款金额 必须
 * @param outRefundNo 系统退款请求交易号 必须
 * @param transactionNo 支付平台交易号 非必须
 * @param refundReason 退款原因 非必须
 */
func HcRefund(outTradeNo string, refundFeeStr string, outRefundNo string, transactionNo string, refundReason string) map[string]interface{} {

	if outRefundNo == "" || outTradeNo == "" || refundFeeStr == "" {
		return map[string]interface{}{"code": "-1", "msg": "缺少参数"}
	}

	crow := model.DB.QueryRowx("select out_trade_no,total_fee,merchant_id from pay_order where out_trade_no=$1 limit 1", outTradeNo)
	if crow == nil {
		return map[string]interface{}{"code": "-1", "msg": "退款失败"}
	}
	payOrder := FormatSQLRowToMap(crow)
	_, cok := payOrder["out_trade_no"]
	if !cok {
		return map[string]interface{}{"code": "-1", "msg": "支付订单不存在"}
	}

	refundFee, _ := strconv.ParseInt(refundFeeStr, 10, 64)
	refundFeeTotal := int64(0)

	if refundFee < 0 {
		return map[string]interface{}{"code": "-1", "msg": "退款金额小于0"}
	}

	rrow := model.DB.QueryRowx(`select out_trade_no,sum(refund_fee) as refund_fee_total
		from refund_order
		where out_trade_no=$1 and refund_status in ('PROCESSING','SUCCESS')
		group by out_trade_no`, outTradeNo)
	if rrow == nil {
		return map[string]interface{}{"code": "-1", "msg": "退款失败"}
	}
	refundOrder := FormatSQLRowToMap(rrow)

	_, ok := refundOrder["out_trade_no"]
	if ok {
		refundFeeTotal = refundOrder["refund_fee_total"].(int64)
	}

	totalFee := payOrder["total_fee"].(int64)

	if totalFee < refundFee+refundFeeTotal {
		return map[string]interface{}{"code": "-1", "msg": "总退款金额大于支付金额"}
	}

	noncestr := hcpay.GenerateNonceString(32)

	var m map[string]string
	m = make(map[string]string, 0)
	m["service_code"] = "hcpay.trade.refund"
	m["merchant_id"] = hcMerchantID[payOrder["merchant_id"].(string)]
	m["partner_id"] = partnerID
	m["nonce_str"] = noncestr
	m["sign_type"] = "MD5"
	m["transaction_id"] = transactionNo
	m["out_trade_no"] = outTradeNo
	m["out_refund_no"] = outRefundNo
	m["refund_fee"] = refundFeeStr
	sign := hcpay.GetSign(m, partnerKey)
	m["sign"] = sign

	insertSQL := `INSERT INTO refund_order (
		out_trade_no,refund_fee,refund_reason,out_refund_no,refund_status)
		VALUES
		($1,$2,$3,$4,$5)`

	updateRefundSQL := `update refund_order set
	refund_status=$2,refund_result=$3,refund_trade_no=$4,updated_time=LOCALTIMESTAMP where out_trade_no=$1`

	_, err := model.DB.Exec(insertSQL, outTradeNo, refundFee, refundReason, outRefundNo, "PROCESSING")
	if err != nil {
		return map[string]interface{}{"code": "-1", "msg": err.Error()}
	}

	resData := request("POST", m)
	mjson, _ := json.Marshal(resData)
	refundResult := string(mjson)

	if resData["code"] == "error" {
		_, err2 := model.DB.Exec("update refund_order set refund_status=$2,refund_result=$3 where out_trade_no=$1", outTradeNo, "FAIL", refundResult)
		if err2 != nil {
			return map[string]interface{}{"code": "-1", "msg": err2.Error()}
		}
		return map[string]interface{}{"code": "-1", "msg": resData["msg"]}
	}
	refundTradeNo := resData["refund_id"]
	_, err2 := model.DB.Exec(updateRefundSQL, outTradeNo, "SUCCESS", refundResult, refundTradeNo)
	if err2 != nil {
		return map[string]interface{}{"code": "-1", "msg": err2.Error()}
	}
	return map[string]interface{}{"code": "200", "data": resData}
}

//QueryHcRefund 退款查询
/**
 * @param outRefundNo 系统退款请求交易号 必须
 * @param transactionNo 支付平台交易号 非必须
 * @param outTradeNo 系统交易号 非必须
 * @param refundTradeNo 支付平台退款流水号 非必须
 */
func QueryHcRefund(outRefundNo string, transactionNo string, outTradeNo string, refundTradeNo string) map[string]interface{} {

	if outRefundNo == "" {
		return map[string]interface{}{"code": "-1", "msg": "缺少参数"}
	}

	crow := model.DB.QueryRowx(`select ro.out_refund_no,po.merchant_id from refund_order ro
	left join pay_order po on po.out_trade_no = ro.out_trade_no
	where ro.out_refund_no=$1 limit 1`, outRefundNo)
	if crow == nil {
		return map[string]interface{}{"code": "-1", "msg": "退款查询失败"}
	}
	refundOrder := FormatSQLRowToMap(crow)
	_, cok := refundOrder["out_refund_no"]
	if !cok {
		return map[string]interface{}{"code": "-1", "msg": "退款订单不存在"}
	}

	noncestr := hcpay.GenerateNonceString(32)

	var m map[string]string
	m = make(map[string]string, 0)
	m["service_code"] = "hcpay.trade.refundquery"
	m["merchant_id"] = hcMerchantID[refundOrder["merchant_id"].(string)]
	m["partner_id"] = partnerID
	m["nonce_str"] = noncestr
	m["sign_type"] = "MD5"
	m["transaction_id"] = transactionNo
	m["out_trade_no"] = outTradeNo
	m["out_refund_no"] = outRefundNo
	m["refund_id"] = refundTradeNo
	sign := hcpay.GetSign(m, partnerKey)
	m["sign"] = sign

	resData := request("POST", m)

	updateRefundSQL := `update refund_order set
	refund_status=$2,refund_result=$3,refund_trade_no=$4,refund_success_time=$5,updated_time=LOCALTIMESTAMP where out_refund_no=$1`

	if resData["code"] == "error" {
		return map[string]interface{}{"code": "-1", "msg": resData["msg"]}
	}
	mjson, _ := json.Marshal(resData)
	refundResult := string(mjson)
	refundCount := resData["refund_count"]
	refundTradeNoP := resData["refund_id_0"]
	refundStatus := resData["refund_status_0"]
	refundSuccessTime := resData["refund_success_time_0"]
	if refundCount.(float64) == 0 {
		refundStatus = "FAIL"
	}
	_, err2 := model.DB.Exec(updateRefundSQL, outRefundNo, refundStatus, refundResult, refundTradeNoP, refundSuccessTime)
	if err2 != nil {
		return map[string]interface{}{"code": "-1", "msg": err2.Error()}
	}
	return map[string]interface{}{"code": "200", "data": resData}
}

//CreateHcOrder 创建订单
/**
 * @param outTradeNo 系统交易号 必须
 * @param merchantID 支付类型 wx 微信 ali 支付宝 必须
 * @param payMode 支付模式 weixin_f2f –微信当面付 alipay_f2f – 支付宝当面付 mybank_weixin_f2f 网商微信当面付 mybank_alipay_f2f 网商支付宝当面付 必须
 * @param totalFee 费用 必须
 * @param body 商品描述 必须
 * @param orderType 订单类型 normal 统一下单; public 公众号下单 必须
 * @param businessType 业务类型 必须
 * @param requestIP 请求地址 必须
 * @param goodsDetail 商品详情 非必须
 * @param deviceInfo 设备信息 非必须
 * @param subjectType 主体类型 H5支付必传(android、ios、h5)
 * @param subjectID 主体ID(对应应用id或网站首页) H5支付必传,H5支付传http://www.hizpay.com; android 传入ap对应的 id;IOS传入对应的id
 * @param subjectName 主体名称 H5支付必传网站名称或app名称
 * @param openID 用户标识 所有微信wap必填，支付宝传入buyer_id
 * @param logonAccount 登陆账号 支付宝渠道时输入,openid 和 logon_account不能同时为空
 */
func CreateHcOrder(outTradeNo string, merchantID string, payMode string, totalFee string, body string, orderType string, businessType string, requestIP string, goodsDetail string, deviceInfo string, subjectType string, subjectID string, subjectName string, openID string, logonAccount string) map[string]interface{} {

	if outTradeNo == "" || payMode == "" || totalFee == "" || body == "" || merchantID == "" || orderType == "" || (openID == "" && logonAccount == "") {
		return map[string]interface{}{"code": "-1", "msg": "缺少参数"}
	}
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
		return map[string]interface{}{"code": "-1", "msg": resData["msg"]}
	}
	return map[string]interface{}{"code": "200", "data": resData}
}

//QueryHcOrder 订单查询
/**
 * @param outTradeNo 系统交易号 非必须
 * @param transactionNo 支付平台交易号 非必须
 */
func QueryHcOrder(outTradeNo string, transactionNo string) map[string]interface{} {

	if outTradeNo == "" {
		return map[string]interface{}{"code": "-1", "msg": "缺少参数"}
	}

	selectSQL := `select out_trade_no,order_status,merchant_id,created_time from pay_order where out_trade_no=$1`
	row := model.DB.QueryRowx(selectSQL, outTradeNo)
	if row == nil {
		return map[string]interface{}{"code": "-1", "msg": "支付订单查询失败"}
	}

	payOrder := FormatSQLRowToMap(row)
	_, ok := payOrder["out_trade_no"]
	if !ok {
		return map[string]interface{}{"code": "-1", "msg": "支付订单不存在"}
	}

	noncestr := hcpay.GenerateNonceString(32)

	var m map[string]string
	m = make(map[string]string, 0)
	m["service_code"] = "hcpay.trade.orderquery"
	m["merchant_id"] = hcMerchantID[payOrder["merchant_id"].(string)]
	m["partner_id"] = partnerID
	m["nonce_str"] = noncestr
	m["sign_type"] = "MD5"
	m["transaction_id"] = transactionNo
	m["out_trade_no"] = outTradeNo
	sign := hcpay.GetSign(m, partnerKey)
	m["sign"] = sign

	updateSQL := `update pay_order set
		trade_no=$2,order_status=$3,openid=$4,pay_time=$5,updated_time=LOCALTIMESTAMP where out_trade_no=$1`

	createdTime := payOrder["created_time"].(time.Time)
	orderStatus := payOrder["order_status"].(string)

	//超时未支付的订单撤销
	if createdTime.Before(time.Now().Add(-30*time.Second)) && (orderStatus == "USERPAYING" || orderStatus == "NOTPAY") {
		fmt.Println("***********撤销************")
		FaceToFaceCancel(outTradeNo)
	}

	resData := request("POST", m)

	if resData["code"] == "error" {
		return map[string]interface{}{"code": "-1", "msg": resData["msg"]}
	}
	tradeNo := resData["transaction_id"]
	openID := resData["openid"]
	payTime := resData["payed_time"]
	orderStatusHc := resData["trade_status"]

	_, err := model.DB.Exec(updateSQL, outTradeNo, tradeNo, orderStatusHc, openID, payTime)
	if err != nil {
		fmt.Println("err ===", err)
		return map[string]interface{}{"code": "-1", "msg": err.Error()}
	}
	return map[string]interface{}{"code": "200", "data": resData}
}

//QueryOrder 查询平台订单状态 定时任务用
/**
 * @param outTradeNo 系统交易号 非必须
 */
func QueryOrder(outTradeNo string, merchantID string) map[string]interface{} {

	if outTradeNo == "" || merchantID == "" {
		return map[string]interface{}{"code": "-1", "msg": "缺少参数"}
	}
	noncestr := hcpay.GenerateNonceString(32)

	var m map[string]string
	m = make(map[string]string, 0)
	m["service_code"] = "hcpay.trade.orderquery"
	m["merchant_id"] = hcMerchantID[merchantID]
	m["partner_id"] = partnerID
	m["nonce_str"] = noncestr
	m["sign_type"] = "MD5"
	m["out_trade_no"] = outTradeNo
	sign := hcpay.GetSign(m, partnerKey)
	m["sign"] = sign

	updateSQL := `update pay_order set
		trade_no=$2,order_status=$3,openid=$4,pay_time=$5,updated_time=LOCALTIMESTAMP where out_trade_no=$1`

	resData := request("POST", m)

	if resData["code"] == "error" {
		return map[string]interface{}{"code": "-1", "msg": resData["msg"]}
	}
	tradeNo := resData["transaction_id"]
	openID := resData["openid"]
	payTime := resData["payed_time"]
	orderStatusHc := resData["trade_status"]

	_, err := model.DB.Exec(updateSQL, outTradeNo, tradeNo, orderStatusHc, openID, payTime)
	if err != nil {
		fmt.Println("err ===", err)
		return map[string]interface{}{"code": "-1", "msg": err.Error()}
	}
	return map[string]interface{}{"code": "200", "data": resData}
}

//HcOrderClose 订单关闭
/**
 * @param outTradeNo 系统交易号 非必须
 * @param merchantID 支付类型 wx 微信 ali 支付宝 必须
 */
func HcOrderClose(outTradeNo string) map[string]interface{} {

	if outTradeNo == "" {
		return map[string]interface{}{"code": "-1", "msg": "缺少参数"}
	}

	selectSQL := `select out_trade_no,merchant_id from pay_order where out_trade_no=$1`

	row := model.DB.QueryRowx(selectSQL, outTradeNo)
	if row == nil {
		return map[string]interface{}{"code": "-1", "msg": "支付订单查询失败"}
	}

	payOrder := FormatSQLRowToMap(row)
	_, ok := payOrder["out_trade_no"]
	if !ok {
		return map[string]interface{}{"code": "-1", "msg": "支付订单不存在"}
	}

	noncestr := hcpay.GenerateNonceString(32)

	var m map[string]string
	m = make(map[string]string, 0)
	m["service_code"] = "hcpay.trade.close"
	m["merchant_id"] = hcMerchantID[payOrder["merchant_id"].(string)]
	m["partner_id"] = partnerID
	m["nonce_str"] = noncestr
	m["sign_type"] = "MD5"
	m["out_trade_no"] = outTradeNo
	sign := hcpay.GetSign(m, partnerKey)
	m["sign"] = sign

	updateSQL := `update pay_order set order_status=$2,updated_time=LOCALTIMESTAMP where out_trade_no=$1`

	resData := request("POST", m)

	if resData["code"] == "error" {
		return map[string]interface{}{"code": "-1", "msg": resData["msg"]}
	}
	_, err := model.DB.Exec(updateSQL, outTradeNo, "CLOSE")
	if err != nil {
		fmt.Println("err ===", err)
		return map[string]interface{}{"code": "-1", "msg": err.Error()}
	}
	return map[string]interface{}{"code": "200", "data": resData}
}

//DownloadBill 下载对账单
/**
 * @param merchantID 支付类型 wx 微信 ali 支付宝 必须
 * @param payChannel 账单渠道 weixin 微信; alipay 支付宝; swiftpass_weixin CIB微信; swiftpass_alipay CIB支付宝; ylznew_alipay 易联众支付宝; ylznew_weixin 易联众微信; mybank_weixin 网商微信渠道; mybank_alipay 网商支付宝渠道 必须
 * @param billDate 账单日期 yyyy-MM-dd 必须
 * @param numPerPage 每页的记录数 10 必须
 * @param pageNo 页码 从1开始 必须
 */
func DownloadBill(merchantID string, payChannel string, billDate string, numPerPage string, pageNo string) map[string]interface{} {

	if merchantID == "" || payChannel == "" || billDate == "" || numPerPage == "" || pageNo == "" {
		return map[string]interface{}{"code": "-1", "msg": "缺少参数"}
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
		return map[string]interface{}{"code": "-1", "msg": resData["msg"]}
	}
	return map[string]interface{}{"code": "200", "data": resData}
}

//request 请求方法
func request(method string, m map[string]string) map[string]interface{} {
	post, er := json.Marshal(m)
	if er != nil {
		fmt.Println("er=========", er)
		return nil
	}

	var jsonStr = []byte(post)
	fmt.Println("请求参数**************", bytes.NewBuffer(jsonStr))

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
	fmt.Println("支付平台返回结果*****************", results)

	if results["return_code"] == "0" {
		if results["result_code"] == "0" {
			return results
		}
		if results["result_code"] == "2" {
			return map[string]interface{}{"code": "error", "msg": results["result_msg"], "data": results["result_code"]}
		}
		return map[string]interface{}{"code": "error", "msg": results["result_msg"]}
	}
	return map[string]interface{}{"code": "error", "msg": results["return_msg"]}
}
