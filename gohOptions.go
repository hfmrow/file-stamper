// gohOptions.go

// Source file auto-generated on Sun, 14 Jul 2019 13:07:22 using Gotk3ObjHandler v1.3 ©2019 H.F.M

/*
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os/exec"

	gi "github.com/hfmrow/fileStamper/gtk3Import"
)

// App infos
var Name = "fileStamper"
var Vers = "v1.0"
var Descr = "This software allow you to clone timestamp between files"
var Creat = "H.F.M"
var YearCreat = "2019"
var LicenseShort = "This program comes with absolutely no warranty.\nSee the The MIT License (MIT) for details:\nhttps://opensource.org/licenses/mit-license.php"
var LicenseAbrv = "License (MIT)"
var Repository = "github.com/fileStamper"

// Vars declarations
var absoluteRealPath, optFilename = getAbsRealPath()
var mainOptions *MainOpt
var err error
var tempDir string
var currentOutFilesList []string
var currentInFilesList []string

// Translations
var devMode bool
var translate = new(MainTranslate)
var sts map[string]string

var nl = "\n"
var execCmd *exec.Cmd
var statusbar = new(gi.StatusBar)

type MainOpt struct {
	/* Public, will be saved and restored */
	About         *gi.AboutInfos
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
	opt.About = new(gi.AboutInfos)

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
	if jsonData, err = json.Marshal(&opt); err == nil {
		if err = json.Indent(&out, jsonData, "", "\t"); err == nil {
			err = ioutil.WriteFile(optFilename, out.Bytes(), 0644)
		}
	}
	return err
}
