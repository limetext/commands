// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	. "github.com/limetext/backend"
)

type (
	CloseView struct {
		DefaultCommand
	}

	NextView struct {
		DefaultCommand
	}

	PrevView struct {
		DefaultCommand
	}

	SetFileType struct {
		DefaultCommand
		Syntax string
	}
)

func (c *CloseView) Run(w *Window) error {
	w.ActiveView().Close()
	return nil
}

func (c *NextView) Run(w *Window) error {
	for i, v := range w.Views() {
		if v == w.ActiveView() {
			i++
			if i == len(w.Views()) {
				i = 0
			}
			w.SetActiveView(w.Views()[i])
			break
		}
	}

	return nil
}

func (c *PrevView) Run(w *Window) error {
	for i, v := range w.Views() {
		if v == w.ActiveView() {
			if i == 0 {
				i = len(w.Views())
			}
			i--
			w.SetActiveView(w.Views()[i])
			break
		}
	}

	return nil
}

func (c *SetFileType) Run(v *View, e *Edit) error {
	v.SetSyntaxFile(c.Syntax)
	return nil
}

func init() {
	register([]Command{
		&CloseView{},
		&NextView{},
		&PrevView{},
		&SetFileType{},
	})
}
