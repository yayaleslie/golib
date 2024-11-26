package strings

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func IIf(b bool, n, m string) string {
	if b {
		return n
	}

	return m
}

func Select(n, m string) string {
	return IIf(n != "", n, m)
}

func InArray(val string, list []string) (exists bool, index int) {
	exists, index = false, -1

	for i, v := range list {
		if val == v {
			exists, index = true, i
			break
		}
	}

	return
}

func ListToSet(list []string) map[string]bool {
	set := make(map[string]bool, len(list))

	for _, v := range list {
		set[v] = true
	}

	return set
}

func SetToList(set map[string]bool) []string {
	list := make([]string, 0, len(set))

	for v := range set {
		list = append(list, v)
	}

	return list
}

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func ShortNumStr(n int64) string {
	m, s := int64(0), make([]rune, 0, 8)

	for n > 0 {
		n, m = n/int64(len(letterRunes)), n%int64(len(letterRunes))
		s = append(s, letterRunes[m])
	}

	return string(s)
}

func Template(t string, params map[string]interface{}) string {
	pairs := make([]string, 0, 2*len(params))

	for k, v := range params {
		pairs = append(pairs, "{"+k+"}")

		if s, ok := v.(string); ok {
			pairs = append(pairs, s)
		} else {
			pairs = append(pairs, fmt.Sprintf("%v", v))
		}
	}

	return strings.NewReplacer(pairs...).Replace(t)
}

func ObjToString(obj interface{}) string {
	str, _ := json.Marshal(obj)
	return string(str)
}
