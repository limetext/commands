// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"github.com/limetext/backend"
	"github.com/limetext/text"
)

type (
	// SingleSelection command merges multiple cursors
	// into a single one.
	SingleSelection struct {
		backend.DefaultCommand
	}
	// SelectAll command selects the whole buffer of
	// the current file.
	SelectAll struct {
		backend.DefaultCommand
	}
)

// Run executes the SingleSelection command.
func (c *SingleSelection) Run(v *backend.View, e *backend.Edit) error {
	/*
		Correct behavior of SingleSelect:
			- Remove all selection regions but the first.
	*/

	r := v.Sel().Get(0)
	v.Sel().Clear()
	v.Sel().Add(r)
	return nil
}

// Run executes the SelectAll command.
func (c *SelectAll) Run(v *backend.View, e *backend.Edit) error {
	/*
		Correct behavior of SelectAll:
			- Select a single region of (0, view.buffersize())
	*/

	r := text.Region{0, v.Size()}
	v.Sel().Clear()
	v.Sel().Add(r)
	return nil
}

func init() {
	register([]backend.Command{
		&SingleSelection{},
		&SelectAll{},
	})
}
