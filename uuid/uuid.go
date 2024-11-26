package uuid

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenUUID(key ...string) string {
	var k string
	if len(key) > 0 {
		k = key[0]
	}
	return GenerateUuid(k, 1)
}

func GenerateUuid(key string, repeat int) string {
	if repeat <= 0 || repeat > 20 {
		repeat = 20
	}

	rand.Seed(int64(time.Now().Nanosecond()))
	i := 0
	sign := ""
	for i < repeat {
		num := rand.Intn(1000000) + 1000000
		u4 := uuid.NewMD5(uuid.New(), []byte(key+strconv.FormatUint(uint64(num), 10)))

		sign += strings.Replace(u4.String(), "-", "", -1)
		i++
	}

	return sign
}
