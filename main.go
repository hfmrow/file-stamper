// main.go

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
	"fmt"

	gimc "github.com/hfmrow/gtk3Import/misc"
)

func main() {
	/* Be or not to be ... in dev mode ... */
	devMode = true

	/* Build directory for tempDir */
	doTempDir = false

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

	/* Init AboutBox */
	mainOptions.About.InitFillInfos(
		mainObjects.MainWindow,
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

	/* Translate init. */
	translate = MainTranslateNew(absoluteRealPath+mainOptions.LanguageFilename, devMode)
	sts = translate.Sentences

	/* Update gtk conctrols with stored values into mainOptions */
	mainOptions.UpdateObjects()

	/* Statusbar init. */
	statusbar = gimc.StatusBarStructureNew(mainObjects.Statusbar, []string{sts["sbInFiles"], sts["sbOutFiles"], sts["sbStatus"]})

	/* init dnd */
	initDropSets() //
}
