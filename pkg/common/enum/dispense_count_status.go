//go:generate enumer -type=DispenseCountStatus -json -sql -transform=snake -trimprefix=DispenseCountStatus
// go get github.com/dmarkham/enumer
package enum

type DispenseCountStatus int

const (
	DispenseCountStatusNormal DispenseCountStatus = iota + 1
	DispenseCountStatusUsed
)
