// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"testing"

	"github.com/limetext/backend"
)

func TestNewWindow(t *testing.T) {
	ed := backend.GetEditor()
	l := len(ed.Windows())
	ed.CommandHandler().RunWindowCommand(ed.ActiveWindow(), "new_window", nil)

	if len(ed.Windows()) != l+1 {
		t.Errorf("Expected %d window, but got %d", l+1, len(ed.Windows()))
	}
}

func TestCloseAll(t *testing.T) {
	ed := backend.GetEditor()

	w := ed.NewWindow()
	defer w.Close()

	ed.CommandHandler().RunWindowCommand(w, "new_file", nil)
	ed.CommandHandler().RunWindowCommand(w, "new_file", nil)
	ed.CommandHandler().RunWindowCommand(w, "new_file", nil)

	ed.CommandHandler().RunWindowCommand(w, "close_all", nil)

	if len(w.Views()) != 0 {
		t.Errorf("Expected no views, but got %d", len(w.Views()))
	}
}

func TestCloseWindow(t *testing.T) {
	ed := backend.GetEditor()
	w := ed.NewWindow()
	l := len(ed.Windows())
	ed.CommandHandler().RunWindowCommand(w, "close_window", nil)

	if len(ed.Windows()) != l-1 {
		t.Errorf("Expected %d window, but got %d", l-1, len(ed.Windows()))
	}
}

func TestNewAppWindow(t *testing.T) {
	ed := backend.GetEditor()
	l := len(ed.Windows())
	ed.CommandHandler().RunApplicationCommand("new_window", nil)

	if len(ed.Windows()) != l+1 {
		t.Errorf("Expected %d window, but got %d", l+1, len(ed.Windows()))
	}
}

func TestCloseAppWindow(t *testing.T) {
	ed := backend.GetEditor()
	_ = ed.NewWindow()
	l := len(ed.Windows())
	ed.CommandHandler().RunApplicationCommand("close_window", nil)

	if len(ed.Windows()) != l-1 {
		t.Errorf("Expected %d window, but got %d", l-1, len(ed.Windows()))
	}
}
