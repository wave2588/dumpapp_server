package main

import (
	"context"
	util2 "dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/util"
	"fmt"
)

func main() {
	ctx := context.Background()

	email := ""
	password := ""

	id := util.MustGenerateID(ctx)
	fmt.Println(id)
	util2.PanicIf(impl.DefaultAccountDAO.Insert(ctx, &models.Account{
		ID:       id,
		Email:    email,
		Phone:    "",
		Password: password,
		Status:   0,
	}))

	util2.PanicIf(impl.DefaultMemberIDEncryptionDAO.Insert(ctx, &models.MemberIDEncryption{
		MemberID: id,
		Code:     util.MustGenerateCode(ctx, 10),
	}))
}
