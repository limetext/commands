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
	// Insert Command inserts the given characters, at all
	// of the current selection locations, possibly replacing
	// text if the selection area covers one or more characters.
	Insert struct {
		backend.DefaultCommand
		// The characters to insert
		Characters string
	}

	// LeftDelete Command deletes characters to the left of the
	// current selection or the current selection if it is not empty.
	LeftDelete struct {
		backend.DefaultCommand
	}

	// RightDelete Command deletes characters to the right of the
	// current selection or the current selection if it is not empty.
	RightDelete struct {
		backend.DefaultCommand
	}

	// DeleteWord Command deletes one word to right or left
	// depending on forward variable.
	DeleteWord struct {
		backend.DefaultCommand
		Forward bool
	}
)

// Run executes the Insert command.
func (c *Insert) Run(v *backend.View, e *backend.Edit) error {
	sel := v.Sel()
	for i := 0; i < sel.Len(); i++ {
		r := sel.Get(i)
		if r.Size() == 0 {
			v.Insert(e, r.B, c.Characters)
		} else {
			v.Replace(e, r, c.Characters)
		}
	}
	return nil
}

// Run executes the LeftDelete command.
func (c *LeftDelete) Run(v *backend.View, e *backend.Edit) error {
	trimSpace := false
	tabSize := 4
	if t := v.Settings().Bool("translate_tabs_to_spaces", false); t {
		if t = v.Settings().Bool("use_tab_stops", true); t {
			trimSpace = true
			tabSize = v.Settings().Int("tab_size", 4)
		}
	}

	sel := v.Sel()
	hasNonEmpty := sel.HasNonEmpty()
	i := 0
	for {
		l := sel.Len()
		if i >= l {
			break
		}
		r := sel.Get(i)
		if r.A == r.B && !hasNonEmpty {
			if trimSpace {
				_, col := v.RowCol(r.A)
				prevCol := r.A - (col - (col-tabSize+(tabSize-1))&^(tabSize-1))
				if prevCol < 0 {
					prevCol = 0
				}
				d := v.SubstrR(text.Region{A: prevCol, B: r.A})
				i := len(d) - 1
				for r.A > prevCol && i >= 0 && d[i] == ' ' {
					r.A--
					i--
				}
			}
			if r.A == r.B {
				r.A--
			}
		}
		v.Erase(e, r)
		if sel.Len() != l {
			continue
		}
		i++
	}
	return nil
}

// Run executes the RightDelete command.
func (c *RightDelete) Run(v *backend.View, e *backend.Edit) error {
	sel := v.Sel()
	hasNonEmpty := sel.HasNonEmpty()
	i := 0
	for {
		l := sel.Len()
		if i >= l {
			break
		}
		r := sel.Get(i)
		if r.A == r.B && !hasNonEmpty {
			r.B++
		}
		v.Erase(e, r)
		if sel.Len() != l {
			continue
		}
		i++
	}
	return nil
}

// Run executes the DeleteWord command.
func (c *DeleteWord) Run(v *backend.View, e *backend.Edit) error {
	var class int
	if c.Forward {
		class = backend.CLASS_WORD_END | backend.CLASS_PUNCTUATION_END | backend.CLASS_LINE_START
	} else {
		class = backend.CLASS_WORD_START | backend.CLASS_PUNCTUATION_START | backend.CLASS_LINE_END | backend.CLASS_LINE_START
	}

	sel := v.Sel()
	var rs []text.Region
	for i := 0; i < sel.Len(); i++ {
		r := sel.Get(i)
		if r.Empty() {
			p := c.findByClass(r.A, class, v)
			if c.Forward {
				r = text.Region{A: r.A, B: p}
			} else {
				r = text.Region{A: p, B: r.A}
			}
		}
		rs = append(rs, r)
	}
	sel.Clear()
	sel.AddAll(rs)
	if c.Forward {
		backend.GetEditor().CommandHandler().RunTextCommand(v, "right_delete", nil)
	} else {
		backend.GetEditor().CommandHandler().RunTextCommand(v, "left_delete", nil)
	}
	return nil
}

func (c *DeleteWord) findByClass(point int, class int, v *backend.View) int {
	var end, d int
	if c.Forward {
		d = 1
		end = v.Size()
		if point > end {
			point = end
		}
		s := v.Substr(text.Region{A: point, B: point + 2})
		if strings.Contains(s, "\t") && strings.Contains(s, " ") {
			class = backend.CLASS_WORD_START | backend.CLASS_PUNCTUATION_START | backend.CLASS_LINE_END
		}
	} else {
		d = -1
		end = 0
		if point < end {
			point = end
		}
		s := v.Substr(text.Region{A: point - 2, B: point})
		if strings.Contains(s, "\t") && strings.Contains(s, " ") {
			class = backend.CLASS_WORD_END | backend.CLASS_PUNCTUATION_END | backend.CLASS_LINE_START
		}
	}
	point += d
	for ; point != end; point += d {
		if v.Classify(point)&class != 0 {
			return point
		}
	}
	return point
}

func init() {
	register([]backend.Command{
		&Insert{},
		&LeftDelete{},
		&RightDelete{},
		&DeleteWord{},
	})
}
