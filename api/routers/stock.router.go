package routers


import (
	"github.com/arnavmahajan630/cfcli/api"
	"github.com/arnavmahajan630/cfcli/api/controller"
	"github.com/arnavmahajan630/cfcli/api/models"
)

func StockUser(username string) error{
     var stockres models.StockResult
	 err := controller.StockResult(username , &stockres)
	 if err != nil {return err}
	 api.StockResult = stockres
	 return nil

}