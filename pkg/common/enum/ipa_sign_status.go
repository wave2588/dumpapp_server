//go:generate enumer -type=IpaSignStatus -json -sql -transform=snake -trimprefix=IpaSignStatus
// go get github.com/dmarkham/enumer
package enum

type IpaSignStatus int

const (
	IpaSignStatusUnprocessed IpaSignStatus = iota + 1 /// 未签名
	IpaSignStatusProcessing                           /// 签名中
	IpaSignStatusSuccess                              /// 签名成功
	IpaSignStatusFail                                 /// 签名失败
)
