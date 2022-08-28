package impl

import (
	"context"
	"fmt"
	"io/ioutil"

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

func (c *FileController) ListPlistFolder(ctx context.Context) {
}

func (c *FileController) GetPlistFileURL(ctx context.Context, key string) string {
	return fmt.Sprintf("https://dumpapp.com/plist/%s", key)
}

func (c *FileController) PutFileToLocal(ctx context.Context, path, key string, data []byte) error {
	fileName := fmt.Sprintf("%s/%s", path, key)
	return ioutil.WriteFile(fileName, data, 0o644)
}
