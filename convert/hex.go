package convert

import (
	"strconv"
)

// HexToInt 16进制转换10进制 int
func HexToInt(hexStr string) (int, error) {
	val, err := strconv.ParseInt(hexStr, 16, 64)
	return int(val), err
}

// HexToInt16 16进制转换10进制 int64
func HexToInt16(hexStr string) (int64, error) {
	val, err := strconv.ParseInt(hexStr, 16, 64)
	return val, err
}
