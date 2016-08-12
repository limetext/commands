// Copyright 2016 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"testing"

	"github.com/limetext/backend"
)

func TestIncreaseFontSize(t *testing.T) {
	ed := backend.GetEditor()
	w := ed.NewWindow()
	defer w.Close()
	w.Settings().Set("font_size", 12)

	previousFontSize := w.Settings().Int("font_size")
	ed.CommandHandler().RunWindowCommand(w, "increase_font_size", backend.Args{})
	if val := w.Settings().Int("font_size"); val != previousFontSize+1 {
		t.Errorf("increase_font_size should have increased the font by 1")
	}
}

func TestDecrementFontSize(t *testing.T) {
	ed := backend.GetEditor()
	w := ed.NewWindow()
	defer w.Close()
	w.Settings().Set("font_size", 12)

	previousFontSize := w.Settings().Int("font_size")
	ed.CommandHandler().RunWindowCommand(w, "decrease_font_size", backend.Args{})
	if val := w.Settings().Int("font_size"); val != previousFontSize-1 {
		t.Errorf("increase_font_size should have increased the font by 1")
	}
}
