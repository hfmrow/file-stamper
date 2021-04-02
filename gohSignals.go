// gohSignals.go

/*
	Source file auto-generated on Fri, 02 Apr 2021 15:44:30 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2019-21 hfmrow - file-stamper v1.1 github.com/hfmrow/file-stamper
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
