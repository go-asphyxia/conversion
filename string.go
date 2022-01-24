package conversion

import (
	"reflect"
	"unsafe"
)

func StringToBytesNoCopy(source string) (target []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&source))
	th := (*reflect.SliceHeader)(unsafe.Pointer(&target))

	th.Data = sh.Data
	th.Len = sh.Len
	th.Cap = sh.Len
	return
}