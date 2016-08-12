// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"sort"
	"strings"

	"github.com/limetext/backend"
	"github.com/limetext/text"
)

type (
	// Copy copies the current selection to the clipboard. If there
	// are multiple selections, they are concatenated in order from
	// top to bottom of the file, separated by newlines
	Copy struct {
		backend.DefaultCommand
	}

	// Cut copies the current selection to the clipboard, removing it from the buffer.
	// If there are multiple selections, they are concatenated in order from top to
	// bottom of the file, separated by newlines.
	Cut struct {
		backend.DefaultCommand
	}

	// Paste pastes the contents of the clipboard, overwriting the current selection,
	// if any. If there are multiple selections, the clipboard is split into lines.
	// If the number of lines equals the number of selections, the lines are pasted
	// separately into each selection in order from top to bottom of the file.
	Paste struct {
		backend.DefaultCommand
	}
)

func getRegions(v *backend.View, cut bool) *text.RegionSet {
	rs := &text.RegionSet{}
	regions := v.Sel().Regions()
	sort.Sort(regionSorter(regions))
	rs.AddAll(regions)

	he, ae := rs.HasEmpty(), !rs.HasNonEmpty() || cut
	for _, r := range rs.Regions() {
		if ae && r.Empty() {
			rs.Add(v.FullLineR(r))
		} else if he && r.Empty() {
			rs.Subtract(r)
		}
	}

	return rs
}

func getSelSubstrs(v *backend.View, rs *text.RegionSet) []string {
	var add, s1 string
	s := make([]string, len(rs.Regions()))
	for i, r := range rs.Regions() {
		add = ""
		s1 = v.Substr(r)
		if !v.Sel().HasNonEmpty() && !strings.HasSuffix(s1, "\n") {
			add = "\n"
		}
		s[i] = s1 + add
	}
	return s
}

// Run executes the Copy command.
func (c *Copy) Run(v *backend.View, e *backend.Edit) error {
	rs := getRegions(v, false)
	s := getSelSubstrs(v, rs)

	backend.GetEditor().SetClipboard(strings.Join(s, "\n"))

	return nil
}

// Run executes the Cut command.
func (c *Cut) Run(v *backend.View, e *backend.Edit) error {
	s := getSelSubstrs(v, getRegions(v, false))

	rs := getRegions(v, true)
	regions := rs.Regions()
	sort.Sort(sort.Reverse(regionSorter(regions)))
	for _, r := range regions {
		v.Erase(e, r)
	}

	backend.GetEditor().SetClipboard(strings.Join(s, "\n"))

	return nil
}

// Run executes the Paste command
func (c *Paste) Run(v *backend.View, e *backend.Edit) error {
	// TODO: Paste the entire line on the line before the cursor if a
	//		 line was autocopied.

	ed := backend.GetEditor()

	rs := &text.RegionSet{}
	regions := v.Sel().Regions()
	sort.Sort(sort.Reverse(regionSorter(regions)))
	rs.AddAll(regions)
	for _, r := range rs.Regions() {
		v.Replace(e, r, ed.GetClipboard())
	}

	return nil
}

func init() {
	register([]backend.Command{
		&Copy{},
		&Cut{},
		&Paste{},
	})
}
