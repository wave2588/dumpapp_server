package impl

import (
	"context"
	"encoding/json"

	"dumpapp_server/pkg/common/enum"
	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	pkgErr "github.com/pkg/errors"
)

type AdminDumpOrderController struct {
	adminDumpOrderDAO dao.AdminDumpOrderDAO
}

var DefaultAdminDumpOrderController *AdminDumpOrderController

func init() {
	DefaultAdminDumpOrderController = NewAdminDumpOrderController()
}

func NewAdminDumpOrderController() *AdminDumpOrderController {
	return &AdminDumpOrderController{
		adminDumpOrderDAO: impl.DefaultAdminDumpOrderDAO,
	}
}

/// demanderID 需求者 id
func (c *AdminDumpOrderController) Upsert(ctx context.Context, demanderID, ipaID int64, ipaName, ipaVersion, ipaBundleID, ipaAppStoreLink string) error {
	order, err := c.adminDumpOrderDAO.GetByIpaIDIpaVersion(ctx, ipaID, ipaVersion)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		return err
	}

	if order == nil {
		bizExt := dao.AdminDumpOrderBizExt{
			IpaName:         ipaName,
			IpaBundleID:     ipaBundleID,
			IpaAppStoreLink: ipaAppStoreLink,
			DemanderIDs:     []int64{demanderID},
		}
		data, err := json.Marshal(bizExt)
		if err != nil {
			return err
		}
		return c.adminDumpOrderDAO.Insert(ctx, &models.AdminDumpOrder{
			DemanderID: demanderID,
			IpaID:      ipaID,
			IpaVersion: ipaVersion,
			IpaBizExt:  string(data),
			Status:     enum.AdminDumpOrderStatusProgressing,
		})
	}
	var bizExt dao.AdminDumpOrderBizExt
	err = json.Unmarshal([]byte(order.IpaBizExt), &bizExt)
	if err != nil {
		return err
	}
	bizExt.DemanderIDs = append(bizExt.DemanderIDs, demanderID)
	return c.adminDumpOrderDAO.Update(ctx, order)
}

func (c *AdminDumpOrderController) Progressed(ctx context.Context, operatorID, ipaID int64, ipaVersion string) error {
	data, err := c.adminDumpOrderDAO.GetByIpaIDIpaVersion(ctx, ipaID, ipaVersion)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		return err
	}
	/// 如果没有此条记录, 则过滤不做处理.
	if data == nil {
		return nil
	}
	data.OperatorID = operatorID
	data.Status = enum.AdminDumpOrderStatusProgressed
	return c.adminDumpOrderDAO.Update(ctx, data)
}

func (c *AdminDumpOrderController) Delete(ctx context.Context, operatorID, ipaID int64, ipaVersion string) error {
	data, err := c.adminDumpOrderDAO.GetByIpaIDIpaVersion(ctx, ipaID, ipaVersion)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		return err
	}
	/// 如果没有此条记录, 则过滤不做处理.
	if data == nil {
		return nil
	}
	data.OperatorID = operatorID
	data.Status = enum.AdminDumpOrderStatusDeleted
	return c.adminDumpOrderDAO.Update(ctx, data)
}
