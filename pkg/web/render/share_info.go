package render

import (
	"context"
	"fmt"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
)

type ShareInfo struct {
	QQGroupURL     string `json:"qq_group_url"`
	QQGroupNum     int64  `json:"qq_group_num"`
	TelegramURL    string `json:"telegram_url"`
	WechatURL      string `json:"wechat_url"`
	Content        string `json:"content"`
	AppTutorialURL string `json:"app_tutorial_url"` /// app 使用教程
}

func MustRenderShareInfo(ctx context.Context, loginID int64) *ShareInfo {
	inviteCodeMap, err := impl.DefaultMemberInviteCodeDAO.BatchGetByMemberID(ctx, []int64{loginID})
	util.PanicIf(err)

	inviteURL := "https://www.dumpapp.com/app"
	if inviteCode, ok := inviteCodeMap[loginID]; ok {
		inviteURL = fmt.Sprintf(constant.InviteURL, inviteCode.Code)
	}

	return &ShareInfo{
		Content:        fmt.Sprintf("App 多开、分身、去广告、签名推广.....只有你想不到，没有我们做不到，更多好玩 App 功能尽在 DumpApp，快来加入吧！%s", inviteURL),
		QQGroupURL:     "https://jq.qq.com/?_wv=1027&k=xVqlWqEc",
		QQGroupNum:     763789550,
		TelegramURL:    "https://t.me/+VGGU8RYVqDo1NTg1",
		WechatURL:      "https://work.weixin.qq.com/u/vc3a10ae3518beb870?v=3.1.23.79300&src=wx",
		AppTutorialURL: "https://g89s5y6zts.feishu.cn/docx/doxcnUUe4ti2rYSNz535iB9NFsh",
	}
}
