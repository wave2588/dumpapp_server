package main

import (
	"dumpapp_server/cmd/cron/update_ipa"
	"github.com/roylee0704/gron"
	"github.com/roylee0704/gron/xtime"
)

var c = gron.New()

func main() {

	var (
		daily = gron.Every(1 * xtime.Day)
	)

	/// 定时更新 ipa
	r := &update_ipa.UpdateIpa{}
	c.Add(daily.At("00:00"), r)

	c.Start()

	for {

	}
}
