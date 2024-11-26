package convert

import (
	"errors"
	"fmt"
	"strconv"
)

var Error = errors.New("convert failed")

func ToUint(data interface{}) (uint, error) {
	switch data.(type) {
	case uint8:
		return uint(data.(uint8)), nil
	case uint16:
		return uint(data.(uint16)), nil
	case uint32:
		return uint(data.(uint32)), nil
	case uint64:
		return uint(data.(uint64)), nil
	case int:
		return uint(data.(int)), nil
	case int8:
		return uint(data.(int8)), nil
	case int16:
		return uint(data.(int16)), nil
	case int32:
		return uint(data.(int32)), nil
	case int64:
		return uint(data.(int64)), nil
	case float32:
		return uint(data.(float32)), nil
	case float64:
		return uint(data.(float64)), nil
	case string:
		i, e := strconv.Atoi(data.(string))
		return uint(i), e
	default:
		return 0, Error
	}
}

func ToUint8(data interface{}) (uint8, error) {
	switch data.(type) {
	case uint8:
		return uint8(data.(uint8)), nil
	case uint16:
		return uint8(data.(uint16)), nil
	case uint32:
		return uint8(data.(uint32)), nil
	case uint64:
		return uint8(data.(uint64)), nil
	case int:
		return uint8(data.(int)), nil
	case int8:
		return uint8(data.(int8)), nil
	case int16:
		return uint8(data.(int16)), nil
	case int32:
		return uint8(data.(int32)), nil
	case int64:
		return uint8(data.(int64)), nil
	case float32:
		return uint8(data.(float32)), nil
	case float64:
		return uint8(data.(float64)), nil
	case string:
		i, e := strconv.Atoi(data.(string))
		return uint8(i), e
	default:
		return 0, Error
	}
}

func ToUint16(data interface{}) (uint16, error) {
	switch data.(type) {
	case uint8:
		return uint16(data.(uint8)), nil
	case uint16:
		return uint16(data.(uint16)), nil
	case uint32:
		return uint16(data.(uint32)), nil
	case uint64:
		return uint16(data.(uint64)), nil
	case int:
		return uint16(data.(int)), nil
	case int8:
		return uint16(data.(int8)), nil
	case int16:
		return uint16(data.(int16)), nil
	case int32:
		return uint16(data.(int32)), nil
	case int64:
		return uint16(data.(int64)), nil
	case float32:
		return uint16(data.(float32)), nil
	case float64:
		return uint16(data.(float64)), nil
	case string:
		i, e := strconv.ParseInt(data.(string), 10, 64)
		return uint16(i), e
	default:
		return 0, Error
	}
}

func ToUint32(data interface{}) (uint32, error) {
	switch data.(type) {
	case uint8:
		return uint32(data.(uint8)), nil
	case uint16:
		return uint32(data.(uint16)), nil
	case uint32:
		return uint32(data.(uint32)), nil
	case uint64:
		return uint32(data.(uint64)), nil
	case int:
		return uint32(data.(int)), nil
	case int8:
		return uint32(data.(int8)), nil
	case int16:
		return uint32(data.(int16)), nil
	case int32:
		return uint32(data.(int32)), nil
	case int64:
		return uint32(data.(int64)), nil
	case float32:
		return uint32(data.(float32)), nil
	case float64:
		return uint32(data.(float64)), nil
	case string:
		i, e := strconv.ParseInt(data.(string), 10, 64)
		return uint32(i), e
	default:
		return 0, Error
	}
}

func ToUint64(data interface{}) (uint64, error) {
	switch data.(type) {
	case uint8:
		return uint64(data.(uint8)), nil
	case uint16:
		return uint64(data.(uint16)), nil
	case uint32:
		return uint64(data.(uint32)), nil
	case uint64:
		return uint64(data.(uint64)), nil
	case int:
		return uint64(data.(int)), nil
	case int8:
		return uint64(data.(int8)), nil
	case int16:
		return uint64(data.(int16)), nil
	case int32:
		return uint64(data.(int32)), nil
	case int64:
		return uint64(data.(int64)), nil
	case float32:
		return uint64(data.(float32)), nil
	case float64:
		return uint64(data.(float64)), nil
	case string:
		i, e := strconv.ParseInt(data.(string), 10, 64)
		return uint64(i), e
	default:
		return 0, Error
	}
}

func ToInt(data interface{}) (int, error) {
	switch data.(type) {
	case uint8:
		return int(data.(uint8)), nil
	case uint16:
		return int(data.(uint16)), nil
	case uint32:
		return int(data.(uint32)), nil
	case uint64:
		return int(data.(uint64)), nil
	case int:
		return int(data.(int)), nil
	case int8:
		return int(data.(int8)), nil
	case int16:
		return int(data.(int16)), nil
	case int32:
		return int(data.(int32)), nil
	case int64:
		return int(data.(int64)), nil
	case float32:
		return int(data.(float32)), nil
	case float64:
		return int(data.(float64)), nil
	case string:
		return strconv.Atoi(data.(string))
	default:
		return 0, Error
	}
}

