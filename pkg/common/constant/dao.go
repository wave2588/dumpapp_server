package constant

import (
	"encoding/json"

	"dumpapp_server/pkg/common/util"
)

type TransactionKey string

const TransactionKeyTxn TransactionKey = "txn"

type AdminDumpOrderBizExt struct {
	IpaName         string  `json:"ipa_name"`
	IpaBundleID     string  `json:"ipa_bundle_id"`
	IpaAppStoreLink string  `json:"ipa_app_store_link"`
	DemanderIDs     []int64 `json:"demander_ids"`
}

type SearchCount struct {
	IpaID int64
	Name  string
	Count int64
}

type IpaVersionBizExt struct {
	DescribeURL *string `json:"describe_url,omitempty"`
}

func (d *IpaVersionBizExt) String() string {
	data, err := json.Marshal(d)
	util.PanicIf(err)
	return string(data)
}
