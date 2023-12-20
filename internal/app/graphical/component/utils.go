package component

import (
	"github.com/progrium/macdriver/macos/foundation"
)

var MI float64 = 6.0

func RectOf(x, y, width, height float64) foundation.Rect {
	return foundation.Rect{Origin: foundation.Point{X: x, Y: y}, Size: foundation.Size{Width: width, Height: height}}
}
