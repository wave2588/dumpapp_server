//go:generate enumer -type=MemberPayExpenseRecordStatus -json -sql -transform=snake -trimprefix=MemberPayExpenseRecordStatus
// go get github.com/dmarkham/enumer
package enum

type MemberPayExpenseRecordStatus int

const (
	MemberPayExpenseRecordStatusAdd    MemberPayExpenseRecordStatus = iota + 1 /// 增加
	MemberPayExpenseRecordStatusReduce                                         /// 减少
)
