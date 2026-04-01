package theme

import "github.com/fatih/color"

type RGB struct {
	R int
	G int
	B int
}
type Colors struct {
	Pokemon   *RGB
	Prompt    *RGB
	Header    *RGB
	Info      *RGB
	Highlight *RGB
	Location  *RGB
	Success   *RGB
	Warning   *RGB
	Error     *RGB
}

type CLITheme struct {
	BaseColors *Colors

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

type CLIThemeFunc struct {
	BaseColors *Colors

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

var DefaultColors = Colors{
	Pokemon:   &RGB{R: 245, G: 194, B: 231}, // Pink
	Prompt:    &RGB{R: 203, G: 166, B: 247}, // Purple
	Header:    &RGB{R: 203, G: 166, B: 247}, // Purple
	Info:      &RGB{R: 137, G: 180, B: 250}, // Blue
	Highlight: &RGB{R: 250, G: 179, B: 135}, // Orange
	Location:  &RGB{R: 148, G: 226, B: 213}, // Teal
	Success:   &RGB{R: 166, G: 227, B: 161}, // Green
	Warning:   &RGB{R: 249, G: 226, B: 175}, // Yellow
	Error:     &RGB{R: 243, G: 139, B: 168}, // Red
}

func LoadTheme(customColors ...*Colors) *CLITheme {
	activeColors := DefaultColors
	if len(customColors) > 0 && customColors[0] != nil {
		activeColors = overloadColors(&activeColors, customColors[0])
	}
	return &CLITheme{
		BaseColors: &activeColors,

		Pokemon:   color.RGB(activeColors.Pokemon.R, activeColors.Pokemon.G, activeColors.Pokemon.B),
		Prompt:    color.RGB(activeColors.Prompt.R, activeColors.Prompt.G, activeColors.Prompt.B),
		Header:    color.RGB(activeColors.Header.R, activeColors.Header.G, activeColors.Header.B).Add(color.BgBlack),
		Info:      color.RGB(activeColors.Info.R, activeColors.Info.G, activeColors.Info.B),
		Highlight: color.RGB(activeColors.Highlight.R, activeColors.Highlight.G, activeColors.Highlight.B),
		Location:  color.RGB(activeColors.Location.R, activeColors.Location.G, activeColors.Location.B),
		Success:   color.RGB(activeColors.Success.R, activeColors.Success.G, activeColors.Success.B),
		Warning:   color.RGB(activeColors.Warning.R, activeColors.Warning.G, activeColors.Warning.B),
		Error:     color.RGB(activeColors.Error.R, activeColors.Error.G, activeColors.Error.B),
	}
}

func LoadThemeFunc(customColors ...*Colors) *CLIThemeFunc {
	activeColors := DefaultColors
	if len(customColors) > 0 && customColors[0] != nil {
		activeColors = overloadColors(&activeColors, customColors[0])
	}
	return &CLIThemeFunc{
		BaseColors: &activeColors,

		Pokemon:   color.RGB(activeColors.Pokemon.R, activeColors.Pokemon.G, activeColors.Pokemon.B).SprintFunc(),
		Prompt:    color.RGB(activeColors.Prompt.R, activeColors.Prompt.G, activeColors.Prompt.B).SprintFunc(),
		Header:    color.RGB(activeColors.Header.R, activeColors.Header.G, activeColors.Header.B).Add(color.BgBlack).SprintFunc(),
		Info:      color.RGB(activeColors.Info.R, activeColors.Info.G, activeColors.Info.B).SprintFunc(),
		Highlight: color.RGB(activeColors.Highlight.R, activeColors.Highlight.G, activeColors.Highlight.B).SprintFunc(),
		Location:  color.RGB(activeColors.Location.R, activeColors.Location.G, activeColors.Location.B).SprintFunc(),
		Success:   color.RGB(activeColors.Success.R, activeColors.Success.G, activeColors.Success.B).SprintFunc(),
		Warning:   color.RGB(activeColors.Warning.R, activeColors.Warning.G, activeColors.Warning.B).SprintFunc(),
		Error:     color.RGB(activeColors.Error.R, activeColors.Error.G, activeColors.Error.B).SprintFunc(),
	}
}

func overloadColors(activeColors, custom *Colors) Colors {
	if custom.Pokemon != nil {
		activeColors.Pokemon = custom.Pokemon
	}
	if custom.Prompt != nil {
		activeColors.Prompt = custom.Prompt
	}
	if custom.Header != nil {
		activeColors.Header = custom.Header
	}
	if custom.Info != nil {
		activeColors.Info = custom.Info
	}
	if custom.Highlight != nil {
		activeColors.Highlight = custom.Highlight
	}
	if custom.Location != nil {
		activeColors.Location = custom.Location
	}
	if custom.Success != nil {
		activeColors.Success = custom.Success
	}
	if custom.Warning != nil {
		activeColors.Warning = custom.Warning
	}
	if custom.Error != nil {
		activeColors.Error = custom.Error
	}
	return *activeColors
}
