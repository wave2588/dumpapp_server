//go:generate enumer -type=MemberVipOrderStatus -json -sql -transform=snake -trimprefix=MemberVipOrderStatus
// go get github.com/dmarkham/enumer
package enum

/// 已经没有再用了, 为了保证 DAO 层不报错才留着
type MemberVipOrderStatus int

const (
	MemberVipOrderStatusPending MemberVipOrderStatus = iota + 1 /// 未支付
	MemberVipOrderStatusPaid                                    /// 已支付
)
