package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func makeDirs(r *bufio.Reader, args ...string) {
	if len(args) == 0 {
		fmt.Println("error: directories list not provided")
		return
	}
	for _, p := range args {
		err := os.MkdirAll(path.Join(currentWorkspace, p), os.ModeDir)
		if err != nil {
			fmt.Println("error on make directory", err)
		}
	}
}

var currentWorkspace string = ""

func setWorkspace(r *bufio.Reader, args ...string) {
	if len(args) == 2 {
		attr := args[0]
		workspace := args[1]
		if attr == "-c" {
			makeDirs(r, workspace)
			currentWorkspace = workspace
		}
	} else if len(args) == 1 {
		currentWorkspace = args[0]
	}
	fmt.Printf("workspace set to %s\n", currentWorkspace)
}

var fsCommands = []Cmd{
	{Name: "w", Description: "switch to workspace", F: setWorkspace},
	{Name: "mkd", Description: "make new directory or directories", F: makeDirs},
	// {Name: "mkf", Description: "make new file", f: makeFile},
	// {Name: "sf", Description: "show file info", f: showFileInfo},
	// {Name: "sd", Description: "show dir info", f: showDirInfo},
}

func FSCmd(r *bufio.Reader, args ...string) {

	commands := make(map[string]Cmd)
	for _, cmd := range fsCommands {
		commands[cmd.Name] = cmd
	}
	for {

		fmt.Printf("[fs]>")
		line, _ := r.ReadString('\n')
		rcmd := strings.Trim(line, "\r\n")
		if rcmd == "" {
			fmt.Println("exit fs")
			break
		}

		attrs := strings.Split(rcmd, " ")
		if cmd, ok := commands[attrs[0]]; ok {
			cmd.F(r, attrs[1:]...)
		}
	}
}
