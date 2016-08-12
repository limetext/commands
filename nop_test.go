// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"testing"

	"github.com/limetext/backend"
)

func TestRunApplication(t *testing.T) {
	nopApplication := NopApplication{}

	if nopApplication.Run() != nil {
		t.Error("No op application command running returns not nil")
	}

}

func TestRunNopWindow(t *testing.T) {
	nopWindow := NopWindow{}

	if nopWindow.Run(&backend.Window{}) != nil {
		t.Error("No op window command running returns not nil")
	}
}

func TestRunNopText(t *testing.T) {
	nopText := NopText{}

	if nopText.Run(&backend.View{}, &backend.Edit{}) != nil {
		t.Error("No op text command running returns not nil")
	}
}
