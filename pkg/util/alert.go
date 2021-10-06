package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/util"
)

func SendWeiXinBotV2(ctx context.Context, keyID string, message string, receivers []string) {
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": "DumpApp - 应用名称：<font color=\"info\">微信</font>\n>" +
				"应用名称:<font color=\"comment\">微信</font>\n" +
				"应用版本:<font color=\"comment\">微信</font>\n" +
				"bundleID:<font color=\"comment\">11111</font>\n" +
				"邮箱:<font color=\"comment\">zhanghaibo</font>\n" +
				"手机号:<font color=\"comment\">15711367321</font>",
			"mentioned_list": receivers,
		},
	}
	SendWeiXinBot(ctx, keyID, data, receivers)
}

func SendWeiXinBot(ctx context.Context, keyID string, data interface{}, receivers []string) {
	jsonStr, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", keyID), bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second}

	resp, err := client.Do(req)
	util.PanicIf(err)
	defer resp.Body.Close() // nolint
}
