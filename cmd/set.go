package cmd

import (

	"github.com/spf13/cobra"
	"github.com/arnavmahajan630/cfcli/ui"
)

var setCmd = &cobra.Command{
	Use:   "set <username>",
	Short: "Set your default Codeforces username",
	Long: `Save your preferred Codeforces handle locally so you can
run commands like 'cfcli profile' without needing to pass it every time.`,
	Args: cobra.ExactArgs(1), // ensure user provides exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		ui.ChnageUname(username)

	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
