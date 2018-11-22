package walkcam

import (
	"image"
	"log"
	"os"
	"runtime"

	"github.com/go-gl/mathgl/mgl32"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/exp/sensor"
	"golang.org/x/mobile/gl"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/exp/shiny/widget"
	"golang.org/x/exp/shiny/widget/node"

	"github.com/xlab/android-go/android"
	"github.com/xlab/android-go/app"
	"github.com/xlab/android-go/gles"
)

type Point struct {
	X int
	Y int
}

type Flusher interface {
	// If priotity == true, will force the action imediately,
	// otherwise sceduled to run at next convenience,
	// in otherwords, it is added to the normal runque.
	Flush(priotity bool) error
}

// WindowContext should hold the currrent state of a surface,
// for example the camera preview inside a walkcam.FloatingWindow.
type WindowContext interface {
	Flusher
	Active() bool
	SwapContext(ctx interface{}) error
}

type FloatingWindow struct {
	Config   *android.Configuration
	Position Point
	WinCtx   WindowContext
}

func (f FloatingWindow) CurrentPosition() Point {
	return f.Position
}

func (f *FloatingWindow) SetPosition(x int, y int) error {
	f.Position.X = x
	f.Position.Y = y
	return nil
}

func NewFloatingWindow() *FloatingWindow {
	var conf = android.ConfigurationNew()
	win := &FloatingWindow{Config: conf, Position: Point{0, 0}}
	return win
}
