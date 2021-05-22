//go:generate enumer -type=MemberVipOrderStatus -json -sql -transform=snake -trimprefix=MemberVipOrderStatus
// go get github.com/dmarkham/enumer
package enum

type MemberVipOrderStatus int

const (
	MemberVipOrderStatusPending MemberVipOrderStatus = iota + 1 /// 未支付
	MemberVipOrderStatusPaid                                    /// 已支付
)
