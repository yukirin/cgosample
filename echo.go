package main

/*
#include <stdlib.h>
#include "echo.h"
*/
import "C"

import (
	"flag"
	"strings"
	"unsafe"
)

func main() {
	flag.Parse()
	cs := C.CString(strings.Join(flag.Args(), " "))
	defer C.free(unsafe.Pointer(cs))
	C.echo(cs)
}
