/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/arnavmahajan630/cfcli/api"
	"github.com/arnavmahajan630/cfcli/api/routers"
	"github.com/arnavmahajan630/cfcli/config"
	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _ := config.LoadConfig()
		var username string

		if len(args) == 0 {
			username = cfg.Username
			fmt.Println("Fetching your profile ...")
		} else {
			username = args[0]
			fmt.Printf("🔍 Fetching profile for: %s\n", username)
		}
		start := time.Now()
		if err := routers.HandleProfileCommand(username); err != nil {
			fmt.Println(err)
		}
		
		elapsed := time.Since(start)
		fmt.Printf("⏱ Elapsed: %.2f seconds\n", elapsed.Seconds())
		fmt.Println(api.UserProfile)
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// profileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// profileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
