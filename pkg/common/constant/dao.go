package constant

import (
	"encoding/json"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
)

type TransactionKey = string

const TransactionKeyTxn TransactionKey = "txn"

type AdminDumpOrderBizExt struct {
	IpaName         string  `json:"ipa_name"`
	IpaBundleID     string  `json:"ipa_bundle_id"`
	IpaAppStoreLink string  `json:"ipa_app_store_link"`
	DemanderIDs     []int64 `json:"demander_ids"`
	IsOld           bool    `json:"is_old"`
}

type SearchCount struct {
	IpaID int64
	Count int64
}

type IpaSignBizExt struct {
	IpaVersionID int64        `json:"ipa_version_id"`
	IpaVersion   string       `json:"ipa_version"`
	IpaType      enum.IpaType `json:"ipa_type"`
	TokenPath    string       `json:"token_path"`
}

func (d *IpaSignBizExt) String() string {
	data, err := json.Marshal(d)
	util.PanicIf(err)
	return string(data)
}
