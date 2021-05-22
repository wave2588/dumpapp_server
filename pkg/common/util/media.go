package util

import (
	"fmt"
	"strings"

	"dumpapp_server/pkg/config"
)

func GetImageUrl(token string) string {
	if strings.HasPrefix(token, "http") {
		return token
	}
	return fmt.Sprintf("%s/%s", config.DumpConfig.AppConfig.TencentCosIpaHost, token)
}
