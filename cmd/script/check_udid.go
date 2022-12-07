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

	email := "multiplediaries@qq.com"
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

			fmt.Println(device.Udid, cer.P12IsActive, tm.Format("2006-01-02 15:04:05"))
		}
	}
}

var udids = []string{
	"00008110-001C10E83601801E",
	"00008120-000C55811180C01E",
	"00008110-001A691A11B9801E",
	"00008110-000929410ED3801E",
	"00008110-000524AC0A93801E",
	"00008101-001245183668001E",
	"9d8547b9721faec21f6488830d22a1ec1081f9b7",
	"00008110-001A18460EBB801E",
	"00008101-00022C3214C1001E",
}
