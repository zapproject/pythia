package qt

import (
	"context"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"

	"github.com/zapproject/pythia/config"
)

var cfg *config.Config
var ctx context.Context
var widget *widgets.QGroupBox

var window *widgets.QMainWindow

func App() *widgets.QApplication {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window = widgets.NewQMainWindow(nil, 0)

	window.SetMinimumSize2(980, 800)

	window.SetWindowTitle("Pythia")

	toolbar := startUp()

	// default page is set to config
	showConfig()

	window.SetCentralWidget(widget)
	window.AddToolBar(core.Qt__LeftToolBarArea, toolbar)

	// set up File menu bar
	fileMenu := window.MenuBar().AddMenu2("&File")
	fileMenu.AddAction("File Option")

	// set up settings menu bar
	settingsMenu := window.MenuBar().AddMenu2("&Settings")
	settingsMenu.AddAction("Settings Option")

	// set up help menu bar
	helpMenu := window.MenuBar().AddMenu2("&Help")
	helpMenu.AddAction("Help Option")

	window.Show()

	return app
}

// func createStatusWidget(){
// 	return
// }
