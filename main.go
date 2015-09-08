package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Enter your app's name: ")
	var appName string
	fmt.Scanln(&appName)

	exitIf(createDirStructure(appName))

	// a := MustAsset("templates/Makefile")
	// fmt.Print(string(a))
}

func exitIf(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createDirStructure(app string) error {
	if app == "" {
		return fmt.Errorf("App name cannot be blank.")
	}
	sep := string(filepath.Separator)
	return os.MkdirAll("."+sep+app+sep+"src"+sep+app, 0777)
}
