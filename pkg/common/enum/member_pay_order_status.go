//go:generate enumer -type=MemberPayOrderStatus -json -sql -transform=snake -trimprefix=MemberPayOrderStatus
// go get github.com/dmarkham/enumer
package enum

type MemberPayOrderStatus int

const (
	MemberPayOrderStatusPending MemberPayOrderStatus = iota + 1 /// 未支付
	MemberPayOrderStatusPaid                                    /// 已支付
)
