package main

import (
	"dumpapp_server/cmd/cron/conclusion"
	"dumpapp_server/cmd/cron/delete_ipa"
	"dumpapp_server/cmd/cron/delete_plist"
	"github.com/robfig/cron/v3"
)

func main() {

	c := cron.New(cron.WithSeconds())

	/// 每晚 0 点删除临时 ipa
	c.AddFunc("00 00 00 * * ?", func() {
		delete_ipa.Run()
	})

	/// 每晚 0 点删除 plist
	c.AddFunc("00 00 00 * * ?", func() {
		delete_plist.Run()
	})

	/// 每晚 0 点推送总结
	c.AddFunc("00 00 00 * * ?", func() {
		conclusion.Run()
	})

	//spec := "*/2 * * * * ?" //cron表达式，每秒一次
	//c.AddFunc(spec, func() {
	//	sign_ipa.Run()
	//})

	c.Start()
	select {}
}
