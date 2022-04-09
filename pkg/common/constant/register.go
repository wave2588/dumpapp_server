package constant

import (
	"strings"

	"dumpapp_server/pkg/util"
)

var SupportedEmailAddress = []string{
	"qq.com",
	"vip.qq.com",
	"163.com",
	"126.com",
	"gmail.com",
	"icloud.com",
	"foxmail.com",
	"sina.com",
	"me.com",
	"aliyun.com",
}

func CheckEmailValid(email string) bool {
	/// 为了防止临时邮箱, 类似这种  jocktmp+r7g599@gmail.com
	if strings.Contains(email, "+") {
		return false
	}

	address := strings.Split(email, "@")
	if len(address) < 2 {
		return false
	}
	suffix := address[1]
	return util.IsContainStrings(SupportedEmailAddress, suffix)
}
