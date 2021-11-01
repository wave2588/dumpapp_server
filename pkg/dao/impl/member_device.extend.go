package impl

import (
	"context"

	"dumpapp_server/pkg/dao/models"
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
