package math

import (
	"testing"
)

func TestFloatFix(t *testing.T) {
	f := FloatFloor(2.99931, 4)
	t.Log(f)
}
