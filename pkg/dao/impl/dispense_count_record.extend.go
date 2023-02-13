package impl

import (
	"context"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/dao/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (d *DispenseCountRecordDAO) BatchGetByObjectIDsAndRecordType(ctx context.Context, objectIDs []int64, recordType enum.DispenseCountRecordType) (map[int64][]*models.DispenseCountRecord, error) {
	qs := []qm.QueryMod{
		models.DispenseCountRecordWhere.ObjectID.IN(objectIDs),
		models.DispenseCountRecordWhere.Type.EQ(recordType),
	}
	data, err := models.DispenseCountRecords(qs...).All(ctx, d.mysqlPool)
	if err != nil {
		return nil, err
	}

	result := make(map[int64][]*models.DispenseCountRecord)
	for _, datum := range data {
		result[datum.ObjectID] = append(result[datum.ObjectID], datum)
	}
	return result, nil
}

func (d *DispenseCountRecordDAO) BatchGetCountByObjectIDs(ctx context.Context, objectIDs []int64) (map[int64]int64, error) {
	qs := []qm.QueryMod{
		qm.Select(models.DispenseCountRecordColumns.ObjectID, "count(*) as count"),
		qm.From("dispense_count_record"),
		models.DispenseCountRecordWhere.ObjectID.IN(objectIDs),
		qm.GroupBy(models.DispenseCountRecordColumns.ObjectID),
	}

	var data []struct {
		ObjectID int64 `boil:"object_id"`
		Count    int64 `boil:"count"`
	}
	err := models.NewQuery(qs...).Bind(ctx, d.mysqlPool, &data)
	if err != nil {
		return nil, err
	}

	result := make(map[int64]int64)
	for _, datum := range data {
		result[datum.ObjectID] = datum.Count
	}
	return result, nil
}
