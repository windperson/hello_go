package TDD

import (
	"time"
)

func InInterval(interval Interval, target int) bool {
	var afterBegin = false
	if interval.IncludeBegin {
		afterBegin = target >= interval.Begin
	} else {
		afterBegin = target > interval.Begin
	}
	var beforeEnd = false
	if interval.IncludeEnd {
		beforeEnd = target <= interval.End
	} else {
		beforeEnd = target < interval.End
	}

	return afterBegin && beforeEnd
}

type Interval struct {
	Begin        int
	IncludeBegin bool
	End          int
	IncludeEnd   bool
}

func GetNowString() string {
	return time.Now().Format(time.RFC3339Nano)
}
