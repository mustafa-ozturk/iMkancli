package main

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	BorderColor         lipgloss.Color
	StatusTextColor     lipgloss.Color
	SelectedTitleColor  lipgloss.Color
	SelectedDescColor   lipgloss.Color
	SelectedBorderColor lipgloss.Color
	ColumnTitleColor    lipgloss.Color
	ColumnTitleBgColor  lipgloss.Color
}

var DefaultTheme = Theme{
	BorderColor:         lipgloss.Color("#4C566A"), // Nord3 (Dark gray)
	StatusTextColor:     lipgloss.Color("#D8DEE9"), // Nord4 (Light gray)
	SelectedTitleColor:  lipgloss.Color("#88C0D0"), // Nord8 (Cyan highlight)
	SelectedDescColor:   lipgloss.Color("#4E7685"), // Nord8 muted
	SelectedBorderColor: lipgloss.Color("#88C0D0"), // Nord8 (Cyan selection bar)
	ColumnTitleColor:    lipgloss.Color("#ECEFF4"), // Nord6 (Bright white)
	ColumnTitleBgColor:  lipgloss.Color("#5E81AC"), // Nord10 (Muted blue background)
}
