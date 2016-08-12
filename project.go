// Copyright 2016 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"fmt"

	"github.com/limetext/backend"
)

type (

	// SaveProjectAs command enables us to save the project
	// as a text file, which can then be imported
	// into lime using PromptOpenProject command.
	SaveProjectAs struct {
		backend.DefaultCommand
	}

	// PromptOpenProject command enables us to open the
	// project file saved using the SaveProjectAs command.
	PromptOpenProject struct {
		backend.DefaultCommand
	}

	// CloseProject command enables us to close an existing
	// open project.
	CloseProject struct {
		backend.DefaultCommand
	}

	// PromptAddFolder adds a folder to the existing
	// opened project.
	PromptAddFolder struct {
		backend.DefaultCommand
	}

	// CloseFolderList removes the folder list from the
	// opened project.
	CloseFolderList struct {
		backend.DefaultCommand
	}
)

// Run executes the SaveProjectAs command.
func (c *SaveProjectAs) Run(w *backend.Window) error {
	dir := viewDirectory(w.ActiveView())
	fe := backend.GetEditor().Frontend()
	files := fe.Prompt("Save file", dir, backend.PROMPT_SAVE_AS)
	if len(files) == 0 {
		return nil
	}

	name := files[0]
	if err := w.Project().SaveAs(name); err != nil {
		fe.ErrorMessage(fmt.Sprintf("Failed to save as %s:%s", name, err))
		return err
	}
	return nil
}

// Run executes the PromptOpenProject command.
func (c *PromptOpenProject) Run(w *backend.Window) error {
	dir := viewDirectory(w.ActiveView())
	fe := backend.GetEditor().Frontend()
	files := fe.Prompt("Open file", dir, 0)
	if len(files) == 0 {
		return nil
	}
	if p := w.OpenProject(files[0]); p == nil {
		err := fmt.Errorf("Unable to read project %s", files[0])
		fe.ErrorMessage(err.Error())
		return err
	}
	return nil
}

// Run executes the CloseProject command.
func (c *CloseProject) Run(w *backend.Window) error {
	w.Project().Close()
	return nil
}

// Run executes the PromptAddFolder command.
func (c *PromptAddFolder) Run(w *backend.Window) error {
	dir := viewDirectory(w.ActiveView())
	fe := backend.GetEditor().Frontend()
	folders := fe.Prompt("Open file", dir, backend.PROMPT_ONLY_FOLDER|backend.PROMPT_SELECT_MULTIPLE)
	for _, folder := range folders {
		w.Project().AddFolder(folder)
	}
	return nil
}

// Run executes the CloseFolderList command.
func (c *CloseFolderList) Run(w *backend.Window) error {
	for _, folder := range w.Project().Folders() {
		w.Project().RemoveFolder(folder)
	}
	return nil
}

func init() {
	register([]backend.Command{
		&SaveProjectAs{},
		&PromptOpenProject{},
		&CloseProject{},
		&PromptAddFolder{},
		&CloseFolderList{},
	})
}
