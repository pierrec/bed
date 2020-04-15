package testpkg

import "time"

type MyTime time.Time

type Time struct {
	Time   time.Time
	MyTime MyTime
}

type TimePtr struct {
	Time   *time.Time
	MyTime *MyTime
}

type TimeSlice struct {
	Time   []time.Time
	MyTime []MyTime
}

type TimeArray struct {
	Time   [4]time.Time
	MyTime [4]MyTime
}

type TimeMap struct {
	TimeKey   map[time.Time]bool
	MyTimeKey map[MyTime]bool
	Time      map[int]time.Time
	MyTime    map[int]MyTime
}
