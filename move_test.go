// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"reflect"
	"testing"

	"github.com/limetext/backend"
	"github.com/limetext/text"
)

type MoveTest struct {
	in      []text.Region
	by      string
	extend  bool
	forward bool
	exp     []text.Region
	args    backend.Args
}

func runMoveTest(tests []MoveTest, t *testing.T, text string) {
	ed := backend.GetEditor()
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
		args := backend.Args{"by": test.by, "extend": test.extend, "forward": test.forward}
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
	inpuText := "Hello World!\nTest123123\nAbrakadabra\n"
	tests := []MoveTest{
		{
			[]text.Region{{1, 1}, {3, 3}, {6, 6}},
			"characters",
			false,
			true,
			[]text.Region{{2, 2}, {4, 4}, {7, 7}},
			nil,
		},
		{
			[]text.Region{{1, 1}, {3, 3}, {6, 6}},
			"characters",
			false,
			false,
			[]text.Region{{0, 0}, {2, 2}, {5, 5}},
			nil,
		},
		{
			[]text.Region{{1, 1}, {3, 3}, {10, 6}},
			"characters",
			false,
			true,
			[]text.Region{{2, 2}, {4, 4}, {7, 7}},
			nil,
		},
		{
			[]text.Region{{1, 1}, {3, 3}, {10, 6}},
			"characters",
			false,
			false,
			[]text.Region{{0, 0}, {2, 2}, {5, 5}},
			nil,
		},
		{
			[]text.Region{{1, 1}, {3, 3}, {10, 6}},
			"characters",
			true,
			true,
			[]text.Region{{1, 2}, {3, 4}, {10, 7}},
			nil,
		},
		{
			[]text.Region{{1, 1}, {3, 3}, {10, 6}},
			"characters",
			true,
			false,
			[]text.Region{{1, 0}, {3, 2}, {10, 5}},
			nil,
		},
		{
			[]text.Region{{1, 3}, {3, 5}, {10, 7}},
			"characters",
			true,
			true,
			[]text.Region{{1, 6}, {10, 8}},
			nil,
		},
		{
			[]text.Region{{1, 1}},
			"stops",
			true,
			true,
			[]text.Region{{1, 5}},
			backend.Args{"word_end": true},
		},
		{
			[]text.Region{{1, 1}},
			"stops",
			false,
			true,
			[]text.Region{{6, 6}},
			backend.Args{"word_begin": true},
		},
		{
			[]text.Region{{6, 6}},
			"stops",
			false,
			false,
			[]text.Region{{0, 0}},
			backend.Args{"word_begin": true},
		},
		{
			[]text.Region{{34, 34}},
			"lines",
			false,
			false,
			[]text.Region{{23, 23}},
			nil,
		},
		{
			[]text.Region{{23, 23}},
			"lines",
			false,
			false,
			[]text.Region{{10, 10}},
			nil,
		},
		{
			[]text.Region{{100, 100}},
			"lines",
			false,
			false,
			[]text.Region{{24, 24}},
			nil,
		},
		{
			[]text.Region{{12, 12}},
			"lines",
			false,
			true,
			[]text.Region{{23, 23}},
			nil,
		},
		{
			[]text.Region{{35, 35}},
			"lines",
			false,
			false,
			[]text.Region{{23, 23}},
			nil,
		},
		{
			[]text.Region{{1, 1}},
			"words",
			true,
			true,
			[]text.Region{{1, 6}},
			nil,
		},
		{
			[]text.Region{{5, 5}},
			"words",
			false,
			true,
			[]text.Region{{6, 6}},
			nil,
		},
		{
			[]text.Region{{6, 6}, {13, 15}},
			"words",
			false,
			true,
			[]text.Region{{12, 12}, {23, 23}},
			nil,
		},
		{
			[]text.Region{{13, 13}},
			"words",
			false,
			false,
			[]text.Region{{12, 12}},
			nil,
		},
		{
			[]text.Region{{1, 1}},
			"word_ends",
			true,
			true,
			[]text.Region{{1, 5}},
			nil,
		},
		{
			[]text.Region{{5, 5}},
			"word_ends",
			false,
			true,
			[]text.Region{{11, 11}},
			nil,
		},
		{
			[]text.Region{{11, 11}, {13, 15}},
			"word_ends",
			false,
			true,
			[]text.Region{{12, 12}, {23, 23}},
			nil,
		},
		{
			[]text.Region{{13, 13}},
			"word_ends",
			false,
			false,
			[]text.Region{{12, 12}},
			nil,
		},
		{
			[]text.Region{{1, 1}},
			"subwords",
			true,
			true,
			[]text.Region{{1, 6}},
			nil,
		},
		{
			[]text.Region{{5, 5}},
			"subwords",
			false,
			true,
			[]text.Region{{6, 6}},
			nil,
		},
		{
			[]text.Region{{6, 6}, {13, 15}},
			"subwords",
			false,
			true,
			[]text.Region{{11, 11}, {23, 23}},
			nil,
		},
		{
			[]text.Region{{13, 13}},
			"subwords",
			false,
			false,
			[]text.Region{{12, 12}},
			nil,
		},
		{
			[]text.Region{{1, 1}},
			"subword_ends",
			true,
			true,
			[]text.Region{{1, 5}},
			nil,
		},
		{
			[]text.Region{{5, 5}},
			"subword_ends",
			false,
			true,
			[]text.Region{{6, 6}},
			nil,
		},
		{
			[]text.Region{{6, 6}, {13, 15}},
			"subword_ends",
			false,
			true,
			[]text.Region{{11, 11}, {23, 23}},
			nil,
		},
		{
			[]text.Region{{13, 13}},
			"subword_ends",
			false,
			false,
			[]text.Region{{12, 12}},
			nil,
		},
		// Try moving outside the buffer
		{
			[]text.Region{{0, 0}},
			"lines",
			false,
			false,
			[]text.Region{{0, 0}},
			nil,
		},
		{
			[]text.Region{{36, 36}},
			"lines",
			false,
			true,
			[]text.Region{{36, 36}},
			nil,
		},
		{
			[]text.Region{{0, 0}},
			"characters",
			false,
			false,
			[]text.Region{{0, 0}},
			nil,
		},
		{
			[]text.Region{{36, 36}},
			"characters",
			false,
			true,
			[]text.Region{{36, 36}},
			nil,
		},
	}
	runMoveTest(tests, t, inpuText)
}

