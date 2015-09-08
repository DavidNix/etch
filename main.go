package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type config struct {
	appName string
	author  string
	year    string
}

func main() {
	fmt.Println("Enter your app's name: ")
	var appName string
	fmt.Scanln(&appName)

	// fmt.Println("Enter author's name for license: ")
	// var author string
	// fmt.Scanln(&author)

	dir := createDirStructure(appName)
	exitIf(os.Chdir(dir))

	writeStatic("Makefile")
	writeStatic("dev.env")

	fmt.Println(appName + "'s workspace is complete.")
}

func exitIf(err error) {
	if err != nil {
		fmt.Println(err, "Exiting...")
		os.Exit(1)
	}
}

func createDirStructure(name string) string {
	if name == "" {
		exitIf(fmt.Errorf("App name cannot be blank."))
	}
	sep := string(filepath.Separator)
	path := "." + sep + name + sep + "src" + sep + name
	exitIf(os.MkdirAll(path, 0777))
	return path
}

func template(name string) []byte {
	templ, err := Asset("templates/" + name)
	exitIf(err)
	return templ
}

func writeStatic(name string) {
	exitIf(ioutil.WriteFile(name, template(name), 0644))
}
