// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"fmt"
	"strings"

	"github.com/limetext/backend"
	"github.com/limetext/text"
	"github.com/limetext/util"
)

type (
	// Move Command moves the current selection.
	Move struct {
		backend.DefaultCommand
		// Specifies the type of "move" operation.
		By MoveByType
		// Whether the current selection should be extended or not.
		Extend bool
		// Whether to move forward or backwards.
		Forward bool
		// Used together with By=Stops, extends "word_separators" defined by settings.
		Separators string
		// Used together with By=Stops, go to word begin.
		WordBegin bool
		// Used together with By=Stops, go to word end.
		WordEnd bool
		// Used together with By=Stops, go to punctuation begin.
		PunctBegin bool
		// Used together with By=Stops, go to punctuation end.
		PunctEnd bool
		// Used together with By=Stops, go to an empty line.
		EmptyLine bool
		// Used together with By=Stops, TODO: ???
		ClipToLine bool
	}

	// MoveByType Specifies the type of "move" operation.
	MoveByType int

	// MoveToType Specifies the type of "move_to" operation to perform.
	MoveToType int

	// MoveTo Command moves or extends the current selection to the specified location.
	MoveTo struct {
		backend.DefaultCommand
		// The type of "move_to" operation to perform.
		To MoveToType
		// Whether the current selection should be extended or not.
		Extend bool
	}

	// ScrollLines Command moves the viewpoint "Amount" lines from the current location.
	ScrollLines struct {
		backend.BypassUndoCommand
		// The number of lines to scroll (positive or negative direction).
		Amount int
	}
)

const (
	// BOL is Beginning of line.
	BOL MoveToType = iota
	// EOL is End of line
	EOL
	//BOF is Beginning of file.
	BOF
	// EOF is End of file.
	EOF
	// Brackets >-> Current level close bracket.
	Brackets
)

const (
	// Characters by Characters.
	Characters MoveByType = iota
	// Stops will move by Stops (TODO(.): what exactly is a stop?).
	Stops
	// Lines will move by Lines.
	Lines
	// Words will move by Words.
	Words
	// WordEnds will move by Word Ends.
	WordEnds
	// SubWords will move by Sub Words.
	SubWords
	// SubWordEnds will Move by Sub Word Ends.
	SubWordEnds
	// Pages will move by Page.
	Pages
)

func moveAction(v *backend.View, extend, scroll bool, transform func(r text.Region) int) {
	fe := backend.GetEditor().Frontend()
	sel := v.Sel()
	rs := sel.Regions()
	bs := v.Size()
	for i := range rs {
		row1, _ := v.RowCol(rs[i].B)
		rs[i].B = transform(rs[i])

		if rs[i].B < 0 {
			rs[i].B = 0
		} else if rs[i].B > bs {
			// Yes > the size, and not size-1 because the cursor being at "size"
			// is the position it will be at when we are appending
			// to the buffer.
			rs[i].B = bs
		}
		row2, _ := v.RowCol(rs[i].B)

		if scroll && row2 != row1 && !fe.VisibleRegion(v).Contains(rs[i].B) {
			scrollLine(v, row1-row2)
		}
		if !extend {
			rs[i].A = rs[i].B
		}
	}
	sel.Clear()
	sel.AddAll(rs)
}

// Set will define the type of move
func (mt *MoveToType) Set(v interface{}) error {
	switch to := v.(string); to {
	case "eol":
		*mt = EOL
	case "bol":
		*mt = BOL
	case "bof":
		*mt = BOF
	case "eof":
		*mt = EOF
	case "brackets":
		*mt = Brackets
	default:
		return fmt.Errorf("move_to: Unimplemented 'to' type: %s", to)
	}
	return nil
}

