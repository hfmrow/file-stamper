// main.go

// Source file auto-generated on Sun, 14 Jul 2019 13:07:22 using Gotk3ObjHandler v1.3 ©2019 H.F.M

/*
	fileStamper v1.0 ©2019 H.F.M

	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:

	Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
	associated documentation files (the "Software"), to dealin the Software without restriction,
	including without limitation the rights to use, copy, modify, merge, publish, distribute,
	sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all copies or
	substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
	NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
	NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
	DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT
	OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

import (
	"fmt"

	gi "github.com/hfmrow/fileStamper/gtk3Import"
)

func main() {
	devMode = true

	/* Set to true when you choose using embedded assets functionality */
	assetsDeclarationsUseEmbedded(!devMode)

	/* Init Options */
	mainOptions = new(MainOpt)
	mainOptions.Init()

	/* Read Options */
	err = mainOptions.Read()
	if err != nil {
		fmt.Printf("%s\n%v\n", "Options file not found or error on parsing.", err)
	}

	/* Init AboutBox */
	mainOptions.About.InitFillInfos(
		"About "+Name,
		Name,
		Vers,
		Creat,
		YearCreat,
		LicenseAbrv,
		LicenseShort,
		Repository,
		Descr,
		filestamper300x29,
		signSelect20)

	/* Init gtk display */
	mainStartGtk(fmt.Sprintf("%s %s  %s %s %s",
		Name,
		Vers,
		"©"+YearCreat,
		Creat,
		LicenseAbrv),
		mainOptions.MainWinWidth,
		mainOptions.MainWinHeight, true)
}

func mainApplication() {

	/* Translate init. */
	translate = MainTranslateNew(absoluteRealPath+mainOptions.LanguageFilename, devMode)
	sts = translate.Sentences

	/* Update gtk conctrols with stored values into mainOptions */
	mainOptions.UpdateObjects()

	/* Statusbar init. */
	statusbar = gi.StatusBarNew(mainObjects.Statusbar, []string{sts["sbInFiles"], sts["sbOutFiles"], sts["sbStatus"]})

	/* init dnd */
	initDropSets() //
}
