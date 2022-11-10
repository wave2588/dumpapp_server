package constant

import (
	"fmt"

	"dumpapp_server/pkg/common/util"
	errors2 "dumpapp_server/pkg/errors"
	"github.com/spf13/cast"
)

func GetInstallAppCDKeySuffix(cerLevel int) string {
	switch cast.ToInt64(cerLevel) {
	case CertificateIDL1:
		return "L1"
	case CertificateIDL2:
		return "L2"
	case CertificateIDL3:
		return "L3"
	}
	return ""
}

func GetInstallAppCerPrice(cerLevel int64) int64 {
	switch cast.ToInt64(cerLevel) {
	case CertificateIDL1:
		return CertificatePriceL1
	case CertificateIDL2:
		return CertificatePriceL2
	case CertificateIDL3:
		return CertificatePriceL3
	}
	/// 如果 cerLevel 不对则直接报错
	util.PanicIf(errors2.UnproccessableError(fmt.Sprintf("未识别的 cdkeyPriceID: %d", cerLevel)))
	return CertificatePriceL1
}
