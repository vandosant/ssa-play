package looper

import "testing"

func TestLooper(t *testing.T) {
	s := Looper(3)

	if s != "3" {
		t.Error("woops")
	}

	s = Looper(0)

	if s != "WHAT IS THIS YA JABRONEY" {
		t.Error("woops")
	}
}
