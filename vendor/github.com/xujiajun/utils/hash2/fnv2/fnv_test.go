package fnv2

import (
	"testing"
)

func TestHash32(t *testing.T) {
	expected := 2621965227
	text := "xujiajun"
	hash := Hash32(text)
	if hash != expected {
		t.Errorf("returned unexpected value : got %v want %v", hash, expected)
	}
}
