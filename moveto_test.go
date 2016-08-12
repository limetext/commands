// Copyright 2016 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"reflect"
	"testing"

	"github.com/limetext/backend"
	"github.com/limetext/text"
)

type MoveToTest struct {
	in     []text.Region
	to     string
	extend bool
	exp    []text.Region
}

func runMoveToTest(tests []MoveToTest, t *testing.T, text string) {
	ed := backend.GetEditor()
	w := ed.NewWindow()
	defer w.Close()

	v := w.NewFile()
	defer func() {
		v.SetScratch(true)
		v.Close()
	}()

	e := v.BeginEdit()

	v.Insert(e, 0, text)
	v.EndEdit(e)

	for i, test := range tests {
		v.Sel().Clear()
		for _, r := range test.in {
			v.Sel().Add(r)
		}
		args := backend.Args{"to": test.to, "extend": test.extend}
		ed.CommandHandler().RunTextCommand(v, "move_to", args)
		if sr := v.Sel().Regions(); !reflect.DeepEqual(sr, test.exp) {
			t.Errorf("Test %d failed. Expected %v, but got %v: %+v", i, test.exp, sr, test)
		}
	}
}

func TestMoveTo(t *testing.T) {
	/*
	   Correct behavior of MoveTo:
	       - Moves each cursor directly to the indicated position.
	       - If extend, the selection will be extended in the direction of movement
	*/

	singleCursor := []text.Region{{16, 16}}

	sameLineCursors := []text.Region{{16, 16}, {17, 17}}
	sameLineCursorsReversed := []text.Region{{17, 17}, {16, 16}}

	diffLineCursors := []text.Region{{3, 3}, {17, 17}}
	diffLineCursorsReversed := []text.Region{{17, 17}, {3, 3}}

	singleForwardSelection := []text.Region{{15, 18}}
	singleBackwardSelection := []text.Region{{18, 15}}

	sameLineForwardSelections := []text.Region{{15, 18}, {20, 21}}
	sameLineForwardSelectionsReversed := []text.Region{{20, 21}, {15, 18}}
	sameLineBackwardSelections := []text.Region{{18, 15}, {21, 20}}
	sameLineBackwardSelectionsReversed := []text.Region{{21, 20}, {18, 15}}
	sameLineForwardThenBackwardSelections := []text.Region{{15, 18}, {21, 20}}
	sameLineForwardThenBackwardSelectionsReversed := []text.Region{{21, 20}, {15, 18}}
	sameLineBackwardThenForwardSelections := []text.Region{{18, 15}, {20, 21}}
	sameLineBackwardThenForwardSelectionsReversed := []text.Region{{20, 21}, {18, 15}}

	diffLineForwardSelections := []text.Region{{4, 6}, {20, 21}}
	diffLineForwardSelectionsReversed := []text.Region{{20, 21}, {4, 6}}
	diffLineBackwardSelections := []text.Region{{6, 4}, {21, 20}}
	diffLineBackwardSelectionsReversed := []text.Region{{21, 20}, {6, 4}}
	diffLineForwardThenBackwardSelections := []text.Region{{4, 6}, {21, 20}}
	diffLineForwardThenBackwardSelectionsReversed := []text.Region{{21, 20}, {4, 6}}
	diffLineBackwardThenForwardSelections := []text.Region{{6, 4}, {20, 21}}
	diffLineBackwardThenForwardSelectionsReversed := []text.Region{{20, 21}, {6, 4}}

	inputText := "Hello World!\nTest123123\nAbrakadabra\n"
	vbufflen := 36

	tests := []MoveToTest{
		// BOF move
		{
			singleCursor,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			sameLineCursors,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			sameLineCursorsReversed,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			diffLineCursors,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			diffLineCursorsReversed,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			singleForwardSelection,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			singleBackwardSelection,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			sameLineForwardSelections,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			sameLineForwardSelectionsReversed,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			sameLineBackwardSelections,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			sameLineBackwardSelectionsReversed,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			sameLineForwardThenBackwardSelections,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			sameLineForwardThenBackwardSelectionsReversed,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			sameLineBackwardThenForwardSelections,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			sameLineBackwardThenForwardSelectionsReversed,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			diffLineForwardSelections,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			diffLineForwardSelectionsReversed,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			diffLineBackwardSelections,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			diffLineBackwardSelectionsReversed,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			diffLineForwardThenBackwardSelections,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			diffLineForwardThenBackwardSelectionsReversed,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			diffLineBackwardThenForwardSelections,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},
		{
			diffLineBackwardThenForwardSelectionsReversed,
			"bof",
			false,
			[]text.Region{{0, 0}},
		},

		// BOF extend
		{
			singleCursor,
			"bof",
			true,
			[]text.Region{{16, 0}},
		},
		{
			sameLineCursors,
			"bof",
			true,
			[]text.Region{{17, 0}},
		},
		{
			sameLineCursorsReversed,
			"bof",
			true,
			[]text.Region{{17, 0}},
		},
		{
			diffLineCursors,
			"bof",
			true,
			[]text.Region{{17, 0}},
		},
		{
			diffLineCursorsReversed,
			"bof",
			true,
			[]text.Region{{17, 0}},
		},
		{
			singleForwardSelection,
			"bof",
			true,
			[]text.Region{{15, 0}},
		},
		{
			singleBackwardSelection,
			"bof",
			true,
			[]text.Region{{18, 0}},
		},
		{
			sameLineForwardSelections,
			"bof",
			true,
			[]text.Region{{20, 0}},
		},
		{
			sameLineForwardSelectionsReversed,
			"bof",
			true,
			[]text.Region{{20, 0}},
		},
		{
			sameLineBackwardSelections,
			"bof",
			true,
			[]text.Region{{21, 0}},
		},
		{
			sameLineBackwardSelectionsReversed,
			"bof",
			true,
			[]text.Region{{21, 0}},
		},
		{
			sameLineForwardThenBackwardSelections,
			"bof",
			true,
			[]text.Region{{21, 0}},
		},
		{
			sameLineForwardThenBackwardSelectionsReversed,
			"bof",
			true,
			[]text.Region{{21, 0}},
		},
		{
			sameLineBackwardThenForwardSelections,
			"bof",
			true,
			[]text.Region{{20, 0}},
		},
		{
			sameLineBackwardThenForwardSelectionsReversed,
			"bof",
			true,
			[]text.Region{{20, 0}},
		},
		{
			diffLineForwardSelections,
			"bof",
			true,
			[]text.Region{{20, 0}},
		},
		{
			diffLineForwardSelectionsReversed,
			"bof",
			true,
			[]text.Region{{20, 0}},
		},
		{
			diffLineBackwardSelections,
			"bof",
			true,
			[]text.Region{{21, 0}},
		},
		{
			diffLineBackwardSelectionsReversed,
			"bof",
			true,
			[]text.Region{{21, 0}},
		},
		{
			diffLineForwardThenBackwardSelections,
			"bof",
			true,
			[]text.Region{{21, 0}},
		},
		{
			diffLineForwardThenBackwardSelectionsReversed,
			"bof",
			true,
			[]text.Region{{21, 0}},
		},
		{
			diffLineBackwardThenForwardSelections,
			"bof",
			true,
			[]text.Region{{20, 0}},
		},
		{
			diffLineBackwardThenForwardSelectionsReversed,
			"bof",
			true,
			[]text.Region{{20, 0}},
		},

		// EOF move
		{
			singleCursor,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			sameLineCursors,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			sameLineCursorsReversed,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			diffLineCursors,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			diffLineCursorsReversed,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			singleForwardSelection,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			singleBackwardSelection,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			sameLineForwardSelections,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			sameLineForwardSelectionsReversed,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			sameLineBackwardSelections,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			sameLineBackwardSelectionsReversed,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			sameLineForwardThenBackwardSelections,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			sameLineForwardThenBackwardSelectionsReversed,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			sameLineBackwardThenForwardSelections,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			sameLineBackwardThenForwardSelectionsReversed,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			diffLineForwardSelections,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			diffLineForwardSelectionsReversed,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			diffLineBackwardSelections,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			diffLineBackwardSelectionsReversed,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			diffLineForwardThenBackwardSelections,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			diffLineForwardThenBackwardSelectionsReversed,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			diffLineBackwardThenForwardSelections,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},
		{
			diffLineBackwardThenForwardSelectionsReversed,
			"eof",
			false,
			[]text.Region{{vbufflen, vbufflen}},
		},

		// EOF extend
		{
			singleCursor,
			"eof",
			true,
			[]text.Region{{16, vbufflen}},
		},
		{
			sameLineCursors,
			"eof",
			true,
			[]text.Region{{16, vbufflen}},
		},
		{
			sameLineCursorsReversed,
			"eof",
			true,
			[]text.Region{{16, vbufflen}},
		},
		{
			diffLineCursors,
			"eof",
			true,
			[]text.Region{{3, vbufflen}},
		},
		{
			diffLineCursorsReversed,
			"eof",
			true,
			[]text.Region{{3, vbufflen}},
		},
		{
			singleForwardSelection,
			"eof",
			true,
			[]text.Region{{15, vbufflen}},
		},
		{
			singleBackwardSelection,
			"eof",
			true,
			[]text.Region{{18, vbufflen}},
		},
		{
			sameLineForwardSelections,
			"eof",
			true,
			[]text.Region{{15, vbufflen}},
		},
		{
			sameLineForwardSelectionsReversed,
			"eof",
			true,
			[]text.Region{{15, vbufflen}},
		},
		{
			sameLineBackwardSelections,
			"eof",
			true,
			[]text.Region{{18, vbufflen}},
		},
		{
			sameLineBackwardSelectionsReversed,
			"eof",
			true,
			[]text.Region{{18, vbufflen}},
		},
		{
			sameLineForwardThenBackwardSelections,
			"eof",
			true,
			[]text.Region{{15, vbufflen}},
		},
		{
			sameLineForwardThenBackwardSelectionsReversed,
			"eof",
			true,
			[]text.Region{{15, vbufflen}},
		},
		{
			sameLineBackwardThenForwardSelections,
			"eof",
			true,
			[]text.Region{{18, vbufflen}},
		},
		{
			sameLineBackwardThenForwardSelectionsReversed,
			"eof",
			true,
			[]text.Region{{18, vbufflen}},
		},
		{
			diffLineForwardSelections,
			"eof",
			true,
			[]text.Region{{4, vbufflen}},
		},
		{
			diffLineForwardSelectionsReversed,
			"eof",
			true,
			[]text.Region{{4, vbufflen}},
		},
		{
			diffLineBackwardSelections,
			"eof",
			true,
			[]text.Region{{6, vbufflen}},
		},
		{
			diffLineBackwardSelectionsReversed,
			"eof",
			true,
			[]text.Region{{6, vbufflen}},
		},
		{
			diffLineForwardThenBackwardSelections,
			"eof",
			true,
			[]text.Region{{4, vbufflen}},
		},
		{
			diffLineForwardThenBackwardSelectionsReversed,
			"eof",
			true,
			[]text.Region{{4, vbufflen}},
		},
		{
			diffLineBackwardThenForwardSelections,
			"eof",
			true,
			[]text.Region{{6, vbufflen}},
		},
		{
			diffLineBackwardThenForwardSelectionsReversed,
			"eof",
			true,
			[]text.Region{{6, vbufflen}},
		},

		// BOL move
		{
			singleCursor,
			"bol",
			false,
			[]text.Region{{13, 13}},
		},
		{
			sameLineCursors,
			"bol",
			false,
			[]text.Region{{13, 13}},
		},
		{
			sameLineCursorsReversed,
			"bol",
			false,
			[]text.Region{{13, 13}},
		},
		{
			diffLineCursors,
			"bol",
			false,
			[]text.Region{{0, 0}, {13, 13}},
		},
		{
			diffLineCursorsReversed,
			"bol",
			false,
			[]text.Region{{13, 13}, {0, 0}},
		},
		{
			singleForwardSelection,
			"bol",
			false,
			[]text.Region{{13, 13}},
		},
		{
			singleBackwardSelection,
			"bol",
			false,
			[]text.Region{{13, 13}},
		},
		{
			sameLineForwardSelections,
			"bol",
			false,
			[]text.Region{{13, 13}},
		},
		{
			sameLineForwardSelectionsReversed,
			"bol",
			false,
			[]text.Region{{13, 13}},
		},
		{
			sameLineBackwardSelections,
			"bol",
			false,
			[]text.Region{{13, 13}},
		},
		{
			sameLineBackwardSelectionsReversed,
			"bol",
			false,
			[]text.Region{{13, 13}},
		},
		{
			sameLineForwardThenBackwardSelections,
			"bol",
			false,
			[]text.Region{{13, 13}},
		},
		{
			sameLineForwardThenBackwardSelectionsReversed,
			"bol",
			false,
			[]text.Region{{13, 13}},
		},
		{
			sameLineBackwardThenForwardSelections,
			"bol",
			false,
			[]text.Region{{13, 13}},
		},
		{
			sameLineBackwardThenForwardSelectionsReversed,
			"bol",
			false,
			[]text.Region{{13, 13}},
		},
		{
			diffLineForwardSelections,
			"bol",
			false,
			[]text.Region{{0, 0}, {13, 13}},
		},
		{
			diffLineForwardSelectionsReversed,
			"bol",
			false,
			[]text.Region{{13, 13}, {0, 0}},
		},
		{
			diffLineBackwardSelections,
			"bol",
			false,
			[]text.Region{{0, 0}, {13, 13}},
		},
		{
			diffLineBackwardSelectionsReversed,
			"bol",
			false,
			[]text.Region{{13, 13}, {0, 0}},
		},
		{
			diffLineForwardThenBackwardSelections,
			"bol",
			false,
			[]text.Region{{0, 0}, {13, 13}},
		},
		{
			diffLineForwardThenBackwardSelectionsReversed,
			"bol",
			false,
			[]text.Region{{13, 13}, {0, 0}},
		},
		{
			diffLineBackwardThenForwardSelections,
			"bol",
			false,
			[]text.Region{{0, 0}, {13, 13}},
		},
		{
			diffLineBackwardThenForwardSelectionsReversed,
			"bol",
			false,
			[]text.Region{{13, 13}, {0, 0}},
		},

		// BOL extend
		{
			singleCursor,
			"bol",
			true,
			[]text.Region{{16, 13}},
		},
		{
			sameLineCursors,
			"bol",
			true,
			[]text.Region{{17, 13}},
		},
		{
			sameLineCursorsReversed,
			"bol",
			true,
			[]text.Region{{17, 13}},
		},
		{
			diffLineCursors,
			"bol",
			true,
			[]text.Region{{3, 0}, {17, 13}},
		},
		{
			diffLineCursorsReversed,
			"bol",
			true,
			[]text.Region{{17, 13}, {3, 0}},
		},
		{
			singleForwardSelection,
			"bol",
			true,
			[]text.Region{{15, 13}},
		},
		{
			singleBackwardSelection,
			"bol",
			true,
			[]text.Region{{18, 13}},
		},
		{
			sameLineForwardSelections,
			"bol",
			true,
			[]text.Region{{20, 13}},
		},
		{
			sameLineForwardSelectionsReversed,
			"bol",
			true,
			[]text.Region{{20, 13}},
		},
		{
			sameLineBackwardSelections,
			"bol",
			true,
			[]text.Region{{21, 13}},
		},
		{
			sameLineBackwardSelectionsReversed,
			"bol",
			true,
			[]text.Region{{21, 13}},
		},
		{
			sameLineForwardThenBackwardSelections,
			"bol",
			true,
			[]text.Region{{21, 13}},
		},
		{
			sameLineForwardThenBackwardSelectionsReversed,
			"bol",
			true,
			[]text.Region{{21, 13}},
		},
		{
			sameLineBackwardThenForwardSelections,
			"bol",
			true,
			[]text.Region{{20, 13}},
		},
		{
			sameLineBackwardThenForwardSelectionsReversed,
			"bol",
			true,
			[]text.Region{{20, 13}},
		},
		{
			diffLineForwardSelections,
			"bol",
			true,
			[]text.Region{{4, 0}, {20, 13}},
		},
		{
			diffLineForwardSelectionsReversed,
			"bol",
			true,
			[]text.Region{{20, 13}, {4, 0}},
		},
		{
			diffLineBackwardSelections,
			"bol",
			true,
			[]text.Region{{6, 0}, {21, 13}},
		},
		{
			diffLineBackwardSelectionsReversed,
			"bol",
			true,
			[]text.Region{{21, 13}, {6, 0}},
		},
		{
			diffLineForwardThenBackwardSelections,
			"bol",
			true,
			[]text.Region{{4, 0}, {21, 13}},
		},
		{
			diffLineForwardThenBackwardSelectionsReversed,
			"bol",
			true,
			[]text.Region{{21, 13}, {4, 0}},
		},
		{
			diffLineBackwardThenForwardSelections,
			"bol",
			true,
			[]text.Region{{6, 0}, {20, 13}},
		},
		{
			diffLineBackwardThenForwardSelectionsReversed,
			"bol",
			true,
			[]text.Region{{20, 13}, {6, 0}},
		},

		// EOL move
		{
			singleCursor,
			"eol",
			false,
			[]text.Region{{23, 23}},
		},
		{
			sameLineCursors,
			"eol",
			false,
			[]text.Region{{23, 23}},
		},
		{
			sameLineCursorsReversed,
			"eol",
			false,
			[]text.Region{{23, 23}},
		},
		{
			diffLineCursors,
			"eol",
			false,
			[]text.Region{{12, 12}, {23, 23}},
		},
		{
			diffLineCursorsReversed,
			"eol",
			false,
			[]text.Region{{23, 23}, {12, 12}},
		},
		{
			singleForwardSelection,
			"eol",
			false,
			[]text.Region{{23, 23}},
		},
		{
			singleBackwardSelection,
			"eol",
			false,
			[]text.Region{{23, 23}},
		},
		{
			sameLineForwardSelections,
			"eol",
			false,
			[]text.Region{{23, 23}},
		},
		{
			sameLineForwardSelectionsReversed,
			"eol",
			false,
			[]text.Region{{23, 23}},
		},
		{
			sameLineBackwardSelections,
			"eol",
			false,
			[]text.Region{{23, 23}},
		},
		{
			sameLineBackwardSelectionsReversed,
			"eol",
			false,
			[]text.Region{{23, 23}},
		},
		{
			sameLineForwardThenBackwardSelections,
			"eol",
			false,
			[]text.Region{{23, 23}},
		},
		{
			sameLineForwardThenBackwardSelectionsReversed,
			"eol",
			false,
			[]text.Region{{23, 23}},
		},
		{
			sameLineBackwardThenForwardSelections,
			"eol",
			false,
			[]text.Region{{23, 23}},
		},
		{
			sameLineBackwardThenForwardSelectionsReversed,
			"eol",
			false,
			[]text.Region{{23, 23}},
		},
		{
			diffLineForwardSelections,
			"eol",
			false,
			[]text.Region{{12, 12}, {23, 23}},
		},
		{
			diffLineForwardSelectionsReversed,
			"eol",
			false,
			[]text.Region{{23, 23}, {12, 12}},
		},
		{
			diffLineBackwardSelections,
			"eol",
			false,
			[]text.Region{{12, 12}, {23, 23}},
		},
		{
			diffLineBackwardSelectionsReversed,
			"eol",
			false,
			[]text.Region{{23, 23}, {12, 12}},
		},
		{
			diffLineForwardThenBackwardSelections,
			"eol",
			false,
			[]text.Region{{12, 12}, {23, 23}},
		},
		{
			diffLineForwardThenBackwardSelectionsReversed,
			"eol",
			false,
			[]text.Region{{23, 23}, {12, 12}},
		},
		{
			diffLineBackwardThenForwardSelections,
			"eol",
			false,
			[]text.Region{{12, 12}, {23, 23}},
		},
		{
			diffLineBackwardThenForwardSelectionsReversed,
			"eol",
			false,
			[]text.Region{{23, 23}, {12, 12}},
		},

		// EOL extend
		{
			singleCursor,
			"eol",
			true,
			[]text.Region{{16, 23}},
		},
		{
			sameLineCursors,
			"eol",
			true,
			[]text.Region{{16, 23}},
		},
		{
			sameLineCursorsReversed,
			"eol",
			true,
			[]text.Region{{16, 23}},
		},
		{
			diffLineCursors,
			"eol",
			true,
			[]text.Region{{3, 12}, {17, 23}},
		},
		{
			diffLineCursorsReversed,
			"eol",
			true,
			[]text.Region{{17, 23}, {3, 12}},
		},
		{
			singleForwardSelection,
			"eol",
			true,
			[]text.Region{{15, 23}},
		},
		{
			singleBackwardSelection,
			"eol",
			true,
			[]text.Region{{18, 23}},
		},
		{
			sameLineForwardSelections,
			"eol",
			true,
			[]text.Region{{15, 23}},
		},
		{
			sameLineForwardSelectionsReversed,
			"eol",
			true,
			[]text.Region{{15, 23}},
		},
		{
			sameLineBackwardSelections,
			"eol",
			true,
			[]text.Region{{18, 23}},
		},
		{
			sameLineBackwardSelectionsReversed,
			"eol",
			true,
			[]text.Region{{18, 23}},
		},
		{
			sameLineForwardThenBackwardSelections,
			"eol",
			true,
			[]text.Region{{15, 23}},
		},
		{
			sameLineForwardThenBackwardSelectionsReversed,
			"eol",
			true,
			[]text.Region{{15, 23}},
		},
		{
			sameLineBackwardThenForwardSelections,
			"eol",
			true,
			[]text.Region{{18, 23}},
		},
		{
			sameLineBackwardThenForwardSelectionsReversed,
			"eol",
			true,
			[]text.Region{{18, 23}},
		},
		{
			diffLineForwardSelections,
			"eol",
			true,
			[]text.Region{{4, 12}, {20, 23}},
		},
		{
			diffLineForwardSelectionsReversed,
			"eol",
			true,
			[]text.Region{{20, 23}, {4, 12}},
		},
		{
			diffLineBackwardSelections,
			"eol",
			true,
			[]text.Region{{6, 12}, {21, 23}},
		},
		{
			diffLineBackwardSelectionsReversed,
			"eol",
			true,
			[]text.Region{{21, 23}, {6, 12}},
		},
		{
			diffLineForwardThenBackwardSelections,
			"eol",
			true,
			[]text.Region{{4, 12}, {21, 23}},
		},
		{
			diffLineForwardThenBackwardSelectionsReversed,
			"eol",
			true,
			[]text.Region{{21, 23}, {4, 12}},
		},
		{
			diffLineBackwardThenForwardSelections,
			"eol",
			true,
			[]text.Region{{6, 12}, {20, 23}},
		},
		{
			diffLineBackwardThenForwardSelectionsReversed,
			"eol",
			true,
			[]text.Region{{20, 23}, {6, 12}},
		},
	}

	runMoveToTest(tests, t, inputText)

	tests = []MoveToTest{
		// Brackets move
		{
			[]text.Region{{1, 1}},
			"brackets",
			false,
			[]text.Region{{1, 1}},
		},
		{
			[]text.Region{{7, 7}},
			"brackets",
			false,
			[]text.Region{{31, 31}},
		},
		{
			[]text.Region{{5, 5}},
			"brackets",
			false,
			[]text.Region{{32, 32}},
		},
		{
			[]text.Region{{9, 9}, {14, 14}},
			"brackets",
			false,
			[]text.Region{{31, 31}, {16, 16}},
		},
		{
			[]text.Region{{16, 16}},
			"brackets",
			false,
			[]text.Region{{13, 13}},
		},
		{
			[]text.Region{{32, 32}},
			"brackets",
			false,
			[]text.Region{{5, 5}},
		},
		{
			[]text.Region{{32, 12}},
			"brackets",
			false,
			[]text.Region{{17, 17}},
		},
		{
			[]text.Region{{35, 35}},
			"brackets",
			false,
			// Sublime text 3 result is []text.Region{{35, 35}}
			// but i think this is a bug
			[]text.Region{{38, 38}},
		},
		{
			[]text.Region{{8, 9}, {10, 10}},
			"brackets",
			false,
			[]text.Region{{31, 31}},
		},
		{
			[]text.Region{{33, 33}},
			"brackets",
			false,
			// Sublime text 3 result is []text.Region{{33, 33}}
			// but i think this is a bug
			[]text.Region{{39, 39}},
		},
		{
			[]text.Region{{36, 36}},
			"brackets",
			false,
			[]text.Region{{36, 36}},
		},
		{
			[]text.Region{{38, 38}},
			"brackets",
			false,
			[]text.Region{{34, 34}},
		},
		// Breackets extend
		{
			[]text.Region{{8, 9}, {10, 10}},
			"brackets",
			true,
			[]text.Region{{8, 31}},
		},
		{
			[]text.Region{{17, 17}},
			"brackets",
			true,
			[]text.Region{{17, 24}},
		},
		{
			[]text.Region{{1, 1}, {14, 14}},
			"brackets",
			true,
			[]text.Region{{1, 1}, {14, 16}},
		},
	}

	runMoveToTest(tests, t, "test (moveto(abc){\ntada}Test123)1{23)a}baraba")
}
