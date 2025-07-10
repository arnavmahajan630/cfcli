package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/arnavmahajan630/cfcli/api/fetchuser/models"
	"github.com/arnavmahajan630/cfcli/api/fetchuser/service"
)

// Generic decoder for responses
func decodeCFResponse[T any](body []byte) ([]T, error) {
	var base models.BaseResponse
	if err := json.Unmarshal(body, &base); err != nil {
		return nil, fmt.Errorf("failed to decode base response: %v", err)
	}
	if base.Status != "OK" {
		return nil, fmt.Errorf("API error: %s", base.Comment)
	}

	var result []T
	if err := json.Unmarshal(base.Result, &result); err != nil {
		return nil, fmt.Errorf("failed to decode result: %v", err)
	}
	return result, nil
}

func FetchUserInfo(username string, profile *models.UserProfile) error {
	resp1, err := http.Get("https://codeforces.com/api/user.info?handles=" + username)
	if err != nil {
		return fmt.Errorf("❌ Failed to fetch user.info: %w", err)
	}
	defer resp1.Body.Close()
	body1, _ := io.ReadAll(resp1.Body)
	users, err := decodeCFResponse[models.CFUser](body1)
	if err != nil || len(users) == 0 {
		return fmt.Errorf("❌ user.info error: %w", err)
	}
	service.PopulateFromInfo(&users[0], profile)
	return nil;
}


func FetchUserStatus(username string, profile *models.UserProfile) error {
	resp2, err := http.Get("https://codeforces.com/api/user.status?handle=" + username)
	if err != nil {
		return fmt.Errorf("⚠️ Failed to fetch user.status: %w", err)
	}
	defer resp2.Body.Close()
	body2, _ := io.ReadAll(resp2.Body)
	subs, err := decodeCFResponse[models.CFSubmission](body2)
	if err != nil {
		return fmt.Errorf("⚠️ user.status decode failed: %w", err)
	}
	service.PopulateFromSubmissions(subs, profile)
	return nil;

}

func FetchUserRatings(username string, profile *models.UserProfile) error {
	resp3, err := http.Get("https://codeforces.com/api/user.rating?handle=" + username)
	if err != nil {
		return fmt.Errorf("⚠️ Failed to fetch user.rating: %w", err)
	}
	defer resp3.Body.Close()
	body3, _ := io.ReadAll(resp3.Body)
	ratings, err := decodeCFResponse[models.CFContestRating](body3)
	if err != nil {
		return fmt.Errorf("⚠️ user.rating decode failed: %w", err)
	}

	service.PopulateFromRatings(ratings, profile)
	return nil;

}
