// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import . "github.com/limetext/backend"

type (
	NopApplication struct {
		BypassUndoCommand
	}

	NopWindow struct {
		BypassUndoCommand
	}

	NopText struct {
		BypassUndoCommand
	}
)

func (c *NopApplication) Run() error {
	return nil
}

func (c *NopApplication) IsChecked() bool {
	return false
}

func (c *NopWindow) Run(w *Window) error {
	return nil
}

func (c *NopText) Run(v *View, e *Edit) error {
	return nil
}

func init() {
	registerByName([]namedCmd{
		{"nop", &NopApplication{}},
		{"nop", &NopWindow{}},
		{"nop", &NopText{}},
	})
}
