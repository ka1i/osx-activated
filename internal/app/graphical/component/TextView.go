package component

import (
	"github.com/progrium/macdriver/macos/appkit"
)

func ViewSize(text string, font string, size float64) (width, height float64) {
	t := appkit.NewTextViewWithFrame(RectOf(0, 0, 0, 0))
	t.SetString(text)
	t.SetFont(appkit.Font_FontWithNameSize(font, size))
	t.LayoutManager().EnsureLayoutForTextContainer(t.TextContainer())
	rect := t.LayoutManager().UsedRectForTextContainer(t.TextContainer())
	return rect.Size.Width + MI*2, rect.Size.Height
}

func TextView(text string, font string, size float64, viewW, viewH float64) appkit.TextView {
	t := appkit.NewTextViewWithFrame(RectOf(MI, 0, viewW, viewH))
	t.SetString(text)
	t.SetFont(appkit.Font_FontWithNameSize(font, size))
	t.SetTextColor(appkit.Color_ColorWithRedGreenBlueAlpha(1, 1, 1, 0.85))
	t.SetDrawsBackground(true)
	t.SetBackgroundColor(appkit.Color_ColorWithRedGreenBlueAlpha(0.33, 0.33, 0.33, 0.33))
	t.SetEditable(false)
	t.SetSelectable(false)
	t.SetImportsGraphics(false)
	t.SetDrawsBackground(false)
	return t
}
