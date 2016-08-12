# commands

[![Build Status](https://travis-ci.org/limetext/commands.svg?branch=master)](https://travis-ci.org/limetext/commands)
[![Coverage Status](https://img.shields.io/coveralls/limetext/commands.svg?branch=master)](https://coveralls.io/r/limetext/commands?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/limetext/commands)](https://goreportcard.com/report/github.com/limetext/commands)
[![GoDoc](https://godoc.org/github.com/limetext/commands?status.svg)](https://godoc.org/github.com/limetext/commands)

This package contains commands made accessible to frontends and plugins. They come in three flavours: [`ApplicationCommand`](https://godoc.org/github.com/limetext/backend#ApplicationCommand)s, [`WindowCommand`](https://godoc.org/github.com/limetext/backend#WindowCommand)s, and [`TextCommand`](https://godoc.org/github.com/limetext/backend#TextCommand)s.

# Goals

The goal for release 1.0 is to implement all of the commands exposed by Sublime Text 3. See the [commands label](https://github.com/limetext/commands/issues?q=is%3Aopen+is%3Aissue+label%3Acommand) for outstanding commands to be implemented.


## Brief Overview of Commands

A command is a type implementing one of the command interfaces.

```go
type (
    DoSomething struct {
        backend.DefaultCommand
    }
)
```

Each command has a `Run` method which is executed when the command is invoked.

```go
func (c *DoSomething) Run(v *backend.View, e *backend.Edit) error {
    // Do something!
}
```

Commands need to be registered with the backend via the `init` function before it can be executed by plugins.

```go
func init() {
    register([]backend.Command{
        &DoSomething{},
    })
}
```


# Implementing Commands

If you are interested in implementing a command, see the [Implementing commands wiki page](https://github.com/limetext/lime/wiki/Implementing-commands).


# Other References

* Lime's [command interface API documentation](http://godoc.org/github.com/limetext/backend#Command)
* Sublime Text 3's unofficial (and [open source](https://github.com/guillermooo/sublime-undocs/)!) [command documentation](http://docs.sublimetext.info/en/sublime-text-3/extensibility/commands.html)
