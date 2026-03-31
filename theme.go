package main

import "github.com/fatih/color"

type CLITheme struct {
	Prompt    *color.Color
	Header    *color.Color
	Highlight *color.Color
	Location  *color.Color
	Success   *color.Color
	Error     *color.Color
	Info      *color.Color
}

type CLIThemeFunc struct {
	Prompt    func(a ...any) string
	Header    func(a ...any) string
	Info      func(a ...any) string
	Highlight func(a ...any) string
	Location  func(a ...any) string
	Success   func(a ...any) string
	Error     func(a ...any) string
}

var Theme = CLITheme{
	Prompt:    color.RGB(203, 166, 247),                    // Purple
	Header:    color.RGB(203, 166, 247).Add(color.BgBlack), // Purple
	Highlight: color.RGB(250, 179, 135),                    // Orange
	Location:  color.RGB(148, 226, 213),                    // Teal
	Info:      color.RGB(137, 180, 250),                    // Blue
	Success:   color.RGB(166, 227, 161),                    // Green
	Error:     color.RGB(243, 139, 168),                    // Red
}

var ThemeFunc = CLIThemeFunc{
	Prompt:    color.RGB(203, 166, 247).SprintFunc(),                    // Purple
	Header:    color.RGB(203, 166, 247).Add(color.BgBlack).SprintFunc(), // Purple
	Info:      color.RGB(137, 180, 250).SprintFunc(),                    // Blue
	Highlight: color.RGB(250, 179, 135).SprintFunc(),                    // Orange
	Location:  color.RGB(148, 226, 213).SprintFunc(),                    // Teal
	Success:   color.RGB(166, 227, 161).SprintFunc(),                    // Green
	Error:     color.RGB(243, 139, 168).SprintFunc(),                    // Red
}
