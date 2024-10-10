package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello Fyne")
	// 设置窗口的初始大小 (宽度 x 高度)
	myWindow.Resize(fyne.NewSize(400, 130))

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter your name")

	label := widget.NewLabel("")

	button := widget.NewButton("Say Hello", func() {
		label.SetText("Hello " + entry.Text)
	})

	myWindow.SetContent(container.NewVBox(
		entry,
		button,
		label,
	))

	myWindow.ShowAndRun()
}
