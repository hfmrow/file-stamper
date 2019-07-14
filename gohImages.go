// gohImages.go

// Source file auto-generated on Sun, 14 Jul 2019 13:07:22 using Gotk3ObjHandler v1.3 ©2019 H.F.M

/*
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

/************************************************************/
/* Images declarations, used to initialize objects with it */
/* The functions: setImage, setWinIcon and setButtonImage */
/* accept both kind of variables: filename or []byte     */
/* content in case of using embedded binary data. The   */
/* variables names are the same. You can use function  */
/* "func assetsDeclarationsUseEmbedded(bool)"         */
/* to toggle between filenames and embedded binary.  */
/****************************************************/
func assignImages() {
	setButtonImage(mainObjects.ButtonApply, apply32)
	setButtonImage(mainObjects.ButtonExit, logOut20)
	setButtonImage(mainObjects.ButtonInFiles, folderTimeCalOut48)
	setButtonImage(mainObjects.ButtonOutFiles, folderTimeCalIn48)
	setButtonImage(mainObjects.ButtonUndo, counterclockwiseArrow20)
	setImage(mainObjects.ImageMainTop, filestamper500x48)
	setWinIcon(mainObjects.MainWindow, folderTime48)
}

// Assets var declarations, this step permit to make a "bridge" between the differents
// types used: (string or []byte) and to simply switch from one to another.
var apply32 interface{}                 // assets/images/Apply-32.png
var counterclockwiseArrow20 interface{} // assets/images/Counterclockwise-arrow-20.png
var filestamper300x29 interface{}       // assets/images/FileStamper-300x29.png
var filestamper500x48 interface{}       // assets/images/FileStamper-500x48.png
var folderTime48 interface{}            // assets/images/Folder-time-48.png
var folderTimeCalIn48 interface{}       // assets/images/Folder-time-cal-in-48.png
var folderTimeCalOut48 interface{}      // assets/images/Folder-time-cal-out-48.png
var logOut20 interface{}                // assets/images/Log-Out-20.png
var mainGlade interface{}               // assets/glade/main.glade
var signCancel20 interface{}            // assets/images/Sign-cancel-20.png
var signSelect20 interface{}            // assets/images/Sign-Select-20.png
