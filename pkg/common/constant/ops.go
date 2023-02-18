package constant

var OpsAuthNameMap = map[string]int64{
	"x-auth-headerv2": 1580838004182749184,
}

var OpsAuthMemberIDMap = map[int64]string{
	1580838004182749184: "x-auth-headerv2",
	1431956649500741632: "x-auth-headerv2",
	1399256949626769408: "x-auth-headerv2",
}

// / 招聘的代理
var OpsAuthMemberIDMapV2 = map[int64]string{
	1569917700480700416: "x-auth-headerv2",
}

func CheckAllOpsByMemberID(memberID int64) bool {
	_, isV1Ok := OpsAuthMemberIDMap[memberID]
	_, isV2Ok := OpsAuthMemberIDMapV2[memberID]
	return isV1Ok || isV2Ok
}
