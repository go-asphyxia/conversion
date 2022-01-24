package conversion

import (
	"reflect"
	"unsafe"
)

func BytesToStringNoCopy(source []byte) (target string) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&source))
	th := (*reflect.StringHeader)(unsafe.Pointer(&target))
  
	th.Data = sh.Data
	th.Len = sh.Len
	return
}
