// Copyright 2016 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import "github.com/limetext/backend"

type (
	// IncreaseFontSize command increases the font size by 1.
	IncreaseFontSize struct {
		backend.DefaultCommand
	}
	// DecreaseFontSize command decreases the font size by 1.
	DecreaseFontSize struct {
		backend.DefaultCommand
	}
)

// Run executes the IncreaseFontSize command.
func (i *IncreaseFontSize) Run(w *backend.Window) error {
	fontSize := w.Settings().Int("font_size")
	fontSize++
	w.Settings().Set("font_size", fontSize)
	return nil
}

// Run executes the DecreaseFontSize command.
func (d *DecreaseFontSize) Run(w *backend.Window) error {
	fontSize := w.Settings().Int("font_size")
	fontSize--
	w.Settings().Set("font_size", fontSize)
	return nil
}

func init() {
	register([]backend.Command{
		&IncreaseFontSize{},
		&DecreaseFontSize{},
	})
}
