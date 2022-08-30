//go:generate enumer -type=DispenseCountRecordType -json -sql -transform=snake -trimprefix=DispenseCountRecordType
// go get github.com/dmarkham/enumer
package enum

type DispenseCountRecordType int

const (
	DispenseCountRecordTypePay            DispenseCountRecordType = iota + 1 /// D 币兑换的次数
	DispenseCountRecordTypeInstallSignIpa                                    /// 安装扣费
	DispenseCountRecordTypeAdminPresented                                    /// 管理员添加的
	DispenseCountRecordTypeAdminDeleted                                      /// 管理员删除的
)
