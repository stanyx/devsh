package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"devsh/cmd"
)

var commands []cmd.Cmd

func InitCommands() []cmd.Cmd {
	return []cmd.Cmd{
		{Name: "sci", Description: "provide execution of mathematic expression by Python", F: cmd.PyExec},
		{Name: "fs", Description: "subshell for file system commands", F: cmd.FSCmd},
		{Name: "create", Description: "create blank project with Git support", F: cmd.CreateProject},
		{Name: "clone", Description: "clone project with Git support"},
		{Name: "select", Description: "activate project", F: cmd.SelectProject},
		{Name: "size", Description: "show current project size"},
		{Name: "snippet", Description: "blank"},
		{Name: "saveall", Description: "blank", F: cmd.SaveAll},
		{Name: "archive", Description: "blank"},
		{Name: "undo", Description: "blank"},
		{Name: "backup", Description: "blank"},
		{Name: "dbstart", Description: "blank"},
		{Name: "dbcheck", Description: "blank"},
		{Name: "sql", Description: "blank"},
		{Name: "get", Description: "send GET http request to provided url", F: cmd.SendGETRequest},
		{Name: "post", Description: "send POST http request to provided url"},
		{Name: "deploy", Description: "blank"},
		{Name: "cachestart", Description: "blank"},
		{Name: "help", Description: "show list of available commands", F: showHelp},
	}
}

func showHelp(r *bufio.Reader, args ...string) {
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
	commandMap := make(map[string]cmd.Cmd)
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
			if cmd.F != nil {
				cmd.F(reader)
			}
		} else {
			fmt.Println("unknown command")
		}

		fmt.Println()
		fmt.Printf(">>")
	}

}