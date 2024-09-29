package tui

import "github.com/charmbracelet/lipgloss"

func panelStyleDefault(width int, height int) lipgloss.Style {
	modelStyle := lipgloss.NewStyle().
		Width(width).
		Height(height).
        PaddingLeft(1).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#CFD6ED"))
	return modelStyle
}



func panelStyleFocused(width int, height int) lipgloss.Style {
	modelStyle := panelStyleDefault(width, height).
		BorderForeground(lipgloss.Color("#1034A6"))
	return modelStyle
}

