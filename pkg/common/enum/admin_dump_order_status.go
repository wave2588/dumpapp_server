//go:generate enumer -type=AdminDumpOrderStatus -json -sql -transform=snake -trimprefix=AdminDumpOrderStatus
// go get github.com/dmarkham/enumer
package enum

type AdminDumpOrderStatus int

const (
	AdminDumpOrderStatusProgressing AdminDumpOrderStatus = iota + 1 /// 处理中
	AdminDumpOrderStatusProgressed                                  /// 未处理
	AdminDumpOrderStatusDeleted                                     /// 已删除
)
