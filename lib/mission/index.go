package mission

import (
	"log"

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
		i++
		log.Println("start", i)
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
