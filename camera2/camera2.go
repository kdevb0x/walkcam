package walkcam // import "walkcam"

import "Java/android/content/Context"
import "Java/android/hardware/camera2/CameraCaptureSession"
import "Java/android/hardware/camera2/CameraDevice"
import "Java/android/hardware/camera2/CameraManager"
import "Java/android/hardware/camera2/CameraMetadata"
import "Java/android/hardware/camera2/CaptureFailure"
import "Java/android/hardware/camera2/CaptureRequest"
import "Java/android/hardware/camera2/CaptureResult"
import "Java/android/view/WindowMan"
import "Java/java/lang"

import (
	"runtime"

	"github.com/xlab/android-go/android"
	"github.com/xlab/android-go/app"
)
