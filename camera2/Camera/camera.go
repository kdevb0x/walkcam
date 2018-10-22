package walkcam // import "walkcam"

import (
	"Java/android/content/Context"
	"Java/android/hardware/camera2/CameraCaptureSession"
	"Java/android/hardware/camera2/CameraDevice"
	"Java/android/hardware/camera2/CameraManager"
	"Java/android/hardware/camera2/CameraMetadata"
	"Java/android/hardware/camera2/CaptureFailure"
	"Java/android/hardware/camera2/CaptureRequest"
	"Java/android/hardware/camera2/CaptureResult"
	"Java/android/view/WindowManager"
	"Java/java/lang"
)

// in order to generate the reverse-binding, need to run:
// 'gomobile bind -bootclasspath /opt/android-sdk/platforms/android-22/android.jar gopkg'
import (
	"runtime"

	"github.com/xlab/android-go/android"
	"github.com/xlab/android-go/app"
)

type CamManager struct {
	CameraManager
}

var manager = CamManager(Context.GetSystemService(Context.CAMERA_SERVICE))
