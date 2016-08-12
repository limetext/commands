// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"path/filepath"
	"testing"

	"github.com/limetext/backend"
)

func TestNewFile(t *testing.T) {
	ed := backend.GetEditor()

	w := ed.NewWindow()
	defer w.Close()

	l := len(w.Views())

	ed.CommandHandler().RunWindowCommand(w, "new_file", nil)

	if len(w.Views()) != l+1 {
		t.Errorf("Expected %d views, but got %d", l+1, len(w.Views()))
	}

	for _, v := range w.Views() {
		v.SetScratch(true)
		v.Close()
	}
}

func TestOpenFile(t *testing.T) {
	var fe scfe
	const testPath = "file_test.go"
	fe.files = []string{testPath}

	ed := backend.GetEditor()
	ed.SetFrontend(&fe)
	w := ed.NewWindow()
	defer w.Close()

	l := len(w.Views())
	ed.CommandHandler().RunWindowCommand(w, "prompt_open_file", nil)
	if len(w.Views()) != l+1 {
		t.Fatalf("Expected %d views, but got %d", l+1, len(w.Views()))
	}
	exp, err := filepath.Abs(testPath)
	if err != nil {
		exp = testPath
	}
	if w.Views()[l].FileName() != exp {
		t.Errorf("Expected %s as FileName, but got %s", testPath, w.Views()[l].FileName())
	}
}
