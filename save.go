// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"fmt"

	"github.com/limetext/backend"
)

type (

	// Save command writes the currently
	// opened file to the disk.
	Save struct {
		backend.DefaultCommand
	}

	// PromptSaveAs command lets us save
	// the currently active
	// file with a different name.
	PromptSaveAs struct {
		backend.DefaultCommand
	}

	// SaveAll command saves all the open files to the disk.
	SaveAll struct {
		backend.DefaultCommand
	}
)

// Run executes the Save command.
func (c *Save) Run(v *backend.View, e *backend.Edit) error {
	err := v.Save()
	if err != nil {
		backend.GetEditor().Frontend().ErrorMessage(fmt.Sprintf("Failed to save %s:n%s", v.FileName(), err))
		return err
	}
	return nil
}

// Run executes the PromptSaveAs command.
func (c *PromptSaveAs) Run(v *backend.View, e *backend.Edit) error {
	dir := viewDirectory(v)
	fe := backend.GetEditor().Frontend()
	files := fe.Prompt("Save file", dir, backend.PROMPT_SAVE_AS)
	if len(files) == 0 {
		return nil
	}

	name := files[0]
	if err := v.SaveAs(name); err != nil {
		fe.ErrorMessage(fmt.Sprintf("Failed to save as %s:%s", name, err))
		return err
	}
	return nil
}

// Run executes the SaveAll command.
func (c *SaveAll) Run(w *backend.Window) error {
	fe := backend.GetEditor().Frontend()
	for _, v := range w.Views() {
		if err := v.Save(); err != nil {
			fe.ErrorMessage(fmt.Sprintf("Failed to save %s:n%s", v.FileName(), err))
			return err
		}
	}
	return nil
}

func init() {
	register([]backend.Command{
		&Save{},
		&PromptSaveAs{},
		&SaveAll{},
	})
}
