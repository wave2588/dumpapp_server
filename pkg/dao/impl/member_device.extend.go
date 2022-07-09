package impl

import (
	"context"

	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/dao/models"
	pkgErr "github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *MemberDeviceDAO) BatchGetByMemberIDs(ctx context.Context, memberIDs []int64) (map[int64][]*models.MemberDevice, error) {
	qs := []qm.QueryMod{
		models.MemberDeviceWhere.MemberID.IN(memberIDs),
	}
	data, err := models.MemberDevices(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}
	result := make(map[int64][]*models.MemberDevice)
	for _, datum := range data {
		result[datum.MemberID] = append(result[datum.MemberID], datum)
	}
	return result, nil
}

func (d *MemberDeviceDAO) GetByMemberIDUdidSafe(ctx context.Context, memberID int64, udid string) (*models.MemberDevice, error) {
	device, err := d.GetByMemberIDUdid(ctx, memberID, udid)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		return nil, err
	}
	return device, nil
}
