// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"fmt"

	. "github.com/limetext/backend"
)

type (
	Save struct {
		DefaultCommand
	}

	SaveAll struct {
		DefaultCommand
	}
)

func (c *Save) Run(v *View, e *Edit) error {
	err := v.Save()
	if err != nil {
		GetEditor().Frontend().ErrorMessage(fmt.Sprintf("Failed to save %s:n%s", v.FileName(), err))
		return err
	}
	return nil
}

func (c *SaveAll) Run(w *Window) error {
	for _, v := range w.Views() {
		if err := v.Save(); err != nil {
			GetEditor().Frontend().ErrorMessage(fmt.Sprintf("Failed to save %s:n%s", v.FileName(), err))
			return err
		}
	}
	return nil
}

func init() {
	register([]Command{
		&Save{},
		&SaveAll{},
	})
}
