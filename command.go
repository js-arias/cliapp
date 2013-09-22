// Copyright (c) 2013, J. Salvador Arias <jsalarias@csnat.unt.edu.ar>
// All rights reserved.
// Distributed under BSD-style license that can be found in the LICENSE file.

package cliapp

import (
	"fmt"
	"os"
	"strings"
)

// A command is a hosted command
type Command struct {
	// Run runs the command.
	// The argument list is the set of unparsed arguments, that is the
	// arguments unparsed by the flag package.
	Run func(c *Command, args []string)

	// Name is the command's name.
	Name string

	// Short is a short, single line description of the command.
	Short string

	// Long is a long description of the command.
	Long string
}

// ErrStr returns an error description from the command
func (c *Command) ErrStr(err interface{}) {
	fmt.Fprintf(os.Stdout, "%s: error: %v\n", c.Name, err)
}

func (c *Command) Help() {
	fmt.Fprintf(os.Stdout, "%s - %s\n%s\n", c.Name, c.Short, strings.TrimSpace(c.Long))
}

func (a *App) help(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stdout, "List of commands:\n")
		for _, c := range a.Command {
			fmt.Fprintf(os.Stdout, "    %-11s %s\n", c.Name, c.Short)
		}
		fmt.Fprintf(os.Stdout, "Type 'help <command>' for more information about a command.\n")
		return
	}
	if len(args) > 1 {
		fmt.Fprintf(os.Stdout, "help: too many arguments\n")
		return
	}
	cmd, ok := a.cmdAlias[args[0]]
	if !ok {
		fmt.Fprintf(os.Stdout, "unknown argument\n")
	}
	cmd.Help()
}
