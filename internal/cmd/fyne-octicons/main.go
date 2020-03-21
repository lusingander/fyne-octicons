package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

const (
	title = "Fyne Octicons"
)

func main() {
	a := app.New()
	w := a.NewWindow(title)
	w.Resize(fyne.NewSize(800, 400))
	img := widget.NewIcon(nil)
	list := widget.NewVBox()
	for _, icon := range icons {
		i := icon
		button := widget.NewButton(i.name, func() {
			img.SetResource(i.resource)
		})
		list.Append(button)
	}
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewAdaptiveGridLayout(2),
			widget.NewScrollContainer(list),
			widget.NewScrollContainer(img),
		),
	)
	w.ShowAndRun()
}
