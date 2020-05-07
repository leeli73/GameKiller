package time2

import (
	"testing"
	"time"
)

func TestStartAndEnd(t *testing.T) {
	Start()
	time.Sleep(time.Second)
	PrintEnd() //1.000249228s
}
