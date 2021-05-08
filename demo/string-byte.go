package demo

import (
	"reflect"
	"unsafe"
)

// 以下是标准库的转换方式
// string 转 []byte
func StringToByte(str string) []byte {
	return []byte(str)
}

// []byte 转 string
func ByteToString(b []byte) string {
	return string(b)
}

// 以下是强转换方式
// string 转 []byte
func StringToByteObligatory(str string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// []byte 转 string
func ByteToStringObligatory(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
