package encoding2

import (
	"testing"
)

type Demo struct {
	Id int
}

func TestWriteAndRead(t *testing.T) {
	expected := 1
	demo1 := Demo{Id: expected}

	filename := "demo"

	err := Write(demo1, filename)
	if err != nil {
		t.Failed()
	}

	var demo1Raw Demo

	Read(&demo1Raw, filename)

	if id := demo1Raw.Id; id != expected {
		t.Errorf("returned unexpected value : got %v want %v", id, expected)
	}
}
