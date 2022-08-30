package impl

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"dumpapp_server/pkg/config"
)

type FileController struct{}

var DefaultFileController *FileController

func init() {
	DefaultFileController = NewFileController()
}

func NewFileController() *FileController {
	return &FileController{}
}

func (c *FileController) GetPlistFolderPath(ctx context.Context) string {
	if config.DumpConfig.AppConfig.Env == config.DumpEnvProduction {
		return "/home/wave/smash/web/plist"
	}
	return "/Users/wave/Downloads/plist"
}

func (c *FileController) GetPlistFileURL(ctx context.Context, key string) string {
	return fmt.Sprintf("https://dumpapp.com/plist/%s", key)
}

func (c *FileController) PutFileToLocal(ctx context.Context, path, key string, data []byte) error {
	fileName := fmt.Sprintf("%s/%s", path, key)
	return ioutil.WriteFile(fileName, data, 0o644)
}

func (c *FileController) ListFolder(ctx context.Context, path string) ([]string, error) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	for _, info := range fileInfos {
		result = append(result, info.Name())
	}
	return result, nil
}

func (c *FileController) DeleteFile(ctx context.Context, path string) error {
	_ = os.Remove(path)
	return nil
}
