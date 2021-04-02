// gohAssets.go

/*
	Source file auto-generated on Fri, 02 Apr 2021 15:44:30 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE
*/

package main

import (
	"embed"
	"log"
)

//go:embed assets/glade
//go:embed assets/images
var embeddedFiles embed.FS

// This functionality does not require explicit encoding of the files, at each
// compilation, the files are inserted into the resulting binary. Thus, updating
// assets is only required when new files are added to be embedded in order to
// create and declare the variables to which the files are linked.
// assetsDeclarationsUseEmbedded: Use native Go 'embed' package to include files
// content at runtime.
func assetsDeclarationsUseEmbedded(embedded ...bool) {
	mainGlade = readEmbedFile("assets/glade/main.glade")
	apply32 = readEmbedFile("assets/images/Apply-32.png")
	counterclockwiseArrow20 = readEmbedFile("assets/images/Counterclockwise-arrow-20.png")
	filestamper300x29 = readEmbedFile("assets/images/FileStamper-300x29.png")
	filestamper500x48 = readEmbedFile("assets/images/FileStamper-500x48.png")
	folderTime48 = readEmbedFile("assets/images/Folder-time-48.png")
	folderTimeCalIn48 = readEmbedFile("assets/images/Folder-time-cal-in-48.png")
	folderTimeCalOut48 = readEmbedFile("assets/images/Folder-time-cal-out-48.png")
	logOut20 = readEmbedFile("assets/images/Log-Out-20.png")
	signCancel20 = readEmbedFile("assets/images/Sign-cancel-20.png")
	signSelect20 = readEmbedFile("assets/images/Sign-Select-20.png")
}

// readEmbedFile: read 'embed' file system and return []byte data.
func readEmbedFile(filename string) (out []byte) {
	var err error
	out, err = embeddedFiles.ReadFile(filename)
	if err != nil {
		log.Printf("Unable to read embedded file: %s, %v\n", filename, err)
	}
	return
}
