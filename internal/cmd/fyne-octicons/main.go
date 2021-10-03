package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	title = "Fyne Octicons"
)

func main() {
	a := app.New()
	w := a.NewWindow(title)
	w.Resize(fyne.NewSize(800, 400))
	img := widget.NewIcon(nil)
	list := container.NewVBox()
	for _, icon := range icons {
		i := icon
		button := widget.NewButton(i.name, func() {
			img.SetResource(i.resource)
		})
		list.Add(button)
	}
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewAdaptiveGridLayout(2),
			container.NewScroll(list),
			container.NewScroll(img),
		),
	)
	w.ShowAndRun()
}