func ToInt8(data interface{}) (int8, error) {
	switch data.(type) {
	case uint8:
		return int8(data.(uint8)), nil
	case uint16:
		return int8(data.(uint16)), nil
	case uint32:
		return int8(data.(uint32)), nil
	case uint64:
		return int8(data.(uint64)), nil
	case int:
		return int8(data.(int)), nil
	case int8:
		return int8(data.(int8)), nil
	case int16:
		return int8(data.(int16)), nil
	case int32:
		return int8(data.(int32)), nil
	case int64:
		return int8(data.(int64)), nil
	case float32:
		return int8(data.(float32)), nil
	case float64:
		return int8(data.(float64)), nil
	case string:
		i, e := strconv.ParseInt(data.(string), 10, 64)
		return int8(i), e
	default:
		return 0, Error
	}
}

func ToInt16(data interface{}) (int16, error) {
	switch data.(type) {
	case uint8:
		return int16(data.(uint8)), nil
	case uint16:
		return int16(data.(uint16)), nil
	case uint32:
		return int16(data.(uint32)), nil
	case uint64:
		return int16(data.(uint64)), nil
	case int:
		return int16(data.(int)), nil
	case int8:
		return int16(data.(int8)), nil
	case int16:
		return int16(data.(int16)), nil
	case int32:
		return int16(data.(int32)), nil
	case int64:
		return int16(data.(int64)), nil
	case float32:
		return int16(data.(float32)), nil
	case float64:
		return int16(data.(float64)), nil
	case string:
		i, e := strconv.ParseInt(data.(string), 10, 64)
		return int16(i), e
	default:
		return 0, Error
	}
}

func ToInt32(data interface{}) (int32, error) {
	switch data.(type) {
	case uint8:
		return int32(data.(uint8)), nil
	case uint16:
		return int32(data.(uint16)), nil
	case uint32:
		return int32(data.(uint32)), nil
	case uint64:
		return int32(data.(uint64)), nil
	case int:
		return int32(data.(int)), nil
	case int8:
		return int32(data.(int8)), nil
	case int16:
		return int32(data.(int16)), nil
	case int32:
		return int32(data.(int32)), nil
	case int64:
		return int32(data.(int64)), nil
	case float32:
		return int32(data.(float32)), nil
	case float64:
		return int32(data.(float64)), nil
	case string:
		i, e := strconv.ParseInt(data.(string), 10, 64)
		return int32(i), e
	default:
		return 0, Error
	}
}

func ToInt64(data interface{}) (int64, error) {
	switch data.(type) {
	case uint8:
		return int64(data.(uint8)), nil
	case uint16:
		return int64(data.(uint16)), nil
	case uint32:
		return int64(data.(uint32)), nil
	case uint64:
		return int64(data.(uint64)), nil
	case int:
		return int64(data.(int)), nil
	case int8:
		return int64(data.(int8)), nil
	case int16:
		return int64(data.(int16)), nil
	case int32:
		return int64(data.(int32)), nil
	case int64:
		return int64(data.(int64)), nil
	case float32:
		return int64(data.(float32)), nil
	case float64:
		return int64(data.(float64)), nil
	case string:
		i, e := strconv.ParseInt(data.(string), 10, 64)
		return int64(i), e
	default:
		return 0, Error
	}
}

func ToFloat32(data interface{}) (float32, error) {
	switch data.(type) {
	case uint8:
		return float32(data.(uint8)), nil
	case uint16:
		return float32(data.(uint16)), nil
	case uint32:
		return float32(data.(uint32)), nil
	case uint64:
		return float32(data.(uint64)), nil
	case int:
		return float32(data.(int)), nil
	case int8:
		return float32(data.(int8)), nil
	case int16:
		return float32(data.(int16)), nil
	case int32:
		return float32(data.(int32)), nil
	case int64:
		return float32(data.(int64)), nil
	case float32:
		return data.(float32), nil
	case float64:
		return float32(data.(float64)), nil
	case string:
		i, e := strconv.ParseFloat(data.(string), 64)
		return float32(i), e
	default:
		return 0, Error
	}
}

