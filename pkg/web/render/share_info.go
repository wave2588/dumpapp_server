package render

type ShareInfo struct {
	QQGroupURL     string `json:"qq_group_url"`
	QQGroupNum     int64  `json:"qq_group_num"`
	TelegramURL    string `json:"telegram_url"`
	WechatURL      string `json:"wechat_url"`
	Content        string `json:"content"`
	AppTutorialURL string `json:"app_tutorial_url"` /// app 使用教程
}

func MustRenderShareInfo() *ShareInfo {
	return &ShareInfo{
		Content:        "这是一段分享文案",
		QQGroupURL:     "https://jq.qq.com/?_wv=1027&k=ItOGYFPG",
		QQGroupNum:     763789550,
		TelegramURL:    "https://t.me/+VGGU8RYVqDo1NTg1",
		WechatURL:      "https://work.weixin.qq.com/u/vc3a10ae3518beb870?v=3.1.23.79300&src=wx",
		AppTutorialURL: "https://g89s5y6zts.feishu.cn/docx/doxcnUUe4ti2rYSNz535iB9NFsh",
	}
}
