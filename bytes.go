package conversion

import (
	"reflect"
	"unsafe"
)

func BytesToString(source []byte) (target string) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&source))
	th := (*reflect.StringHeader)(unsafe.Pointer(&target))
  
	th.Data = bh.Data
	th.Len = bh.Len
  
	return
}
