// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	. "github.com/limetext/backend"
)

type (
	Undo struct {
		BypassUndoCommand
		hard bool
	}
	Redo struct {
		BypassUndoCommand
		hard bool
	}
)

func (c *Undo) Run(v *View, e *Edit) error {
	v.UndoStack().Undo(c.hard)
	return nil
}

func (c *Redo) Run(v *View, e *Edit) error {
	v.UndoStack().Redo(c.hard)
	return nil
}

func init() {
	register([]Command{
		&Undo{hard: true},
		&Redo{hard: true},
	})

	registerByName([]namedCmd{
		{"soft_undo", &Undo{}},
		{"soft_redo", &Redo{}},
	})
}
