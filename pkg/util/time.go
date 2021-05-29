package util

import "time"

// 获取两个时间相差的天数，0表同一天，正数表t1>t2，负数表t1<t2
func GetDiffDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

// 获取t1和t2的相差天数，单位：秒，0表同一天，正数表t1>t2，负数表t1<t2
func GetDiffDaysBySecond(t1, t2 int64) int {
	time1 := time.Unix(t1, 0)
	time2 := time.Unix(t2, 0)

	// 调用上面的函数
	return GetDiffDays(time1, time2)
}
