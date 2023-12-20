package views

import (
	"fmt"
	"log"

	"github.com/kai1/osx-activated/internal/app/graphical/component"
	"github.com/kai1/osx-activated/pkg/version"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/objc"
)

func PreviewWindow() *appkit.Window {
	watermark := fmt.Sprintf("Build Version:%s", version.Version.ToString())
	var fontSize float64 = 10

	viewW, viewH := component.ViewSize(watermark, "Monaco", fontSize)

	windowSize := component.RectOf(3, 3, viewW, viewH)

	log.Printf("Create WaterMark :%v\n", windowSize)

	t := component.TextView(watermark, "Monaco", fontSize, viewW, viewH)

	c := appkit.NewViewWithFrame(component.RectOf(0, 0, viewW, viewH))

	objc.Call[objc.Void](c, objc.Sel("setBackgroundColor:"), appkit.Color_ColorWithRedGreenBlueAlpha(0.1, 0.1, 0.1, 0.223))

	c.SetWantsLayer(true)
	c.Layer().SetCornerRadius(16.0)
	c.AddSubviewPositionedRelativeTo(t, appkit.WindowAbove, nil)

	w := appkit.NewWindowWithContentRectStyleMaskBackingDefer(component.RectOf(0, 0, 0, 0),
		appkit.WindowStyleMaskBorderless, appkit.BackingStoreBuffered, false)
	objc.Retain(&w)

	w.SetContentView(c)
	w.SetTitlebarAppearsTransparent(true)
	w.SetTitleVisibility(appkit.WindowTitleHidden)
	w.SetOpaque(false)
	w.SetBackgroundColor(appkit.Color_ClearColor())
	w.SetLevel(appkit.MainMenuWindowLevel + 2)
	w.SetFrameDisplay(windowSize, true)
	w.MakeKeyAndOrderFront(nil)

	return &w
}
