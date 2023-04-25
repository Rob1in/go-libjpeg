//go:build !dynamic
package jpeg

/*
Note the contents of include/ are extracted from jpeg-turbo 2.1.5.1.
Contents of lib/ have been built from that source as well.
Libraries are actually copies of libturbojpeg.a as built for each architecture.
*/

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo linux,amd64 LDFLAGS: ${SRCDIR}/lib/libjpeg_linux_amd64.a
#cgo linux,arm64 LDFLAGS: ${SRCDIR}/lib/libjpeg_linux_arm64.a
#cgo darwin,amd64 LDFLAGS: ${SRCDIR}/lib/libjpeg_darwin_amd64.a
#cgo darwin,arm64 LDFLAGS: ${SRCDIR}/lib/libjpeg_darwin_arm64.a
*/
import "C"
