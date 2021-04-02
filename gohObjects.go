// gohObjects.go

/*
	Source file auto-generated on Fri, 02 Apr 2021 15:44:30 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2019-21 hfmrow - file-stamper v1.1 github.com/hfmrow/file-stamper
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
