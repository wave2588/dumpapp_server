package render

import (
	"context"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
)

type Config struct {
	AdminBusy      bool   `json:"admin_busy"`
	DailyFreeCount int64  `json:"daily_free_count"`
	CerSource      string `json:"cer_source"`

	Announcement Announcement `json:"announcement"`
}

type Announcement struct {
	IsShow      bool    `json:"is_show"`
	Description *string `json:"description,omitempty"`
}

type ConfigRender struct {
	loginID       int64
	includeFields []string

	config *Config

	adminConfigDAO dao.AdminConfigInfoDAO
}

func NewConfigRender(loginID int64) *ConfigRender {
	f := &ConfigRender{
		loginID: loginID,

		adminConfigDAO: impl.DefaultAdminConfigInfoDAO,
	}
	return f
}

func (f *ConfigRender) Render(ctx context.Context) *Config {
	f.fetch(ctx)
	return f.config
}

func (f *ConfigRender) fetch(ctx context.Context) {
	config, err := f.adminConfigDAO.GetConfig(ctx)
	util.PanicIf(err)

	c := &Config{
		AdminBusy:      config.BizExt.AdminBusy,
		DailyFreeCount: config.BizExt.DailyFreeCount,
		CerSource:      config.BizExt.CerSource.String(),
	}

	if config.BizExt.Announcement != "" {
		c.Announcement = Announcement{
			IsShow:      true,
			Description: util.StringPtr(config.BizExt.Announcement),
		}
	}

	f.config = c
}
