package hello

// NOTE: There should be NO space between the comments and the `import "C"` line.
// The -ldl is sometimes necessary to fix linker errors about `dlsym`.

/*
#cgo linux LDFLAGS: ./pkg/test-rust/lib/hello/target/release/libhello.so -ldl
#cgo linux CFLAGS: -I../../pkg/test-rust/lib
#cgo darwin LDFLAGS: ./pkg/test-rust/lib/hello/target/release/libhello.dylib -ldl
#cgo darwin CFLAGS: -I../../pkg/test-rust/lib
#cgo windows LDFLAGS: ./pkg/test-rust/lib/hello/target/release/hello.dll
#cgo windows CFLAGS: -I../../pkg/test-rust/lib
#include "hello.h"
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
