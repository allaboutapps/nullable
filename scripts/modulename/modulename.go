// +build scripts

//  This program prints the current project's module name
// It can be invoked by running go run scripts/modulename/modulename.go
package main

import (
	"fmt"
	"log"

	"allaboutapps.dev/tpa/portal-backend/internal/util"
	"allaboutapps.dev/tpa/portal-backend/scripts"
)

// https://blog.carlmjohnson.net/post/2016-11-27-how-to-use-go-generate/

var (
	PROJECT_ROOT  = util.GetProjectRootDir()
	PATH_MOD_FILE = PROJECT_ROOT + "/go.mod"
)

// get all functions in above handler packages
// that match Get*, Put*, Post*, Patch*, Delete*
func main() {
	baseModuleName, err := scripts.GetModuleName(PATH_MOD_FILE)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(baseModuleName)
}
