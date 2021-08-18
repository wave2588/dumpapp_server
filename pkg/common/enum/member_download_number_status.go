//go:generate enumer -type=MemberDownloadNumberStatus -json -sql -transform=snake -trimprefix=MemberDownloadNumberStatus
// go get github.com/dmarkham/enumer
package enum

type MemberDownloadNumberStatus int

const (
	MemberDownloadNumberStatusNormal MemberDownloadNumberStatus = iota + 1 /// 未使用
	MemberDownloadNumberStatusUsed                                         /// 已使用
)
