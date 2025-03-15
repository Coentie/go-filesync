package main

import (
	"fmt"
	"github.com/coentie/filesync/packages/paths"
	"github.com/coentie/filesync/packages/upload"
	"os"
)

var pathManagement = "manage"
var sync = "sync"

func main() {

	paths.Bootstrap()

	if os.Args == nil || len(os.Args) < 2 {
		fmt.Println("missing arguments. Options:")
		fmt.Println("sync - syncing files")
		fmt.Println("manage - Manage paths for syncing.")
		return
	}

	if os.Args[1] == pathManagement {
		paths.Manage()
		return
	}

	if os.Args[1] == sync {
		upload.Upload()
	}
}
