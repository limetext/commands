// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import . "github.com/limetext/backend"

type (
	// The ToggleSettingCommand toggles the value of a setting,
	// making it false when it was true or true when it was false.
	ToggleSetting struct {
		BypassUndoCommand
		Setting string
	}

	// The SetSettingCommand set the value of a setting.
	SetSetting struct {
		BypassUndoCommand
		Setting string
		Value   interface{}
	}

	ToggleSideBar struct {
		toggleSetting
	}

	ToggleStatusBar struct {
		toggleSetting
	}

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

// helper struct for commands that just toggle a simple setting
type toggleSetting struct {
	// the setting name which we are going to toggle
	name string
	BypassUndoCommand
}

func (c *ToggleSetting) Run(v *View, e *Edit) error {
	setting := c.Setting
	prev, boolean := v.Settings().Get(setting, false).(bool)
	// if the setting was non-boolean, it is set to true, else it is toggled
	v.Settings().Set(setting, !boolean || !prev)
	return nil
}

func (c *SetSetting) Run(v *View, e *Edit) error {
	setting := c.Setting
	v.Settings().Set(setting, c.Value)
	return nil
}

func (t *toggleSetting) Run(w *Window) error {
	res, ok := w.Settings().Get(t.name, false).(bool)
	w.Settings().Set(t.name, !ok || !res)
	return nil
}

func toggle(name string) toggleSetting {
	return toggleSetting{name: name}
}

func init() {
	register([]Command{
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
