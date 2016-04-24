// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	. "github.com/limetext/backend"
)

type (
	NewWindow struct {
		DefaultCommand
	}

	CloseAll struct {
		DefaultCommand
	}

	CloseWindow struct {
		DefaultCommand
	}

	NewWindowApp struct {
		DefaultCommand
	}

	CloseWindowApp struct {
		DefaultCommand
	}
)

func (c *NewWindow) Run(w *Window) error {
	ed := GetEditor()
	ed.SetActiveWindow(ed.NewWindow())
	return nil
}

func (c *CloseAll) Run(w *Window) error {
	w.CloseAllViews()
	return nil
}

func (c *CloseWindow) Run(w *Window) error {
	ed := GetEditor()
	ed.ActiveWindow().Close()
	return nil
}

func (c *NewWindowApp) Run() error {
	ed := GetEditor()
	ed.SetActiveWindow(ed.NewWindow())
	return nil
}

func (c *CloseWindowApp) Run() error {
	ed := GetEditor()
	ed.ActiveWindow().Close()
	return nil
}

func (c *NewWindowApp) IsChecked() bool {
	return false
}

func (c *CloseWindowApp) IsChecked() bool {
	return false
}

func init() {
	register([]Command{
		&NewWindow{},
		&CloseAll{},
		&CloseWindow{},
	})

	registerByName([]namedCmd{
		{"new_window", &NewWindowApp{}},
		{"close_window", &CloseWindowApp{}},
	})
}
