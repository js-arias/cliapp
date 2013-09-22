// Copyright (c) 2013, J. Salvador Arias <jsalarias@csnat.unt.edu.ar>
// All rights reserved.
// Distributed under BSD-style license that can be found in the LICENSE file.

// Package cliapp implements a command line (CLI) interface application.
package cliapp

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// App is a CLI application
type App struct {
	// Subject is a list of subjects (and therefore, of commands)
	Command []*Command

	// prompt is the used prompt
	prompt string

	// command aliases
	cmdAlias map[string]*Command
}

// New returns a new app
func New(cmds []*Command) *App {
	return &App{
		Command:  cmds,
		cmdAlias: make(map[string]*Command),
	}
}

// Run runs the command loop.
func (a *App) Run() {
	for _, c := range a.Command {
		a.AddAlias(c.Name, c)
	}
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprintf(os.Stdout, "[%s]$ ", a.prompt)
		line, err := in.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		args := strings.Fields(line)
		if (len(args) == 0) || (args[0][0] == '#') || (args[0][0] == ';') {
			continue
		}
		if (args[0] == "help") || (args[0] == "?") {
			a.help(args[1:])
			continue
		}
		c, ok := a.cmdAlias[args[0]]
		if !ok {
			fmt.Fprintf(os.Stdout, "%s: unknown command\n", args[0])
			continue
		}
		c.Run(c, args[1:])
	}
}

// Finish the application
func (a *App) Exit() {
	os.Exit(0)
}

// Set Prompt
func (a *App) SetPrompt(p string) {
	a.prompt = p
}

// add alias
func (a *App) AddAlias(alias string, cmd *Command) {
	if _, ok := a.cmdAlias[alias]; ok {
		panic("alias " + alias + " already defined")
	}
	a.cmdAlias[alias] = cmd
}
