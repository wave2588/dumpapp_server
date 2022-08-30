//go:generate enumer -type=MemberPayCountUse -json -sql -transform=snake -trimprefix=MemberPayCountUse
// go get github.com/dmarkham/enumer
package enum

type MemberPayCountUse int

const (
	MemberPayCountUseIpa         MemberPayCountUse = iota + 1 /// 买的 ipa
	MemberPayCountUseCertificate                              /// 买证书了
	MemberPayCountUseAdminDelete                              /// 管理员删除
	MemberPayCountUseDispense
)
