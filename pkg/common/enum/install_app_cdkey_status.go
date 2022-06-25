//go:generate enumer -type=InstallAppCDKeyStatus -json -sql -transform=snake -trimprefix=InstallAppCDKeyStatus
// go get github.com/dmarkham/enumer
package enum

type InstallAppCDKeyStatus int

const (
	InstallAppCDKeyStatusNormal      InstallAppCDKeyStatus = iota + 1 /// 未使用
	InstallAppCDKeyStatusUsed                                         /// 已使用
	InstallAppCDKeyStatusAdminDelete                                  /// 管理员删除
)
