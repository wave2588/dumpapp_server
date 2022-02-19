//go:generate enumer -type=MemberPayCountSource -json -sql -transform=snake -trimprefix=MemberPayCountSource
// go get github.com/dmarkham/enumer
package enum

type MemberPayCountSource int

const (
	MemberPayCountSourceNormal           MemberPayCountSource = iota + 1 /// 正常支付
	MemberPayCountSourcePayForFree                                       /// 多买多送
	MemberPayCountSourceAdminPresented                                   /// 管理员添加
	MemberPayCountSourceInvitedPresented                                 /// 邀请赠送
)
