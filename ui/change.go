package ui

import (
	"fmt"
	"time"

	"github.com/arnavmahajan630/cfcli/config"
	"github.com/arnavmahajan630/cfcli/ui/utils"
)

func ChnageUname(username string) {
	// Save config
		cfg := config.Config{Username: username}
		err := config.SaveConfig(cfg)
		if err != nil {
			fmt.Println("❌ Failed to save config:", err)
			return
		}

		fmt.Printf("✅ Username set to: %s\n", username)
		fmt.Println("Reloading..")
		time.Sleep(1*time.Second)
		utils.ClearScreen()
		printBanner(username)

}