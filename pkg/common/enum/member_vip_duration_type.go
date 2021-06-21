//go:generate enumer -type=MemberVipDurationType -json -sql -transform=snake -trimprefix=MemberVipDurationType
// go get github.com/dmarkham/enumer
package enum

type MemberVipDurationType int

const (
	MemberVipDurationTypeOne MemberVipDurationType = iota + 1
	MemberVipDurationTypeTwo
	MemberVipDurationTypeThree
)
