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
#cgo linux pkg-config: -ldl

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

int get_function_ptr_from_handle(void *lib_handle, void **function_ptr, const char *symbol_name) {
    dlerror();
    void *function = dlsym(lib_handle, symbol_name);
    char *error = dlerror();
    if (error != NULL) {
        fprintf(stderr, "Error while looking up symbol %s: %s\n", symbol_name, error);
        return -1;
    }
    *function_ptr = function;
    return 0;
}

int populate_symbols_table_from_handle(void *lib_handle, SymbolsTable *symbols){
	int ret_code = 0;
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_start_compress_ptr, "jpeg_start_compress");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_CreateCompress_ptr, "jpeg_CreateCompress");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_std_error_ptr, "jpeg_std_error");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_destroy_compress_ptr, "jpeg_destroy_compress");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_set_defaults_ptr, "jpeg_set_defaults");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_set_quality_ptr, "jpeg_set_quality");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_simple_progression_ptr, "jpeg_simple_progression");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_finish_compress_ptr, "jpeg_finish_compress");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_write_raw_data_ptr, "jpeg_write_raw_data");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_write_scanlines_ptr, "jpeg_write_scanlines");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_CreateDecompress_ptr, "jpeg_CreateDecompress");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_resync_to_restart_ptr, "jpeg_resync_to_restart");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_destroy_decompress_ptr, "jpeg_destroy_decompress");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_read_header_ptr, "jpeg_read_header");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_start_decompress_ptr, "jpeg_start_decompress");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_read_raw_data_ptr, "jpeg_read_raw_data");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_finish_decompress_ptr, "jpeg_finish_decompress");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_calc_output_dimensions_ptr, "jpeg_calc_output_dimensions");
	ret_code |= get_function_ptr_from_handle(lib_handle, (void**)&symbols->jpeg_read_scanlines_ptr, "jpeg_read_scanlines");
	return ret_code;
}

int init(void* handle){
	int ret_code = populate_symbols_table_from_handle(handle, &symbols_table);
	if (ret_code != 0) {
        return -1;
    }
    return 0;
}

*/
import "C"
import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"unsafe"
)

// Y/Cb/Cr Planes
const (
	Y  = 0
	Cb = 1
	Cr = 2
)

var HasLibJpeg bool
var ErrLibjpegNotFound = errors.New("Unable to load libjpeg or resolve all symbols")

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
	var locs []string
	switch runtime.GOOS {
	case "darwin":
		locs = []string{"libjpeg.dylib"}
		addBrewPath(&locs, "libjpeg.dylib")
	case "linux":
		locs = []string{"libjpeg.so"}
		addBrewPath(&locs, "libjpeg.so")
	case "windows":
		//TODO
		locs = []string{"libjpeg.dll"}
	default:
		fmt.Println("Unknown operating system")
		return
	}

	for _, l := range locs {
		loc := C.CString(l)
		defer C.free(unsafe.Pointer(loc))
		handle := C.dlopen(loc, C.RTLD_NOW)
		if handle == nil {
			fmt.Println("Couldn't open handle for libjpeg at location:", l)
			continue
		}
		if err := C.init(handle); err == 0 {
			HasLibJpeg = true
			break
		}
	}
}

func addBrewPath(slicePtr *[]string, filename string) {

	repoCmd := exec.Command("brew", "--repo")
	repoOutput, err := repoCmd.Output()
	if err != nil {
		return
	}
	repoPath := strings.TrimSpace(string(repoOutput))
	libPath := filepath.Join(repoPath, "lib", filename)
	if filepath.IsAbs(libPath) {
		*slicePtr = append(*slicePtr, libPath)
	}
}
