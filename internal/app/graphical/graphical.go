package graphical

import (
	"runtime"

	"github.com/kai1/osx-activated/internal/app/graphical/views"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/objc"
)

const (
	MenuTitle = "☘️"
)

func init() {
	runtime.LockOSThread()
}

func UserInterface() {
	cocoa.TerminateAfterWindowsClose = false

	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		win := views.ActivateWindow()

		// Status Bar
		statusBar := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		statusBar.Retain()
		statusBar.Button().SetTitle(MenuTitle)
		statusBar.SetMenu(*MenuSystem(win))
	})

	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyAccessory)
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}

func MenuSystem(win *cocoa.NSWindow) *cocoa.NSMenu {
	menuEnabled := cocoa.NSMenuItem_New()
	menuEnabled.Retain()
	menuEnabled.SetTitle("Enabled")
	menuEnabled.SetState(1)
	menuEnabled.SetAction(objc.Sel("enabled:"))
	cocoa.DefaultDelegateClass.AddMethod("enabled:", func(_ objc.Object) {
		if win.IsVisible() {
			win.Send("orderOut:", win)
			menuEnabled.SetState(0)
		} else {
			win.Send("orderFront:", win)
			menuEnabled.SetState(1)
		}
	})

	menuQuit := cocoa.NSMenuItem_New()
	menuQuit.SetTitle("Quit")
	menuQuit.SetAction(objc.Sel("terminate:"))

	// add menuitem to menu view
	menu := cocoa.NSMenu_New()
	menu.SetAutoenablesItems(false)
	menu.AddItem(menuEnabled)
	menu.AddItem(cocoa.NSMenuItem_Separator())
	menu.AddItem(menuQuit)

	return &menu
}
