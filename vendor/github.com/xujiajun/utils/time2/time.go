package time2

import (
	"fmt"
	"time"
)

var startTime time.Time

func Start() {
	startTime = time.Now()
}

func End() time.Duration {
	return time.Since(startTime)
}

func PrintEnd() {
	fmt.Println(End())
}
