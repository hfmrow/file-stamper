// gohOptions.go

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
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os/exec"

	gidg "github.com/hfmrow/gtk3Import/dialog"
	gimc "github.com/hfmrow/gtk3Import/misc"
)

// App infos
var Name = "file-stamper"
var Vers = "v1.1"
var Descr = "This software allow you to clone timestamp between files"
var Creat = "hfmrow"
var YearCreat = "2019-21"
var LicenseShort = "This program comes with absolutely no warranty.\nSee the The MIT License (MIT) for details:\nhttps://opensource.org/licenses/mit-license.php"
var LicenseAbrv = "License (MIT)"
var Repository = "github.com/hfmrow/file-stamper"

// Vars declarations
var absoluteRealPath, optFilename = getAbsRealPath()
var mainOptions *MainOpt
var err error
var tempDir string
var doTempDir bool
var currentOutFilesList []string
var currentInFilesList []string

// Translations
var devMode bool

var nl = "\n"
var execCmd *exec.Cmd
var statusbar = new(gimc.StatusBar)

type MainOpt struct {
	/* Public, will be saved and restored */
	About         *gidg.AboutInfos
	MainWinWidth  int
	MainWinHeight int

	LanguageFilename string
	AtimeMod         bool
	MtimeMod         bool
	CtimeMod         bool

	SubDir         bool
	Directory      bool
	PrevInFilesTs  timeList
	PrevOutFilesTs timeList

	/* Private, will NOT be saved */
}

// Main options initialisation
func (opt *MainOpt) Init() {
	/* Aboutbox initialisation */
	opt.About = new(gidg.AboutInfos)

	opt.LanguageFilename = "assets/lang/eng.lang"

	opt.MainWinWidth = 320
	opt.MainWinHeight = 240
}

// Variables -> Objects.
func (opt *MainOpt) UpdateObjects() {
	mainObjects.MainWindow.Resize(opt.MainWinWidth, opt.MainWinHeight)

	mainObjects.CheckbuttonCtime.SetSensitive(false)

	mainObjects.CheckbuttonAtime.SetActive(opt.AtimeMod)
	mainObjects.CheckbuttonMtime.SetActive(opt.MtimeMod)
	mainObjects.CheckbuttonCtime.SetActive(opt.CtimeMod)
	mainObjects.CheckbuttonSubdir.SetActive(opt.SubDir)
	mainObjects.CheckbuttonIncludeDir.SetActive(opt.Directory)
}

// Objects -> Variables.
func (opt *MainOpt) UpdateOptions() {
	opt.MainWinWidth, opt.MainWinHeight = mainObjects.MainWindow.GetSize()

	opt.AtimeMod = mainObjects.CheckbuttonAtime.GetActive()
	opt.MtimeMod = mainObjects.CheckbuttonMtime.GetActive()
	opt.CtimeMod = mainObjects.CheckbuttonCtime.GetActive()
	opt.SubDir = mainObjects.CheckbuttonSubdir.GetActive()
	opt.Directory = mainObjects.CheckbuttonIncludeDir.GetActive()
}

// Read Options from file
func (opt *MainOpt) Read() (err error) {
	var textFileBytes []byte
	if textFileBytes, err = ioutil.ReadFile(optFilename); err == nil {
		err = json.Unmarshal(textFileBytes, &opt)
	}
	return err
}

// Write Options to file
func (opt *MainOpt) Write() (err error) {
	var jsonData []byte
	var out bytes.Buffer
	opt.UpdateOptions()

	opt.About.DlgBoxStruct = nil // remove dialog object before saving

	if jsonData, err = json.Marshal(&opt); err == nil {
		if err = json.Indent(&out, jsonData, "", "\t"); err == nil {
			err = ioutil.WriteFile(optFilename, out.Bytes(), 0644)
		}
	}
	return err
}
