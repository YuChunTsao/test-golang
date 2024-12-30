package hello

// NOTE: There should be NO space between the comments and the `import "C"` line.
// The -ldl is sometimes necessary to fix linker errors about `dlsym`.

/*
#cgo LDFLAGS: ./pkg/test-rust/lib/hello/target/release/libhello.so -ldl
#include "../../pkg/test-rust/lib/hello.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

func Hello() {
	str1 := C.CString("world")
	str2 := C.CString("this is code from the static library")
	defer C.free(unsafe.Pointer(str1))
	defer C.free(unsafe.Pointer(str2))

	C.hello(str1)
	C.hello(str2)
}
