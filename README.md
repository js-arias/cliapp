cliapp
======

Package cliapp implements a command line (CLI) interface application.

Quick usage
-----------

    go get github.com/js-arias/cliapp

First, the list of commands must be set, then a new application will
be created with that list of commands. This guarantee that the list of
commands will be shown in the set order.

Appart of their name, commands can have aliases (set up with SetAlias
method of App type), if an alias is repeated, the program will panic.

After program initialization, then the command loop can run.

    var cmd = []*cliapp.Command{
        cd,
        ls,
        mv,
    }
    var app = cliapp.New(cmd)
    
    func main() {
    	// initialization
    	app.Run()
    }
    
To finish the program, the method Exit must be called.

Use the method SetPrompt to set up the content of the prompt box.

Authorship and license
----------------------

Copyright (c) 2013, J. Salvador Arias <jsalarias@csnat.unt.edu.ar>
All rights reserved.
Distributed under BSD-style license that can be found in the LICENSE file.

