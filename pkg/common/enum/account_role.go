//go:generate enumer -type=AccountRole -json -sql -transform=snake -trimprefix=AccountRole
// go get github.com/dmarkham/enumer
package enum

type AccountRole int

const (
	AccountRoleNone        AccountRole = iota + 1 ///没有角色
	AccountRoleInfluential                        /// 大V
	AccountRoleAgent                              /// 代理商
)
