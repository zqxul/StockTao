package util

import (
	"reflect"
	"time"
	"unsafe"
)

// GetValue => get value from point type value
func GetValue(value reflect.Value) interface{} {
	switch value.Interface().(type) {
	case *string:
		return *(*string)(unsafe.Pointer(value.Pointer()))
	case *int:
		return *(*int)(unsafe.Pointer(value.Pointer()))
	case *int8:
		return *(*int8)(unsafe.Pointer(value.Pointer()))
	case *int16:
		return *(*int16)(unsafe.Pointer(value.Pointer()))
	case *int32:
		return *(*int32)(unsafe.Pointer(value.Pointer()))
	case *int64:
		return *(*int64)(unsafe.Pointer(value.Pointer()))
	case *uint:
		return *(*uint)(unsafe.Pointer(value.Pointer()))
	case *uint8:
		return *(*uint8)(unsafe.Pointer(value.Pointer()))
	case *uint16:
		return *(*uint16)(unsafe.Pointer(value.Pointer()))
	case *uint32:
		return *(*uint32)(unsafe.Pointer(value.Pointer()))
	case *uint64:
		return *(*uint64)(unsafe.Pointer(value.Pointer()))
	case *uintptr:
		return *(*uintptr)(unsafe.Pointer(value.Pointer()))
	case *time.Time:
		return *(*time.Time)(unsafe.Pointer(value.Pointer()))
	case *float32:
		return *(*float32)(unsafe.Pointer(value.Pointer()))
	case *bool:
		return *(*bool)(unsafe.Pointer(value.Pointer()))
	case *complex64:
		return *(*complex64)(unsafe.Pointer(value.Pointer()))
	case *complex128:
		return *(*complex128)(unsafe.Pointer(value.Pointer()))
	}
	return nil
}
