package main

import (
	"context"
	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"fmt"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {

	ctx := context.Background()

	cerID := cast.ToInt64(1552340805081894912)
	cerPrice := cast.ToInt64(0)

	offset := 0
	limit := 100
	hasNext := true

	for hasNext {
		filter := []qm.QueryMod{
			models.MemberPayCountRecordWhere.Type.EQ(enum.MemberPayCountRecordTypeBuyCertificate),
		}
		ids, err := impl.DefaultMemberPayCountRecordDAO.ListIDs(ctx, offset, limit, filter, nil)
		util.PanicIf(err)

		offset += len(ids)
		hasNext = limit == len(ids)

		resp, err := impl.DefaultMemberPayCountRecordDAO.BatchGet(ctx, ids)
		util.PanicIf(err)

		for _, record := range resp {
			//MemberPayCountRecordBizExtObjectTypeCertificate
			if record.BizExt.ObjectType != datatype.MemberPayCountRecordBizExtObjectTypeCertificate {
				continue
			}
			if record.BizExt.ObjectID == cerID {
				cerPrice = record.Count
				break
			}
		}
	}

	fmt.Println("证书价格--->: ", cerPrice)
}
