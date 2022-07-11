//go:generate enumer -type=MemberPayOrderPlatform -json -sql -transform=snake -trimprefix=MemberPayOrderPlatform
// go get github.com/dmarkham/enumer
package enum

type MemberPayOrderPlatform int

const (
	MemberPayOrderPlatformWeb MemberPayOrderPlatform = iota + 1
	MemberPayOrderPlatformIOS
)
