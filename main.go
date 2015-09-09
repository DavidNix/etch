package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type config struct {
	AppName string
	Author  string
	Year    int
}

func main() {
	conf := setup()

	dir := createDirStructure(conf.AppName)
	exitIf(os.Chdir(dir))

	writeStatic("Makefile")
	writeStatic("dev.env")
	writeTemplate("LICENSE", conf)
	writeTemplate("tmux", conf)

	fmt.Println(conf.AppName + "'s workspace is complete!")
}

func exitIf(err error) {
	if err != nil {
		fmt.Println(err, "Exiting...")
		os.Exit(1)
	}
}

func setup() config {
	conf := config{}

	in := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your app's name: ")
	line, err := in.ReadString('\n')
	exitIf(err)
	conf.AppName = strings.TrimSpace(line)

	in.Reset(os.Stdin)

	fmt.Println("Enter author's name for license: ")
	line, err = in.ReadString('\n')
	exitIf(err)
	conf.Author = strings.TrimSpace(line)

	conf.Year = time.Now().Year()

	return conf
}

func createDirStructure(name string) string {
	if name == "" {
		exitIf(fmt.Errorf("App name cannot be blank."))
	}
	sep := string(os.PathSeparator)
	path := "." + sep + name + sep + "src" + sep + name
	exitIf(os.MkdirAll(path, 0777))
	return path
}

func memoTemplate(name string) []byte {
	templ, err := Asset("templates/" + name)
	exitIf(err)
	return templ
}

func writeStatic(name string) {
	exitIf(ioutil.WriteFile(name, memoTemplate(name), 0644))
}

func writeTemplate(name string, c config) {
	f, e := os.Create(name)
	exitIf(e)
	defer f.Close()
	tmpl, err := template.New(name).Parse(string(memoTemplate(name)))
	exitIf(err)
	exitIf(tmpl.Execute(f, c))
}
