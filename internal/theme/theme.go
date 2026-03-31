package theme

import "github.com/fatih/color"

type CLITheme struct {
	Pokemon   *color.Color
	Prompt    *color.Color
	Header    *color.Color
	Info      *color.Color
	Highlight *color.Color
	Location  *color.Color
	Success   *color.Color
	Warning   *color.Color
	Error     *color.Color
}

func (c *CLITheme) LoadThemeFunc() any {
	panic("unimplemented")
}

type CLIThemeFunc struct {
	Pokemon   func(a ...any) string
	Prompt    func(a ...any) string
	Header    func(a ...any) string
	Info      func(a ...any) string
	Highlight func(a ...any) string
	Location  func(a ...any) string
	Success   func(a ...any) string
	Warning   func(a ...any) string
	Error     func(a ...any) string
}

var Theme = CLITheme{
	Pokemon:   color.RGB(245, 194, 231),                    // Pink
	Prompt:    color.RGB(203, 166, 247),                    // Purple
	Header:    color.RGB(203, 166, 247).Add(color.BgBlack), // Purple
	Highlight: color.RGB(250, 179, 135),                    // Orange
	Location:  color.RGB(148, 226, 213),                    // Teal
	Info:      color.RGB(137, 180, 250),                    // Blue
	Success:   color.RGB(166, 227, 161),                    // Green
	Warning:   color.RGB(249, 226, 175),                    // Yellow
	Error:     color.RGB(243, 139, 168),                    // Red
}

var ThemeFunc = CLIThemeFunc{
	Pokemon:   color.RGB(245, 194, 231).SprintFunc(),                    // Pink
	Prompt:    color.RGB(203, 166, 247).SprintFunc(),                    // Purple
	Header:    color.RGB(203, 166, 247).Add(color.BgBlack).SprintFunc(), // Purple
	Info:      color.RGB(137, 180, 250).SprintFunc(),                    // Blue
	Highlight: color.RGB(250, 179, 135).SprintFunc(),                    // Orange
	Location:  color.RGB(148, 226, 213).SprintFunc(),                    // Teal
	Success:   color.RGB(166, 227, 161).SprintFunc(),                    // Green
	Warning:   color.RGB(249, 226, 175).SprintFunc(),                    // Yellow
	Error:     color.RGB(243, 139, 168).SprintFunc(),                    // Red
}

func LoadTheme() *CLITheme {
	return &Theme
}

func LoadThemeFunc() *CLIThemeFunc {
	return &ThemeFunc
}
