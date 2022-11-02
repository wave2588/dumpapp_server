package constant

var OpsAuthNameMap = map[string]int64{
	"x-auth-header": 1390315185599680512,
}

var OpsAuthMemberIDMap = map[int64]string{
	1390315185599680512: "x-auth-header",
	1431956649500741632: "x-auth-header",
	1399256949626769408: "x-auth-header",
}

/// 招聘的代理
var OpsAuthMemberIDMapV2 = map[int64]string{
	1569917700480700416: "x-auth-header",
}

func CheckAllOpsByMemberID(memberID int64) bool {
	_, isV1Ok := OpsAuthMemberIDMap[memberID]
	_, isV2Ok := OpsAuthMemberIDMapV2[memberID]
	return isV1Ok || isV2Ok
}
