package render

import (
	"context"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
)

type Config struct {
	AdminBusy bool `json:"admin_busy"`
}

type ConfigRender struct {
	loginID       int64
	includeFields []string

	config *Config

	configDAO dao.AdminConfigDAO
}

func NewConfigRender(loginID int64) *ConfigRender {
	f := &ConfigRender{
		loginID: loginID,

		configDAO: impl.DefaultAdminConfigDAO,
	}
	return f
}

func (f *ConfigRender) Render(ctx context.Context) *Config {
	f.fetch(ctx)
	return f.config
}

func (f *ConfigRender) fetch(ctx context.Context) {
	busy, err := f.configDAO.GetAdminBusy(ctx)
	util.PanicIf(err)

	f.config = &Config{
		busy,
	}
}
