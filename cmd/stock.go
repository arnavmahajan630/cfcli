package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/arnavmahajan630/cfcli/api"
	"github.com/arnavmahajan630/cfcli/api/routers"
	"github.com/spf13/cobra"
)

// stockCmd represents the stock command
var stockCmd = &cobra.Command{
	Use:   "stock <username>",
	Short: "Show the last attempted problem by the user",
	Long:  "Fetches the most recent problem submission by the given Codeforces username.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("❌ Username is required. Usage: cfcli stock <username>")
			os.Exit(1)
		}

		username := args[0]
		err := routers.StockUser(username)
		if err != nil {
			fmt.Printf("❌ Failed to fetch stock info: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Stocking " + username + "...\n")
		time.Sleep(2 * time.Second)
		fmt.Println(api.StockResult)
	},
}

func init() {
	rootCmd.AddCommand(stockCmd)
}
