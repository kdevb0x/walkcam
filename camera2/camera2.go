package main // import "walkcam"

import (
	"Java/android/hardware/Camera"
	"Java/android/hardware/CameraInfo"
	"Java/android/hardware/camera2/CameraCaptureSession"
	"Java/android/hardware/camera2/CameraDevice"
	"Java/android/hardware/camera2/CameraManager"
	"Java/android/hardware/camera2/CameraMetadata"
	"Java/android/hardware/camera2/CaptureFailure"
	"Java/android/hardware/camera2/CaptureRequest"
	"Java/android/hardware/camera2/CaptureResult"
)
import "Java/java/lang"

import (
	"runtime"
	"unsafe"

	"github.com/xlab/android-go/android"
	"github.com/xlab/android-go/app"
)