func ToFloat64(data interface{}) (float64, error) {
	switch data.(type) {
	case uint8:
		return float64(data.(uint8)), nil
	case uint16:
		return float64(data.(uint16)), nil
	case uint32:
		return float64(data.(uint32)), nil
	case uint64:
		return float64(data.(uint64)), nil
	case int:
		return float64(data.(int)), nil
	case int8:
		return float64(data.(int8)), nil
	case int16:
		return float64(data.(int16)), nil
	case int32:
		return float64(data.(int32)), nil
	case int64:
		return float64(data.(int64)), nil
	case float32:
		return float64(data.(float32)), nil
	case float64:
		return data.(float64), nil
	case string:
		return strconv.ParseFloat(data.(string), 64)
	default:
		return 0, Error
	}
}

func ToString(data interface{}) (string, error) {
	switch data.(type) {
	case int:
		return strconv.Itoa(data.(int)), nil
	case int8:
		return strconv.FormatInt(int64(data.(int8)), 10), nil
	case int16:
		return strconv.FormatInt(int64(data.(int16)), 10), nil
	case int32:
		return strconv.FormatInt(int64(data.(int32)), 10), nil
	case int64:
		return strconv.FormatInt(data.(int64), 10), nil
	case uint:
		return strconv.FormatUint(uint64(data.(uint)), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(data.(uint8)), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(data.(uint16)), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(data.(uint32)), 10), nil
	case uint64:
		return strconv.FormatUint(data.(uint64), 10), nil
	//case float32:
	//	return strconv.FormatFloat(float64(data.(float32)), 'f', -1, 32), nil
	//case float64:
	//	return strconv.FormatFloat(data.(float64), 'f', -1, 32), nil
	case string:
		return data.(string), nil
	case []byte:
		return string(data.([]byte)), nil
	default:
		s := fmt.Sprintf("%v", data)
		return s, nil
	}
}

func ToStringNoNull(data interface{}) (string, error) {
	if data == nil {
		return "", nil
	}
	return ToString(data)
}

// 强制转换
func ForceToUint(data interface{}) uint {
	res, e := ToUint(data)
	if e != nil {
		//log.Printf("[ForceToUint] convert failed | err: %s", e.Error())
	}
	return res
}

func ForceToUint8(data interface{}) uint8 {
	res, e := ToUint8(data)
	if e != nil {
		//log.Printf("[ForceToUint8] convert failed | err: %s", e.Error())
	}
	return res
}

func ForceToUint16(data interface{}) uint16 {
	res, e := ToUint16(data)
	if e != nil {
		//log.Printf("[ForceToUint16] convert failed | err: %s", e.Error())
	}
	return res
}

func ForceToUint32(data interface{}) uint32 {
	res, e := ToUint32(data)
	if e != nil {
		//log.Printf("[ForceToUint32] convert failed | err: %s", e.Error())
	}
	return res
}

func ForceToUint64(data interface{}) uint64 {
	res, e := ToUint64(data)
	if e != nil {
		//log.Printf("[ForceToUint64] convert failed | err: %s", e.Error())
	}
	return res
}

func ForceToInt(data interface{}) int {
	res, e := ToInt(data)
	if e != nil {
		//log.Printf("[ForceToInt] convert failed | err: %s", e.Error())
	}
	return res
}

func ForceToInt8(data interface{}) int8 {
	res, e := ToInt8(data)
	if e != nil {
		//log.Printf("[ForceToInt8] convert failed | err: %s", e.Error())
	}
	return res
}

func ForceToInt16(data interface{}) int16 {
	res, e := ToInt16(data)
	if e != nil {
		//log.Printf("[ForceToInt16] convert failed | err: %s", e.Error())
	}
	return res
}

func ForceToInt32(data interface{}) int32 {
	res, e := ToInt32(data)
	if e != nil {
		//log.Printf("[ForceToInt32] convert failed | err: %s", e.Error())
	}
	return res
}

func ForceToInt64(data interface{}) int64 {
	res, e := ToInt64(data)
	if e != nil {
		//log.Printf("[ForceToInt64] convert failed | err: %s", e.Error())
	}
	return res
}

func ForceToFloat32(data interface{}) float32 {
	res, e := ToFloat32(data)
	if e != nil {
		//log.Printf("[ForceToFloat32] convert failed | err: %s", e.Error())
	}
	return res
}

func ForceToFloat64(data interface{}) float64 {
	res, e := ToFloat64(data)
	if e != nil {
		//log.Printf("[ForceToFloat64] convert failed | err: %s", e.Error())
	}
	return res
}

func ForceToString(data interface{}) string {
	res, e := ToString(data)
	if e != nil {
		//log.Printf("[ForceToString] convert failed | err: %s", e.Error())
	}
	return res
}

func ForceToStringNoNull(data interface{}) string {
	res, e := ToStringNoNull(data)
	if e != nil {
		//log.Printf("[ForceToStringNoNull] convert failed | err: %s", e.Error())
	}
	return res
}

func ToBool(data interface{}) (bool, error) {
	switch data.(type) {
	case bool:
		return data.(bool), nil
	default:
		return false, Error
	}
}
