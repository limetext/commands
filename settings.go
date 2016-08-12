// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import "github.com/limetext/backend"

type (
	// ToggleSetting Command toggles the value of a setting,
	// making it false when it was true or true when it was false.
	ToggleSetting struct {
		backend.BypassUndoCommand
		Setting string
	}

	// SetSetting Command set the value of a setting.
	SetSetting struct {
		backend.BypassUndoCommand
		Setting string
		Value   interface{}
	}

	// ToggleSideBar Command enables us to toggle the sidebar
	// when the sidebar is visible, it'll be made invisible
	// and vice versa.
	ToggleSideBar struct {
		toggleSetting
	}

	// ToggleStatusBar command enables us to toggle the status bar.
	ToggleStatusBar struct {
		toggleSetting
	}

	// ToggleFullScreen command enables us to toggle full screen.
	ToggleFullScreen struct {
		toggleSetting
	}

	ToggleDistractionFree struct {
		toggleSetting
	}

	ToggleMinimap struct {
		toggleSetting
	}

	ToggleTabs struct {
		toggleSetting
	}
)

// helper struct for commands that just toggle a simple setting.
type toggleSetting struct {
	// the setting name which we are going to toggle.
	name string
	backend.BypassUndoCommand
}

// Run executes the ToggleSetting command.
func (c *ToggleSetting) Run(v *backend.View, e *backend.Edit) error {
	setting := c.Setting
	prev, boolean := v.Settings().Get(setting, false).(bool)
	// if the setting was non-boolean, it is set to true, else it is toggled
	v.Settings().Set(setting, !boolean || !prev)
	return nil
}

// Run executes the SetSetting command.
func (c *SetSetting) Run(v *backend.View, e *backend.Edit) error {
	setting := c.Setting
	v.Settings().Set(setting, c.Value)
	return nil
}

func (t *toggleSetting) Run(w *backend.Window) error {
	res, ok := w.Settings().Get(t.name, false).(bool)
	w.Settings().Set(t.name, !ok || !res)
	return nil
}

func toggle(name string) toggleSetting {
	return toggleSetting{name: name}
}

func init() {
	register([]backend.Command{
		&ToggleSetting{},
		&SetSetting{},
		&ToggleSideBar{toggle("show_side_bar")},
		&ToggleStatusBar{toggle("show_status_bar")},
		&ToggleFullScreen{toggle("show_full_screen")},
		&ToggleDistractionFree{toggle("show_distraction_free")},
		&ToggleMinimap{toggle("show_minimap")},
		&ToggleTabs{toggle("show_tabs")},
	})
}
