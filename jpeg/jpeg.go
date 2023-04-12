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


void get_function_ptr(void **function_ptr, const char *symbol_name) {
  // Clear any previous errors
  dlerror();

  // Use dlsym to get the function pointer for the specified symbol name
  void *function = dlsym(RTLD_DEFAULT, symbol_name);

  // Check for errors
  char *error = dlerror();
  if (error != NULL) {
      fprintf(stderr, "Error while looking up symbol %s: %s\n", symbol_name, error);
      return;
  }

  // Assign the function pointer to the specified function pointer variable
  *function_ptr = function;
}

void populate_symbols_table(SymbolsTable *symbols){
	//symbols for compression
	get_function_ptr((void**)&symbols->jpeg_start_compress_ptr, "jpeg_start_compress");
	get_function_ptr((void**)&symbols->jpeg_CreateCompress_ptr, "jpeg_CreateCompress");
	get_function_ptr((void**)&symbols->jpeg_std_error_ptr, "jpeg_std_error");
	get_function_ptr((void**)&symbols->jpeg_destroy_compress_ptr, "jpeg_destroy_compress");
	get_function_ptr((void**)&symbols->jpeg_set_defaults_ptr, "jpeg_set_defaults");
	get_function_ptr((void**)&symbols->jpeg_set_quality_ptr, "jpeg_set_quality");
	get_function_ptr((void**)&symbols->jpeg_simple_progression_ptr, "jpeg_simple_progression");
	get_function_ptr((void**)&symbols->jpeg_finish_compress_ptr, "jpeg_finish_compress");
	get_function_ptr((void**)&symbols->jpeg_write_raw_data_ptr, "jpeg_write_raw_data");
	get_function_ptr((void**)&symbols->jpeg_write_scanlines_ptr, "jpeg_write_scanlines");

	//symbols for decompression
	get_function_ptr((void**)&symbols->jpeg_CreateDecompress_ptr, "jpeg_CreateDecompress");
	get_function_ptr((void**)&symbols->jpeg_resync_to_restart_ptr, "jpeg_resync_to_restart");
	get_function_ptr((void**)&symbols->jpeg_destroy_decompress_ptr, "jpeg_destroy_decompress");
	get_function_ptr((void**)&symbols->jpeg_read_header_ptr, "jpeg_read_header");
	get_function_ptr((void**)&symbols->jpeg_start_decompress_ptr, "jpeg_start_decompress");
	get_function_ptr((void**)&symbols->jpeg_read_raw_data_ptr, "jpeg_read_raw_data");
	get_function_ptr((void**)&symbols->jpeg_finish_decompress_ptr, "jpeg_finish_decompress");
	get_function_ptr((void**)&symbols->jpeg_calc_output_dimensions_ptr, "jpeg_calc_output_dimensions");
	get_function_ptr((void**)&symbols->jpeg_read_scanlines_ptr, "jpeg_read_scanlines");

}


void init(){
	populate_symbols_table(&symbols_table);
}

*/
import "C"
import "fmt"

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

	//Find the library
	libs := []string{"/opt/homebrew/lib/libjpeg.8.2.2.dylib",
		"libjpeg.8.2.2.dylib@", "libjpeg62.dylib"}

	_, err := GetHandle(libs)
	if err != nil {
		panic("TU DOIS INSTALLER LIBJPEG")
	}
	fmt.Println(C.RTLD_GLOBAL)
	C.init()
}
