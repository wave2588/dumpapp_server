//go:generate enumer -type=MemberSignIpaDownloadCountStatus -json -sql -transform=snake -trimprefix=MemberSignIpaDownloadCountStatus
// go get github.com/dmarkham/enumer
package enum

type MemberSignIpaDownloadCountStatus int

const (
	MemberSignIpaDownloadCountStatusNormal MemberSignIpaDownloadCountStatus = iota + 1 /// 未使用
	MemberSignIpaDownloadCountStatusUsed                                               /// 已使用
)
