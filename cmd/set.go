package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/arnavmahajan630/cfcli/config"
)

var setCmd = &cobra.Command{
	Use:   "set <username>",
	Short: "Set your default Codeforces username",
	Long: `Save your preferred Codeforces handle locally so you can
run commands like 'cfcli profile' without needing to pass it every time.`,
	Args: cobra.ExactArgs(1), // ensure user provides exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]

		// Save config
		cfg := config.Config{Username: username}
		err := config.SaveConfig(cfg)
		if err != nil {
			fmt.Println("❌ Failed to save config:", err)
			return
		}

		fmt.Printf("✅ Username set to: %s\n", username)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
