package constant

import (
	"fmt"

	"dumpapp_server/pkg/common/util"
	errors2 "dumpapp_server/pkg/errors"
)

const (
	InstallAppCDKeyPriceIDL1 = 1
	InstallAppCDKeyPriceIDL2 = 2
	InstallAppCDKeyPriceIDL3 = 3
)

func GetInstallAppCDKeySuffix(cerLevel int) string {
	switch cerLevel {
	case InstallAppCDKeyPriceIDL1:
		return "L1"
	case InstallAppCDKeyPriceIDL2:
		return "L2"
	case InstallAppCDKeyPriceIDL3:
		return "L3"
	}
	return ""
}

func GetInstallAppCerPrice(cerLevel int64) int64 {
	switch cerLevel {
	case InstallAppCDKeyPriceIDL1:
		return CertificatePriceL1
	case InstallAppCDKeyPriceIDL2:
		return CertificatePriceL2
	case InstallAppCDKeyPriceIDL3:
		return CertificatePriceL3
	}
	/// 如果 cerLevel 不对则直接报错
	util.PanicIf(errors2.UnproccessableError(fmt.Sprintf("未识别的 cdkeyPriceID: %d", cerLevel)))
	return CertificatePriceL1
}
