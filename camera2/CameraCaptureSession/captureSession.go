package CameraCaptureSession

import (
	"android.hardware.camera2.CameraDevice"
)

// public abstract int setRepeatingRequest (CaptureRequest request,
//                CameraCaptureSession.CaptureCallback listener,
//                Handler handler)

// to stop call stopReating(), or abortCaptures()

type CaptureCallback struct {
}

type StateCallback struct {
}

func CaptureCallback() *CaptureCallback {
	return &CaptureCallback{}
}

/*
// Unsure if explicit method creation in go is necessary, or if gobind takes care of it.

func (c *CaptureCallback) onCaptureBufferLost(session CameraCaptureSession, request CaptureRequest, target Surface, frameNumber int) {
}

*/
