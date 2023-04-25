//go:build !dynamic

package jpeg

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo linux,amd64 LDFLAGS: ${SRCDIR}/lib/libjpeg_linux_amd64.a
#cgo linux,arm64 LDFLAGS: ${SRCDIR}/lib/libjpeg_linux_arm64.a
#cgo darwin,amd64 LDFLAGS: ${SRCDIR}/lib/libjpeg_darwin_amd64.a
#cgo darwin,arm64 LDFLAGS: ${SRCDIR}/lib/libjpeg_darwin_arm64.a
*/
import "C"
