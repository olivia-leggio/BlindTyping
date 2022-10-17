package main

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"

	"github.com/olivia-leggio/BlindTyping/pkg/textdata"
)

func openFilePicker(parent fyne.Window) {

	dialog.ShowFileOpen(func(path fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, parent)
		}
		if path != nil {
			textdata.NewFile(path.URI().Path())
		}
	}, parent)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Files")

	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() {
			myApp.Quit()
		}),
		fyne.NewMenuItem("Import", func() { openFilePicker(myWindow) }),
	)

	mainMenu := fyne.NewMainMenu(
		fileMenu,
	)

	myWindow.SetMainMenu(mainMenu)

	processedDir, _ := filepath.Abs("data/processed")
	fileDir, _ := os.Open(processedDir)
	defer fileDir.Close()
	existingFiles, _ := fileDir.Readdirnames(0)
	for _, file := range existingFiles {
		fmt.Println(file)
	}

	text := canvas.NewText("Bigger Test", color.Black)
	text.Alignment = fyne.TextAlignCenter

	box := container.NewVBox(text)

	myWindow.SetContent(box)

	myWindow.Resize(fyne.NewSize(900, 900))
	myWindow.ShowAndRun()
}
