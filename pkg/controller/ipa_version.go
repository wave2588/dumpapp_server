package controller

import "context"

type IpaVersionController interface {
	Delete(ctx context.Context, ID int64) error
	/// 获取下载链接
	GetDownloadURL(ctx context.Context, ID, loginID int64) (string, error)
}
