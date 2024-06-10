package gui


import (
	"fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
	"image/color"
)


type CustomTheme struct {
	fyne.Theme
}


func (custom_theme CustomTheme) BackgroundColor() color.Color {
	return color.RGBA{R:40, G:42, B:54, A: 225}
}


func (custom_theme CustomTheme) ButtonColor() color.Color {
	return color.RGBA{R:64, G:224, B:208, A:255}
}


func (custom_theme CustomTheme) PrimaryColor() color.Color {
	return color.RGBA{R:64, G:224, B:208, A:225}
}

func(custom_theme CustomTheme) HoverColor() color.Color {
	return color.RGBA{R:72, G:61, B:139, A:255}
}


func GuiRun() {
	myApp := app.New()
	myApp.Settings().SetTheme(&CustomTheme{})

	main_window := myApp.NewWindow("Database migration tool")
	main_window.Resize(fyne.NewSize(600,600))

	argument_entry := widget.NewEntry()
	argument_entry.SetPlaceHolder("Enter args")


	submit_button := widget.NewButton("Submit", func(){})

	form := container.NewVBox(
		canvas.NewText("Enter args", color.White),
		argument_entry,
		submit_button,
	)

	background := canvas.NewRectangle(color.RGBA{R:40, G: 42, B: 54, A: 255})
	content := container.NewMax(background, form)

	main_window.SetContent(content)
	main_window.ShowAndRun()

}