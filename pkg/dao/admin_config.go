package dao

import "context"

type AdminConfigDAO interface {
	SetAdminBusy(ctx context.Context, busy bool) error
	GetAdminBusy(ctx context.Context) (bool, error)
}
