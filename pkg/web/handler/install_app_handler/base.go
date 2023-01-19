package install_app_handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	"github.com/go-playground/form"
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
