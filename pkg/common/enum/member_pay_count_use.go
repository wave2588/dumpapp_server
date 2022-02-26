//go:generate enumer -type=MemberPayCountUse -json -sql -transform=snake -trimprefix=MemberPayCountUse
// go get github.com/dmarkham/enumer
package enum

type MemberPayCountUse int

const (
	MemberPayCountUseIpa MemberPayCountUse = iota + 1
	MemberPayCountUseCertificate
)
