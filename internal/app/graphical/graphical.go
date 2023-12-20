package graphical

import (
	"github.com/kai1/osx-activated/internal/app/graphical/views"
	"github.com/progrium/macdriver/macos"
	"github.com/progrium/macdriver/macos/appkit"
)

func UserInterface() {

	macos.RunApp(func(a appkit.Application, ad *appkit.ApplicationDelegate) {

		views.ActivateWindow()

		views.PreviewWindow()
	})
}
