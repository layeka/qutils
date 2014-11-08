package qutils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func SafeInt64(o interface{}) int64 {
	switch val := o.(type) {
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return int64(val)
	}
	panic(fmt.Sprintf("cannot convert a (type %v) to type int64", reflect.TypeOf(o)))
}
func SafeUInt64(o interface{}) uint64 {
	switch val := o.(type) {
	case uint:
		return uint64(val)
	case uint8:
		return uint64(val)
	case uint16:
		return uint64(val)
	case uint32:
		return uint64(val)
	case uint64:
		return uint64(val)
	case uintptr:
		return uint64(val)
	}
	panic(fmt.Sprintf("cannot convert a (type %v) to type uint64", reflect.TypeOf(o)))
}

func ParseInt64(o interface{}) (int64, error) {
	switch val := o.(type) {
	case string:
		{
			return strconv.ParseInt(val, 10, 0)
		}
	case int, int8, int16, int32, int64:
		return SafeInt64(val), nil
	case uint, uint8, uint16, uint32:
		{
			return int64(SafeUInt64(val)), nil
		}
	case uint64:
		{
			ival := int64(val)
			if strconv.FormatUint(val, 10) == strconv.FormatInt(ival, 10) {
				return ival, nil
			}
		}
	case uintptr:
		{
			ival := int64(val)
			if strconv.FormatUint(uint64(val), 10) == strconv.FormatInt(ival, 10) {
				return ival, nil
			}
		}
	case float32, float64:
		{
			fval := SafeFloat64(val)
			ival := int64(fval)
			if strconv.FormatFloat(fval, 'f', -1, 64) == strconv.FormatInt(ival, 10) {
				return ival, nil
			}
		}
	}
	return 0, errors.New(fmt.Sprintf("cannot convert a (type %v) to type int64", reflect.TypeOf(o)))
}

func SafeFloat64(o interface{}) float64 {
	switch val := o.(type) {
	case float32:
		return float64(val)
	case float64:
		return val
	}
	panic(fmt.Sprintf("cannot convert a (type %v) to type float64", reflect.TypeOf(o)))
}
