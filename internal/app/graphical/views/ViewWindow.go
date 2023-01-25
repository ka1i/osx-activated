package views

import (
	"fmt"
	"log"

	"github.com/kai1/osx-activated/internal/app/graphical/component"
	"github.com/kai1/osx-activated/pkg/version"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
)

func ActivateWindow() *cocoa.NSWindow {
	watermark := "Activate macOS\nGo to Settings to activate macOS."
	var fontSize float64 = 12

	viewW, viewH := component.ViewSize(watermark, "Monaco", fontSize)

	screen := cocoa.NSScreen_Main().Frame().Size

	windowSize := core.Rect(screen.Width-15-viewW, 15, viewW, viewH)

	log.Printf("Create WaterMark :%v\n", windowSize)

	t := component.TextView(watermark, "Monaco", fontSize, viewW, viewH)

	c := cocoa.NSView_Init(core.Rect(0, 0, viewW, viewH))
	// text background
	c.SetBackgroundColor(cocoa.Color(0.1, 0.1, 0.1, 0.223))
	c.SetWantsLayer(true)
	c.Layer().SetCornerRadius(16.0)
	c.AddSubviewPositionedRelativeTo(t, cocoa.NSWindowAbove, nil)

	w := cocoa.NSWindow_Init(windowSize,
		cocoa.NSBorderlessWindowMask,
		cocoa.NSBackingStoreBuffered,
		false,
	)
	w.SetContentView(c)
	w.SetTitlebarAppearsTransparent(true)
	w.SetTitleVisibility(cocoa.NSWindowTitleHidden)
	w.SetOpaque(false)
	w.SetBackgroundColor(cocoa.NSColor_Clear())
	w.SetLevel(cocoa.NSMainMenuWindowLevel + 2)
	w.SetFrameDisplay(windowSize, true)
	w.MakeKeyAndOrderFront(nil)

	return &w
}

func PreviewWindow() *cocoa.NSWindow {
	watermark := fmt.Sprintf("Build Version:%s", version.Version.ToString())
	var fontSize float64 = 10

	viewW, viewH := component.ViewSize(watermark, "Monaco", fontSize)

	windowSize := core.Rect(3, 3, viewW, viewH)

	log.Printf("Create WaterMark :%v\n", windowSize)

	t := component.TextView(watermark, "Monaco", fontSize, viewW, viewH)

	c := cocoa.NSView_Init(core.Rect(0, 0, viewW, viewH))
	// text background
	c.SetBackgroundColor(cocoa.Color(0.1, 0.1, 0.1, 0.223))
	c.SetWantsLayer(true)
	c.Layer().SetCornerRadius(16.0)
	c.AddSubviewPositionedRelativeTo(t, cocoa.NSWindowAbove, nil)

	w := cocoa.NSWindow_Init(windowSize,
		cocoa.NSBorderlessWindowMask,
		cocoa.NSBackingStoreBuffered,
		false,
	)
	w.SetContentView(c)
	w.SetTitlebarAppearsTransparent(true)
	w.SetTitleVisibility(cocoa.NSWindowTitleHidden)
	w.SetOpaque(false)
	w.SetBackgroundColor(cocoa.NSColor_Clear())
	w.SetLevel(cocoa.NSMainMenuWindowLevel + 2)
	w.SetFrameDisplay(windowSize, true)
	w.MakeKeyAndOrderFront(nil)

	return &w
}
