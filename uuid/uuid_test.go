package uuid

import (
	"testing"
)

func TestGenerateUuid(t *testing.T) {
	key := "yaya"
	repeat := 2
	token := GenerateUuid(key, repeat)

	t.Log("len:", len(token))
	t.Log("token:", token)
}
