package main

import (
	"dumpapp_server/cmd/cron/delete_ipa"
	"github.com/roylee0704/gron"
)

var c = gron.New()

func main() {

	//var (
	//	daily = gron.Every(1 * xtime.Day)
	//)

	///// 定时更新 ipa
	//r := &update_ipa.UpdateIpa{}
	//c.Add(daily.At("00:00"), r)
	//
	//c.Start()
	//

	delete_ipa.Run()
}
