package cgo_4

import "C"

func HelloCgo() {
	C.puts(C.CString("hello cgo \n"))
}
