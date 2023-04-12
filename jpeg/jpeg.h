#pragma once

#include <stdio.h>
#include <stdlib.h>
#include <setjmp.h>
#include "jpeglib.h"
#include "jerror.h"

// the dimension multiple to which data buffers should be aligned.
#define ALIGN_SIZE 16
//struct my_symbol_table {
//
//}
struct my_error_mgr {
	struct jpeg_error_mgr pub;
	jmp_buf jmpbuf;
};

#if defined(_WIN32) && !defined(__CYGWIN__)
// setjmp/longjmp occasionally crashes on Windows
// see https://github.com/golang/go/issues/13672
// use __builtin_setjmp/longjmp for workaround
#undef setjmp
#define setjmp(b)		__builtin_setjmp(b)
#undef longjmp
#define longjmp(b, c)	__builtin_longjmp((b), 1)	// __builtin_longjmp accepts only `1` as the secound argument
#endif

void error_longjmp(j_common_ptr cinfo);


typedef struct {
	void (*jpeg_start_compress_ptr)(j_compress_ptr, boolean);
	void (*jpeg_CreateCompress_ptr)(j_compress_ptr, int, size_t);
	struct jpeg_error_mgr * (*jpeg_std_error_ptr)(struct jpeg_error_mgr *err);
	void (*jpeg_destroy_compress_ptr)(j_compress_ptr);
	void (*jpeg_set_defaults_ptr)(j_compress_ptr);
	void (*jpeg_set_quality_ptr)(j_compress_ptr, int, boolean);
	void (*jpeg_simple_progression_ptr)(j_compress_ptr);
	void (*jpeg_finish_compress_ptr)(j_compress_ptr);
	JDIMENSION (*jpeg_write_raw_data_ptr)(j_compress_ptr, JSAMPIMAGE, JDIMENSION);
	JDIMENSION (*jpeg_write_scanlines_ptr) (j_compress_ptr, JSAMPARRAY, JDIMENSION);

	//for decompress
	void (*jpeg_CreateDecompress_ptr)(j_decompress_ptr, int, size_t);
    void (*jpeg_destroy_decompress_ptr)(j_decompress_ptr);
    int (*jpeg_read_header_ptr)(j_decompress_ptr, boolean);
    boolean (*jpeg_start_decompress_ptr)(j_decompress_ptr);
    JDIMENSION (*jpeg_read_raw_data_ptr)(j_decompress_ptr, JSAMPIMAGE, JDIMENSION);
    boolean (*jpeg_finish_decompress_ptr)(j_decompress_ptr);
    void (*jpeg_calc_output_dimensions_ptr)(j_decompress_ptr);
    JDIMENSION (*jpeg_read_scanlines_ptr)(j_decompress_ptr, JSAMPARRAY, JDIMENSION);
    //for source Manager
    boolean (*jpeg_resync_to_restart_ptr)(j_decompress_ptr, int);

} SymbolsTable;