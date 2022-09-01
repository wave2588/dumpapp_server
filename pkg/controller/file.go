package controller

import "context"

type FileController interface {
	/// plist 文件存储
	GetPlistFolderPath(ctx context.Context) string
	GetPlistFileURL(ctx context.Context, key string) string

	PutFileToLocal(ctx context.Context, path, key string, data []byte) error

	ListFolder(ctx context.Context, path string) ([]string, error)
	DeleteFile(ctx context.Context, path string) error
}