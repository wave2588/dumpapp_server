package statistics_dispense

import (
	"context"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"sort"
)

func Run() {
	fmt.Println("statistics_dispense")
	run()
}

type CountItem struct {
	ObjectID, Count int64
}

func run() {

	ctx := context.Background()

	offset := 0
	limit := 500
	hasNext := true

	memberSignIpaMap := make(map[int64]*models.MemberSignIpa)

	for hasNext {
		ids, err := impl.DefaultMemberSignIpaDAO.ListIDs(ctx, offset, limit, []qm.QueryMod{models.MemberSignIpaWhere.IsDelete.EQ(false)}, nil)
		util.PanicIf(err)

		offset += len(ids)
		hasNext = limit == len(ids)

		data, err := impl.DefaultMemberSignIpaDAO.BatchGet(ctx, ids)
		util.PanicIf(err)

		for id, ipa := range data {
			memberSignIpaMap[id] = ipa
		}
	}

	memberSignIpaIDs := make([]int64, 0)
	accountIDs := make([]int64, 0)
	for id, signIpa := range memberSignIpaMap {
		memberSignIpaIDs = append(memberSignIpaIDs, id)
		accountIDs = append(accountIDs, signIpa.MemberID)
	}

	accountMap, err := impl.DefaultAccountDAO.BatchGet(ctx, accountIDs)
	util.PanicIf(err)

	countMap, err := impl.DefaultDispenseCountRecordDAO.BatchGetCountByObjectIDs(ctx, memberSignIpaIDs)
	util.PanicIf(err)

	counts := make([]*CountItem, 0)
	for id, count := range countMap {
		counts = append(counts, &CountItem{
			ObjectID: id,
			Count:    count,
		})
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Count > counts[j].Count
	})

	isAll := false
	resultCounts := make([]*CountItem, 0)
	for _, count := range counts {
		if count.Count > 20 {
			resultCounts = append(resultCounts, count)
		}
		if count.Count > 50 {
			isAll = true
		}
	}

	nameStr := ""
	for _, count := range resultCounts {
		memberSignIpa, ok := memberSignIpaMap[count.ObjectID]
		if !ok {
			continue
		}
		account, ok := accountMap[memberSignIpa.MemberID]
		if !ok {
			continue
		}

		bucket := config.DumpConfig.AppConfig.LingshulianMemberSignIpaBucket
		url, _ := impl2.DefaultLingshulianController.GetURL(ctx, bucket, memberSignIpa.IpaFileToken)
		nameStr += fmt.Sprintf("用户邮箱: %s\nID: %d\n名称: %s\n分发次数: %d\n ipa 下载链接: [%s](%s)\n", account.Email, memberSignIpa.ID, memberSignIpa.BizExt.IpaName, count.Count, url, url)
		nameStr += "<font color=\"info\">-----------------------</font>\n"
	}

	if nameStr == "" {
		nameStr = fmt.Sprintf("暂无")
	}
	contentStr := fmt.Sprintf("<font color=\"warning\">分发 ipa 统计：</font>\n>")
	token := "e999db44-93f2-4515-9445-748e6c849e34"

	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": contentStr + nameStr,
		},
	}
	util2.SendWeiXinBot(ctx, token, data, []string{})

	if isAll {
		data = map[string]interface{}{
			"msgtype": "text",
			"text": map[string]interface{}{
				"content":        "",
				"mentioned_list": []string{"@all"},
			},
		}
		util2.SendWeiXinBot(ctx, token, data, []string{})
	}
}