// Run executes the MoveTo command.
func (c *MoveTo) Run(v *backend.View, e *backend.Edit) error {
	switch c.To {
	case EOL:
		moveAction(v, c.Extend, true, func(r text.Region) int {
			line := v.Line(r.B)
			return line.B
		})
	case BOL:
		moveAction(v, c.Extend, true, func(r text.Region) int {
			line := v.Line(r.B)
			return line.A
		})
	case BOF:
		moveAction(v, c.Extend, true, func(r text.Region) int {
			return 0
		})
	case EOF:
		moveAction(v, c.Extend, true, func(r text.Region) int {
			return v.Size()
		})
	case Brackets:
		moveAction(v, c.Extend, true, func(r text.Region) (pos int) {
			var (
				of          int
				co          = 1
				str, br, rv string
				opening     = "([{"
				closing     = ")]}"
			)
			pos = r.B

			// next and before character
			n := v.Substr(text.Region{r.B, r.B + 1})
			b := v.Substr(text.Region{r.B, r.B - 1})
			if strings.ContainsAny(n, opening) {
				// TODO: Maybe it's better to use sth like view.FindByClass or even
				// view.FindByClass() function itself instead of getting whole text
				// and looping through it. With using view.FindByClass() function
				// backward we won't need to reverse the text anymore
				str = v.Substr(text.Region{r.B + 1, v.Size()})
				br = n
				rv = revert(n)
				of = 2
			} else if strings.ContainsAny(b, closing) {
				// TODO: same as above
				str = v.Substr(text.Region{0, r.B - 1})
				br = b
				rv = revert(b)
				str = reverse(str)
				co = -1
				of = -2
			} else if strings.ContainsAny(n, closing) {
				// TODO: same as above
				str = v.Substr(text.Region{0, r.B - 1})
				br = n
				rv = revert(n)
				str = reverse(str)
				co = -1
				of = -1
			} else {
				// TODO: same as above
				str = v.Substr(text.Region{r.B, v.Size()})
				bef := v.Substr(text.Region{0, r.B})
				if p := strings.LastIndexAny(bef, opening); p == -1 {
					return
				} else {
					br = string(bef[p])
					rv = revert(br)
				}
			}
			count := 1
			for i, c := range str {
				if ch := string(c); ch == br {
					count++
				} else if ch == rv {
					count--
				}
				if count == 0 {
					return i*co + r.B + of
				}
			}
			return
		})
	default:
		return fmt.Errorf("move_to: Unimplemented 'to' action: %d", c.To)
	}
	return nil
}

// Set the type of move.
func (m *MoveByType) Set(v interface{}) error {
	switch by := v.(string); by {
	case "lines":
		*m = Lines
	case "characters":
		*m = Characters
	case "stops":
		*m = Stops
	case "words":
		*m = Words
	case "word_ends":
		*m = WordEnds
	case "subwords":
		*m = SubWords
	case "subword_ends":
		*m = SubWordEnds
	case "pages":
		*m = Pages
	default:
		return fmt.Errorf("move: Unimplemented 'by' action: %s", by)
	}
	return nil
}

