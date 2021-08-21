//go:generate enumer -type=MemberDownloadOrderStatus -json -sql -transform=snake -trimprefix=MemberDownloadOrderStatus
// go get github.com/dmarkham/enumer
package enum

type MemberDownloadOrderStatus int

const (
	MemberDownloadOrderStatusPending MemberDownloadOrderStatus = iota + 1 /// 未支付
	MemberDownloadOrderStatusPaid                                         /// 已支付
)
