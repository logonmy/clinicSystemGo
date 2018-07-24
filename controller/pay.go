package controller

import (
	"clinicSystemGo/lib/hcpay"

	"github.com/kataras/iris"
)

//CreateHcOrderWeb 创建订单
func CreateHcOrderWeb(ctx iris.Context) {
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
	outTradeNo := ctx.PostValue("out_trade_no")
	businessType := ctx.PostValue("business_type")
	outTradeNo = hcpay.CreateTradeNo(20)
	// outTradeNo = "TRZ" + hcpay.CreateTradeNo(20)

	if outTradeNo == "" || payMode == "" || totalFee == "" || body == "" || merchantID == "" || orderType == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	requestIP := "47.93.206.157"

	resData := CreateHcOrder(outTradeNo, merchantID, payMode, totalFee, body, orderType, businessType, requestIP, goodsDetail, deviceInfo, subjectType, subjectID, subjectName, openID, logonAccount)

	if resData["code"] != "200" {
		ctx.JSON(iris.Map{"code": "-1", "msg": resData["msg"]})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": resData})
}

//QueryHcOrderWeb 订单查询
func QueryHcOrderWeb(ctx iris.Context) {
	outTradeNo := ctx.PostValue("out_trade_no")
	transactionNo := ctx.PostValue("trade_no")

	if outTradeNo == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	resData := QueryHcOrder(outTradeNo, transactionNo)

	if resData["code"] != "200" {
		ctx.JSON(iris.Map{"code": "-1", "msg": resData["msg"]})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": resData})
}

//QueryHcOrderTest 订单查询测试
func QueryHcOrderTest(ctx iris.Context) {
	outTradeNo := ctx.PostValue("out_trade_no")
	merchantID := ctx.PostValue("merchant_id")

	if outTradeNo == "" || merchantID == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	resData := QueryOrder(outTradeNo, merchantID)

	ctx.JSON(iris.Map{"data": resData})
}

//HcRefundWeb 退款
func HcRefundWeb(ctx iris.Context) {
	outTradeNo := ctx.PostValue("out_trade_no")
	transactionNo := ctx.PostValue("trade_no")
	refundFeeStr := ctx.PostValue("refund_fee")
	refundReason := ctx.PostValue("refund_reason")
	outRefundNo := ctx.PostValue("out_refund_no")

	outRefundNo = hcpay.CreateTradeNo(20)
	// outRefundNo = "TRZ"+hcpay.CreateTradeNo(20)

	if outRefundNo == "" || outTradeNo == "" || refundFeeStr == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	resData := HcRefund(outTradeNo, refundFeeStr, outRefundNo, transactionNo, refundReason)

	if resData["code"] != "200" {
		ctx.JSON(iris.Map{"code": "-1", "msg": resData["msg"]})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": resData})
}

//QueryHcRefundWeb 退款查询
func QueryHcRefundWeb(ctx iris.Context) {
	outTradeNo := ctx.PostValue("out_trade_no")
	transactionNo := ctx.PostValue("trade_no")
	outRefundNo := ctx.PostValue("out_refund_no")
	refundTradeNo := ctx.PostValue("refund_trade_no")

	if outRefundNo == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	resData := QueryHcRefund(outRefundNo, transactionNo, outTradeNo, refundTradeNo)

	if resData["code"] != "200" {
		ctx.JSON(iris.Map{"code": "-1", "msg": resData["msg"]})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": resData})
}

//HcOrderCloseWeb 订单关闭
func HcOrderCloseWeb(ctx iris.Context) {
	outTradeNo := ctx.PostValue("out_trade_no")
	merchantID := ctx.PostValue("merchant_id")

	if merchantID == "" || outTradeNo == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	resData := HcOrderClose(outTradeNo)

	if resData["code"] != "200" {
		ctx.JSON(iris.Map{"code": "-1", "msg": resData["msg"]})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": resData})
}

//FaceToFaceWeb 当面付
func FaceToFaceWeb(ctx iris.Context) {
	payMode := ctx.PostValue("pay_mode")
	businessType := ctx.PostValue("business_type")
	totalFee := ctx.PostValue("total_fee")
	body := ctx.PostValue("body")
	goodsDetail := ctx.PostValue("goods_detail")
	deviceInfo := ctx.PostValue("device_info")
	authCode := ctx.PostValue("auth_code")
	merchantID := ctx.PostValue("merchant_id")
	outTradeNo := ctx.PostValue("out_trade_no")

	outTradeNo = hcpay.CreateTradeNo(20)
	// outTradeNo = "TRZ" + hcpay.CreateTradeNo(20)

	if outTradeNo == "" || businessType == "" || payMode == "" || totalFee == "" || body == "" || merchantID == "" || authCode == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}
	requestIP := "47.93.206.157"
	resData := FaceToFace(outTradeNo, authCode, merchantID, payMode, businessType, totalFee, body, requestIP, goodsDetail, deviceInfo)

	if resData["code"] != "200" {
		ctx.JSON(iris.Map{"code": "-1", "msg": resData["msg"]})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": resData})
}

//FaceToFaceCancelWeb 当面付支付撤销
func FaceToFaceCancelWeb(ctx iris.Context) {
	outTradeNo := ctx.PostValue("out_trade_no")

	if outTradeNo == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	resData := FaceToFaceCancel(outTradeNo)

	if resData["code"] != "200" {
		ctx.JSON(iris.Map{"code": "-1", "msg": resData["msg"]})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": resData})
}

//DownloadBillWeb 下载对账单
func DownloadBillWeb(ctx iris.Context) {
	merchantID := ctx.PostValue("merchant_id")
	payChannel := ctx.PostValue("pay_channel")
	billDate := ctx.PostValue("bill_date")
	numPerPage := ctx.PostValue("num_per_page")
	pageNo := ctx.PostValue("page_no")

	if merchantID == "" || payChannel == "" || billDate == "" || numPerPage == "" || pageNo == "" {
		ctx.JSON(iris.Map{"code": "-1", "msg": "缺少参数"})
		return
	}

	resData := DownloadBill(merchantID, payChannel, billDate, numPerPage, pageNo)

	if resData["code"] != "200" {
		ctx.JSON(iris.Map{"code": "-1", "msg": resData["msg"]})
		return
	}
	ctx.JSON(iris.Map{"code": "200", "data": resData})
}
