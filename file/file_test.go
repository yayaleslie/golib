package file

import (
	"testing"
)

func TestSizeToString(t *testing.T) {
	s := SizeToString(1023 * 1024 * 1024)

	t.Log(s)
}
