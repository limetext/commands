// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"strings"
	"unicode"

	"github.com/limetext/backend"
)

type (
	// TitleCase Command transforms all selections
	// to be in Title Case.  For instance, the text:
	// "this is some sample text"
	// turns in to:
	// "This Is Some Sample Text".
	TitleCase struct {
		backend.DefaultCommand
	}

	// SwapCase Command transforms all selections
	// so that each character in the selection
	// is the opposite case.  For example, the text:
	// "Hello, World!"
	// turns in to:
	// "hELLO, wORLD!".
	SwapCase struct {
		backend.DefaultCommand
	}

	// UpperCase Command transforms all selections
	// so that each character in the selection
	// is in its upper case equivalent (if any.)
	UpperCase struct {
		backend.DefaultCommand
	}

	// LowerCase Command transforms all selections
	// so that each character in the selection
	// is in its lower case equivalent.
	LowerCase struct {
		backend.DefaultCommand
	}
)

// Run executes the TitleCase command.
func (c *TitleCase) Run(v *backend.View, e *backend.Edit) error {
	sel := v.Sel()
	for i := 0; i < sel.Len(); i++ {
		r := sel.Get(i)
		if r.Size() != 0 {
			t := v.Substr(r)
			v.Replace(e, r, strings.Title(t))
		}
	}
	return nil
}

// Run executes the SwapCase command.
func (c *SwapCase) Run(v *backend.View, e *backend.Edit) error {
	sel := v.Sel()
	for i := 0; i < sel.Len(); i++ {
		r := sel.Get(i)
		if r.Size() == 0 {
			continue
		}
		text := v.Substr(r)
		swapped := make([]rune, 0)
		for _, c := range text {
			if unicode.IsUpper(c) {
				swapped = append(swapped, unicode.ToLower(c))
			} else {
				swapped = append(swapped, unicode.ToUpper(c))
			}
		}
		v.Replace(e, r, string(swapped))
	}
	return nil
}

// Run executes the UpperCase command.
func (c *UpperCase) Run(v *backend.View, e *backend.Edit) error {
	sel := v.Sel()
	for i := 0; i < sel.Len(); i++ {
		r := sel.Get(i)
		if r.Size() != 0 {
			t := v.Substr(r)
			v.Replace(e, r, strings.ToUpper(t))
		}
	}
	return nil
}

// Run executes the LowerCase command.
func (c *LowerCase) Run(v *backend.View, e *backend.Edit) error {
	sel := v.Sel()
	for i := 0; i < sel.Len(); i++ {
		r := sel.Get(i)
		if r.Size() != 0 {
			t := v.Substr(r)
			v.Replace(e, r, strings.ToLower(t))
		}
	}
	return nil
}

func init() {
	register([]backend.Command{
		&TitleCase{},
		&SwapCase{},
		&UpperCase{},
		&LowerCase{},
	})
}
