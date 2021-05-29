package constant

import (
	"dumpapp_server/pkg/common/enum"
	"github.com/spf13/cast"
	"time"
)

const OneMonthPrice float64 = 68

var MemberVipDurationTypeToPrice = map[enum.MemberVipDurationType]int64{
	enum.MemberVipDurationTypeOneMonth:   cast.ToInt64(OneMonthPrice),
	enum.MemberVipDurationTypeThreeMonth: cast.ToInt64(OneMonthPrice * 3 * 0.85), /// 3 个月 85 折
	enum.MemberVipDurationTypeSixMonth:   cast.ToInt64(OneMonthPrice * 12 * 0.7), /// 12 个月 7折
}

var MemberVipDurationTypeToDays = map[enum.MemberVipDurationType]time.Time{
	enum.MemberVipDurationTypeOneMonth:   time.Now().AddDate(0, 1, 0),
	enum.MemberVipDurationTypeThreeMonth: time.Now().AddDate(0, 3, 0),
	enum.MemberVipDurationTypeSixMonth:   time.Now().AddDate(1, 0, 0),
}

var DurationToMemberVipDurationType = map[string]enum.MemberVipDurationType{
	enum.MemberVipDurationTypeOneMonth.String():   enum.MemberVipDurationTypeOneMonth,
	enum.MemberVipDurationTypeThreeMonth.String(): enum.MemberVipDurationTypeThreeMonth,
	enum.MemberVipDurationTypeSixMonth.String():   enum.MemberVipDurationTypeSixMonth,
}

var MemberVipDurationTypeToSubject = map[enum.MemberVipDurationType]string{
	enum.MemberVipDurationTypeOneMonth:   "DumpApp",
	enum.MemberVipDurationTypeThreeMonth: "DumpApp",
	enum.MemberVipDurationTypeSixMonth:   "DumpApp",
}
