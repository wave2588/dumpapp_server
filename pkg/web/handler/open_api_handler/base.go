package open_api_handler

import (
	"context"
	"encoding/json"
	"github.com/go-playground/form"
	"net/http"
	"net/url"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller/impl"
)

func mustGetLoginID(ctx context.Context, r *http.Request) int64 {
	token := r.Header.Get("x-dumpapp-token")
	memberID, err := impl.DefaultMemberIDEncryptionController.GetMemberIDByCode(ctx, token)
	util.PanicIf(err)
	return memberID
}

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
