// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"testing"

	"github.com/limetext/backend"
)

func TestClose(t *testing.T) {
	ed := backend.GetEditor()

	l := len(ed.Windows())
	w := ed.NewWindow()
	ed.CommandHandler().RunWindowCommand(w, "close", nil)
	if got, exp := len(ed.Windows()), l; got != exp {
		t.Error("When there is no view the window should close")
		t.Errorf("Expected %d windows, but got %d", exp, got)
	}

	l = len(w.Views())
	w.NewFile()
	ed.CommandHandler().RunWindowCommand(w, "close", nil)

	if len(w.Views()) != l {
		t.Errorf("Expected %d view, but got %d", l, len(w.Views()))
	}

	for _, v := range w.Views() {
		v.SetScratch(true)
		v.Close()
	}
}

func TestNextView(t *testing.T) {
	ed := backend.GetEditor()
	w := ed.NewWindow()
	defer w.Close()

	v0 := w.NewFile()
	defer func() {
		v0.SetScratch(true)
		v0.Close()
	}()

	v1 := w.NewFile()
	defer func() {
		v1.SetScratch(true)
		v1.Close()
	}()

	v2 := w.NewFile()
	defer func() {
		v2.SetScratch(true)
		v2.Close()
	}()

	v3 := w.NewFile()
	defer func() {
		v3.SetScratch(true)
		v3.Close()
	}()

	w.SetActiveView(v1)

	ed.CommandHandler().RunWindowCommand(w, "next_view", nil)

	av := w.ActiveView()
	if av != v2 {
		t.Error("Expected to get v2, but didn't")
	}

	w.SetActiveView(v3)

	ed.CommandHandler().RunWindowCommand(w, "next_view", nil)

	av = w.ActiveView()
	if av != v0 {
		t.Error("Expected to get v0, but didn't")
	}
}

func TestPrevView(t *testing.T) {
	ed := backend.GetEditor()
	w := ed.NewWindow()
	defer w.Close()

	v0 := w.NewFile()
	defer func() {
		v0.SetScratch(true)
		v0.Close()
	}()

	v1 := w.NewFile()
	defer func() {
		v1.SetScratch(true)
		v1.Close()
	}()

	v2 := w.NewFile()
	defer func() {
		v2.SetScratch(true)
		v2.Close()
	}()

	v3 := w.NewFile()
	defer func() {
		v3.SetScratch(true)
		v3.Close()
	}()

	w.SetActiveView(v2)

	ed.CommandHandler().RunWindowCommand(w, "prev_view", nil)

	av := w.ActiveView()
	if av != v1 {
		t.Error("Expected to get v1, but didn't")
	}

	w.SetActiveView(v0)

	ed.CommandHandler().RunWindowCommand(w, "prev_view", nil)

	av = w.ActiveView()
	if av != v3 {
		t.Error("Expected to get v3, but didn't")
	}
}
