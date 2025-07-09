package ui

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/arnavmahajan630/cfcli/config"
	"github.com/charmbracelet/lipgloss"
)

func ShowWelcome() {
	cfg, err := config.LoadConfig()

	// If config is missing or username not set, trigger onboarding
	if err != nil || strings.TrimSpace(cfg.Username) == "" {
		cfg = runOnboarding()
		config.SaveConfig(cfg)
		clearScreen()
	}

	// print the real banner with username
	printBanner(cfg.Username)
}

func runOnboarding() config.Config {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Println("👋 Welcome to CFCLI!")
	fmt.Print("→ Please enter a username to get started: ")

	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	return config.Config{Username: username}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printBanner(username string) {
	mainStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFD700")).
		Bold(true)

	infoStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FFAA"))

	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Italic(true)

	banner := `
   ____ _ _____ _     
  / ___| |_   _| |    
 | |   | | | | | |    
 | |___| | | | | |___ 
  \____|_| |_| |_____|`

	fmt.Println(mainStyle.Render(banner))
	fmt.Println()
	fmt.Println(infoStyle.Render("👤 Username     : " + username))
	fmt.Println(infoStyle.Render("📦 Version      : v1.0.0"))
	fmt.Println(infoStyle.Render("⚡ CLI Ready    : Type `cfcli --help` for commands"))
	fmt.Println()
	fmt.Println(footerStyle.Render("🔗 https://buymeacoffee.com/arnav630   ❤️  Made with love by arnav"))
}
