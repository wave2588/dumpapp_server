package main

import (
	"dumpapp_server/cmd/cron/conclusion"
	"dumpapp_server/cmd/cron/delete_ipa"
	"github.com/robfig/cron/v3"
)

func main() {

	c := cron.New(cron.WithSeconds())

	/// 每晚 0 点删除临时 ipa
	c.AddFunc("00 00 00 * * ?", func() {
		delete_ipa.Run()
	})

	/// 每晚 0 点推送总结
	c.AddFunc("00 00 00 * * ?", func() {
		conclusion.Run()
	})

	c.Start()
	select {}
}
