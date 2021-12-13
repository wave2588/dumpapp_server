//go:generate enumer -type=IpaType -json -sql -transform=snake -trimprefix=IpaType
// go get github.com/dmarkham/enumer
package enum

type IpaType int

const (
	IpaTypeNormal IpaType = iota + 1 /// 正常破解包
	IpaTypeCrack                     /// 破解带插件
)
