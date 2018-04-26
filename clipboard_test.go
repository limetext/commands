// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"testing"

	. "github.com/limetext/backend"
	"github.com/limetext/text"
)

type (
	dummyClipboard struct {
		cachedText   string
		autoExpanded bool
	}

	copyTest struct {
		buf             string
		clip            string
		autoExpanded    bool
		regions         []text.Region
		expClip         string
		expAutoExpanded bool
		expBuf          string
	}
)

func (c *dummyClipboard) Get() (string, bool) {
	return c.cachedText, c.autoExpanded
}

func (c *dummyClipboard) Set(s string, b bool) {
	c.cachedText = s
	c.autoExpanded = b
}

// TODO: Also test where the cursors end up.
func runClipboardTest(command string, tests *[]copyTest, t *testing.T) {
	ed := GetEditor()

	w := ed.NewWindow()
	defer w.Close()

	cb := &dummyClipboard{}
	ed.UseClipboard(cb)

	for i, test := range *tests {
		v := w.NewFile()
		defer func() {
			v.SetScratch(true)
			v.Close()
		}()

		edit := v.BeginEdit()
		v.Insert(edit, 0, test.buf)
		v.EndEdit(edit)
		v.Sel().Clear()

		cb.Set(test.clip, test.autoExpanded)

		for _, r := range test.regions {
			v.Sel().Add(r)
		}

		ed.CommandHandler().RunTextCommand(v, command, nil)

		clip, auto := cb.Get()

		if clip != test.expClip {
			t.Errorf("Test %d: Expected clipboard to be %q, but got %q", i, test.expClip, clip)
		}

		if auto != test.expAutoExpanded {
			t.Errorf("Test %d: Expected the clipboard's auto-expanded flag to be %t, but got %q", i, test.expAutoExpanded, auto)
		}

		b := v.Substr(text.Region{A: 0, B: v.Size()})

		if b != test.expBuf {
			t.Errorf("Test %d: Expected buffer to be %q, but got %q", i, test.expBuf, b)
		}
	}
}

func TestCopy(t *testing.T) {
	tests := []copyTest{
		{
			"test string",
			"",
			false,
			[]text.Region{{1, 3}},
			"es",
			false,
			"test string",
		},
		{
			"test\nstring",
			"",
			false,
			[]text.Region{{3, 6}},
			"t\ns",
			false,
			"test\nstring",
		},
		{
			"test string",
			"",
			false,
			[]text.Region{{3, 3}},
			"test string\n",
			true,
			"test string",
		},
		{
			"test string",
			"",
			false,
			[]text.Region{{1, 3}, {5, 6}},
			"es\ns",
			false,
			"test string",
		},
		{
			"test\nstring",
			"",
			false,
			[]text.Region{{1, 3}, {5, 6}},
			"es\ns",
			false,
			"test\nstring",
		},
		{
			"test\nstring",
			"",
			false,
			[]text.Region{{1, 1}, {7, 7}},
			"test\n\nstring\n",
			true,
			"test\nstring",
		},
		{
			"test\nstring",
			"",
			false,
			[]text.Region{{3, 6}, {9, 10}},
			"t\ns\nn",
			false,
			"test\nstring",
		},
		{
			"test string",
			"",
			false,
			[]text.Region{{5, 6}, {1, 3}},
			"es\ns",
			false,
			"test string",
		},
		{
			"test string",
			"",
			false,
			[]text.Region{{1, 1}, {6, 7}},
			"t",
			false,
			"test string",
		},
		{
			"test string\n",
			"",
			false,
			[]text.Region{{0, 12}},
			"test string\n",
			false,
			"test string\n",
		},
		{
			"test\nstring\n",
			"",
			false,
			[]text.Region{{0, 12}},
			"test\nstring\n",
			false,
			"test\nstring\n",
		},
		{
			"test\nstring\n",
			"",
			false,
			[]text.Region{{0, 5}, {6, 12}},
			"test\n\ntring\n",
			false,
			"test\nstring\n",
		},
		{
			"test\n\nstring\n",
			"",
			false,
			[]text.Region{{0, 6}, {7, 13}},
			"test\n\n\ntring\n",
			false,
			"test\n\nstring\n",
		},
	}

	runClipboardTest("copy", &tests, t)
}

func TestCut(t *testing.T) {
	tests := []copyTest{
		{
			"test string",
			"",
			false,
			[]text.Region{{1, 3}},
			"es",
			false,
			"tt string",
		},
		{
			"test\nstring",
			"",
			false,
			[]text.Region{{3, 6}},
			"t\ns",
			false,
			"testring",
		},
		{
			"test string",
			"",
			false,
			[]text.Region{{3, 3}},
			"test string\n",
			true,
			"",
		},
		{
			"test string",
			"",
			false,
			[]text.Region{{5, 6}, {1, 3}},
			"es\ns",
			false,
			"tt tring",
		},
		{
			"test\nstring",
			"",
			false,
			[]text.Region{{1, 3}, {5, 6}},
			"es\ns",
			false,
			"tt\ntring",
		},
		{
			"test\nstring",
			"",
			false,
			[]text.Region{{1, 1}, {7, 7}},
			"test\n\nstring\n",
			true,
			"",
		},
		{
			"test\nstring",
			"",
			false,
			[]text.Region{{3, 6}, {9, 10}},
			"t\ns\nn",
			false,
			"testrig",
		},
		{
			"test string",
			"",
			false,
			[]text.Region{{5, 6}, {1, 3}},
			"es\ns",
			false,
			"tt tring",
		},
		{
			"test string",
			"",
			false,
			[]text.Region{{6, 7}, {1, 1}},
			"t",
			false,
			"",
		},
		{
			"test\nstring",
			"",
			false,
			[]text.Region{{1, 1}, {6, 7}},
			"t",
			false,
			"sring",
		},
		{
			"test string\n",
			"",
			false,
			[]text.Region{{0, 12}},
			"test string\n",
			false,
			"",
		},
		{
			"test\nstring\n",
			"",
			false,
			[]text.Region{{0, 12}},
			"test\nstring\n",
			false,
			"",
		},
		{
			"test\nstring\n",
			"",
			false,
			[]text.Region{{0, 5}, {6, 12}},
			"test\n\ntring\n",
			false,
			"s",
		},
		{
			"test\n\nstring\n",
			"",
			false,
			[]text.Region{{0, 6}, {7, 13}},
			"test\n\n\ntring\n",
			false,
			"s",
		},
	}

	runClipboardTest("cut", &tests, t)
}

