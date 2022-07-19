package constant

import "time"

var (
	MemberTypeGuest    = "guest"
	MemberTypeRegister = "register"

	AppOpsAuthNameHeaderKey = "Ops-Auth-Name"
)

var (
	MemberIDKey = "member_id"
	RemoteIP    = "remote_ip"
)

var HOST = "https://dumpapp.com/api"

// var HOST = "http://10.14.9.188:1996/api"

var (
	OneMinuteTTL = time.Minute
	OneHourTTL   = 60 * OneMinuteTTL
	OneDayTTL    = 24 * OneHourTTL
	ThreeDayTTL  = 3 * OneDayTTL
)
