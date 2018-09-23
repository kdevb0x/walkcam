package droidcamera // import "droidcamera"

// #cgo pkg-config: libcamera2ndk
// #include <sys/cdefs.h>
//
// #include "NdkCameraError.h"
// #include "NdkCameraMetadata.h"
// #include "NdkCameraDevice.h"

import "C"

import (
	"unsafe"
)

// CameraManager is instance of C Camera
type CameraManager struct {
	Manager C.ACameraManager
}

func CreateACameraManager() *CameraManager {
	cCameraManager := C.calloc(1, C.sizeof_ACameraManager)
	defer C.free(cCameraManager)
	camManager := (*C.ACameraManager)(cCameraManager)

}
func (m *CameraManager) Delete() {

}
