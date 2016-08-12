// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import "github.com/limetext/backend"

type (
	// NopApplication performs NOP.
	NopApplication struct {
		backend.BypassUndoCommand
	}
	// NopWindow performs NOP.
	NopWindow struct {
		backend.BypassUndoCommand
	}
	// NopText performs NOP.
	NopText struct {
		backend.BypassUndoCommand
	}
)

// Run executes the NopApplication command.
func (c *NopApplication) Run() error {
	return nil
}

// IsChecked represents if the command
// contains a checkbox in the frontend
func (c *NopApplication) IsChecked() bool {
	return false
}

// Run executes the NopWindow command.
func (c *NopWindow) Run(w *backend.Window) error {
	return nil
}

// Run executes the NopText command.
func (c *NopText) Run(v *backend.View, e *backend.Edit) error {
	return nil
}

func init() {
	registerByName([]namedCmd{
		{"nop", &NopApplication{}},
		{"nop", &NopWindow{}},
		{"nop", &NopText{}},
	})
}
