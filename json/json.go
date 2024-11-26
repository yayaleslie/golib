package json

import (
	"fmt"
	"log"

	json "github.com/bytedance/sonic"
)

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func UnmarshalString(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}

func Convert(a, b interface{}) error {
	data, e := Marshal(a)
	if e != nil {
		return e
	}
	return Unmarshal(data, b)
}

func Byte(v interface{}) []byte {
	b, e := Marshal(v)
	if e != nil {
		log.Println(e)
	}
	return b
}

func String(v interface{}) string {
	b, e := Marshal(v)
	if e != nil {
		log.Println(e)
	}
	return string(b)
}

func Print(v interface{}) {
	fmt.Println(String(v))
}
