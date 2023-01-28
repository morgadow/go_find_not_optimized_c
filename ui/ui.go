package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/morgadow/go_find_not_optimized_c/finder"
)

// ui size
const Width = 400
const Height = 175

var worker = finder.NewTool() // static tool struct

// function for select file button
func CBSelectFile(parent fyne.Window, start *widget.Button, lbl *widget.Entry) {
	dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
		worker.SetSourceFile(uc.URI().Path())
		CBSetSource(worker.GetSourceFile(), start, lbl)
	}, parent)
}

// function for select folder button
func CBSelectFolder(parent fyne.Window, start *widget.Button, lbl *widget.Entry) {
	dialog.ShowFolderOpen(func(lu fyne.ListableURI, err error) {
		worker.SetSourceFolder(lu.Path())
		CBSetSource(worker.GetSourceFolder(), start, lbl)
	}, parent)
}

// function for start button
func CBStartSearching(parent fyne.Window) {
	_, err := worker.FindNonOptimizedFunctions()
	if err != nil {
		dialog.ShowError(err, parent)
	}
}

// callback for setting source file after a file or folder was selected
func CBSetSource(source string, start *widget.Button, lbl *widget.Entry) {
	if source != "" {
		start.Enable()
		lbl.SetText(source)
	}
}

func updateUI(status *widget.Label, prog *widget.ProgressBar) {
	for {
		status.SetText(worker.GetStatus())
		prog.SetValue(worker.GetProgress())
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {

	// new application
	a := app.New()
	win := a.NewWindow("Find not optimized C")
	win.Resize(fyne.NewSize(Width, Height))
	defer win.Close()

	res, err := fyne.LoadResourceFromPath("C:/workspace/go/src/github.com/morgadow/go_find_not_optimized_c/icon.png")
	if err != nil {
		fmt.Println(err)
	}
	win.SetIcon(res)

	// select optimiztions
	desc := widget.NewLabel("Select accepted optimizations:")
	check := widget.NewCheckGroup(finder.PossibleOptimizations, func(s []string) {
		worker.SetAcceptedOptimizations(s)
	})
	check.Horizontal = true

	// layout with currently selected file
	lbl := widget.NewLabel("Source:")
	sourceLbl := widget.NewEntry()
	sourceLbl.Enable()
	lbls := container.New(layout.NewFormLayout(), lbl, sourceLbl)

	// button layout
	startBtn := widget.NewButton("Start Searching", func() { CBStartSearching(win) })
	startBtn.Disable()
	fileBtn := widget.NewButton("Single File", func() { CBSelectFile(win, startBtn, sourceLbl) })
	folderBtn := widget.NewButton("Select Folder", func() { CBSelectFolder(win, startBtn, sourceLbl) })
	btns := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), fileBtn, layout.NewSpacer(), startBtn, layout.NewSpacer(), folderBtn, layout.NewSpacer())

	// progressbar and status label
	progress := widget.NewProgressBar()
	state := widget.NewLabel("")

	win.SetContent(container.New(layout.NewVBoxLayout(), desc, check, btns, lbls, layout.NewSpacer(), progress, canvas.NewLine(color.Black), state))
	go updateUI(state, progress) // start update task
	win.ShowAndRun()
}
