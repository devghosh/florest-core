package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

var projectPath = flag.String("projectPath", "", "path to project directory e.g /home/user/myapp")

const (
	florest_core  = "github.com/jabong/florest-core"
	newapp_script = "scripts/newapp.sh"
)

func main() {
	flag.Parse()
	if *projectPath == "" {
		flag.PrintDefaults()
		return
	}

	// check if gopath is set
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		fmt.Println("Please, set $GOPATH environment variable\n")
		return
	}

	//Support gopaths with multiple directories
	dirs := strings.Split(gopath, ":")
	var fullPath string
	var err error
	found := false
	for _, d := range dirs {
		fullPath = path.Join(d, "src", florest_core)
		if _, err = os.Stat(fullPath); err == nil {
			found = true
			break
		}
	}
	// return if florest was not found
	if found == false {
		fmt.Println(florest_core + " not found")
		return
	}
	// florest-core is present, make the new app now
	stdout, err := exec.Command("/bin/sh", path.Join(fullPath, newapp_script), fullPath, *projectPath).CombinedOutput()
	// print the output, will print the error also
	fmt.Println(string(stdout))
}
