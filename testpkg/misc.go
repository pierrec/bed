package testpkg

import "time"

type MyTime time.Time

type Misc struct {
	Time   time.Time
	MyTime MyTime
}
