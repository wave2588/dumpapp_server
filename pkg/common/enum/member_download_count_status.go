//go:generate enumer -type=MemberDownloadCountStatus -json -sql -transform=snake -trimprefix=MemberDownloadCountStatus
// go get github.com/dmarkham/enumer
package enum

type MemberDownloadCountStatus int

const (
	MemberDownloadCountStatusNormal MemberDownloadCountStatus = iota + 1 /// 未使用
	MemberDownloadCountStatusUsed                                        /// 已使用
)
