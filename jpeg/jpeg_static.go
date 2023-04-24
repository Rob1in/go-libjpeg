//go:build !dynamic
// +build !dynamic

package jpeg

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo linux, arm64 LDFLAGS: -L${SRCDIR}/lib/libjpeg_linux_arm64.a
#cgo darwin, arm64 LDFLAGS: -${SRCDIR}/lib/libjpeg_darwin_arm64.a

*/
//import "C"
