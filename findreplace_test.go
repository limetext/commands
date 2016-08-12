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

type findTest struct {
	text string
	in   []text.Region
	exp  []text.Region
	fw   bool
}

func runFindTest(tests []findTest, t *testing.T, commands ...string) {
	ed := backend.GetEditor()
	w := ed.NewWindow()
	defer w.Close()

	v := w.NewFile()
	defer func() {
		v.SetScratch(true)
		v.Close()
	}()

	for i, test := range tests {
		e := v.BeginEdit()
		v.Insert(e, 0, test.text)
		v.EndEdit(e)
		v.Sel().Clear()
		for _, r := range test.in {
			v.Sel().Add(r)
		}

		v.Settings().Set("find_wrap", test.fw)

		for _, command := range commands {
			ed.CommandHandler().RunTextCommand(v, command, nil)
		}
		if sr := v.Sel().Regions(); !reflect.DeepEqual(sr, test.exp) {
			t.Errorf("Test %d: Expected %s, but got %s", i, test.exp, sr)
		}
		e = v.BeginEdit()
		v.Erase(e, text.Region{0, v.Size()})
		v.EndEdit(e)
	}
}

func TestFindUnderExpand(t *testing.T) {
	tests := []findTest{
		{
			"Hello World!\nTest123123\nAbrakadabra\n",
			[]text.Region{{0, 0}},
			[]text.Region{{0, 5}},
			true,
		},
		{
			"Hello World!\nTest123123\nAbrakadabra\n",
			[]text.Region{{19, 20}},
			[]text.Region{{19, 20}, {22, 23}},
			true,
		},
	}

	runFindTest(tests, t, "find_under_expand")
}

func TestFindNext(t *testing.T) {
	tests := []findTest{
		{
			"Hello World!\nTest123123\nAbrakadabra\n",
			[]text.Region{{17, 20}},
			[]text.Region{{17, 20}},
			true,
		},
		// test find_wrap setting true
		{
			"Hello World!\nTest123123\nAbrakadabra\n",
			[]text.Region{{21, 23}},
			[]text.Region{{18, 20}},
			true,
		},
		// test find_wrap setting false
		{
			"Hello World!\nTest123123\nAbrakadabra\n",
			[]text.Region{{21, 23}},
			[]text.Region{{21, 23}},
			false,
		},
	}

	runFindTest(tests, t, "find_under_expand", "find_next")
}

type replaceTest struct {
	cursors []text.Region
	in      string
	exp     string
	fw      bool
}

func runReplaceTest(tests []replaceTest, t *testing.T, commands ...string) {
	ed := backend.GetEditor()
	w := ed.NewWindow()
	defer w.Close()

	v := w.NewFile()
	defer func() {
		v.SetScratch(true)
		v.Close()
	}()

	for i, test := range tests {
		e := v.BeginEdit()
		v.Insert(e, 0, test.in)
		v.EndEdit(e)
		v.Sel().Clear()

		for _, r := range test.cursors {
			v.Sel().Add(r)
		}

		v.Settings().Set("find_wrap", test.fw)

		replaceText = "f"
		for _, command := range commands {
			ed.CommandHandler().RunTextCommand(v, command, nil)
		}
		if out := v.Substr(text.Region{0, v.Size()}); out != test.exp {
			t.Errorf("Test %d failed: %s, %+v", i, out, test)
		}
		e = v.BeginEdit()
		v.Erase(e, text.Region{0, v.Size()})
		v.EndEdit(e)
	}
}

func TestReplaceNext(t *testing.T) {
	tests := []replaceTest{
		{
			[]text.Region{{1, 1}, {2, 2}, {3, 3}},
			"abc abc bac abc abc",
			"abc f bac abc abc",
			true,
		},
		{
			[]text.Region{{0, 0}, {4, 4}, {8, 8}, {12, 13}},
			"abc abc bac abc abc",
			"abc abc bac abc f",
			true,
		},
		{
			[]text.Region{{12, 13}, {8, 8}, {4, 4}, {1, 0}},
			"abc abc bac abc abc",
			"abc abc bac abc f",
			true,
		},
		{
			[]text.Region{{15, 15}},
			"abc abc bac abc abc",
			"abc abc bac abc f",
			true,
		},
		{
			[]text.Region{{0, 0}},
			"abc abc bac abc abc",
			"abc f bac abc abc",
			true,
		},
		// test find_wrap setting true
		{
			[]text.Region{{16, 19}},
			"abc abc bac abc abc",
			"f abc bac abc abc",
			true,
		},
		// test find_wrap setting false
		{
			[]text.Region{{16, 19}},
			"abc abc bac abc abc",
			"abc abc bac abc abc",
			false,
		},
	}

	runReplaceTest(tests, t, "find_under_expand", "replace_next")
}