func TestMoveByStops(t *testing.T) {
	inputText := "Hello WorLd!\nTest12312{\n\n3Stop (testing) tada}\n Abr_akad[abra"
	tests := []MoveTest{
		{
			[]text.Region{{45, 45}},
			"stops",
			false,
			true,
			[]text.Region{{56, 56}},
			backend.Args{"word_end": true},
		},
		{
			[]text.Region{{45, 45}},
			"stops",
			false,
			true,
			[]text.Region{{46, 46}},
			backend.Args{"word_end": true, "separators": ""},
		},
		{
			[]text.Region{{8, 8}},
			"stops",
			false,
			true,
			[]text.Region{{24, 24}},
			backend.Args{"empty_line": true},
		},
		{
			[]text.Region{{0, 0}},
			"stops",
			false,
			true,
			[]text.Region{{4, 4}},
			backend.Args{"word_begin": true, "separators": "l"},
		},
		{
			[]text.Region{{58, 58}},
			"stops",
			false,
			true,
			[]text.Region{{61, 61}},
			backend.Args{"punct_begin": true},
		},
		{
			[]text.Region{{7, 7}},
			"stops",
			false,
			true,
			[]text.Region{{12, 12}},
			backend.Args{"punct_end": true},
		},
	}
	runMoveTest(tests, t, inputText)
}

func TestMoveWhenWeHaveTabs(t *testing.T) {
	inputText := "\ttype qmlfrontend struct {\n\t\tstatus_message string"
	tests := []MoveTest{
		{
			[]text.Region{{35, 35}},
			"lines",
			false,
			false,
			[]text.Region{{11, 11}},
			nil,
		},
		{
			[]text.Region{{11, 11}},
			"lines",
			false,
			true,
			[]text.Region{{35, 35}},
			nil,
		},
	}
	runMoveTest(tests, t, inputText)
}

type scfe struct {
	show          text.Region
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
func (f *scfe) VisibleRegion(v *backend.View) text.Region {
	s := v.Line(v.TextPoint(3*3, 1))
	e := v.Line(v.TextPoint(6*3, 1))
	return text.Region{s.Begin(), e.End()}
}
func (f *scfe) Show(v *backend.View, r text.Region) {
	f.show = r
}
func (f *scfe) Prompt(title, dir string, flags int) []string {
	return f.files
}

func TestScrollLines(t *testing.T) {
	var fe scfe
	ed := backend.GetEditor()
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
	ch.RunTextCommand(v, "scroll_lines", backend.Args{"amount": 0})

	if c := v.Line(v.TextPoint(3*3, 1)); fe.show.Begin() != c.Begin() {
		t.Errorf("Expected %v, but got %v", c, fe.show)
	}

	ch.RunTextCommand(v, "scroll_lines", backend.Args{"amount": 1})
	if c := v.Line(v.TextPoint(3*3-1, 1)); fe.show.Begin() != c.Begin() {
		t.Errorf("Expected %v, but got %v", c, fe.show)
	}
	t.Log(fe.VisibleRegion(v), v.Line(v.TextPoint(6*3+1, 1)))
	ch.RunTextCommand(v, "scroll_lines", backend.Args{"amount": -1})
	if c := v.Line(v.TextPoint(6*3+1, 1)); fe.show.Begin() != c.Begin() {
		t.Errorf("Expected %v, but got %v", c, fe.show)
	}
}
