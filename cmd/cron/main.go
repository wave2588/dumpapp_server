package main

import (
	"dumpapp_server/cmd/cron/conclusion"
	"dumpapp_server/cmd/cron/delete_ipa"
	"github.com/robfig/cron/v3"
)

func main() {

	c := cron.New(cron.WithSeconds())

	c.AddFunc("00 00 20 * * ?", func() {
		conclusion.Run()
	})

	c.AddFunc("00 00 00 * * ?", func() {
		delete_ipa.Run()
	})

	c.Start()
	select {}
}
