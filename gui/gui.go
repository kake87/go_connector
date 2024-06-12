package gui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// CustomTheme представляет собой пользовательскую тему
type CustomTheme struct {
	fyne.Theme
}

// BackgroundColor возвращает цвет фона
func (custom_theme CustomTheme) BackgroundColor() color.Color {
	return color.RGBA{R: 40, G: 42, B: 54, A: 255}
}

// ButtonColor возвращает цвет кнопок
func (custom_theme CustomTheme) ButtonColor() color.Color {
	return color.RGBA{R: 64, G: 224, B: 208, A: 255}
}

// PrimaryColor возвращает основной цвет
func (custom_theme CustomTheme) PrimaryColor() color.Color {
	return color.RGBA{R: 64, G: 224, B: 208, A: 255}
}

// HoverColor возвращает цвет при наведении
func (custom_theme CustomTheme) HoverColor() color.Color {
	return color.RGBA{R: 72, G: 61, B: 139, A: 255}
}

// Color возвращает цвет по указанному идентификатору и тематике
func (custom_theme CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return custom_theme.BackgroundColor()
	case theme.ColorNameButton:
		return custom_theme.ButtonColor()
	case theme.ColorNamePrimary:
		return custom_theme.PrimaryColor()
	case theme.ColorNameHover:
		return custom_theme.HoverColor()
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

// Font возвращает шрифт по указанному идентификатору и варианту
func (custom_theme CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

// Icon возвращает иконку по указанному идентификатору
func (custom_theme CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

// Size возвращает размер по указанному идентификатору
func (custom_theme CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
}

func createForm(dbType string) *fyne.Container {
	username_entry := widget.NewEntry()
	username_entry.SetPlaceHolder("Имя пользователя")

	password_entry := widget.NewEntry()
	password_entry.SetPlaceHolder("Пароль пользователя")

	hostname_entry := widget.NewEntry()
	hostname_entry.SetPlaceHolder("IP-адрес хоста")

	port_entry := widget.NewEntry()
	port_entry.SetPlaceHolder("Номер порта")

	dbname_entry := widget.NewEntry()
	dbname_entry.SetPlaceHolder("Подключаемая БД")

	usernamePasswordContainer := container.NewGridWithColumns(2,
		username_entry,
		password_entry,
	)
	return container.NewVBox(
		canvas.NewText(fmt.Sprintf("%s Конфигурации", dbType), color.White),
		usernamePasswordContainer,
		hostname_entry,
		port_entry,
		dbname_entry,
	)

}

func GuiRun() {
	myApp := app.New()
	myApp.Settings().SetTheme(&CustomTheme{})

	main_window := myApp.NewWindow("Database migration tool")
	main_window.Resize(fyne.NewSize(600, 600))

	db_types := []string{"MySQL", "PostgreSQL", "Windows Server"}
	fromDB_picker := widget.NewSelect(db_types, func(selected string) {})
	toDB_picker := widget.NewSelect(db_types, func(selected string) {})

	first_form := createForm("Настройка ")
	second_form := createForm("Куда залить")

	submit_button := widget.NewButton("Миграция данных", func() {
		fmt.Println("Начало миграции")
	})

	formContainer := container.NewVBox(
		canvas.NewText("Выберите исходную БД", color.White),
		fromDB_picker,
		first_form,
		widget.NewSeparator(),
		canvas.NewText("Выберите целевую БД", color.White),
		toDB_picker,
		second_form,
		submit_button,
	)

	background := canvas.NewRectangle(color.RGBA{R: 40, G: 42, B: 54, A: 255})
	content := container.NewMax(background, formContainer)

	main_window.SetContent(content)
	main_window.ShowAndRun()

}