// Run executes the Move command.
func (c *Move) Run(v *backend.View, e *backend.Edit) error {
	p := util.Prof.Enter("move.run.action")
	defer p.Exit()

	switch c.By {
	case Characters:
		moveAction(v, c.Extend, true, func(r text.Region) int {
			dir := 1
			if !c.Forward {
				dir = -1
			}
			return r.B + dir
		})
	case Stops:
		moveAction(v, c.Extend, true, func(in text.Region) int {
			tmp := v.Settings().String("word_separators", backend.DEFAULT_SEPARATORS)
			defer v.Settings().Set("word_separators", tmp)
			v.Settings().Set("word_separators", c.Separators)

			classes := 0
			if c.WordBegin {
				classes |= backend.CLASS_WORD_START
			}
			if c.WordEnd {
				classes |= backend.CLASS_WORD_END
			}
			if c.PunctBegin {
				classes |= backend.CLASS_PUNCTUATION_START
			}
			if c.PunctEnd {
				classes |= backend.CLASS_PUNCTUATION_END
			}
			if c.EmptyLine {
				classes |= backend.CLASS_EMPTY_LINE
			}
			return v.FindByClass(in.B, c.Forward, classes)
		})
	case Lines:
		moveAction(v, c.Extend, true, func(in text.Region) int {
			r, col := v.RowCol(in.B)
			fromLine := v.Line(v.TextPoint(r, 0))
			if !c.Forward {
				r--
			} else {
				r++
			}
			// the line we are moving to
			toLine := v.Line(v.TextPoint(r, 0))
			col = findCol(v, in, fromLine, toLine, col)

			return v.TextPoint(r, col)
		})
	case Words:
		moveAction(v, c.Extend, true, func(in text.Region) int {
			return v.FindByClass(in.B, c.Forward, backend.CLASS_WORD_START|
				backend.CLASS_LINE_END|backend.CLASS_LINE_START)
		})
	case WordEnds:
		moveAction(v, c.Extend, true, func(in text.Region) int {
			return v.FindByClass(in.B, c.Forward, backend.CLASS_WORD_END|
				backend.CLASS_LINE_END|backend.CLASS_LINE_START)
		})
	case SubWords:
		moveAction(v, c.Extend, true, func(in text.Region) int {
			return v.FindByClass(in.B, c.Forward, backend.CLASS_SUB_WORD_START|
				backend.CLASS_WORD_START|backend.CLASS_PUNCTUATION_START|
				backend.CLASS_LINE_END|backend.CLASS_LINE_START)
		})
	case SubWordEnds:
		moveAction(v, c.Extend, true, func(in text.Region) int {
			return v.FindByClass(in.B, c.Forward, backend.CLASS_SUB_WORD_END|
				backend.CLASS_WORD_END|backend.CLASS_PUNCTUATION_END|
				backend.CLASS_LINE_END|backend.CLASS_LINE_START)
		})
	case Pages:
		fe := backend.GetEditor().Frontend()
		vr := fe.VisibleRegion(v)
		ls := v.Lines(vr)
		l := len(ls)
		moveAction(v, c.Extend, false, func(in text.Region) int {
			row, col := v.RowCol(in.B)
			fromLine := v.Line(v.TextPoint(row, 0))
			dir := -1
			if c.Forward {
				dir = 1
			}

			row += dir * (l - 1)
			toLine := v.Line(v.TextPoint(row, 0))
			col = findCol(v, in, fromLine, toLine, col)

			return v.TextPoint(row, col)
		})

		dir := -1
		if c.Forward {
			dir = 1
		}
		row, _ := v.RowCol(ls[0].Begin())

		row += dir * (l - 1)
		if row < 0 {
			row = 0
		}
		if end, _ := v.RowCol(v.Size()); row+l-1 > end {
			row = end - l + 1
		}

		end := row + l - 1

		a, b := v.TextPoint(row, 0), v.TextPoint(end, 0)
		r := v.Line(b)
		b = r.End()

		fe.Show(v, text.Region{a, b})
	}
	return nil
}

// If there is tabs in the fromLine, buffer counts them as 1 character but we
// need to count them as tab_size from settings and find the column we should
// move to in toLine based on that
func findCol(v *backend.View, in, fromLine, toLine text.Region, col int) int {
	size := v.Settings().Int("tab_size", 4)
	fromTabs := strings.Count(v.Substr(text.Region{fromLine.Begin(), in.B}), "\t")
	col += fromTabs * (size - 1)

	tab := strings.Repeat("\t", size)
	toText := strings.Replace(v.Substr(toLine), "\t", tab, -1)
	if col > len(toText) {
		col = len(toText)
	}
	toText = toText[:col]
	toTabs := strings.Count(toText, "\t")

	col -= (size - 1) * (toTabs / size)
	mod := toTabs % size
	col -= mod
	if mod > size/2 {
		col += 1
	}

	if col > toLine.Size() {
		col = toLine.Size()
	}

	return col
}

// Default returns the default seprators.
func (c *Move) Default(key string) interface{} {
	if key == "separators" {
		return backend.DEFAULT_SEPARATORS
	}
	return nil
}

func revert(c string) string {
	switch c {
	case "(":
		return ")"
	case ")":
		return "("
	case "[":
		return "]"
	case "]":
		return "["
	case "{":
		return "}"
	case "}":
		return "{"
	}
	return ""
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func (c *ScrollLines) Run(v *backend.View, e *backend.Edit) error {
	scrollLine(v, c.Amount)

	return nil
}

func scrollLine(v *backend.View, amount int) {
	fe := backend.GetEditor().Frontend()
	vr := fe.VisibleRegion(v)

	r1, _ := v.RowCol(vr.Begin())
	r2, _ := v.RowCol(vr.End())
	r1 -= amount
	r2 -= amount

	a := v.TextPoint(r1, 0)
	b := v.TextPoint(r2, 0)
	r := v.Line(b)
	b = r.End()

	fe.Show(v, text.Region{a, b})
}

func init() {
	register([]backend.Command{
		&Move{},
		&MoveTo{},
		&ScrollLines{},
	})
}
