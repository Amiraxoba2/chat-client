package internal

import "github.com/charmbracelet/lipgloss"

var Important = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#ff0"))

var Success = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#0f0"))

var Error = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#f00"))
