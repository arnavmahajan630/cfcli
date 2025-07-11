package service

import (
	"time"

	"github.com/arnavmahajan630/cfcli/api/models"
)

	func PopulateStockresult(username string, subs []models.CFSubmission, stockres * models.StockResult) {
		desired := subs[0];

		stockres.ProblemName = desired.Problem.Name
		stockres.Handle = username
		stockres.Verdict = desired.Verdict
		stockres.ContestID = desired.Problem.ContestID
		stockres.TimeFormatted = formatToLocal(desired.CreationTimeSeconds)
	}

	func formatToLocal(ts int64) string{
		return time.Unix(ts, 0).Local().Format("2006 - 01- 02 15:04 MST")
	}