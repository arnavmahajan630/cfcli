package ui

import (
	"github.com/arnavmahajan630/cfcli/config"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var banner = `
  _________  ___________ _________  .____     .___ 
\_   ___ \ \_   _____/ \_   ___ \ |    |    |   |
/    \  \/  |    __)   /    \  \/ |    |    |   |
\     \____ |     \    \     \____|    |___ |   |
 \______  / \___  /     \______  /|_______ \|___|
        \/      \/             \/         \/     `

var (
	mainStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFD700")).
		Bold(true)

	infoStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FFAA")).
		Bold(true)

	warningStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF5F5F")).
		Bold(true)

)

func ShowWelcome() {
	fmt.Println(mainStyle.Render(banner))
	fmt.Println()

	cfg, err := config.LoadConfig()

	if err != nil || strings.TrimSpace(cfg.Username) == "" {
		fmt.Println(warningStyle.Render("❗ Username not set."))
		fmt.Println("→ Please run: cfcli set <your_username>")
	} else {
		fmt.Println(infoStyle.Render("👤 Username     : " + cfg.Username))

		fmt.Println(infoStyle.Render("📦 Version      : v1.0.0"))
		fmt.Println(infoStyle.Render("⚡ CLI Ready    : Type `cfcli --help` for commands"))
	}
}
