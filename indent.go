// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"strings"

	"github.com/limetext/backend"
	"github.com/limetext/text"
)

type (
	// Indent Command increments indentation of selection.
	Indent struct {
		backend.DefaultCommand
	}

	// Unindent Command decrements indentation of selection.
	Unindent struct {
		backend.DefaultCommand
	}
)

// Run executes the Indent command.
func (c *Indent) Run(v *backend.View, e *backend.Edit) error {
	indent := "\t"
	if t := v.Settings().Bool("translate_tabs_to_spaces", false); t {
		indent = strings.Repeat(" ", v.Settings().Int("tab_size", 4))
	}
	sel := v.Sel()

	for i := 0; i < sel.Len(); i++ {
		r := sel.Get(i)
		startRow, _ := v.RowCol(r.Begin())
		endRow, _ := v.RowCol(r.End())
		for row := startRow; row <= endRow; row++ {
			// Insert an indent at the beginning of the line
			pos := v.TextPoint(row, 0)
			v.Insert(e, pos, indent)
		}
	}
	return nil
}

// Run executes the Unindent command.
func (c *Unindent) Run(v *backend.View, e *backend.Edit) error {
	tabSize := v.Settings().Int("tab_size", 4)
	sel := v.Sel()
	for i := 0; i < sel.Len(); i++ {
		r := sel.Get(i)
		startRow, _ := v.RowCol(r.Begin())
		endRow, _ := v.RowCol(r.End())
		for row := startRow; row <= endRow; row++ {
			pos := v.TextPoint(row, 0)
			// Get the first at the beginning of the line (as many as defined by tab_size)
			sub := v.Substr(text.Region{pos, pos + tabSize})
			if len(sub) == 0 {
				continue
			}
			toRemove := 0
			if sub[0] == byte('\t') {
				// Case 1: the first character is a tab, remove only it
				toRemove = 1
			} else if sub[0] == byte(' ') {
				// Case 2: the first character is a space, we remove as much spaces as we can
				toRemove = 1
				for toRemove < len(sub) && sub[toRemove] == byte(' ') {
					toRemove++
				}
			}
			if toRemove > 0 {
				v.Erase(e, text.Region{pos, pos + toRemove})
			}
		}
	}
	return nil
}

func init() {
	register([]backend.Command{
		&Indent{},
		&Unindent{},
	})
}
