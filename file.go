// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"os"
	"os/user"
	"path"

	. "github.com/limetext/backend"
)

type (
	NewFile struct {
		DefaultCommand
	}

	PromptOpenFile struct {
		DefaultCommand
	}
)

func (c *NewFile) Run(w *Window) error {
	ed := GetEditor()
	ed.ActiveWindow().NewFile()
	return nil
}

func (o *PromptOpenFile) Run(w *Window) error {
	dir := "/"
	if v := w.ActiveView(); v != nil {
		p := path.Dir(v.FileName())
		if _, err := os.Stat(p); err == nil {
			dir = p
		} else if usr, err := user.Current(); err != nil {
			dir = usr.HomeDir
		}
	}

	fe := GetEditor().Frontend()
	files := fe.Prompt("Open file", dir, SelectMultiple)
	for _, file := range files {
		w.OpenFile(file, 0)
	}
	return nil
}

func init() {
	register([]Command{
		&NewFile{},
		&PromptOpenFile{},
	})
}
