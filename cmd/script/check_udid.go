package main

import (
	"context"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/web/render"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	email := "863612913@qq.com"
	accountMap, err := impl.DefaultAccountDAO.BatchGetByEmail(ctx, []string{email})
	util.PanicIf(err)

	account, ok := accountMap[email]
	if !ok {
		fmt.Println("没有 account")
		return
	}

	devices, err := impl.DefaultMemberDeviceDAO.GetByMemberIDAndUDIDs(ctx, account.ID, udids)
	util.PanicIf(err)

	devideIDs := make([]int64, 0)
	for _, device := range devices {
		devideIDs = append(devideIDs, device.ID)
	}

	cerIDMap, err := impl.DefaultCertificateV2DAO.ListIDsByDeviceIDs(ctx, devideIDs)
	util.PanicIf(err)

	cerIDs := make([]int64, 0)
	for _, int64s := range cerIDMap {
		for _, i := range int64s {
			cerIDs = append(cerIDs, i)
		}
	}

	cerMap := render.NewCertificateRender(cerIDs, account.ID, render.CertificateDefaultRenderFields...).RenderMap(ctx)

	for _, device := range devices {
		cerIDs := cerIDMap[device.ID]
		for _, d := range cerIDs {
			cer, ok := cerMap[d]
			if !ok {
				fmt.Println("没找到")
				continue
			}

			tm := time.Unix(cer.CreatedAt, 0)

			fmt.Println(device.Udid, cer.P12IsActive, tm.Format("2006-01-02 15:04:05"), cer.Level)
		}
	}
}

var udids = []string{
	"645a8fa88ea640e8b8fe9e9b06e29ba46818d1a7",
	"00008030-000A5DD926FA802E",
}
