// fileStamperFunc.go

// Source file auto-generated on Sat, 09 Mar 2019 03:18:51 using Gotk3ObjHandler v1.0 ©2019 H.F.M

/*
	fileStamper v1.0 ©2019 H.F.M

	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type timeList []fileTimeInfos

type fileTimeInfos struct {
	Filename string
	Atime    time.Time
	Mtime    time.Time
	Ctime    time.Time
	Btime    time.Time // actually unused
}

// updateSb: update statusbar informations
func updateSb(state ...string) {
	statusbar.CleanAll()
	if len(state) != 0 {
		statusbar.Set(fmt.Sprintf("%s", state[0]), 2)
	}
	statusbar.Set(fmt.Sprintf("%d", len(currentInFilesList)), 0)
	statusbar.Set(fmt.Sprintf("%d", len(currentOutFilesList)), 1)
}

// chgTimestamp: 2 methodes, case form equal in and out filenames and another when not equal length
func chgTimestamp() (err error) {
	var chgCount, stampCounter int
	var startSubCounter, changed bool
	var errTxt string
	var lenIn = len(mainOptions.PrevInFilesTs)
	var lenOut = len(mainOptions.PrevOutFilesTs)
	if lenIn == lenOut {
		for idx := 0; idx < lenIn; idx++ {
			changed, err = doTimestampChg(mainOptions.PrevInFilesTs[idx], mainOptions.PrevOutFilesTs[idx])
			if err != nil {
				errTxt += err.Error() + "\n"

			}
			if changed {
				chgCount++
				changed = false
			}
		}
	} else {
		for idx := 0; idx < lenIn; idx++ {
			changed, err = doTimestampChg(mainOptions.PrevInFilesTs[idx], mainOptions.PrevOutFilesTs[idx+stampCounter])
			if startSubCounter {
				stampCounter++
			}
			if err != nil {
				errTxt += err.Error() + "\n"

			}
			if changed {
				chgCount++
				changed = false
			}
			if idx == lenOut-1 || idx+stampCounter == lenOut {
				break
			}
			if idx == lenIn-1 {
				idx = -1
				startSubCounter = true
			}
		}
	}
	if len(errTxt) != 0 {
		err = errors.New(errTxt)
	}
	updateSb(fmt.Sprintf("%d "+sts["fileModified"], chgCount))

	return err
}

// doTimestampChg: perform timestamp change in file dest.
func doTimestampChg(item1, item2 fileTimeInfos) (changed bool, err error) {
	var fi os.FileInfo
	var atime, mtime time.Time
	if fi, err = os.Stat(item2.Filename); err == nil && !(fi.IsDir() && !mainObjects.CheckbuttonIncludeDir.GetActive()) {
		//	if !(fi.IsDir() && !mainObjects.CheckbuttonIncludeDir.GetActive()) {
		atime = item2.Atime
		mtime = item2.Mtime
		if mainObjects.CheckbuttonAtime.GetActive() {
			atime = item1.Atime
		}
		if mainObjects.CheckbuttonMtime.GetActive() {
			mtime = item1.Mtime
		}
		return true, os.Chtimes(item2.Filename, atime, mtime)
	}
	//	}
	return false, err
}

// storeAllTimestamp:
func StoreInTimestamp() (err error) {
	if mainObjects.CheckbuttonSubdir.GetActive() {
		if currentInFilesList, err = getSubdir(currentInFilesList); err != nil {
			return err
		}
	}
	mainOptions.PrevInFilesTs, err = storeTimes(currentInFilesList)
	return err
}

func StoreOutTimestamp() (err error) {
	if mainObjects.CheckbuttonSubdir.GetActive() {
		if currentOutFilesList, err = getSubdir(currentOutFilesList); err != nil {
			return err
		}
	}
	mainOptions.PrevOutFilesTs, err = storeTimes(currentOutFilesList)
	return err
}

// getSubdir: retreive files into subdirs and return whole files
func getSubdir(inFiles []string) (outFiles []string, err error) {
	var files []string
	for _, file := range inFiles {
		if fi, err := os.Stat(file); fi.IsDir() && !os.IsNotExist(err) {
			if files, err = ScanSubDir(file, mainObjects.CheckbuttonIncludeDir.GetActive()); err == nil {
				outFiles = append(outFiles, files...)
			}
		}
	}
	return append(inFiles, outFiles...), err
}

// ScanSubDir retrieve files in a specific directory and his sub-directory.
// don't follow symlink (walk)
func ScanSubDir(root string, showDirs ...bool) (files []string, err error) {
	var listDir bool
	if len(showDirs) != 0 {
		listDir = showDirs[0]
	}
	err = filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				files = append(files, path)
			} else if listDir {
				files = append(files, path)
			}
			return nil
		})
	return files, err
}

// storeTimes: store time information into structure
func storeTimes(inList []string) (fileTimeList timeList, err error) {
	for _, file := range inList {
		atime, mtime, ctime, err := getTimes(file)
		if err == nil {
			fileTimeList = append(fileTimeList, fileTimeInfos{Filename: file, Atime: atime, Mtime: mtime, Ctime: ctime})
		}
	}
	return fileTimeList, err
}

// statTimes: get times from file
func getTimes(name string) (atime, mtime, ctime time.Time, err error) {
	var fi os.FileInfo
	fi, err = os.Stat(name)
	if err == nil {
		mtime = fi.ModTime()
		stat := fi.Sys().(*syscall.Stat_t)
		atime = time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
		ctime = time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	}
	return atime, mtime, ctime, err
}

// initDropSets: configure controls to receive dndcontent.
func initDropSets() {
	var targets []gtk.TargetEntry // Build dnd context
	te, err := gtk.TargetEntryNew("text/uri-list", gtk.TARGET_OTHER_APP, 0)
	if err != nil {
		log.Fatal(err)
	}
	targets = append(targets, *te)

	mainObjects.ButtonInFiles.DragDestSet(
		gtk.DEST_DEFAULT_ALL,
		targets,
		gdk.ACTION_COPY)

	mainObjects.ButtonOutFiles.DragDestSet(
		gtk.DEST_DEFAULT_ALL,
		targets,
		gdk.ACTION_COPY)
}

// GetTextEOL: Get EOL from text bytes (CR, LF, CRLF) > string
func GetTextEOL(inTextBytes []byte) (outString string) {
	bCR := []byte{0x0D}
	bLF := []byte{0x0A}
	bCRLF := []byte{0x0D, 0x0A}
	if bytes.Contains(inTextBytes, bCRLF) {
		return string(bCRLF)
	} else if bytes.Contains(inTextBytes, bCR) {
		return string(bCR)
	}
	return string(bLF)
}
