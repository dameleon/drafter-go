package main

/*
#cgo CFLAGS: -I./drafter/src/
#cgo LDFLAGS: -lc++ -ldrafter -lsos -lmarkdownparser -lsnowcrash -lsundown -L"./drafter/build/out/Release/" -L"./drafter/build/out/Release/obj.target/"
#include <stdlib.h>
#include <stdio.h>
#include "drafter.h"
*/
import "C"
import (
	"unsafe"
	"fmt"
	"github.com/k0kubun/pp"
)

func ParseBlueprintTo(source string, options DrafterOptions) (string, error) {
	csrc := C.CString(source)
	defer C.free(unsafe.Pointer(csrc))
	var out string
	cout := C.CString(out)
	defer C.free(unsafe.Pointer(cout))
	copt := C.drafter_options{}
	copt.sourcemap = C.bool(options.sourcemap)
	copt.format = newCDrafterFormat(options.drafterFormat)
	res := newDrafterErrorCode(int(C.drafter_parse_blueprint_to(csrc, &cout, copt)))
	if res != DRAFTER_NO_ERROR {
		return "", fmt.Errorf("parse error: %s", res)
	}
	return C.GoString(cout), nil
}

func newCDrafterFormat(format DrafterFormat) C.drafter_format {
	switch format {
	case DRAFTER_SERIALIZE_YAML:
		return C.DRAFTER_SERIALIZE_YAML
	case DRAFTER_SERIALIZE_JSON:
		return C.DRAFTER_SERIALIZE_JSON
	default:
		panic("undefined drafter_format")
	}
}

func CheckBlueprint(source string) bool {
	csrc := C.CString(source)
	defer C.free(unsafe.Pointer(csrc))
	res := C.drafter_check_blueprint(csrc)
	defer C.drafter_free_result(res)
	return res == nil
}

func Version() uint {
	return uint(C.drafter_version())
}

func VersionString() string {
	return C.GoString(C.drafter_version_string())
}