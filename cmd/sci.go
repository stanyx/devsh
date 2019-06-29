package cmd

import (
	"fmt"
	"log"
	"bufio"
	"sync"
	"strings"
	"strconv"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"runtime"
)


func searchPyBin() string {
	var candidates []string
	if runtime.GOOS == "windows" {
		candidates = append(candidates, "/AppData/Local/Programs/Python")
	} else {
		//TODO using pyenv
	}
	pyPaths := []string{}
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	userPath := usr.HomeDir
	for _, p := range candidates {
		rootPath := filepath.Join(userPath, p)
		filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() && filepath.Dir(path) == rootPath {
				pyPaths = append(pyPaths, path)
			}
			return nil
		})
	}
	versionsWithPaths := make(map[int]string)
	var maxVersion int
	for _, p := range pyPaths {
		basePath := filepath.Base(p)
		strVersion := strings.TrimPrefix(strings.Replace(basePath, "-", "", -1), "Python")
		version, _:= strconv.Atoi(strVersion)
		if version > 0 {
			versionsWithPaths[version] = p
			if version > maxVersion {
				maxVersion = version
			} 
		}
	}
	return versionsWithPaths[maxVersion]
}

func installDeps() {

}

var pyBinSearch sync.Once
var pyPath string

//PyExec execute Python expression and print results
func PyExec(r *bufio.Reader) {

	pyBinSearch.Do(func() {
		pyPath = path.Join(searchPyBin(), "python.exe")
	})

	for {

		line, _ := r.ReadString('\n')
		expression := strings.Trim(line, "\r\n")
		if expression == "" {
			continue
		}

		strCommand := fmt.Sprintf("from math import *; from itertools import *; from functools import *; print(%s)", expression)
		cm := exec.Command(pyPath, "-c", strCommand)
		out, err := cm.Output()
		fmt.Println()
		if err != nil {
			fmt.Println(out)
			fmt.Println(err)
		}
		fmt.Println(string(out))
	}
}
