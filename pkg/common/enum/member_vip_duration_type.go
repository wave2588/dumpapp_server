//go:generate enumer -type=MemberVipDurationType -json -sql -transform=snake -trimprefix=MemberVipDurationType
// go get github.com/dmarkham/enumer
package enum

type MemberVipDurationType int

const (
	MemberVipDurationTypeOneMonth MemberVipDurationType = iota + 1
	MemberVipDurationTypeThreeMonth
	MemberVipDurationTypeSixMonth
)
