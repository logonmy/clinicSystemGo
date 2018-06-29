package hcpay

// new_str {
// 	"auth_code":"134666868728085631",
// 	"body":"测试商品名称",
// 	"create_ip":"47.93.206.157",
// 	"device_info":"支付发起设备的信息",
// 	"goods_detail":"商品详细描述",
// 	"merchant_id":"1260939901",
// 	"nonce_str":"8POfi6pQ6vZo8RUjQrmqIz3hxA7C0oWY",
// 	"out_trade_no":"18180217114381719470",
// 	"partner_id":"GZHXDMD",
// 	"pay_mode":"weixin_f2f",
// 	"service_code":"hcpay.trade.f2f",
// 	"sign":"2923E430FF6E823618B200E3C74CF1DE",
// 	"sign_type":"MD5",
// 	"total_fee":"1"}
// results========= map[
// 	real_total_fee:1
// 	transaction_id:4200000111201806299527506865
// 	openid:o7Pb4sp2FEOlGTq2sOspZoj1AUWw
// 	pay_mode:weixin_f2f
// 	out_trade_no:18180217114381719470
// 	payed_time:20180629212300
// 	result_code:0
// 	sign_type:MD5
// 	sign:CD6BD6BE9CA82A74A396443D8C3D0228
// 	return_msg:OK result_msg:OK return_code:0
// 	nonce_str:ocuy9nabh8kycl2glqwy
// 	pay_order_id:1.8062960108e+18]
// [INFO] 2018/06/29 21:23 200 1.851805533s ::1 POST /hcPay/FaceToFace

// new_str {"merchant_id":"1260939901","nonce_str":"BJZRJawDYbX0NA0EbzseMKCtVO2vFxs4","out_refund_no":"18180981873339659580","out_trade_no":"18180217114381719470","partner_id":"GZHXDMD","refund_fee":"1","service_code":"hcpay.trade.refund","sign":"9BB12B4CB39C76EBD00C4EA9BDD1C1A6","sign_type":"MD5","transaction_id":""}
// results========= map[
// 	refund_fee:1
// 	return_code:0
// 	nonce_str:bkmowvvtcfnemlso7g1r
// 	return_msg:OK
// 	merchant_id:1260939901
// 	out_trade_no:18180217114381719470
// 	result_code:0
// 	sign_type:MD5
// 	out_refund_no:18180981873339659580
// 	sign:D2DFF06C973571437CC06F388E619FFD
// 	refund_id:50000107342018062905256536511]
// [INFO] 2018/06/29 21:28 200 1.772671672s ::1 POST /hcPay/HcRefund

// new_str {"auth_code":"134734637247497075","body":"测试商品名称","create_ip":"47.93.206.157","device_info":"支付发起设备的信息","goods_detail":"商品详细描述","merchant_id":"1260939901","nonce_str":"1dJKz7oN7SBPzb1lGkCrFw6hCIjZF1I1","out_trade_no":"18180570761537812751","partner_id":"GZHXDMD","pay_mode":"weixin_f2f","service_code":"hcpay.trade.f2f","sign":"7D80686EA6AB6D440BAFEDF32070A94C","sign_type":"MD5","total_fee":"1"}
// results========= map[
// 	nonce_str:oll94w2og1ihvqgskqc6
// 	out_trade_no:18180570761537812751
// 	sign:76E68C2ED6E3A91D854BE2C3BBAF8CB3
// 	return_msg:OK
// 	result_code:2
// 	return_code:0
// 	sign_type:MD5
// 	result_msg:OK|USERPAYING-需要用户输入支付密码]
// [INFO] 2018/06/29 21:29 200 770.187847ms ::1 POST /hcPay/FaceToFace

// new_str {"merchant_id":"1260939901","nonce_str":"lqXlPbhSYp41XELwOY2S76Iqd9risR6h","out_trade_no":"18180570761537812751","partner_id":"GZHXDMD","service_code":"hcpay.trade.orderquery","sign":"A657FFD64772221F3B229E7EC929B74A","sign_type":"MD5","transaction_id":""}
// results========= map[
// 	return_msg:OK total_fee:1 sign_type:MD5 sign:2D7CF6AB812FC3C0190077522668B666 biz_channel:99 return_code:0 trade_status:SUCCESS result_code:0 nonce_str:gk4wfw9lgiwq767m9d1h biz_type:default merchant_id:1260939901 payed_time:20180629212941 pay_order_id:1.8062960108e+18 discount_fee:0 transaction_id:4200000134201806297915461850 openid:o7Pb4sv7472Gk_nItMotnNa40WIU pay_mode:weixin_f2f out_trade_no:18180570761537812751 real_total_fee:1]
// [INFO] 2018/06/29 21:31 200 3.672458884s ::1 POST /hcPay/QueryHcOrder
// new_str {"merchant_id":"1260939901","nonce_str":"M9ItCcpBlfj9ohlSL5Upv2k95HFAR3Pl","out_refund_no":"18180514215609145861","out_trade_no":"18180570761537812751","partner_id":"GZHXDMD","refund_fee":"1","service_code":"hcpay.trade.refund","sign":"BFD87ED62768E3AE07D9A676F51FFEDB","sign_type":"MD5","transaction_id":""}
// results========= map[refund_fee:1 result_code:0 return_code:0 nonce_str:odccd3w3sbb781yeom0h return_msg:OK merchant_id:1260939901 out_trade_no:18180570761537812751 out_refund_no:18180514215609145861 sign:CDA64895E38CE63766AD0A295527C699 refund_id:50000207132018062905223708470 sign_type:MD5]
