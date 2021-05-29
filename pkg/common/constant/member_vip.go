package constant

import (
	"dumpapp_server/pkg/common/enum"
	"github.com/spf13/cast"
)

const OneMonthPrice float64 = 68

var MemberVipDurationTypeToPrice = map[enum.MemberVipDurationType]int64{
	enum.MemberVipDurationTypeOneMonth:   cast.ToInt64(OneMonthPrice),
	enum.MemberVipDurationTypeThreeMonth: cast.ToInt64(OneMonthPrice * 3 * 0.85), /// 3 个月 85 折
	enum.MemberVipDurationTypeSixMonth:   cast.ToInt64(OneMonthPrice * 12 * 0.7), /// 12 个月 7折
}

var MemberVipDurationTypeToDays = map[enum.MemberVipDurationType]int{
	enum.MemberVipDurationTypeOneMonth:   30,
	enum.MemberVipDurationTypeThreeMonth: 30 * 3,
	enum.MemberVipDurationTypeSixMonth:   30 * 12,
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