// TODO: Maybe just test pastes as the reverse operation of a cut.
func TestPaste(t *testing.T) {
	tests := []copyTest{
		{
			"test string",
			"test",
			false,
			[]text.Region{{1, 1}},
			"test",
			false,
			"ttestest string",
		},
		{
			"test string",
			"test",
			false,
			[]text.Region{{1, 3}},
			"test",
			false,
			"ttestt string",
		},
		{
			"test\nstring",
			"test",
			false,
			[]text.Region{{3, 6}},
			"test",
			false,
			"testesttring",
		},
		{
			"test string",
			"test",
			false,
			[]text.Region{{1, 3}, {5, 6}},
			"test",
			false,
			"ttestt testtring",
		},
		{
			"test\nstring",
			"test",
			false,
			[]text.Region{{1, 3}, {5, 6}},
			"test",
			false,
			"ttestt\ntesttring",
		},
		{
			"test\nstring",
			"test",
			false,
			[]text.Region{{3, 6}, {9, 10}},
			"test",
			false,
			"testesttritestg",
		},
		{
			"test\nstring",
			"test",
			false,
			[]text.Region{{9, 10}, {3, 6}},
			"test",
			false,
			"testesttritestg",
		},
		{
			"test",
			"abc\n",
			true,
			[]text.Region{{1, 1}},
			"abc\n",
			true,
			"abc\ntest",
		},
		{
			"test",
			"abc\n",
			true,
			[]text.Region{{1, 2}},
			"abc\n",
			true,
			"abc\ntest",
		},
		{
			"test",
			"abc\n",
			true,
			[]text.Region{{1, 1}, {3, 3}},
			"abc\n",
			true,
			"abc\nabc\ntest",
		},
		{
			"test",
			"abc\n",
			true,
			[]text.Region{{1, 1}, {3, 3}},
			"abc\n",
			true,
			"abc\nabc\ntest",
		},
		{
			"test",
			"abc\n",
			true,
			[]text.Region{{1, 2}, {3, 3}},
			"abc\n",
			true,
			"abc\nabc\ntest",
		},
		{
			"test",
			"abc\n",
			true,
			[]text.Region{{1, 2}, {3, 4}},
			"abc\n",
			true,
			"abc\nabc\ntest",
		},
		{
			"test",
			"abc\n\ndef\n",
			true,
			[]text.Region{{1, 1}},
			"abc\n\ndef\n",
			true,
			"abc\n\ndef\ntest",
		},
		{
			"test",
			"abc\n\ndef\n",
			true,
			[]text.Region{{1, 2}},
			"abc\n\ndef\n",
			true,
			"abc\n\ndef\ntest",
		},
		{
			"test",
			"abc\n\ndef\n",
			true,
			[]text.Region{{1, 1}, {3, 3}},
			"abc\n\ndef\n",
			true,
			"abc\n\ndef\nabc\n\ndef\ntest",
		},
		{
			"test",
			"abc\n\ndef\n",
			true,
			[]text.Region{{1, 2}, {3, 3}},
			"abc\n\ndef\n",
			true,
			"abc\n\ndef\nabc\n\ndef\ntest",
		},
		{
			"test",
			"abc\n\ndef\n",
			true,
			[]text.Region{{1, 2}, {3, 4}},
			"abc\n\ndef\n",
			true,
			"abc\n\ndef\nabc\n\ndef\ntest",
		},
		{
			"test",
			"abc\ndef",
			false,
			[]text.Region{{1, 1}},
			"abc\ndef",
			false,
			"tabc\ndefest",
		},
		{
			"test",
			"abc\ndef",
			false,
			[]text.Region{{1, 2}},
			"abc\ndef",
			false,
			"tabc\ndefst",
		},
		{
			"test",
			"abc\ndef",
			false,
			[]text.Region{{1, 1}, {3, 3}},
			"abc\ndef",
			false,
			"tabcesdeft",
		},
		{
			"test",
			"abc\ndef",
			false,
			[]text.Region{{1, 2}, {3, 3}},
			"abc\ndef",
			false,
			"tabcsdeft",
		},
		{
			"test",
			"abc\ndef",
			false,
			[]text.Region{{1, 2}, {3, 4}},
			"abc\ndef",
			false,
			"tabcsdef",
		},
		{
			"test",
			"abc\ndef",
			false,
			[]text.Region{{1, 1}, {3, 3}, {4, 4}},
			"abc\ndef",
			false,
			"tabc\ndefesabc\ndeftabc\ndef",
		},
	}

	runClipboardTest("paste", &tests, t)
}
