package main

import (
	"context"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"fmt"
)

func main() {

	ctx := context.Background()

	ss, err := impl.DefaultAdminConfigDAO.GetDailyFreeCount(ctx)
	util.PanicIf(err)

	fmt.Println("ss-->: ", ss)
}
