// Copyright 2016 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"fmt"

	. "github.com/limetext/backend"
)

type (
	SaveProjectAs struct {
		DefaultCommand
	}

	PromptOpenProject struct {
		DefaultCommand
	}

	CloseProject struct {
		DefaultCommand
	}

	PromptAddFolder struct {
		DefaultCommand
	}

	CloseFolderList struct {
		DefaultCommand
	}
)

func (c *SaveProjectAs) Run(w *Window) error {
	dir := viewDirectory(w.ActiveView())
	fe := GetEditor().Frontend()
	files := fe.Prompt("Save file", dir, PROMPT_SAVE_AS)
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

func (c *PromptOpenProject) Run(w *Window) error {
	dir := viewDirectory(w.ActiveView())
	fe := GetEditor().Frontend()
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

func (c *CloseProject) Run(w *Window) error {
	w.Project().Close()
	return nil
}

func (c *PromptAddFolder) Run(w *Window) error {
	dir := viewDirectory(w.ActiveView())
	fe := GetEditor().Frontend()
	folders := fe.Prompt("Open file", dir, PROMPT_ONLY_FOLDER|PROMPT_SELECT_MULTIPLE)
	for _, folder := range folders {
		w.Project().AddFolder(folder)
	}
	return nil
}

func (c *CloseFolderList) Run(w *Window) error {
	for _, folder := range w.Project().Folders() {
		w.Project().RemoveFolder(folder)
	}
	return nil
}

func init() {
	register([]Command{
		&SaveProjectAs{},
		&PromptOpenProject{},
		&CloseProject{},
		&PromptAddFolder{},
		&CloseFolderList{},
	})
}
