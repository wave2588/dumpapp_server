package main

import (
	"dumpapp_server/cmd/cron/conclusion"
	"github.com/robfig/cron/v3"
)

func main() {

	c := cron.New(cron.WithSeconds())

	spec := "00 00 20 * * ?" /// 每秒一次
	c.AddFunc(spec, func() {
		conclusion.Run()
	})

	//delete_ipa.Run()
	c.Start()
	select {}
}
