package main

import (
	"context"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"fmt"
)

func main() {

	ctx := context.Background()

	data, err := impl.DefaultIpaRankingDAO.GetIpaRankingData(ctx)
	util.PanicIf(err)
	fmt.Println("ss-->: ", data)
}
