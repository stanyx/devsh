package main

import (
	"os"
	"log"
	"fmt"
	"path"
	"bufio"
	"strings"
	"os/exec"
)

type Cmd struct {
	Name string
	Description string
	f func(*bufio.Reader)
}

var commands []Cmd

func InitCommands() []Cmd {
	return []Cmd{
		{Name: "create", Description: "create blank project with Git support", f: createProject},
		{Name: "clone", Description: "clone project with Git support"},
		{Name: "select", Description: "activate project", f: selectProject},
		{Name: "size", Description: "show current project size"},
		{Name: "snippet", Description: "blank"},
		{Name: "saveall", Description: "blank", f: saveAll},
		{Name: "archive", Description: "blank"},
		{Name: "undo", Description: "blank"},
		{Name: "backup", Description: "blank"},
		{Name: "dbstart", Description: "blank"},
		{Name: "dbcheck", Description: "blank"},
		{Name: "sql", Description: "blank"},
		{Name: "uget", Description: "send GET http request to provided url"},
		{Name: "upost", Description: "send POST http request to provided url"},
		{Name: "calc", Description: "blank"},
		{Name: "deploy", Description: "blank"},
		{Name: "cachestart", Description: "blank"},
		{Name: "help", Description: "show list of available commands", f: showHelp},
	}
}

var currentProject string

func createProject(r *bufio.Reader) {
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
	colorTpl := "%s"
	fmt.Print(fmt.Sprintf(colorTpl, "\ncreate empty project"), projectPath)
	currentProject = projectPath
}

func selectProject(r *bufio.Reader) {
	fmt.Println("enter project name: ")
	line, _ := r.ReadString('\n')
	projectName := strings.Trim(line, "\r\n")
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	currentProject = path.Join(currentDir, projectName)
}

func saveAll(r *bufio.Reader) {
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

func showHelp(r *bufio.Reader) {
	fmt.Println("\n[ List or commands ]\n")
	for _, cmd := range commands {
		fmt.Printf("* %s - %s\n", cmd.Name, cmd.Description)
	}
}

func main() {

	welcomeMessage := fmt.Sprintf(`
############# DevSH v%s ##############
#                                       #
#     Powerful All-In-One solution      #
#    for rapid software development     #
#                                       #
##### Enter help to getting started #####

>>`, "0.0.1")

	fmt.Printf(welcomeMessage)
	commands = InitCommands()
	commandMap := make(map[string]Cmd)
	for _, cmd := range commands {
		commandMap[cmd.Name] = cmd
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error", err)
		}
		input := strings.Trim(line, "\r\n")
		if cmd, ok := commandMap[input]; ok {
			if cmd.f != nil {
				cmd.f(reader)
			}
		} else {
			fmt.Println("unknown command")
		}

		fmt.Println()
		fmt.Printf(">>")
	}

}