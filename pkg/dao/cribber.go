package dao

import "context"

type CribberDAO interface {
	IncrRemoteIP(ctx context.Context, ip string) error
	GetRemoteIPIncrCount(ctx context.Context, ip string) (int, error)

	/// 黑名单
	SetMemberIDToBlacklist(ctx context.Context, memberID int64) error
	GetBlacklistByMemberID(ctx context.Context, memberID int64) (bool, error)
}
