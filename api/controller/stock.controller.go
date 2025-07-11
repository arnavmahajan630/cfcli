package controller

import (
	"fmt"

	"github.com/arnavmahajan630/cfcli/api/models"
	"github.com/arnavmahajan630/cfcli/api/service"
)

func StockResult(username string, stock * models.StockResult) error {
	url := "https://codeforces.com/api/user.status?handle=" + username + "&count=1"
	subs , err := service.FetchAndDecode[models.CFSubmission](url)
	if err != nil || len(subs) == 0 {return fmt.Errorf("Failed %w", err) }
	service.PopulateStockresult(username, subs, stock)
    return nil

}