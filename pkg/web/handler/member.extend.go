package handler

import (
	"context"

	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/web/render"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (h *MemberHandler) getMemberAllCertificateByUDID(ctx context.Context, loginID int64, udid string, offset, limit int) ([]*render.Certificate, int64, error) {
	device, err := h.memberDeviceDAO.GetByMemberIDUdidSafe(ctx, loginID, udid)
	if err != nil {
		return nil, 0, err
	}
	if device == nil {
		return []*render.Certificate{}, 0, nil
	}
	filters := []qm.QueryMod{
		models.CertificateV2Where.DeviceID.IN([]int64{device.ID}),
	}
	ids, err := h.certificateDAO.ListIDs(ctx, offset, limit, filters, nil)
	if err != nil {
		return nil, 0, err
	}
	count, err := h.certificateDAO.Count(ctx, filters)
	if err != nil {
		return nil, 0, err
	}
	data := render.NewCertificateRender(ids, loginID, render.CertificateDefaultRenderFields...).RenderSlice(ctx)
	return data, count, nil
}

func (h MemberHandler) getMemberAllCertificate(ctx context.Context, loginID int64, offset, limit int) ([]*render.Certificate, int64, error) {
	deviceMap, err := h.memberDeviceDAO.BatchGetByMemberIDs(ctx, []int64{loginID})
	if err != nil {
		return nil, 0, err
	}

	deviceIDs := make([]int64, 0)
	for _, devices := range deviceMap {
		for _, memberDevice := range devices {
			deviceIDs = append(deviceIDs, memberDevice.ID)
		}
	}

	filters := []qm.QueryMod{
		models.CertificateV2Where.DeviceID.IN(deviceIDs),
	}
	ids, err := h.certificateDAO.ListIDs(ctx, offset, limit, filters, nil)
	if err != nil {
		return nil, 0, err
	}
	count, err := h.certificateDAO.Count(ctx, filters)
	if err != nil {
		return nil, 0, err
	}

	data := render.NewCertificateRender(ids, loginID, render.CertificateDefaultRenderFields...).RenderSlice(ctx)
	return data, count, nil
}
