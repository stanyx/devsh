package cmd

import "bufio"

type Cmd struct {
	Name        string
	Description string
	F           func(*bufio.Reader, ...string)
}
