// fileStamperHandler.go

// Source file auto-generated on Sat, 09 Mar 2019 03:18:51 using Gotk3ObjHandler v1.0 ©2019 H.F.M

/*
	fileStamper v1.0 ©2019 H.F.M

	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	gidg "github.com/hfmrow/gtk3Import/dialog"
)

// ButtonExitClicked:
func ButtonExitClicked() {
	windowDestroy()
}

// ButtonApplyClicked:
func ButtonApplyClicked() {
	if len(mainOptions.PrevOutFilesTs) != 0 {
		if err := chgTimestamp(); err != nil {
			gidg.DialogMessage(mainObjects.MainWindow, "error", sts["errChgTimestamp"], err.Error(), "", sts["ok"])
		}
	} else {
		updateSb(fmt.Sprintf("%s", sts["noToDo"]))
	}
}

// ButtonUndoClicked:
func ButtonUndoClicked() {
	if len(mainOptions.PrevOutFilesTs) != 0 {
		var changed bool
		var chgCount int
		var errTxt string
		var lenOut = len(mainOptions.PrevOutFilesTs)
		if lenOut != 0 {
			for idx := 0; idx < lenOut; idx++ {
				changed, err = doTimestampChg(mainOptions.PrevOutFilesTs[idx], mainOptions.PrevOutFilesTs[idx])
				if err != nil {
					errTxt += err.Error() + "\n"
				}
				if changed {
					chgCount++
				}
			}
			if len(errTxt) != 0 {
				err = errors.New(errTxt)
			}
		}
		if err != nil {
			gidg.DialogMessage(mainObjects.MainWindow, "error", sts["errUndo"], err.Error(), "", sts["ok"])
		}
		updateSb(fmt.Sprintf("%d "+sts["fileRestored"], chgCount))
	} else {
		updateSb(fmt.Sprintf("%s", sts["noToUndo"]))
	}
}

// ButtonOutFilesClicked:
func ButtonOutFilesClicked() {
	mainOptions.PrevOutFilesTs = mainOptions.PrevOutFilesTs[:0]
	currentOutFilesList = []string{}
	updateSb(sts["outRst"])
}

// ButtonInFilesClicked:
func ButtonInFilesClicked() {
	mainOptions.PrevInFilesTs = mainOptions.PrevInFilesTs[:0]
	currentInFilesList = []string{}
	updateSb(sts["inRst"])
}

// ButtonInFilesReceived: Store in files list
func ButtonInFilesReceived(bti *gtk.Button, context *gdk.DragContext, x, y int, selData *gtk.SelectionData, info, time uint) {
	currentInFilesList = []string{}
	// Convert pointer to datas
	data := selData.GetData()
	list := strings.Split(string(data), GetTextEOL(data))
	for _, file := range list {
		if len(file) != 0 {
			if u, err := url.PathUnescape(file); err == nil {
				currentInFilesList = append(currentInFilesList, strings.TrimPrefix(u, "file://"))
			}
		}
	}
	StoreInTimestamp()
	updateSb(sts["inLoaded"])
}

// ButtonOutFilesReceived: Store out files list
func ButtonOutFilesReceived(bto *gtk.Button, context *gdk.DragContext, x, y int, selData *gtk.SelectionData, info, time uint) {
	currentOutFilesList = []string{}
	// Convert pointer to datas
	data := selData.GetData()
	list := strings.Split(string(data), GetTextEOL(data))
	for _, file := range list {
		if len(file) != 0 {
			if u, err := url.PathUnescape(file); err == nil {
				currentOutFilesList = append(currentOutFilesList, strings.TrimPrefix(u, "file://"))

			}
		}
	}
	StoreOutTimestamp()
	updateSb(sts["outLoaded"])
}

// EvenboxClicked DialogBox
func EvenboxClicked() {
	mainOptions.About.Show()
}
