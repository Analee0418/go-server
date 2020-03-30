package utils

import "time"

func NowMilliseconds() int64 {
	return time.Now().UnixNano() / 1e6
}

func NowMillisecondsByTime(t time.Time) int64 {
	return t.UnixNano() / 1e6
}
