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
	// JoinLines removes every new line in the
	// selections and the first new line after.
	JoinLines struct {
		backend.DefaultCommand
	}

	// SelectLines makes the selection fill the lines
	// covered by it currently
	SelectLines struct {
		backend.DefaultCommand
		Forward bool
	}
	// SwapLineUp swaps the currently selected lines with the ones above
	SwapLineUp struct {
		backend.DefaultCommand
	}

	// SwapLineDown swaps the currently selected
	// lines with the ones below
	SwapLineDown struct {
		backend.DefaultCommand
	}

	// SplitSelectionIntoLines will split the current
	// selection into lines.
	SplitSelectionIntoLines struct {
		backend.DefaultCommand
	}
)

// Run executes the JoinLines command.
func (c *JoinLines) Run(v *backend.View, e *backend.Edit) error {
	sel := v.Sel()
	for i := 0; i < sel.Len(); i++ {
		r := sel.Get(i)
		// Removing new line and triming in the selection
		t := v.Substr(r)
		t = strings.Replace(t, "\r", "\n", -1)
		slice := strings.Split(t, "\n")
		t = ""
		for j, s := range slice {
			if j == 0 {
				t += s
				continue
			}
			t += " " + strings.TrimLeft(s, " \t")
		}
		v.Replace(e, r, t)
		// Removing the first new line after selection
		liner := v.FullLine(r.End())
		line := v.Substr(liner)
		line = strings.Replace(line, "\n", "", -1)
		line = strings.Replace(line, "\r", "", -1)
		line = strings.TrimRight(line, " \t")
		// Triming the line after
		nextline := liner.End() + 1
		nextliner := v.FullLine(nextline)
		nline := v.Substr(nextliner)
		if nline != "" {
			v.Replace(e, nextliner, " "+strings.TrimLeft(nline, " \t"))
		}
		v.Replace(e, liner, line)
	}

	return nil
}

// Run executes the SwapLineUp command.
func (c *SwapLineUp) Run(v *backend.View, e *backend.Edit) error {
	sel := v.Sel()
	for i := 0; i < sel.Len(); i++ {
		r := sel.Get(i)
		// Expand to all lines under selection
		fline := v.Line(r.Begin())
		lline := v.Line(r.End())
		r = text.Region{fline.Begin(), lline.End()}
		t := v.Substr(r)
		// Select line before region
		bline := v.Line(r.Begin() - 1)
		bt := v.Substr(bline)
		v.Replace(e, r, bt)
		v.Replace(e, bline, t)
	}

	return nil
}

// Run executes the SwapLineDown command.
func (c *SwapLineDown) Run(v *backend.View, e *backend.Edit) error {
	sel := v.Sel()
	for i := 0; i < sel.Len(); i++ {
		r := sel.Get(i)
		// Expand to all lines under selection
		fline := v.Line(r.Begin())
		lline := v.Line(r.End())
		r = text.Region{fline.Begin(), lline.End()}
		t := v.Substr(r)
		// Select line before region
		nline := v.Line(r.End() + 1)
		nt := v.Substr(nline)
		v.Replace(e, nline, t)
		v.Replace(e, r, nt)
	}

	return nil
}

// Run executes the SelectLines command.
func (c *SelectLines) Run(v *backend.View, e *backend.Edit) error {
	var (
		rs      []text.Region
		line, l text.Region
		d       int
	)

	sel := v.Sel()
	for i := 0; i < sel.Len(); i++ {
		r := sel.Get(i)
		// Get the distance of the selection to the begining of line
		if c.Forward {
			line = v.FullLine(r.End())
			l = v.Line(line.End() + 1)
			d = r.End() - line.Begin()
		} else {
			line = v.FullLine(r.Begin())
			l = v.Line(line.Begin() - 1)
			d = r.Begin() - line.Begin()
		}
		// If the next line lenght is more than the calculated distance
		// Put new region at the exact distance
		// If not put region at the end of the next|before line
		if l.Size() < d {
			rs = append(rs, text.Region{l.End(), l.End()})
		} else {
			rs = append(rs, text.Region{l.Begin() + d, l.Begin() + d})
		}
	}
	v.Sel().AddAll(rs)

	return nil
}

// Run executes the SplitSelectionIntoLines command
func (c *SplitSelectionIntoLines) Run(v *backend.View, e *backend.Edit) error {
	var rs []text.Region

	sel := v.Sel()
	for i := 0; i < sel.Len(); i++ {
		r := sel.Get(i)
		lines := v.Lines(r)
		for i := 0; i < len(lines); i++ {
			if i != 0 {
				// Remove line endings
				r2 := v.FullLine(lines[i-1].End())
				lines[i] = lines[i].Clip(r2)
			}
			rs = append(rs, lines[i].Intersection(r))
		}
	}
	v.Sel().Clear()
	v.Sel().AddAll(rs)

	return nil
}

func init() {
	register([]backend.Command{
		&JoinLines{},
		&SelectLines{},
		&SwapLineUp{},
		&SwapLineDown{},
		&SplitSelectionIntoLines{},
	})
}
