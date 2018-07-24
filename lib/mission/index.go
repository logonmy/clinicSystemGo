package mission

import (
	"clinicSystemGo/controller"
	"clinicSystemGo/model"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron"
)

// Seconds      | Yes        | 0-59            | * / , -
// Minutes      | Yes        | 0-59            | * / , -
// Hours        | Yes        | 0-23            | * / , -
// Day of month | Yes        | 1-31            | * / , - ?
// Month        | Yes        | 1-12 or JAN-DEC | * / , -
// Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

// StartMission 开始定时任务
func StartMission() {
	c := cron.New()
	i := 0
	//每分钟
	c.AddFunc("@every 1m", func() {
		fmt.Println("开始定时任务")
		selectSQL := `select out_trade_no,order_status,merchant_id,created_time from pay_order where order_status in ('NOTPAY','USERPAYING','UNKNOW')`
		rows, _ := model.DB.Queryx(selectSQL)
		if rows == nil {
			fmt.Println("查询订单数据库错误")
		}

		payOrders := controller.FormatSQLRowsToMapArray(rows)

		for _, v := range payOrders {
			createdTime := v["created_time"].(time.Time)
			outTradeNo := v["out_trade_no"].(string)
			merchantID := v["merchant_id"].(string)
			resData := controller.QueryOrder(outTradeNo, merchantID)

			if resData["code"].(string) == "200" {
				results := resData["data"].(map[string]interface{})

				fmt.Println("************订单信息存在**********", outTradeNo)
				if results["trade_status"] != "REFUND" && results["trade_status"] != "CLOSE" && createdTime.Before(time.Now().Add(-30*time.Second)) {
					fmt.Println("***********定时撤销************")
					controller.FaceToFaceCancel(outTradeNo)
				}
			} else if resData["msg"].(string) == "未找到支付订单信息" {
				fmt.Println("************订单信息不存在**********", outTradeNo)
				updateSQL := `update pay_order set order_status=$2,updated_time=LOCALTIMESTAMP where out_trade_no=$1`
				_, err := model.DB.Exec(updateSQL, outTradeNo, "CLOSE")
				if err != nil {
					fmt.Println("定时撤销订单err===", err)
				}
			}
		}
	})

	//每天2点整
	c.AddFunc("0 0 2 * * *", func() {
		i++
		log.Println("start", i)
	})

	//每个月1号凌晨3点
	c.AddFunc("0 0 3 1 * *", func() {
		i++
		log.Println("start", i)
	})

	c.Start()
}
