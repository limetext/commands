// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"reflect"
	"testing"

	. "github.com/limetext/backend"
	. "github.com/limetext/text"
)

type MoveTest struct {
	in      []Region
	by      string
	extend  bool
	forward bool
	exp     []Region
	args    Args
}

func runMoveTest(tests []MoveTest, t *testing.T, text string) {
	ed := GetEditor()
	w := ed.NewWindow()
	v := w.NewFile()

	defer func() {
		v.SetScratch(true)
		w.Close()
	}()

	e := v.BeginEdit()
	v.Insert(e, 0, text)
	v.EndEdit(e)

	for i, test := range tests {
		v.Sel().Clear()
		for _, r := range test.in {
			v.Sel().Add(r)
		}
		args := Args{"by": test.by, "extend": test.extend, "forward": test.forward}
		if test.args != nil {
			for k, v := range test.args {
				args[k] = v
			}
		}
		ed.CommandHandler().RunTextCommand(v, "move", args)
		if sr := v.Sel().Regions(); !reflect.DeepEqual(sr, test.exp) {
			t.Errorf("Test %d failed. Expected %v, but got %v: %+v", i, test.exp, sr, test)
		}
	}
}

func TestMove(t *testing.T) {
	text := "Hello World!\nTest123123\nAbrakadabra\n"
	tests := []MoveTest{
		{
			[]Region{{1, 1}, {3, 3}, {6, 6}},
			"characters",
			false,
			true,
			[]Region{{2, 2}, {4, 4}, {7, 7}},
			nil,
		},
		{
			[]Region{{1, 1}, {3, 3}, {6, 6}},
			"characters",
			false,
			false,
			[]Region{{0, 0}, {2, 2}, {5, 5}},
			nil,
		},
		{
			[]Region{{1, 1}, {3, 3}, {10, 6}},
			"characters",
			false,
			true,
			[]Region{{2, 2}, {4, 4}, {7, 7}},
			nil,
		},
		{
			[]Region{{1, 1}, {3, 3}, {10, 6}},
			"characters",
			false,
			false,
			[]Region{{0, 0}, {2, 2}, {5, 5}},
			nil,
		},
		{
			[]Region{{1, 1}, {3, 3}, {10, 6}},
			"characters",
			true,
			true,
			[]Region{{1, 2}, {3, 4}, {10, 7}},
			nil,
		},
		{
			[]Region{{1, 1}, {3, 3}, {10, 6}},
			"characters",
			true,
			false,
			[]Region{{1, 0}, {3, 2}, {10, 5}},
			nil,
		},
		{
			[]Region{{1, 3}, {3, 5}, {10, 7}},
			"characters",
			true,
			true,
			[]Region{{1, 6}, {10, 8}},
			nil,
		},
		{
			[]Region{{1, 1}},
			"stops",
			true,
			true,
			[]Region{{1, 5}},
			Args{"word_end": true},
		},
		{
			[]Region{{1, 1}},
			"stops",
			false,
			true,
			[]Region{{6, 6}},
			Args{"word_begin": true},
		},
		{
			[]Region{{6, 6}},
			"stops",
			false,
			false,
			[]Region{{0, 0}},
			Args{"word_begin": true},
		},
		{
			[]Region{{34, 34}},
			"lines",
			false,
			false,
			[]Region{{23, 23}},
			nil,
		},
		{
			[]Region{{23, 23}},
			"lines",
			false,
			false,
			[]Region{{10, 10}},
			nil,
		},
		{
			[]Region{{100, 100}},
			"lines",
			false,
			false,
			[]Region{{24, 24}},
			nil,
		},
		{
			[]Region{{1, 1}},
			"words",
			true,
			true,
			[]Region{{1, 6}},
			nil,
		},
		{
			[]Region{{5, 5}},
			"words",
			false,
			true,
			[]Region{{6, 6}},
			nil,
		},
		{
			[]Region{{6, 6}, {13, 15}},
			"words",
			false,
			true,
			[]Region{{12, 12}, {23, 23}},
			nil,
		},
		{
			[]Region{{13, 13}},
			"words",
			false,
			false,
			[]Region{{12, 12}},
			nil,
		},
		{
			[]Region{{1, 1}},
			"word_ends",
			true,
			true,
			[]Region{{1, 5}},
			nil,
		},
		{
			[]Region{{5, 5}},
			"word_ends",
			false,
			true,
			[]Region{{11, 11}},
			nil,
		},
		{
			[]Region{{11, 11}, {13, 15}},
			"word_ends",
			false,
			true,
			[]Region{{12, 12}, {23, 23}},
			nil,
		},
		{
			[]Region{{13, 13}},
			"word_ends",
			false,
			false,
			[]Region{{12, 12}},
			nil,
		},
		{
			[]Region{{1, 1}},
			"subwords",
			true,
			true,
			[]Region{{1, 6}},
			nil,
		},
		{
			[]Region{{5, 5}},
			"subwords",
			false,
			true,
			[]Region{{6, 6}},
			nil,
		},
		{
			[]Region{{6, 6}, {13, 15}},
			"subwords",
			false,
			true,
			[]Region{{11, 11}, {23, 23}},
			nil,
		},
		{
			[]Region{{13, 13}},
			"subwords",
			false,
			false,
			[]Region{{12, 12}},
			nil,
		},
		{
			[]Region{{1, 1}},
			"subword_ends",
			true,
			true,
			[]Region{{1, 5}},
			nil,
		},
		{
			[]Region{{5, 5}},
			"subword_ends",
			false,
			true,
			[]Region{{6, 6}},
			nil,
		},
		{
			[]Region{{6, 6}, {13, 15}},
			"subword_ends",
			false,
			true,
			[]Region{{11, 11}, {23, 23}},
			nil,
		},
		{
			[]Region{{13, 13}},
			"subword_ends",
			false,
			false,
			[]Region{{12, 12}},
			nil,
		},
		// Try moving outside the buffer
		{
			[]Region{{0, 0}},
			"lines",
			false,
			false,
			[]Region{{0, 0}},
			nil,
		},
		{
			[]Region{{36, 36}},
			"lines",
			false,
			true,
			[]Region{{36, 36}},
			nil,
		},
		{
			[]Region{{0, 0}},
			"characters",
			false,
			false,
			[]Region{{0, 0}},
			nil,
		},
		{
			[]Region{{36, 36}},
			"characters",
			false,
			true,
			[]Region{{36, 36}},
			nil,
		},
	}
	runMoveTest(tests, t, text)
}

