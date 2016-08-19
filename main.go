package main

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type config struct {
	AppName string
	Author  string
	Year    int
}

func main() {
	checkDependencies()
	conf := setup()

	if conf.AppName == "" {
		exitIf(errors.New("App name cannot be blank."))
	}

	checkForExistingDir(conf)

	exitIf(os.Mkdir(conf.AppName, 0700))
	exitIf(os.Chdir(conf.AppName))

	exitIf(os.Mkdir("vendor", 0700))

	initGit()

	writeStatic(".git/hooks/pre_commit", 0766)
	writeStatic("Makefile", 0666)
	writeStatic("dev.env", 0666)
	writeStatic(".gitignore", 0666)
	writeTemplate("LICENSE", conf, 0666)
	writeTemplate("tmux", conf, 0766)
	writeTemplate("modd.conf", conf, 0666)

	fmt.Println(conf.AppName + "'s workspace is complete!")
}

func exitIf(err error) {
	if err != nil {
		fmt.Println(err, "\nExiting...")
		os.Exit(1)
	}
}

func checkForExistingDir(c config) {
	if _, err := os.Stat(c.AppName); os.IsExist(err) {
		exitIf(errors.New("Directory \"%s\" already exists"))
	}
}

func checkDependencies() {
	out, err := exec.Command("which", "git").Output()
	if len(out) == 0 || err != nil {
		exitIf(fmt.Errorf("Git not found, is it installed?"))
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

func memoTemplate(name string) []byte {
	templ, err := Asset("templates/" + name)
	exitIf(err)
	return templ
}

func writeStatic(name string, mode os.FileMode) {
	exitIf(ioutil.WriteFile(name, memoTemplate(filepath.Base(name)), mode))
}

func writeTemplate(name string, c config, mode os.FileMode) {
	f, e := os.Create(name)
	exitIf(e)
	exitIf(f.Chmod(mode))
	defer f.Close()
	tmpl, err := template.New(name).Parse(string(memoTemplate(name)))
	exitIf(err)
	exitIf(tmpl.Execute(f, c))
}

func initGit() {
	_, err := exec.Command("git", "init").Output()
	exitIf(err)
}
