package rolling

//event type
const (
	Success = iota
	Failure
	Timeout
)

//HealthCounts rolling window counter
type HealthCounts struct {
	TotalNum int
	ErrNum   int
	ErrPer   int
}

//NewHealthCounts construct new counter
func NewHealthCounts(total int, err int) *HealthCounts {
	var per int

	if total > 0 {
		per = int((float64(err) / float64(total)) * 100)
	} else {
		per = 0
	}
	counts := &HealthCounts{
		TotalNum: total,
		ErrNum:   err,
		ErrPer:   per,
	}

	return counts
}

//Empty emtpy counter
var Empty = NewHealthCounts(0, 0)

//Plus counter + counter
func (count HealthCounts) Plus(eventTypes []int) *HealthCounts {
	updatedTotal := count.TotalNum
	updatedErr := count.ErrNum

	successNum := eventTypes[Success]
	failureNum := eventTypes[Failure]
	timeoutNum := eventTypes[Timeout]

	updatedTotal += (successNum + failureNum + timeoutNum)
	updatedErr += (failureNum + timeoutNum)
	return NewHealthCounts(updatedTotal, updatedErr)
}
