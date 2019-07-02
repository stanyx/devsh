package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

var currentProject string

func CreateProject(r *bufio.Reader, args ...string) {
	fmt.Println("enter project name: ")
	line, _ := r.ReadString('\n')
	projectName := strings.Trim(line, "\r\n")
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	projectPath := path.Join(currentDir, projectName)
	err = os.MkdirAll(projectPath, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err = exec.Command("git", "-C", projectPath, "init").Output()
	if err != nil {
		log.Fatal(err)
	}

	// create .gitignore
	if f, err := os.Create(path.Join(projectPath, ".gitignore")); err != nil {
		log.Fatal(err)
	} else {
		f.Close()
		fmt.Println("create .gitignore - ok")
	}

	// create README.md
	if f, err := os.Create(path.Join(projectPath, "README.md")); err != nil {
		log.Fatal(err)
	} else {
		f.Close()
		fmt.Println("create README.md - ok")
	}

	colorTpl := "%s"
	fmt.Print(fmt.Sprintf(colorTpl, "\ncreate empty project"), projectPath)
	currentProject = projectPath
}

func SelectProject(r *bufio.Reader, args ...string) {
	fmt.Println("enter project name: ")
	line, _ := r.ReadString('\n')
	projectName := strings.Trim(line, "\r\n")
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	currentProject = path.Join(currentDir, projectName)
}

func SaveAll(r *bufio.Reader, args ...string) {
	if currentProject == "" {
		fmt.Println("project not selected")
		return
	}
	out, err := exec.Command("git", "-C", currentProject, "add", "-A").Output()
	if err != nil {
		fmt.Println("add error: ", string(out))
		log.Fatal(err)
	}
	out, err = exec.Command("git", "-C", currentProject, "commit", "-am", "\".\"").Output()
	if err != nil {
		fmt.Println("commit error: ", string(out))
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