func TestMoveByStops(t *testing.T) {
	text := "Hello WorLd!\nTest12312{\n\n3Stop (testing) tada}\n Abr_akad[abra"
	tests := []MoveTest{
		{
			[]Region{{45, 45}},
			"stops",
			false,
			true,
			[]Region{{56, 56}},
			Args{"word_end": true},
		},
		{
			[]Region{{45, 45}},
			"stops",
			false,
			true,
			[]Region{{46, 46}},
			Args{"word_end": true, "separators": ""},
		},
		{
			[]Region{{8, 8}},
			"stops",
			false,
			true,
			[]Region{{24, 24}},
			Args{"empty_line": true},
		},
		{
			[]Region{{0, 0}},
			"stops",
			false,
			true,
			[]Region{{4, 4}},
			Args{"word_begin": true, "separators": "l"},
		},
		{
			[]Region{{58, 58}},
			"stops",
			false,
			true,
			[]Region{{61, 61}},
			Args{"punct_begin": true},
		},
		{
			[]Region{{7, 7}},
			"stops",
			false,
			true,
			[]Region{{12, 12}},
			Args{"punct_end": true},
		},
	}
	runMoveTest(tests, t, text)
}

func TestMoveWhenWeHaveTabs(t *testing.T) {
	text := "\ttype qmlfrontend struct {\n\t\tstatus_message string"
	tests := []MoveTest{
		{
			[]Region{{35, 35}},
			"lines",
			false,
			false,
			[]Region{{11, 11}},
			nil,
		},
		{
			[]Region{{11, 11}},
			"lines",
			false,
			true,
			[]Region{{35, 35}},
			nil,
		},
	}
	runMoveTest(tests, t, text)
}

type scfe struct {
	show          Region
	defaultAction bool
	files         []string
}

func (f *scfe) StatusMessage(msg string) {}
func (f *scfe) ErrorMessage(msg string)  {}
func (f *scfe) MessageDialog(msg string) {}

func (f *scfe) SetDefaultAction(action bool) {
	f.defaultAction = action
}
func (f *scfe) OkCancelDialog(msg string, button string) bool {
	return f.defaultAction
}
func (f *scfe) VisibleRegion(v *View) Region {
	s := v.Line(v.TextPoint(3*3, 1))
	e := v.Line(v.TextPoint(6*3, 1))
	return Region{s.Begin(), e.End()}
}
func (f *scfe) Show(v *View, r Region) {
	f.show = r
}
func (f *scfe) Prompt(title, dir string, flags int) []string {
	return f.files
}

func TestScrollLines(t *testing.T) {
	var fe scfe
	ed := GetEditor()
	ed.SetFrontend(&fe)
	ch := ed.CommandHandler()
	w := ed.NewWindow()
	defer w.Close()

	v := w.NewFile()
	defer func() {
		v.SetScratch(true)
	}()

	e := v.BeginEdit()
	for i := 0; i < 10; i++ {
		v.Insert(e, 0, "Hello World!\nTest123123\nAbrakadabra\n")
	}
	v.EndEdit(e)
	ch.RunTextCommand(v, "scroll_lines", Args{"amount": 0})

	if c := v.Line(v.TextPoint(3*3, 1)); fe.show.Begin() != c.Begin() {
		t.Errorf("Expected %v, but got %v", c, fe.show)
	}

	ch.RunTextCommand(v, "scroll_lines", Args{"amount": 1})
	if c := v.Line(v.TextPoint(3*3-1, 1)); fe.show.Begin() != c.Begin() {
		t.Errorf("Expected %v, but got %v", c, fe.show)
	}
	t.Log(fe.VisibleRegion(v), v.Line(v.TextPoint(6*3+1, 1)))
	ch.RunTextCommand(v, "scroll_lines", Args{"amount": -1})
	if c := v.Line(v.TextPoint(6*3+1, 1)); fe.show.Begin() != c.Begin() {
		t.Errorf("Expected %v, but got %v", c, fe.show)
	}
}
