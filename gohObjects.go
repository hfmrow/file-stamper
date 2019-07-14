// gohObjects.go

// Source file auto-generated on Sun, 14 Jul 2019 13:07:22 using Gotk3ObjHandler v1.3 ©2019 H.F.M

/*
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"github.com/gotk3/gotk3/gtk"
)

// Control over all used objects from glade.
var mainObjects *MainControlsObj

/****************************************/
/* Main structure Declaration. You may */
/* add your own declarations here,    */
/* all is preserved on updates       */
/************************************/
type MainControlsObj struct {
	ButtonApply           *gtk.Button
	ButtonExit            *gtk.Button
	ButtonInFiles         *gtk.Button
	ButtonOutFiles        *gtk.Button
	ButtonUndo            *gtk.Button
	CheckbuttonAtime      *gtk.CheckButton
	CheckbuttonCtime      *gtk.CheckButton
	CheckbuttonIncludeDir *gtk.CheckButton
	CheckbuttonMtime      *gtk.CheckButton
	CheckbuttonSubdir     *gtk.CheckButton
	EvenboxTop            *gtk.EventBox
	ImageMainTop          *gtk.Image
	mainUiBuilder         *gtk.Builder
	MainWindow            *gtk.Window
	Statusbar             *gtk.Statusbar
}

/*******************************************/
/* GtkObjects Initialisation. You may add */
/* your own declarations as you wish,    */
/* best way is to add them by grouping  */
/* objects names. Yours modifications  */
/* will be preserved on updates.      */
/*************************************/
// gladeObjParser: Initialise Gtk3 Objects into mainObjects structure.
func gladeObjParser() {
	mainObjects.ButtonApply = loadObject("ButtonApply").(*gtk.Button)
	mainObjects.ButtonExit = loadObject("ButtonExit").(*gtk.Button)
	mainObjects.ButtonInFiles = loadObject("ButtonInFiles").(*gtk.Button)
	mainObjects.ButtonOutFiles = loadObject("ButtonOutFiles").(*gtk.Button)
	mainObjects.ButtonUndo = loadObject("ButtonUndo").(*gtk.Button)
	mainObjects.CheckbuttonAtime = loadObject("CheckbuttonAtime").(*gtk.CheckButton)
	mainObjects.CheckbuttonCtime = loadObject("CheckbuttonCtime").(*gtk.CheckButton)
	mainObjects.CheckbuttonIncludeDir = loadObject("CheckbuttonIncludeDir").(*gtk.CheckButton)
	mainObjects.CheckbuttonMtime = loadObject("CheckbuttonMtime").(*gtk.CheckButton)
	mainObjects.CheckbuttonSubdir = loadObject("CheckbuttonSubdir").(*gtk.CheckButton)
	mainObjects.EvenboxTop = loadObject("EvenboxTop").(*gtk.EventBox)
	mainObjects.ImageMainTop = loadObject("ImageMainTop").(*gtk.Image)
	mainObjects.MainWindow = loadObject("MainWindow").(*gtk.Window)
	mainObjects.Statusbar = loadObject("Statusbar").(*gtk.Statusbar)
}
