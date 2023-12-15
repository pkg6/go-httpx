package timex

import "time"

func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}
func Usleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}
