package conversion

import (
	"reflect"
	"unsafe"
)

func BytesToString(bytes []byte) (target string) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	th := (*reflect.StringHeader)(unsafe.Pointer(&target))
  
	th.Data = bh.Data
	th.Len = bh.Len
  
	return
}
