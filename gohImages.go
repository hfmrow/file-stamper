// gohImages.go

/*
	Source file auto-generated on Fri, 02 Apr 2021 15:44:30 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2019-21 hfmrow - file-stamper v1.1 github.com/hfmrow/file-stamper
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

/**********************************************************/
/* This section preserve user modifications on update.   */
/* Images declarations, used to initialize objects with */
/* The SetPict() func, accept both kind of variables:  */
/* filename or []byte content in case of using        */
/* embedded binary data. The variables names are the */
/* same. "assetsDeclarationsUseEmbedded(bool)" func */
/* could be used to toggle between filenames and   */
/* embedded binary type. See SetPict()            */
/* declaration to learn more on how to use it.   */
/************************************************/
func assignImages() {
	setButtonImage(mainObjects.ButtonApply, apply32)
	setButtonImage(mainObjects.ButtonExit, logOut20)
	setButtonImage(mainObjects.ButtonInFiles, folderTimeCalOut48)
	setButtonImage(mainObjects.ButtonOutFiles, folderTimeCalIn48)
	setButtonImage(mainObjects.ButtonUndo, counterclockwiseArrow20)
	setImage(mainObjects.ImageMainTop, filestamper500x48)
	setWinIcon(mainObjects.MainWindow, folderTime48)
}

/**********************************************************/
/* This section is rewritten on assets update.           */
/* Assets var declarations, this step permit to make a  */
/* bridge between the differents types used, string or */
/* []byte, and to simply switch from one to another.  */
/*****************************************************/
var mainGlade interface{}               // assets/glade/main.glade
var apply32 interface{}                 // assets/images/Apply-32.png
var counterclockwiseArrow20 interface{} // assets/images/Counterclockwise-arrow-20.png
var filestamper300x29 interface{}       // assets/images/FileStamper-300x29.png
var filestamper500x48 interface{}       // assets/images/FileStamper-500x48.png
var folderTime48 interface{}            // assets/images/Folder-time-48.png
var folderTimeCalIn48 interface{}       // assets/images/Folder-time-cal-in-48.png
var folderTimeCalOut48 interface{}      // assets/images/Folder-time-cal-out-48.png
var logOut20 interface{}                // assets/images/Log-Out-20.png
var signCancel20 interface{}            // assets/images/Sign-cancel-20.png
var signSelect20 interface{}            // assets/images/Sign-Select-20.png
