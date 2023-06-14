package utime

import "time"

var Now = time.Now

func Freeze() {
	now := time.Now()
	Now = func() time.Time {
		return now
	}
}

func UnFreeze() {
	Now = time.Now
}
