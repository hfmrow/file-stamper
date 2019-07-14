// translate.go

// File generated on Sun, 10 Mar 2019 03:38:32 using Gotk3ObjectsTranslate v1.0 2019 H.F.M

/*
* 	This program comes with absolutely no warranty.
*	See the The MIT License (MIT) for details:
*	https://opensource.org/licenses/mit-license.php
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gotk3/gotk3/gtk"
)

// initGtkObjectsText: read translations from structure and set them to objects.
func (trans *MainTranslate) initGtkObjectsText() {
	trans.setTextToGtkObjects(&mainObjects.ButtonApply.Widget, "ButtonApply")
	trans.setTextToGtkObjects(&mainObjects.ButtonExit.Widget, "ButtonExit")
	trans.setTextToGtkObjects(&mainObjects.ButtonInFiles.Widget, "ButtonInFiles")
	trans.setTextToGtkObjects(&mainObjects.ButtonOutFiles.Widget, "ButtonOutFiles")
	trans.setTextToGtkObjects(&mainObjects.ButtonUndo.Widget, "ButtonUndo")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonAtime.Widget, "CheckbuttonAtime")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonCtime.Widget, "CheckbuttonCtime")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonIncludeDir.Widget, "CheckbuttonIncludeDir")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonMtime.Widget, "CheckbuttonMtime")
	trans.setTextToGtkObjects(&mainObjects.CheckbuttonSubdir.Widget, "CheckbuttonSubdir")
	trans.setTextToGtkObjects(&mainObjects.ImageMainTop.Widget, "ImageMainTop")
}

// sentences: some sentences/words used in the application.
var sentences = map[string]string{
	`savef`:           `Save file`,
	`outRst`:          `Output files, reseted`,
	`outLoaded`:       `Output files, stored`,
	`openf`:           `Open file`,
	`fileModified`:    `Files was modified`,
	`sbInFiles`:       `In files:`,
	`inRst`:           `Input files, reseted`,
	`ok`:              `Ok`,
	`sbStatus`:        `Status:`,
	`retry`:           `Retry`,
	`sbOutFiles`:      `Out files:`,
	`allow`:           `Allow`,
	`yes`:             `Yes`,
	`deny`:            `Deny`,
	`inLoaded`:        `Input files, stored`,
	`cancel`:          `Cancel`,
	`errChgTimestamp`: `An error occurred on changing timestamp`,
	`errUndo`:         `An error occurred on undo timestamp`,
	`fileRestored`:    `Files was restored`,
	`no`:              `No`,
	`noToDo`:          `Nothing to do`,
	`noToUndo`:        `Nothing to undo`,
}

// Translations structure with methods
type MainTranslate struct {
	// Public
	ProgInfos    progInfo
	Language     language
	Options      parsingFlags
	ObjectsCount int
	Objects      []object
	Sentences    map[string]string
	// Private
	objectsLoaded bool
}

// MainTranslateNew: Initialise new translation structure and assign language file content to GtkObjects.
// devModeActive, indicate that the new sentences must be added to original language file
func MainTranslateNew(filename string, devModeActive ...bool) (mt *MainTranslate) {
	mt = new(MainTranslate)
	if _, err := os.Stat(filename); err == nil {
		mt.read(filename)
		mt.initGtkObjectsText()
		if len(devModeActive) != 0 {
			if devModeActive[0] {
				mt.Sentences = sentences
				err := mt.write(filename)
				if err != nil {
					fmt.Printf("%s\n%s\n", "Cannot write actual sentences to language file.", err.Error())
				}
			}
		}
	} else {
		fmt.Printf("%s\n%s\n", "Error loading language file !", err.Error())
	}
	return mt
}

// readFile: language file.
func (trans *MainTranslate) read(filename string) (err error) {
	var textFileBytes []byte
	if textFileBytes, err = ioutil.ReadFile(filename); err == nil {
		if err = json.Unmarshal(textFileBytes, &trans); err == nil {
			trans.ProgInfos.GladeXmlFilenameRel, _ = filepath.Rel(filepath.Dir(filename), trans.ProgInfos.GladeXmlFilename)
			trans.objectsLoaded = true
		}
	}
	return err
}

// Write json datas to file
func (trans *MainTranslate) write(filename string) (err error) {
	var out bytes.Buffer
	var jsonData []byte
	if jsonData, err = json.Marshal(&trans); err == nil && trans.objectsLoaded {
		if err = json.Indent(&out, jsonData, "", "\t"); err == nil {
			err = ioutil.WriteFile(filename, out.Bytes(), 0644)
		}
	}
	return err
}

type parsingFlags struct {
	SkipLowerCase  bool
	SkipEmptyLabel bool
	DoBackup       bool
}

type progInfo struct {
	Name                string
	Version             string
	Creat               string
	MainObjStructName   string
	GladeXmlFilename    string
	GladeXmlFilenameRel string
}

type language struct {
	LangNameLong string
	LangNameShrt string
	Author       string
	Date         string
	Updated      string
	Contributors []string
}

type object struct {
	Class   string
	Id      string
	Label   string
	Tooltip string
	Text    string
	Uri     string
	Markup  bool
	Comment string
}

// Define available property within objects
type propObject struct {
	Class   string
	Label   bool
	Tooltip bool
	Markup  bool
	Text    bool
	Uri     bool
}

// Property that exists for GObject ...	(Used for Class)
var propPerObjects = []propObject{
	{Class: "GtkButton", Label: true, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkToggleButton", Label: true, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkLabel", Label: true, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkSpinButton", Label: false, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkEntry", Label: false, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkCheckButton", Label: true, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkProgressBar", Label: false, Tooltip: true, Markup: true, Text: true, Uri: false},
	{Class: "GtkSearchBar", Label: false, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkImage", Label: false, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkRadioButton", Label: true, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkComboBoxText", Label: false, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkComboBox", Label: false, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkLinkButton", Label: true, Tooltip: true, Markup: true, Text: false, Uri: true},
	{Class: "GtkSwitch", Label: false, Tooltip: true, Markup: true, Text: false, Uri: false}}

// setTextToGtkObjects: read translations from structure and set them to object.
// like this: setTextToGtkObjects(&mainObjects.TransLabelHint.Widget, "TransLabelHint")
func (trans *MainTranslate) setTextToGtkObjects(obj *gtk.Widget, objectId string) {
	for _, currObject := range trans.Objects {
		if currObject.Id == objectId {
			for _, props := range propPerObjects {
				if currObject.Class == props.Class {
					if props.Label {
						obj.SetProperty("label", currObject.Label)
					}
					if props.Tooltip && !props.Markup {
						obj.SetProperty("tooltip_text", currObject.Tooltip)
					}
					if props.Tooltip && props.Markup {
						obj.SetProperty("tooltip_markup", currObject.Tooltip)
					}
					if props.Text {
						obj.SetProperty("text", currObject.Text)
					}
					if props.Uri {
						obj.SetProperty("uri", currObject.Uri)
					}
				}
			}
		}
	}
}
