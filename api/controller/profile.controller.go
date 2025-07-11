package controller

import (
	"fmt"

	"github.com/arnavmahajan630/cfcli/api/models"
	"github.com/arnavmahajan630/cfcli/api/service"
)

func FetchUserInfo(username string, profile *models.UserProfile) error {
	url := "https://codeforces.com/api/user.info?handles=" + username
	users, err := service.FetchAndDecode[models.CFUser](url)
	if err != nil || len(users) == 0 {
		return fmt.Errorf("user.info error: %w", err)
	}
	service.PopulateFromInfo(&users[0], profile)
	return nil
}

func FetchUserStatus(username string, profile *models.UserProfile) error {
	url := "https://codeforces.com/api/user.status?handle=" + username
	subs, err := service.FetchAndDecode[models.CFSubmission](url)
	if err != nil || len(subs) == 0 {
		return fmt.Errorf("user.info error: %w", err)
	}
	service.PopulateFromSubmissions(subs, profile)
	return nil

}

func FetchUserRatings(username string, profile *models.UserProfile) error {
	url := "https://codeforces.com/api/user.rating?handle=" + username
	ratings, err := service.FetchAndDecode[models.CFContestRating](url)
	if err != nil {
		return fmt.Errorf("user.rating decode failed: %w", err)
	}

	service.PopulateFromRatings(ratings, profile)
	return nil

}
