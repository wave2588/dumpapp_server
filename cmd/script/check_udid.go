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

			fmt.Println(device.Udid, cer.P12IsActive, tm.Format("2006-01-02 15:04:05"), cer.Level)
		}
	}
}

var udids = []string{
	"00008020-00016DD10A92002E",
	"00008120-001455443E90C01E",
	"00008110-001E4D221110401E",
	"00008110-001A691A11B9801E",
	"00008101-001245183668001E",
	"00008110-001278141169801E",
	"00008110-001A18460EBB801E",
	"00008101-00060CEC2208001E",
	"00008120-001824CA0A33401E",
	"00008101-0014654A3A99001E",
	"00008101-00114D180151001E",
}
