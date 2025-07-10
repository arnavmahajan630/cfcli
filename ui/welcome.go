package ui

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	styles "github.com/arnavmahajan630/cfcli/Styles"
	"github.com/arnavmahajan630/cfcli/config"
)

func ShowWelcome() {
	cfg, err := config.LoadConfig()

	// If config is missing or username not set, trigger onboarding
	if err != nil || strings.TrimSpace(cfg.Username) == "" {
		cfg = runOnboarding()
		err := config.SaveConfig(cfg)
		if err != nil {
			log.Fatal(err)
		}
		clearScreen()
	}

	// print the banner with username
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
	banner := `
  _________  ___________ _________  .____     .___ 
\_   ___ \ \_   _____/ \_   ___ \ |    |    |   |
/    \  \/  |    __)   /    \  \/ |    |    |   |
\     \____ |     \    \     \____|    |___ |   |
 \______  / \___  /     \______  /|_______ \|___|
        \/      \/             \/         \/     `

	// 🖨️ Render
	fmt.Println(styles.TitleStyle.Render(banner))
	fmt.Println()
	fmt.Println(styles.LabelStyle.Render("👤 Username     : ") + styles.ValueStyle.Render(username))
	fmt.Println(styles.LabelStyle.Render("📦 Version      : ") + styles.ValueStyle.Render("v1.0.0"))
	fmt.Println(styles.LabelStyle.Render(styles.ElecStyle.Render("⚡") + " CLI Ready    : ") + styles.ValueStyle.Render("Type `cfcli --help` for commands"))
	fmt.Println()
	fmt.Println(styles.FooterStyle.Render("🔗 https://buymeacoffee.com/arnav630   ❤️  Made with love by arnav"))
}
