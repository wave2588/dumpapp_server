//go:generate enumer -type=MemberPayCountRecordType -json -sql -transform=snake -trimprefix=MemberPayCountRecordType
// go get github.com/dmarkham/enumer
package enum

import (
	"fmt"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/errors"
)

type MemberPayCountRecordType int

const (
	MemberPayCountRecordTypePay              MemberPayCountRecordType = iota + 1 /// 增加: 正常支付
	MemberPayCountRecordTypePayForFree                                           /// 增加: 多买多送
	MemberPayCountRecordTypeAdminPresented                                       /// 增加: 管理员添加
	MemberPayCountRecordTypeInvitedPresented                                     /// 增加: 邀请赠送
	MemberPayCountRecordTypeRebate                                               /// 增加: 被邀请的人充值后返还给邀请人

	MemberPayCountRecordTypeBuyIpa         /// 减少: 购买 ipa
	MemberPayCountRecordTypeBuyCertificate /// 减少: 购买证书
	MemberPayCountRecordTypeAdminDelete    /// 减少: 管理员删除
	MemberPayCountRecordTypeDispense       /// 减少: 兑换了下载次数
)

func ConvertMemberPayCountSourceToRecordType(source MemberPayCountSource) MemberPayCountRecordType {
	convertMap := map[MemberPayCountSource]MemberPayCountRecordType{
		MemberPayCountSourceNormal:           MemberPayCountRecordTypePay,
		MemberPayCountSourcePayForFree:       MemberPayCountRecordTypePayForFree,
		MemberPayCountSourceAdminPresented:   MemberPayCountRecordTypeAdminPresented,
		MemberPayCountSourceInvitedPresented: MemberPayCountRecordTypeInvitedPresented,
		MemberPayCountSourceRebate:           MemberPayCountRecordTypeRebate,
	}
	res, ok := convertMap[source]
	if !ok {
		util.PanicIf(errors.UnproccessableError(fmt.Sprintf("not support member_pay_count source: %s", source)))
	}
	return res
}

func ConvertMemberPayCountUseToRecordType(use MemberPayCountUse) MemberPayCountRecordType {
	convertMap := map[MemberPayCountUse]MemberPayCountRecordType{
		MemberPayCountUseIpa:         MemberPayCountRecordTypeBuyIpa,
		MemberPayCountUseCertificate: MemberPayCountRecordTypeBuyCertificate,
		MemberPayCountUseAdminDelete: MemberPayCountRecordTypeAdminDelete,
		MemberPayCountUseDispense:    MemberPayCountRecordTypeDispense,
	}
	res, ok := convertMap[use]
	if !ok {
		util.PanicIf(errors.UnproccessableError(fmt.Sprintf("not support member_pay_count use: %s", use)))
	}
	return res
}
