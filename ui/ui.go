// WARNING! All changes made in this file will be lost!
package uigen

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
)

type UIUiMainWindow struct {
	Centralwidget *widgets.QWidget
	MainFrame *widgets.QFrame
	BotFrame *widgets.QFrame
	HLine *widgets.QFrame
	VLine *widgets.QFrame
	VersionLbl *widgets.QLabel
	InfoLbl *widgets.QLabel
	FolderBtn *widgets.QPushButton
	FileBtn *widgets.QPushButton
	SrcLbl *widgets.QLabel
	SrcEntryLbl *widgets.QLabel
	Startbtn *widgets.QToolButton
	ProgressBar *widgets.QProgressBar
	Menubar *widgets.QMenuBar
	MenuOptimization *widgets.QMenu
	MenuHelp *widgets.QMenu
	ActionAbout *widgets.QAction
	ActionAcceptAllOptimizations *widgets.QAction
	ActionAcceptO1 *widgets.QAction
	ActionAcceptO2 *widgets.QAction
	ActionAcceptO3 *widgets.QAction
	ActionAcceptO0 *widgets.QAction
	ActionAcceptOfast *widgets.QAction
	ActionAcceptOg *widgets.QAction
	ActionAcceptOs *widgets.QAction
}

func (this *UIUiMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("MainWindow")
	MainWindow.SetWindowModality(core.Qt__NonModal)
	MainWindow.SetGeometry(core.NewQRect4(0, 0, 330, 194))
	var sizePolicy *widgets.QSizePolicy
	sizePolicy = widgets.NewQSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed, widgets.QSizePolicy__DefaultType)
	sizePolicy.SetHorizontalStretch(0)
	sizePolicy.SetVerticalStretch(0)
	sizePolicy.SetHeightForWidth(MainWindow.SizePolicy().HasHeightForWidth())
	MainWindow.SetSizePolicy(sizePolicy)
	MainWindow.SetMinimumSize(core.NewQSize2(330, 194))
	MainWindow.SetMaximumSize(core.NewQSize2(330, 216))
	MainWindow.SetAutoFillBackground(true)
	this.Centralwidget = widgets.NewQWidget(MainWindow, core.Qt__Widget)
	this.Centralwidget.SetObjectName("Centralwidget")
	this.Centralwidget.SetMinimumSize(core.NewQSize2(330, 175))
	this.Centralwidget.SetMaximumSize(core.NewQSize2(330, 175))
	this.MainFrame = widgets.NewQFrame(this.Centralwidget, core.Qt__Widget)
	this.MainFrame.SetObjectName("MainFrame")
	this.MainFrame.SetGeometry(core.NewQRect4(0, 0, 330, 175))
	sizePolicy = widgets.NewQSizePolicy2(widgets.QSizePolicy__Maximum, widgets.QSizePolicy__Maximum, widgets.QSizePolicy__DefaultType)
	sizePolicy.SetHorizontalStretch(0)
	sizePolicy.SetVerticalStretch(0)
	sizePolicy.SetHeightForWidth(this.MainFrame.SizePolicy().HasHeightForWidth())
	this.MainFrame.SetSizePolicy(sizePolicy)
	this.MainFrame.SetMinimumSize(core.NewQSize2(330, 175))
	this.MainFrame.SetMaximumSize(core.NewQSize2(330, 175))
	this.MainFrame.SetFrameShape(widgets.QFrame__StyledPanel)
	this.MainFrame.SetFrameShadow(widgets.QFrame__Raised)
	this.BotFrame = widgets.NewQFrame(this.MainFrame, core.Qt__Widget)
	this.BotFrame.SetObjectName("BotFrame")
	this.BotFrame.SetGeometry(core.NewQRect4(0, 150, 330, 25))
	this.BotFrame.SetMinimumSize(core.NewQSize2(330, 25))
	this.BotFrame.SetMaximumSize(core.NewQSize2(330, 25))
	this.BotFrame.SetFrameShape(widgets.QFrame__StyledPanel)
	this.BotFrame.SetFrameShadow(widgets.QFrame__Raised)
	this.HLine = widgets.NewQFrame(this.BotFrame, core.Qt__Widget)
	this.HLine.SetObjectName("HLine")
	this.HLine.SetGeometry(core.NewQRect4(0, 0, 330, 2))
	this.HLine.SetMinimumSize(core.NewQSize2(330, 2))
	this.HLine.SetMaximumSize(core.NewQSize2(330, 2))
	this.HLine.SetFrameShadow(widgets.QFrame__Sunken)
	this.HLine.SetFrameShape(widgets.QFrame__HLine)
	this.VLine = widgets.NewQFrame(this.BotFrame, core.Qt__Widget)
	this.VLine.SetObjectName("VLine")
	this.VLine.SetGeometry(core.NewQRect4(250, 0, 2, 25))
	this.VLine.SetMinimumSize(core.NewQSize2(2, 25))
	this.VLine.SetMaximumSize(core.NewQSize2(2, 25))
	this.VLine.SetFrameShadow(widgets.QFrame__Sunken)
	this.VLine.SetFrameShape(widgets.QFrame__VLine)
	this.VersionLbl = widgets.NewQLabel(this.BotFrame, core.Qt__Widget)
	this.VersionLbl.SetObjectName("VersionLbl")
	this.VersionLbl.SetGeometry(core.NewQRect4(260, 2, 60, 21))
	this.VersionLbl.SetLayoutDirection(core.Qt__LeftToRight)
	this.VersionLbl.SetAlignment(core.Qt__AlignCenter)
	this.InfoLbl = widgets.NewQLabel(this.BotFrame, core.Qt__Widget)
	this.InfoLbl.SetObjectName("InfoLbl")
	this.InfoLbl.SetGeometry(core.NewQRect4(10, 2, 230, 21))
	this.FolderBtn = widgets.NewQPushButton(this.MainFrame)
	this.FolderBtn.SetObjectName("FolderBtn")
	this.FolderBtn.SetGeometry(core.NewQRect4(25, 20, 85, 30))
	this.FileBtn = widgets.NewQPushButton(this.MainFrame)
	this.FileBtn.SetObjectName("FileBtn")
	this.FileBtn.SetGeometry(core.NewQRect4(220, 20, 85, 30))
	this.SrcLbl = widgets.NewQLabel(this.MainFrame, core.Qt__Widget)
	this.SrcLbl.SetObjectName("SrcLbl")
	this.SrcLbl.SetGeometry(core.NewQRect4(10, 70, 40, 20))
	var font *gui.QFont
	font = gui.NewQFont()
	this.SrcLbl.SetFont(font)
	this.SrcEntryLbl = widgets.NewQLabel(this.MainFrame, core.Qt__Widget)
	this.SrcEntryLbl.SetObjectName("SrcEntryLbl")
	this.SrcEntryLbl.SetGeometry(core.NewQRect4(55, 70, 260, 20))
	font = gui.NewQFont()
	font.SetPointSize(8)
	this.SrcEntryLbl.SetFont(font)
	this.SrcEntryLbl.SetFrameShape(widgets.QFrame__Panel)
	this.SrcEntryLbl.SetFrameShadow(widgets.QFrame__Sunken)
	this.Startbtn = widgets.NewQToolButton(this.MainFrame)
	this.Startbtn.SetObjectName("Startbtn")
	this.Startbtn.SetEnabled(false)
	this.Startbtn.SetGeometry(core.NewQRect4(135, 15, 60, 40))
	font = gui.NewQFont()
	font.SetPointSize(12)
	font.SetWeight(75)
	this.Startbtn.SetFont(font)
	this.ProgressBar = widgets.NewQProgressBar(this.MainFrame)
	this.ProgressBar.SetObjectName("ProgressBar")
	this.ProgressBar.SetGeometry(core.NewQRect4(10, 120, 311, 15))
	this.ProgressBar.SetValue(0)
	this.ProgressBar.SetTextVisible(true)
	MainWindow.SetCentralWidget(this.Centralwidget)
	this.Menubar = widgets.NewQMenuBar(MainWindow)
	this.Menubar.SetObjectName("Menubar")
	this.Menubar.SetGeometry(core.NewQRect4(0, 0, 330, 21))
	this.MenuOptimization = widgets.NewQMenu(this.Menubar)
	this.MenuOptimization.SetObjectName("MenuOptimization")
	this.MenuHelp = widgets.NewQMenu(this.Menubar)
	this.MenuHelp.SetObjectName("MenuHelp")
	MainWindow.SetMenuBar(this.Menubar)
	this.ActionAbout = widgets.NewQAction(MainWindow)
	this.ActionAbout.SetObjectName("actionAbout")
	this.ActionAcceptAllOptimizations = widgets.NewQAction(MainWindow)
	this.ActionAcceptAllOptimizations.SetObjectName("actionAccept_all_Optimizations")
	this.ActionAcceptAllOptimizations.SetCheckable(true)
	this.ActionAcceptAllOptimizations.SetChecked(true)
	this.ActionAcceptO1 = widgets.NewQAction(MainWindow)
	this.ActionAcceptO1.SetObjectName("actionAccept_O1")
	this.ActionAcceptO1.SetCheckable(true)
	this.ActionAcceptO2 = widgets.NewQAction(MainWindow)
	this.ActionAcceptO2.SetObjectName("actionAccept_O2")
	this.ActionAcceptO2.SetCheckable(true)
	this.ActionAcceptO3 = widgets.NewQAction(MainWindow)
	this.ActionAcceptO3.SetObjectName("actionAccept_O3")
	this.ActionAcceptO3.SetCheckable(true)
	this.ActionAcceptO0 = widgets.NewQAction(MainWindow)
	this.ActionAcceptO0.SetObjectName("actionAccept_O0")
	this.ActionAcceptO0.SetCheckable(true)
	this.ActionAcceptOfast = widgets.NewQAction(MainWindow)
	this.ActionAcceptOfast.SetObjectName("actionAccept_Ofast")
	this.ActionAcceptOfast.SetCheckable(true)
	this.ActionAcceptOg = widgets.NewQAction(MainWindow)
	this.ActionAcceptOg.SetObjectName("actionAccept_Og")
	this.ActionAcceptOg.SetCheckable(true)
	this.ActionAcceptOs = widgets.NewQAction(MainWindow)
	this.ActionAcceptOs.SetObjectName("actionAccept_Os")
	this.ActionAcceptOs.SetCheckable(true)
	this.MenuOptimization.QWidget.AddAction(this.ActionAcceptAllOptimizations)
	this.MenuOptimization.QWidget.AddAction(this.ActionAcceptOs)
	this.MenuOptimization.QWidget.AddAction(this.ActionAcceptO1)
	this.MenuOptimization.QWidget.AddAction(this.ActionAcceptO2)
	this.MenuOptimization.QWidget.AddAction(this.ActionAcceptO3)
	this.MenuOptimization.QWidget.AddAction(this.ActionAcceptO0)
	this.MenuOptimization.QWidget.AddAction(this.ActionAcceptOfast)
	this.MenuOptimization.QWidget.AddAction(this.ActionAcceptOg)
	this.MenuHelp.QWidget.AddAction(this.ActionAbout)
	this.Menubar.QWidget.AddAction(this.MenuOptimization.MenuAction())
	this.Menubar.QWidget.AddAction(this.MenuHelp.MenuAction())


    this.RetranslateUi(MainWindow)

}

