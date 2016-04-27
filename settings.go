// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	. "github.com/limetext/backend"
)

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
		BypassUndoCommand
	}

	ToggleStatuseBar struct {
		BypassUndoCommand
	}

	ToggleFullScreen struct {
		BypassUndoCommand
	}

	ToggleDsitractionFree struct {
		BypassUndoCommand
	}
)

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

func (c *ToggleSideBar) Run(w *Window) error {
	res, ok := w.Settings().Get("toggle_sidebar", false).(bool)
	w.Settings().Set("toggle_sidebar", !ok || !res)
	return nil
}

func (c *ToggleStatuseBar) Run(w *Window) error {
	res, ok := w.Settings().Get("toggle_status_bar", false).(bool)
	w.Settings().Set("toggle_status_bar", !ok || !res)
	return nil
}

func (c *ToggleFullScreen) Run(w *Window) error {
	res, ok := w.Settings().Get("toggle_full_screen", false).(bool)
	w.Settings().Set("toggle_fullscreen", !ok || !res)
	return nil
}

func (c *ToggleDsitractionFree) Run(w *Window) error {
	res, ok := w.Settings().Get("toggle_distraction_free", false).(bool)
	w.Settings().Set("toggle_distraction_free", !ok || !res)
	return nil
}

func init() {
	register([]Command{
		&ToggleSetting{},
		&SetSetting{},
		&ToggleSideBar{},
		&ToggleStatuseBar{},
		&ToggleFullScreen{},
		&ToggleDsitractionFree{},
	})
}
