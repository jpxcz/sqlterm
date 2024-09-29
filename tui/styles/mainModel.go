package styles

import "github.com/charmbracelet/lipgloss"


func DatabasesConnectivityViewStyleDefault (width int, height int) lipgloss.Style {
    modelStyle := lipgloss.NewStyle().
            Width(width).
            Height(height).
            Align(lipgloss.Center)

    return modelStyle
}

func DatabasesConnectivityViewStyleFocused (width int, height int) lipgloss.Style {
    modelStyle := DatabasesConnectivityViewStyleDefault(width, height).
        BorderStyle(lipgloss.NormalBorder()).
        BorderForeground(lipgloss.Color("#FF00FF"))
    return modelStyle
}
