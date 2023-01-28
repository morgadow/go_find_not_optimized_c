package ui

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

// bundling all data
type UI struct {
	worker   finder.Tool
	win      fyne.Window
	startBtn *widget.Button
	running  bool
}

func NewUI() *UI {
	return &UI{}
}

func (f *UI) Run() {

	f.running = false
	f.worker = *finder.NewTool()

	// new application
	a := app.New()
	f.win = a.NewWindow("Find not optimized C")
	f.win.Resize(fyne.NewSize(Width, Height))
	defer f.win.Close()

	res, err := fyne.LoadResourceFromPath("C:/workspace/go/src/github.com/morgadow/go_find_not_optimized_c/icon.png")
	if err != nil {
		fmt.Println(err)
	}
	f.win.SetIcon(res)

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
	f.startBtn = widget.NewButton("Start Searching", func() { f.CBStartSearching() })
	f.startBtn.Disable()
	fileBtn := widget.NewButton("Single File", func() { f.CBSelectFile(sourceLbl) })
	folderBtn := widget.NewButton("Select Folder", func() { f.CBSelectFolder(sourceLbl) })
	btns := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), fileBtn, layout.NewSpacer(), f.startBtn, layout.NewSpacer(), folderBtn, layout.NewSpacer())

	// progressbar and status label
	progress := widget.NewProgressBar()
	state := widget.NewLabel("")

	f.win.SetContent(container.New(layout.NewVBoxLayout(), desc, check, canvas.NewLine(color.Gray16{0x4FA4}), btns, lbls, layout.NewSpacer(), canvas.NewLine(color.Black), progress, state))
	go func() { // start update task
		for {
			state.SetText(worker.GetStatus())
			progress.SetValue(worker.GetProgress())
			time.Sleep(50 * time.Millisecond)
		}
	}()

	f.win.ShowAndRun()
}

// function for select file button
func (f *UI) CBSelectFile(lbl *widget.Entry) {
	dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
		worker.SetSourceFile(uc.URI().Path())
		f.CBSetSource(worker.GetSourceFile(), f.startBtn, lbl)
	}, f.win)
}

// function for select folder button
func (f *UI) CBSelectFolder(lbl *widget.Entry) {
	dialog.ShowFolderOpen(func(lu fyne.ListableURI, err error) {
		worker.SetSourceFolder(lu.Path())
		f.CBSetSource(worker.GetSourceFolder(), f.startBtn, lbl)
	}, f.win)
}

// function for start button
func (f *UI) CBStartSearching() {
	f.running = true
	go func() {
		f.startBtn.Disable()
		_, err := worker.FindNonOptimizedFunctions()
		if err != nil {
			dialog.ShowError(err, f.win)
		}
		f.running = false
		f.startBtn.Enable()
	}()
}

// callback for setting source file after a file or folder was selected
func (f *UI) CBSetSource(source string, start *widget.Button, lbl *widget.Entry) {
	if source != "" {
		start.Enable()
		lbl.SetText(source)
	}
}
