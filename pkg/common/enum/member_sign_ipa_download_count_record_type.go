//go:generate enumer -type=MemberSignIpaDownloadCountRecordType -json -sql -transform=snake -trimprefix=MemberSignIpaDownloadCountRecordType
// go get github.com/dmarkham/enumer
package enum

type MemberSignIpaDownloadCountRecordType int

const (
	/// 添加
	MemberSignIpaDownloadCountRecordTypeNormal         MemberSignIpaDownloadCountRecordType = iota + 1 /// D 币兑换的
	MemberSignIpaDownloadCountRecordTypeAdminPresented                                                 /// 管理后台添加的
	MemberSignIpaDownloadCountRecordTypeAdminDeleted                                                   /// 管理后台删除

	/// 消费
	MemberSignIpaDownloadCountRecordTypeInstallSignIpa /// 用户下载 ipa 消费
)
