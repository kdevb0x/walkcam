package main // import "gocamera"

import (
	"os"
	"syscall"

	sc "github.com/donomii/sceneCamera"
	"golang.org/x/sys/unix"
)

func getDevice() *sc.SceneCamera {
	var err error
	if d, err := unix.Open("/dev/video0", unix.O_RDONLY|unix.O_DSYNC, 0755); err == nil {
		return &d
	}
	cam := sc.New()
	return cam
}

func main() {

}
