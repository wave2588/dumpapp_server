package controller

import (
	"context"

	"dumpapp_server/pkg/dao/models"
)

type MemberDownloadController interface {
	GetDownloadNumber(ctx context.Context, loginID int64) (*models.MemberDownloadNumber, error)
}
