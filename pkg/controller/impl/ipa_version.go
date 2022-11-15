package impl

import (
	"context"
	"fmt"
	"strings"
	"time"

	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
)

type IpaVersionController struct {
	ipaVersionDAO  dao.IpaVersionDAO
	tencentCtl     controller.TencentController
	lingshulianCtl controller.LingshulianController
}

var DefaultIpaVersionController *IpaVersionController

func init() {
	DefaultIpaVersionController = NewIpaVersionController()
}

func NewIpaVersionController() *IpaVersionController {
	return &IpaVersionController{
		ipaVersionDAO:  impl.DefaultIpaVersionDAO,
		tencentCtl:     NewTencentController(),
		lingshulianCtl: NewLingshulianController(),
	}
}

func (c *IpaVersionController) Delete(ctx context.Context, ID int64) error {
	ipaVersionMap, err := c.ipaVersionDAO.BatchGet(ctx, []int64{ID})
	if err != nil {
		return err
	}
	ipaVersion, ok := ipaVersionMap[ID]
	if !ok {
		return nil
	}

	if err = c.ipaVersionDAO.Delete(ctx, ID); err != nil {
		return err
	}

	if ipaVersion.BizExt.Storage == "lingshulian" {
		_ = c.lingshulianCtl.Delete(ctx, config.DumpConfig.AppConfig.LingshulianShareIpaBucket, ipaVersion.TokenPath)
	} else {
		_ = c.tencentCtl.DeleteFile(ctx, ipaVersion.TokenPath)
	}

	return nil
}

func (c *IpaVersionController) GetDownloadURL(ctx context.Context, ID, loginID int64) (string, error) {
	ipaVersionMap, err := c.ipaVersionDAO.BatchGet(ctx, []int64{ID})
	if err != nil {
		return "", err
	}
	ipaVersion, ok := ipaVersionMap[ID]
	if !ok {
		return "", nil
	}

	var openURL string
	if ipaVersion.BizExt.Storage == "" || ipaVersion.BizExt.Storage == "cos" {
		openURL, err = c.tencentCtl.GetSignatureURL(ctx, ipaVersion.TokenPath, 30*time.Minute)
		openURL = fmt.Sprintf("%s&member_id=%d", openURL, loginID)
	} else if ipaVersion.BizExt.Storage == "lingshulian" {
		openURL, err = c.lingshulianCtl.GetSignatureURL(ctx, strings.ToUpper(config.DumpConfig.AppConfig.LingshulianShareIpaBucket), ipaVersion.TokenPath, 30*time.Minute)
	}
	return openURL, err
}
