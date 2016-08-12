// Copyright 2016 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"os"
	"testing"

	"github.com/limetext/backend"
)

func TestSaveProjectAs(t *testing.T) {
	const testfile = "testdata/save_project_as"
	defer os.Remove(testfile)

	var fe scfe
	fe.files = []string{testfile}
	ed := backend.GetEditor()
	ed.SetFrontend(&fe)
	w := ed.NewWindow()
	defer w.Close()

	if err := ed.CommandHandler().RunWindowCommand(w, "save_project_as", nil); err != nil {
		t.Errorf("Error running 'save_project_as' command: %s", err)
	}
	if _, err := os.Stat(testfile); err != nil {
		t.Errorf("Expected %s to exist: %s", testfile, err)
	}
}

func TestPromptAddFolder(t *testing.T) {
	const testfolder = "testdata"
	var fe scfe
	fe.files = []string{testfolder}

	ed := backend.GetEditor()
	ed.SetFrontend(&fe)
	w := ed.NewWindow()
	defer w.Close()

	if err := ed.CommandHandler().RunWindowCommand(w, "prompt_add_folder", nil); err != nil {
		t.Errorf("Error running 'prompt_add_folder' command: %s", err)
	}
	folders := w.Project().Folders()
	if len(folders) == 0 {
		t.Fatal("Expected 1 folder in project folders list")
	}
	if got := folders[0]; got != testfolder {
		t.Errorf("Expected %s in project folders, but got %s", testfolder, got)
	}
}

func TestCloseFolderList(t *testing.T) {
	ed := backend.GetEditor()
	w := ed.NewWindow()
	defer w.Close()
	p := w.Project()
	p.AddFolder("testdata")
	p.AddFolder(".")

	if err := ed.CommandHandler().RunWindowCommand(w, "close_folder_list", nil); err != nil {
		t.Errorf("Error running 'close_folder_list' command: %s", err)
	}
	if len(p.Folders()) != 0 {
		t.Errorf("Expected project folders list be empty, but got %s", p.Folders())
	}
}