func (this *UIUiMainWindow) RetranslateUi(MainWindow *widgets.QMainWindow) {
    _translate := core.QCoreApplication_Translate
	MainWindow.SetWindowTitle(_translate("MainWindow", "MainWindow", "", -1))
	this.VersionLbl.SetText(_translate("MainWindow", "", "", -1))
	this.InfoLbl.SetText(_translate("MainWindow", "", "", -1))
	this.FolderBtn.SetText(_translate("MainWindow", "Select Folder", "", -1))
	this.FolderBtn.SetShortcut(_translate("MainWindow", "Ctrl+D", "", -1))
	this.FileBtn.SetText(_translate("MainWindow", "Select File", "", -1))
	this.FileBtn.SetShortcut(_translate("MainWindow", "Ctrl+F", "", -1))
	this.SrcLbl.SetText(_translate("MainWindow", "Source:", "", -1))
	this.SrcEntryLbl.SetText(_translate("MainWindow", "", "", -1))
	this.Startbtn.SetText(_translate("MainWindow", "Go", "", -1))
	this.Startbtn.SetShortcut(_translate("MainWindow", "Ctrl+Return", "", -1))
	this.MenuOptimization.SetTitle(_translate("MainWindow", "Optimization", "", -1))
	this.MenuHelp.SetTitle(_translate("MainWindow", "Help", "", -1))
	this.ActionAbout.SetText(_translate("MainWindow", "About", "", -1))
	this.ActionAcceptAllOptimizations.SetText(_translate("MainWindow", "Accept all Optimizations", "", -1))
	this.ActionAcceptO1.SetText(_translate("MainWindow", "Accept '-O1'", "", -1))
	this.ActionAcceptO2.SetText(_translate("MainWindow", "Accept '-O2'", "", -1))
	this.ActionAcceptO3.SetText(_translate("MainWindow", "Accept '-O3'", "", -1))
	this.ActionAcceptO0.SetText(_translate("MainWindow", "Accept '-O0'", "", -1))
	this.ActionAcceptOfast.SetText(_translate("MainWindow", "Accept '-Ofast'", "", -1))
	this.ActionAcceptOg.SetText(_translate("MainWindow", "Accept '-Og'", "", -1))
	this.ActionAcceptOs.SetText(_translate("MainWindow", "Accept '-Os'", "", -1))
}
