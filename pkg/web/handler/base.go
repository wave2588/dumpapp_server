package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"dumpapp_server/pkg/common/constant"
	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	"github.com/go-playground/form"
	pkgErr "github.com/pkg/errors"
)

var (
	formDecoder *form.Decoder
	formEncoder *form.Encoder // nolint
)

func init() {
	formDecoder = form.NewDecoder()

	formDecoder.RegisterCustomTypeFunc(func(strings []string) (i interface{}, err error) {
		sv := Int64StringSlice{}
		err = sv.Decode(strings[0])
		return sv, err
	}, Int64StringSlice{})

	formEncoder = form.NewEncoder()
}

type Int64StringSlice []int64

func (sv *Int64StringSlice) Decode(text string) (err error) {
	data := []byte(text)
	var values []string
	err = json.Unmarshal(data, &values)
	if err != nil {
		// Fall back to array of integers:
		var values []int64
		if err := json.Unmarshal(data, &values); err != nil {
			return err
		}
		*sv = values
		return nil
	}
	*sv = make([]int64, len(values))
	for i, value := range values {
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		(*sv)[i] = value
	}
	return nil
}

func (sv Int64StringSlice) Encode() (string, error) {
	v, err := json.Marshal(sv)
	return string(v), err
}

func GetIntArgument(r *http.Request, key string, fallback int) int {
	if v, err := getIntArgument(r, key, 32); err != nil {
		return fallback
	} else {
		return int(v)
	}
}

func getIntArgument(r *http.Request, key string, bitSize int) (int64, error) {
	return strconv.ParseInt(r.URL.Query().Get(key), 10, bitSize)
}

func mustGetLoginID(ctx context.Context) int64 {
	account := mustGetLoginAccount(ctx)
	return account.ID
}

func mustGetLoginAccount(ctx context.Context) *models.Account {
	loginID := middleware.MustGetMemberID(ctx)
	account, err := impl.DefaultAccountDAO.Get(ctx, loginID)
	util.PanicIf(err)
	if account.Status == 2 {
		panic(errors.ErrAccountUnusual)
	}
	return account
}

func getLoginID(ctx context.Context) (int64, error) {
	return middleware.GetMemberID(ctx)
}

func GetAccountByLoginID(ctx context.Context, loginID int64) *models.Account {
	account, err := impl.DefaultAccountDAO.Get(ctx, loginID)
	if err != nil {
		if err := pkgErr.Cause(err); err != errors2.ErrNotFound {
			util.PanicIf(err)
		}
	}
	if account == nil {
		panic(errors.UnproccessableError("该邮箱未注册"))
	}
	return account
}

func GetAccountByEmail(ctx context.Context, email string) *models.Account {
	account, err := impl.DefaultAccountDAO.GetByEmail(ctx, email)
	if err != nil {
		if err := pkgErr.Cause(err); err != errors2.ErrNotFound {
			util.PanicIf(err)
		}
	}
	if account == nil {
		panic(errors.ErrNotFoundMember)
	}
	return account
}

func GetAccountByAccount(ctx context.Context, account string) *models.Account {
	/// 说明是邮箱登录
	if strings.Contains(account, "@") {
		return GetAccountByEmail(ctx, account)
	}
	return GetAccountByPhone(ctx, account)
}

func GetAccountByPhone(ctx context.Context, phone string) *models.Account {
	accountMap, err := impl.DefaultAccountDAO.BatchGetByPhones(ctx, []string{phone})
	util.PanicIf(err)
	account := accountMap[phone]
	if account == nil {
		panic(errors.ErrNotFoundMember)
	}
	return account
}

func DefaultSuccessBody(ctx context.Context) interface{} {
	appPlatform, ok := ctx.Value(constant.CtxKeyAppPlatform).(string)
	if !ok {
		return "ok"
	}
	/// 判断是否是 ios 平台
	if appPlatform == "ios" {
		return map[string]bool{
			"success": true,
		}
	}
	return "ok"
}
