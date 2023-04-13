// Package jpeg decodes JPEG image to image.YCbCr using libjpeg (or libjpeg-turbo).
package jpeg

//
// Original codes are bollowed from go-thumber.
// Copyright (c) 2014 pixiv Inc. All rights reserved.
//
// See: https://github.com/pixiv/go-thumber
//

/*
//#cgo windows LDFLAGS: -ljpeg
//#cgo !windows pkg-config: libjpeg
#include <stdlib.h>
#include <stdio.h>
#include <jpeglib.h>
#include <jpeg.h>
#include <dlfcn.h>

static J_COLOR_SPACE getJCS_EXT_RGBA(void) {
#ifdef JCS_ALPHA_EXTENSIONS
	return JCS_EXT_RGBA;
#endif
  return JCS_UNKNOWN;
}

SymbolsTable symbols_table;

void get_function_ptr_from_handle(void *lib_handle, void **function_ptr, const char *symbol_name) {
 dlerror();
 void *function = dlsym(lib_handle, symbol_name);
 char *error = dlerror();
 if (error != NULL) {
     fprintf(stderr, "Error while looking up symbol %s: %s\n", symbol_name, error);
     return;
 }
 *function_ptr = function;
}

void populate_symbols_table_from_handle(void *lib_handle, SymbolsTable *symbols){

	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_start_compress_ptr, "jpeg_start_compress");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_CreateCompress_ptr, "jpeg_CreateCompress");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_std_error_ptr, "jpeg_std_error");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_destroy_compress_ptr, "jpeg_destroy_compress");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_set_defaults_ptr, "jpeg_set_defaults");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_set_quality_ptr, "jpeg_set_quality");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_simple_progression_ptr, "jpeg_simple_progression");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_finish_compress_ptr, "jpeg_finish_compress");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_write_raw_data_ptr, "jpeg_write_raw_data");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_write_scanlines_ptr, "jpeg_write_scanlines");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_CreateDecompress_ptr, "jpeg_CreateDecompress");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_resync_to_restart_ptr, "jpeg_resync_to_restart");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_destroy_decompress_ptr, "jpeg_destroy_decompress");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_read_header_ptr, "jpeg_read_header");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_start_decompress_ptr, "jpeg_start_decompress");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_read_raw_data_ptr, "jpeg_read_raw_data");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_finish_decompress_ptr, "jpeg_finish_decompress");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_calc_output_dimensions_ptr, "jpeg_calc_output_dimensions");
	get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_read_scanlines_ptr, "jpeg_read_scanlines");

}

void init(void* handle){
	populate_symbols_table_from_handle(handle, &symbols_table);
}

*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

// Y/Cb/Cr Planes
const (
	Y  = 0
	Cb = 1
	Cr = 2
)

// DCTMethod is the DCT/IDCT method type.
type DCTMethod C.J_DCT_METHOD

const (
	// DCTISlow is slow but accurate integer algorithm
	DCTISlow DCTMethod = C.JDCT_ISLOW
	// DCTIFast is faster, less accurate integer method
	DCTIFast DCTMethod = C.JDCT_IFAST
	// DCTFloat is floating-point: accurate, fast on fast HW
	DCTFloat DCTMethod = C.JDCT_FLOAT
)

func getJCS_EXT_RGBA() C.J_COLOR_SPACE {
	return C.getJCS_EXT_RGBA()
}

func init() {

	var handle unsafe.Pointer
	var locs []string
	switch runtime.GOOS {
	case "darwin":
		locs = []string{"libfoo.dylib", "/opt/homebrew/lib/libjpeg.dylib"}
		fmt.Println("Running on macOS")
	case "linux":
		locs = []string{"libjpeg.so"}
		fmt.Println("Running on Linux")
	case "windows":
		fmt.Println("Running on Windows")
	default:
		fmt.Println("Unknown operating system")
	}

	for _, l := range locs {
		loc := C.CString(l)
		//defer C.free(unsafe.Pointer(loc))
		handle = C.dlopen(loc, C.RTLD_NOW)
		C.free(unsafe.Pointer(loc))
		if handle == nil {
			fmt.Println("Couldn't find handle at location:", l)
		}
	}

	if handle == nil {
		panic("install libjpeg dependency")
	}
	C.init(handle)
}
