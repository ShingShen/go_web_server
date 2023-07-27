package cronjobs

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/robfig/cron/v3"

	systemService "server/services/dbservice/service/system"
	sqlOperator "server/utils/sqloperator"
)

func OneMin(db sqlOperator.ISqlDB, rdb *redis.Client) {
	go func() {
		c := cron.New()
		i := 1
		c.AddFunc("*/1 * * * *", func() {
			fmt.Println("Running at every 1 min", i)
			i++
		})
		c.Start()
		time.Sleep(time.Minute * 5)
	}()
}

func SevenDays(db sqlOperator.ISqlDB, rdb *redis.Client) {
	go func() {
		c := cron.New()
		c.AddFunc("*/0 0 * * 0", func() {
			systemServiceFactory := &systemService.SystemServiceFactory{}
			getService, _ := systemServiceFactory.GetSystemService(db, "system")
			service := getService.RefreshLoginTokens(rdb)
			fmt.Print(service)
		})
		c.Start()
	}()
}
