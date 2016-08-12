// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"os"
	"os/user"
	"path"

	"github.com/limetext/backend"
)

type (
	// NewFile command creates a new file.
	NewFile struct {
		backend.DefaultCommand
	}
	// PromptOpenFile command prompts opening
	// an existing file from the filesystem.
	PromptOpenFile struct {
		backend.DefaultCommand
	}
)

// Run executes the NewFile command.
func (c *NewFile) Run(w *backend.Window) error {
	ed := backend.GetEditor()
	ed.ActiveWindow().NewFile()
	return nil
}

// Run executes the PromptOpenFile command.
func (o *PromptOpenFile) Run(w *backend.Window) error {
	dir := viewDirectory(w.ActiveView())
	fe := backend.GetEditor().Frontend()
	files := fe.Prompt("Open file", dir, backend.PROMPT_SELECT_MULTIPLE)
	for _, file := range files {
		w.OpenFile(file, 0)
	}
	return nil
}

func viewDirectory(v *backend.View) string {
	if v != nil && v.FileName() != "" {
		p := path.Dir(v.FileName())
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return "/"
}

func init() {
	register([]backend.Command{
		&NewFile{},
		&PromptOpenFile{},
	})
}
