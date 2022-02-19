//go:generate enumer -type=MemberPayCountStatus -json -sql -transform=snake -trimprefix=MemberPayCountStatus
// go get github.com/dmarkham/enumer
package enum

type MemberPayCountStatus int

const (
	MemberPayCountStatusNormal MemberPayCountStatus = iota + 1
	MemberPayCountStatusUsed
	MemberPayCountStatusAdminDelete
)
