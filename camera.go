package walkcam

import (
	"log"
	"os"
	"syscall"
	"unsafe"

	"github.com/xlab/android-go/android"
	"github.com/xlab/android-go/app"
	gl "github.com/xlab/android-go/gles3"
	"gocv.io/x/gocv"
)

type Display struct {
	Window gl.
	SizeX  float64
	SizeY  float64
	Canvas *android.NativeWindow
}

// CameraDevice is a unix camera or video device.
type CameraDevice struct {
	Name string
	Fd   *os.File // the device; prob /dev/video0
}

type AndroidCamera struct {
	DeviceName string
}

func NewAndroidCamera(name string) AndroidCamera {
	return AndroidCamera{name}
}

func (c *AndroidCamera) Init() error {
	cameraManager := android.JNIEnvFindClass()
}
func (c AndroidCamera) OpenDevice(android.Jobject)
