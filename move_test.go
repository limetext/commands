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
			t.Errorf("Test %d: Expected %v, but got %v", i, test.exp, sr)
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
		{
			[]text.Region{{0, 0}, {1, 1}, {2, 2}, {4, 4}, {7, 7}},
			"lines",
			false,
			true,
			[]text.Region{{27, 27}, {28, 28}, {29, 29}, {31, 31}},
			nil,
		},
	}
	runMoveTest(tests, t, inputText)
}

func TestMoveNeedingScroll(t *testing.T) {
	input := "Hello World!\nTest123123\nAbrakadabra\n"
	for i := 0; i < 9; i++ {
		input += input
	}
	tests := []struct {
		MoveTest
		inVr, expVr text.Region
	}{
		{
			MoveTest{
				[]text.Region{{59, 59}},
				"characters",
				false,
				true,
				[]text.Region{{60, 60}},
				nil,
			},
			text.Region{0, 59},
			text.Region{13, 71},
		},
		{
			MoveTest{
				[]text.Region{{1, 1}},
				"characters",
				false,
				false,
				[]text.Region{{0, 0}},
				nil,
			},
			text.Region{0, 59},
			text.Region{0, 59},
		},
		{
			MoveTest{
				[]text.Region{{13, 13}},
				"characters",
				false,
				false,
				[]text.Region{{12, 12}},
				nil,
			},
			text.Region{13, 71},
			text.Region{0, 59},
		},
		{
			MoveTest{
				[]text.Region{{50, 50}},
				"lines",
				false,
				true,
				[]text.Region{{61, 61}},
				nil,
			},
			text.Region{0, 59},
			text.Region{13, 71},
		},
		{
			MoveTest{
				[]text.Region{{17, 17}},
				"lines",
				false,
				true,
				[]text.Region{{28, 28}},
				nil,
			},
			text.Region{0, 59},
			text.Region{0, 59},
		},
		{
			MoveTest{
				[]text.Region{{0, 0}},
				"pages",
				false,
				true,
				[]text.Region{{49, 49}},
				nil,
			},
			text.Region{0, 59},
			text.Region{49, 107},
		},
		{
			MoveTest{
				[]text.Region{{36, 36}},
				"pages",
				false,
				false,
				[]text.Region{{0, 0}},
				nil,
			},
			text.Region{24, 84},
			text.Region{0, 59},
		},
		{
			MoveTest{
				[]text.Region{{24, 24}},
				"pages",
				true,
				true,
				[]text.Region{{24, 72}},
				nil,
			},
			text.Region{0, 59},
			text.Region{49, 107},
		},
	}

	ed := backend.GetEditor()
	var fe front
	ed.SetFrontend(&fe)
	w := ed.NewWindow()
	v := w.NewFile()

	defer func() {
		v.SetScratch(true)
		w.Close()
	}()

	e := v.BeginEdit()
	v.Insert(e, 0, input)
	v.EndEdit(e)

	for i, test := range tests {
		v.Sel().Clear()
		for _, r := range test.in {
			v.Sel().Add(r)
		}
		fe.vr = test.inVr

		args := backend.Args{"by": test.by, "extend": test.extend, "forward": test.forward}
		if test.args != nil {
			for k, v := range test.args {
				args[k] = v
			}
		}
		ed.CommandHandler().RunTextCommand(v, "move", args)

		if sr := v.Sel().Regions(); !reflect.DeepEqual(sr, test.exp) {
			t.Errorf("Test %d: Expected %v, but got %v", i, test.exp, sr)
		}
		if vr := ed.Frontend().VisibleRegion(v); vr.String() != test.expVr.String() {
			t.Errorf("Test %d: Expected visible region %s, but got %s",
				i, test.expVr, vr)
		}
	}
}

type front struct {
	vr            text.Region
	defaultAction bool
	files         []string
}

func (f *front) StatusMessage(msg string) {}
func (f *front) ErrorMessage(msg string)  {}
func (f *front) MessageDialog(msg string) {}

func (f *front) SetDefaultAction(action bool) {
	f.defaultAction = action
}
func (f *front) OkCancelDialog(msg string, button string) bool {
	return f.defaultAction
}
func (f *front) VisibleRegion(v *backend.View) text.Region {
	return f.vr
}
func (f *front) Show(v *backend.View, r text.Region) {
	f.vr = r
}
func (f *front) Prompt(title, dir string, flags int) []string {
	return f.files
}

func TestScrollLines(t *testing.T) {
	var fe front
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

	fe.vr = text.Region{13, 71}

	ch.RunTextCommand(v, "scroll_lines", backend.Args{"amount": -1})
	if fe.vr.A != 24 || fe.vr.B != 84 {
		t.Errorf("Expected %s, but got %s", text.Region{24, 84}, fe.vr)
	}

	ch.RunTextCommand(v, "scroll_lines", backend.Args{"amount": 0})
	if fe.vr.A != 24 || fe.vr.B != 84 {
		t.Errorf("Expected %s, but got %s", text.Region{24, 84}, fe.vr)
	}

	ch.RunTextCommand(v, "scroll_lines", backend.Args{"amount": 1})
	if fe.vr.A != 13 || fe.vr.B != 71 {
		t.Errorf("Expected %s, but got %s", text.Region{13, 71}, fe.vr)
	}
}
