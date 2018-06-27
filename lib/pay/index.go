package pay

import (
	"github.com/kataras/iris"
)

// 创建订单
func Create(ctx iris.Context) {
	payMode := ctx.PostValue("pay_mode")
	subjectType := ctx.PostValue("subject_type")
	subjectId := ctx.PostValue("subject_id")
	subjectName := ctx.PostValue("subject_name")
	openId := ctx.PostValue("openid")
	logonAccount := ctx.PostValue("logon_account")
	totalFee := ctx.PostValue("total_fee")
	body := ctx.PostValue("body")
	goodsDetail := ctx.PostValue("goods_detail")
	deviceInfo := ctx.PostValue("device_info")
}
