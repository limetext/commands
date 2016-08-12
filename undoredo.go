// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"github.com/limetext/backend"
)

type (
	Undo struct {
		backend.BypassUndoCommand
		hard bool
	}

	Redo struct {
		backend.BypassUndoCommand
		hard bool
	}
)

// Run executes the Undo command.
func (c *Undo) Run(v *backend.View, e *backend.Edit) error {
	v.UndoStack().Undo(c.hard)
	return nil
}

// Run executes the Redo command.
func (c *Redo) Run(v *backend.View, e *backend.Edit) error {
	v.UndoStack().Redo(c.hard)
	return nil
}

func init() {
	register([]backend.Command{
		&Undo{hard: true},
		&Redo{hard: true},
	})

	registerByName([]namedCmd{
		{"soft_undo", &Undo{}},
		{"soft_redo", &Redo{}},
	})
}
