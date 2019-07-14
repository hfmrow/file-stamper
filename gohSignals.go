// gohSignals.go

// Source file auto-generated on Sun, 14 Jul 2019 13:07:22 using Gotk3ObjHandler v1.3 ©2019 H.F.M

/*
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

/***********************************************************/
/* Signals & Property Implementations part. You may add   */
/* properties or signals as you wish, best way is to add */
/* them by grouping objects names yours modifications   */
/* will be preserved on updates.                       */
/******************************************************/
// signalsPropHandler: initialise signals used by gtk objects ...
func signalsPropHandler() {
	mainObjects.ButtonApply.Connect("clicked", ButtonApplyClicked)
	mainObjects.ButtonExit.Connect("clicked", ButtonExitClicked)
	mainObjects.ButtonInFiles.Connect("drag-data-received", ButtonInFilesReceived)
	mainObjects.ButtonInFiles.Connect("clicked", ButtonInFilesClicked)
	mainObjects.ButtonOutFiles.Connect("drag-data-received", ButtonOutFilesReceived)
	mainObjects.ButtonOutFiles.Connect("clicked", ButtonOutFilesClicked)
	mainObjects.ButtonUndo.Connect("clicked", ButtonUndoClicked)
	mainObjects.CheckbuttonAtime.Connect("notify", blankNotify)
	mainObjects.CheckbuttonCtime.Connect("notify", blankNotify)
	mainObjects.CheckbuttonIncludeDir.Connect("notify", blankNotify)
	mainObjects.CheckbuttonMtime.Connect("notify", blankNotify)
	mainObjects.CheckbuttonSubdir.Connect("notify", blankNotify)
	mainObjects.EvenboxTop.Connect("button-release-event", EvenboxClicked)
	mainObjects.ImageMainTop.Connect("notify", blankNotify)
	mainObjects.MainWindow.Connect("notify", blankNotify)
	mainObjects.Statusbar.Connect("notify", blankNotify)
}
