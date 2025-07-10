package styles

import "github.com/charmbracelet/lipgloss"

// --- Adaptive Colors ---
var (
	TitleColor  = lipgloss.AdaptiveColor{Light: "#005F87", Dark: "#00D7FF"}    // Blue-ish / Cyan
	LabelColor  = lipgloss.AdaptiveColor{Light: "#5A5A5A", Dark: "#999999"}    // Muted gray
	ValueColor  = lipgloss.AdaptiveColor{Light: "#000000", Dark: "#E0E0E0"}    // Black or soft white
	FooterColor = lipgloss.AdaptiveColor{Light: "#444444", Dark: "#777777"}    // Dim gray
	ElecColor   = lipgloss.AdaptiveColor{Light: "#FFA500", Dark: "#FFCC00"}    // Orange / Bright yellow
)

// --- Styles ---
var (
	TitleStyle = lipgloss.NewStyle().
		Foreground(TitleColor).
		Bold(true)

	LabelStyle = lipgloss.NewStyle().
		Foreground(LabelColor)

	ValueStyle = lipgloss.NewStyle().
		Foreground(ValueColor).
		Bold(true)

	FooterStyle = lipgloss.NewStyle().
		Foreground(FooterColor).
		Italic(true)

	ElecStyle = lipgloss.NewStyle().
		Foreground(ElecColor)
)
