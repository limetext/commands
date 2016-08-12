// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"testing"

	"github.com/limetext/backend"
	"github.com/limetext/text"
)

func TestJoinLines(t *testing.T) {
	tests := []struct {
		text   string
		sel    []text.Region
		expect string
	}{
		{
			"a\n\t  bc",
			[]text.Region{{1, 1}},
			"a bc",
		},
		{
			"abc\r\n\tde",
			[]text.Region{{0, 0}},
			"abc de",
		},
		{
			"testing \t\t\n join",
			[]text.Region{{9, 8}},
			"testing join",
		},
		{
			"test\n join\n command\n whith\n multiple\n regions",
			[]text.Region{{2, 17}, {34, 40}},
			"test join command whith\n multiple regions",
		},
	}

	ed := backend.GetEditor()
	w := ed.NewWindow()
	defer w.Close()

	for i, test := range tests {
		v := w.NewFile()
		defer func() {
			v.SetScratch(true)
			v.Close()
		}()

		e := v.BeginEdit()
		v.Insert(e, 0, test.text)
		v.EndEdit(e)

		v.Sel().Clear()
		v.Sel().AddAll(test.sel)

		ed.CommandHandler().RunTextCommand(v, "join_lines", nil)
		if d := v.Substr(text.Region{0, v.Size()}); d != test.expect {
			t.Errorf("Test %d:\nExcepted: '%s'\nbut got: '%s'", i, test.expect, d)
		}
	}
}

func TestSelectLines(t *testing.T) {
	tests := []struct {
		text    string
		sel     []text.Region
		forward bool
		expect  []text.Region
	}{
		{
			"abc\ndefg",
			[]text.Region{{1, 1}},
			true,
			[]text.Region{{1, 1}, {5, 5}},
		},
		{
			"abcde\nfg",
			[]text.Region{{4, 4}},
			true,
			[]text.Region{{4, 4}, {8, 8}},
		},
		{
			"Testing select lines command\nin\nlime text",
			[]text.Region{{8, 14}, {30, 30}},
			true,
			[]text.Region{{8, 14}, {30, 30}, {31, 31}, {33, 33}},
		},
		{
			"abc\n\ndefg",
			[]text.Region{{6, 6}},
			false,
			[]text.Region{{6, 6}, {4, 4}},
		},
		{
			"Testing select lines command\nin\nlime text",
			[]text.Region{{30, 36}, {29, 29}},
			false,
			[]text.Region{{30, 36}, {29, 29}, {0, 0}, {1, 1}},
		},
	}

	ed := backend.GetEditor()
	w := ed.NewWindow()
	defer w.Close()

	for i, test := range tests {
		v := w.NewFile()
		defer func() {
			v.SetScratch(true)
			v.Close()
		}()

		e := v.BeginEdit()
		v.Insert(e, 0, test.text)
		v.EndEdit(e)

		v.Sel().Clear()
		v.Sel().AddAll(test.sel)

		ed.CommandHandler().RunTextCommand(v, "select_lines", backend.Args{"forward": test.forward})
		// Comparing regions
		d := v.Sel()
		if d.Len() != len(test.expect) {
			t.Errorf("Test %d: Excepted '%d' regions, but got '%d' regions", i, len(test.expect), d.Len())
			t.Errorf("%+v, %+v", test.expect, d.Regions())
		} else {
			var found bool
			for _, r := range test.expect {
				found = false
				for _, r2 := range d.Regions() {
					if r2 == r {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Test %d:\nRegion %+v not found in view regions: %+v", i, r, d.Regions())
				}
			}
		}
	}
}

func TestSwapLine(t *testing.T) {
	type SwapLineTest struct {
		text   string
		sel    []text.Region
		expect string
	}

	uptests := []SwapLineTest{
		{
			"a\nb",
			[]text.Region{{2, 2}},
			"b\na",
		},
		{
			"Testing swap line up\ncommand whit multiple\nregions selected\nTesting swap line up\ncommand whit multiple\nregions selected",
			[]text.Region{{25, 53}, {86, 95}},
			"command whit multiple\nregions selected\nTesting swap line up\ncommand whit multiple\nTesting swap line up\nregions selected",
		},
	}

	dwtests := []SwapLineTest{
		{
			"a\nb",
			[]text.Region{{1, 1}},
			"b\na",
		},
		{
			"Testing swap line up\ncommand whit multiple\nregions selected\nTesting swap line up\ncommand whit multiple\nregions selected",
			[]text.Region{{25, 53}, {86, 95}},
			"Testing swap line up\nTesting swap line up\ncommand whit multiple\nregions selected\nregions selected\ncommand whit multiple",
		},
	}

	ed := backend.GetEditor()
	w := ed.NewWindow()
	defer w.Close()

	for i, test := range uptests {
		v := w.NewFile()
		defer func() {
			v.SetScratch(true)
			v.Close()
		}()

		e := v.BeginEdit()
		v.Insert(e, 0, test.text)
		v.EndEdit(e)

		v.Sel().Clear()
		v.Sel().AddAll(test.sel)

		ed.CommandHandler().RunTextCommand(v, "swap_line_up", nil)
		if d := v.Substr(text.Region{0, v.Size()}); d != test.expect {
			t.Errorf("Test %d:\nExcepted: '%s'\nbut got: '%s'", i, test.expect, d)
		}
	}

	for i, test := range dwtests {
		v := w.NewFile()
		defer func() {
			v.SetScratch(true)
			v.Close()
		}()

		e := v.BeginEdit()
		v.Insert(e, 0, test.text)
		v.EndEdit(e)

		v.Sel().Clear()
		v.Sel().AddAll(test.sel)

		ed.CommandHandler().RunTextCommand(v, "swap_line_down", nil)
		if d := v.Substr(text.Region{0, v.Size()}); d != test.expect {
			t.Errorf("Test %d:\nExcepted: '%s'\nbut got: '%s'", i, test.expect, d)
		}
	}
}

func TestSplitToLines(t *testing.T) {
	tests := []struct {
		text   string
		sel    []text.Region
		expect []text.Region
	}{
		{
			"ab\ncd\nef",
			[]text.Region{{4, 7}},
			[]text.Region{{4, 5}, {6, 7}},
		},
		{
			"ab\ncd\nef",
			[]text.Region{{0, 8}},
			[]text.Region{{0, 2}, {3, 5}, {6, 8}},
		},
		{
			"ab\ncd\nef",
			[]text.Region{{0, 4}, {4, 7}},
			[]text.Region{{0, 2}, {3, 4}, {4, 5}, {6, 7}},
		},
	}

	ed := backend.GetEditor()
	w := ed.NewWindow()
	defer w.Close()

	for i, test := range tests {
		v := w.NewFile()
		defer func() {
			v.SetScratch(true)
			v.Close()
		}()

		e := v.BeginEdit()
		v.Insert(e, 0, test.text)
		v.EndEdit(e)

		v.Sel().Clear()
		v.Sel().AddAll(test.sel)

		ed.CommandHandler().RunTextCommand(v, "split_selection_into_lines", nil)
		// Comparing regions
		d := v.Sel()
		if d.Len() != len(test.expect) {
			t.Errorf("Test %d: Excepted '%d' regions, but got '%d' regions", i, len(test.expect), d.Len())
			t.Errorf("%+v, %+v", test.expect, d.Regions())
		} else {
			var found bool
			for _, r := range test.expect {
				found = false
				for _, r2 := range d.Regions() {
					if r2 == r {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Test %d:\nRegion %+v not found in view regions: %+v", i, r, d.Regions())
				}
			}
		}
	}
}
