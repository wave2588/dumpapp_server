package main

import (
	"context"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"fmt"
)

func main() {
	ctx := context.Background()

	offset := 0
	limit := 100
	hasNext := true

	for hasNext {

		fmt.Println(offset)

		ids, err := impl.DefaultCertificateDeviceDAO.ListIDs(ctx, offset, limit, nil, nil)
		util.PanicIf(err)

		offset += len(ids)
		hasNext = len(ids) == limit

		res, err := impl.DefaultCertificateDeviceDAO.BatchGet(ctx, ids)
		util.PanicIf(err)

		cerIDs := make([]int64, 0)
		for _, device := range res {
			cerIDs = append(cerIDs, device.CertificateID)
		}

		cerMap, err := impl.DefaultCertificateDAO.BatchGet(ctx, cerIDs)
		util.PanicIf(err)

		cerV2Map, err := impl.DefaultCertificateV2DAO.BatchGet(ctx, cerIDs)
		util.PanicIf(err)

		for _, device := range res {
			cer, ok := cerMap[device.CertificateID]
			if !ok {
				fmt.Println("not found", device.CertificateID, device.ID)
				continue
			}
			_, ok = cerV2Map[device.CertificateID]
			if ok {
				fmt.Println("已存在--> ", device.CertificateID)
				continue
			}
			bizExt := &constant.CertificateBizExt{
				V1UDIDBatchNo:       cer.UdidBatchNo,
				V1CerAppleID:        cer.CerAppleid,
				OriginalP12Password: "1",
				NewP12Password:      "dumpapp",
			}
			util.PanicIf(impl.DefaultCertificateV2DAO.Insert(ctx, &models.CertificateV2{
				ID:                         cer.ID,
				DeviceID:                   device.DeviceID,
				P12FileData:                cer.P12FileDate,
				P12FileDataMD5:             cer.P12FileDateMD5,
				ModifiedP12FileDate:        cer.ModifiedP12FileDate,
				MobileProvisionFileData:    cer.MobileProvisionFileData,
				MobileProvisionFileDataMD5: cer.MobileProvisionFileDataMD5,
				Source:                     enum.CertificateSourceV1,
				BizExt:                     bizExt.String(),
				CreatedAt:                  cer.CreatedAt,
				UpdatedAt:                  cer.UpdatedAt,
			}))
		}
	}
}
