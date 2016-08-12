# commands
[![Build Status](https://travis-ci.org/limetext/commands.svg?branch=master)](https://travis-ci.org/limetext/commands)
[![Coverage Status](https://img.shields.io/coveralls/limetext/commands.svg?branch=master)](https://coveralls.io/r/limetext/commands?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/limetext/commands)](https://goreportcard.com/report/github.com/limetext/commands)
[![GoDoc](https://godoc.org/github.com/limetext/commands?status.svg)](https://godoc.org/github.com/limetext/commands)

This package contains commands made accessible to frontends and plugins. They come in three flavours: [`ApplicationCommand`](https://godoc.org/github.com/limetext/backend#ApplicationCommand)s, [`WindowCommand`](https://godoc.org/github.com/limetext/backend#WindowCommand)s, and [`TextCommand`](https://godoc.org/github.com/limetext/backend#TextCommand)s."

#Goals

The goal for release 1.0 is to implement all Sublime Text commands.

##Brief overview of Commands

You first build a type of the command which you are writing. 

    type (
        // JoinLines removes every new line in the
        // selections and the first new line after
        JoinLines struct {
            backend.DefaultCommand
        }
    )

Each command has a `Run` method which is executed when the command is invoked.

    //Run executes the DoSomething command
    func (c *DoSomething) Run(v *backend.View, e *backend.Edit) error {
        //Do something here
    }

Commands need to be registered with the backend via the init function.

    func init() {
        register([]backend.Command{
            &DoSomething{},
        })
    }

#Implementing custom commands

If you are interested in implementing your own command, look into the [Implementing commands wiki page](https://github.com/limetext/lime/wiki/Implementing-commands).


# Other references

* [Lime Command interface Api Documentation](http://godoc.org/github.com/limetext/backend#Command)
* Sublime Text 3 unofficial (and [open source](https://github.com/guillermooo/sublime-undocs/)!) [Command documentation](http://docs.sublimetext.info/en/sublime-text-3/extensibility/commands.html)