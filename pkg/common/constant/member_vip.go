package constant

import (
	"time"

	"dumpapp_server/pkg/common/enum"
)

var MemberVipDurationTypeToPrice = map[enum.MemberVipDurationType]int64{
	enum.MemberVipDurationTypeOne:   18,
	enum.MemberVipDurationTypeTwo:   28,
	enum.MemberVipDurationTypeThree: 38,
}

var DurationToMemberVipDurationType = map[string]enum.MemberVipDurationType{
	enum.MemberVipDurationTypeOne.String():   enum.MemberVipDurationTypeOne,
	enum.MemberVipDurationTypeTwo.String():   enum.MemberVipDurationTypeTwo,
	enum.MemberVipDurationTypeThree.String(): enum.MemberVipDurationTypeThree,
}

var MemberVipDurationTypeToSubject = map[enum.MemberVipDurationType]string{
	enum.MemberVipDurationTypeOne:   "DumpApp",
	enum.MemberVipDurationTypeTwo:   "DumpApp",
	enum.MemberVipDurationTypeThree: "DumpApp",
}

func GetMemberVipDays(durationType enum.MemberVipDurationType) time.Time {
	now := time.Now()
	MemberVipDurationTypeToDays := map[enum.MemberVipDurationType]time.Time{
		enum.MemberVipDurationTypeOne:   now.AddDate(0, 0, 10),
		enum.MemberVipDurationTypeTwo:   now.AddDate(0, 0, 20),
		enum.MemberVipDurationTypeThree: now.AddDate(0, 0, 30),
	}
	return MemberVipDurationTypeToDays[durationType]
}
