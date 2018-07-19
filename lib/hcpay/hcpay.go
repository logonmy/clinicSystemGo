package hcpay

import (
	"bytes"
	"clinicSystemGo/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var partnerID = "GZHXDMD"
var partnerKey = "C1F013D2B90E45E9995E16A3411A6910"
var hcMerchantID = map[string]string{"wx": "1260939901", "ali": "2017081008129270"}
var hcOrdertype = map[string]string{"normal": "hcpay.trade.unifiedorder", "public": "hcpay.trade.registerorder"}

// var url = "https://pay.med.gzhc365.com/api/hcpay/gateway"
var url = "https://upay.med.gzhc365.com/api/hcpay/gateway"

//FaceToFace 当面付
/**
 * @param payMode 支付模式 weixin_f2f –微信当面付 alipay_f2f – 支付宝当面付 mybank_weixin_f2f 网商微信当面付 mybank_alipay_f2f 网商支付宝当面付
 * @param businessType 业务类型
 * @param totalFee 费用
 * @param body 商品描述
 * @param goodsDetail 商品详情
 * @param authCode 付款码
 * @param merchantID 支付类型 wx 微信 ali 支付宝
 * @param requestIP 请求地址
 */
func FaceToFace(payMode string, businessType string, totalFee string, body string, goodsDetail string, deviceInfo string, authCode string, merchantID string, outTradeNo string, requestIP string) map[string]interface{} {
	if outTradeNo == "" || businessType == "" || payMode == "" || totalFee == "" || body == "" || merchantID == "" || authCode == "" || requestIP == "" {
		return map[string]interface{}{"code": "-1", "msg": "缺少参数"}
	}

	noncestr := GenerateNonceString(32)

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
	sign := GetSign(m, partnerKey)
	m["sign"] = sign

	mjson, _ := json.Marshal(m)
	originalData := string(mjson)

	insertSQL := `INSERT INTO pay_order (
		out_trade_no,total_fee,body,order_status,original_data,trade_type,business_type) 
		VALUES 
		($1,$2,$3,$4,$5,$6,$7)`

	updateSQL := `update pay_order set 
	trade_no=$2,order_status=$3,openid=$4,pay_time=$5 where out_trade_no=$1`

	_, err := model.DB.Exec(insertSQL, outTradeNo, totalFee, body, "NOTPAY", originalData, payMode, businessType)
	if err != nil {
		fmt.Println("err ===", err)
		return map[string]interface{}{"code": "-1", "msg": err.Error()}
	}

	resData := request("POST", m)

	if resData["code"] == "error" {
		if resData["data"] != nil && resData["data"].(string) == "USERPAYING" {
			_, err := model.DB.Exec("update pay_order set order_status=$2 where out_trade_no=$1", outTradeNo, "USERPAYING")
			if err != nil {
				fmt.Println("err ===", err)
				return map[string]interface{}{"code": "-1", "msg": err.Error()}
			}
			return map[string]interface{}{"code": "2", "msg": resData["msg"]}
		}
		return map[string]interface{}{"code": "-1", "msg": resData["msg"]}
	} else {
		tradeNo := resData["transaction_id"]
		openID := resData["openid"]
		payTime := resData["payed_time"]

		_, err := model.DB.Exec(updateSQL, outTradeNo, tradeNo, "SUCCESS", openID, payTime)
		if err != nil {
			fmt.Println("err ===", err)
			return map[string]interface{}{"code": "-1", "msg": err.Error()}
		}

		return map[string]interface{}{"code": "200", "data": resData}
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
		if results["result_code"] == "2" {
			return map[string]interface{}{"code": "error", "msg": results["result_msg"], "data": "USERPAYING"}
		}
		return map[string]interface{}{"code": "error", "msg": results["result_msg"]}
	}
	return map[string]interface{}{"code": "error", "msg": results["return_msg"]}
}
