package handler

import (
	"context"
	"dumpapp_server/pkg/common/constant"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

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
		sv := SortValues{}
		err = sv.Decode(strings[0])
		return sv, err
	}, SortValues{})
	formEncoder = form.NewEncoder()
}

type SortValues []interface{}

func (sv *SortValues) Decode(text string) (err error) {
	cursorStr, err := url.QueryUnescape(text)
	if err != nil {
		return err
	}
	var value []interface{}
	if cursorStr != "" {
		if err := json.Unmarshal([]byte(cursorStr), &value); err != nil {
			return err
		}
	}
	*sv = value
	return
}

func (sv SortValues) Encode() (string, error) {
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
	return middleware.MustGetMemberID(ctx)
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
