package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/arnavmahajan630/cfcli/api"
	"github.com/arnavmahajan630/cfcli/api/fetchuser/models"
	"github.com/arnavmahajan630/cfcli/api/fetchuser/service"
)

type baseResponse struct {
	Status  string          `json:"status"`
	Comment string          `json:"comment"`
	Result  json.RawMessage `json:"result"`
}

// Generic decoder for Codeforces responses
func decodeCFResponse[T any](body []byte) ([]T, error) {
	var base baseResponse
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

func HandleProfileCommand(username string) {
	var profile models.UserProfile

	// 1️⃣ user.info
	resp1, err := http.Get("https://codeforces.com/api/user.info?handles=" + username)
	if err != nil {
		fmt.Println("❌ Failed to fetch user.info:", err)
		return
	}
	defer resp1.Body.Close()
	body1, _ := io.ReadAll(resp1.Body)
	users, err := decodeCFResponse[models.CFUser](body1)
	if err != nil || len(users) == 0 {
		fmt.Println("❌ user.info error:", err)
		return
	}
	service.PopulateFromInfo(&users[0], &profile)

	// 2️⃣ user.status
	resp2, err := http.Get("https://codeforces.com/api/user.status?handle=" + username)
	if err != nil {
		fmt.Println("⚠️ Failed to fetch user.status:", err)
	} else {
		defer resp2.Body.Close()
		body2, _ := io.ReadAll(resp2.Body)
		subs, err := decodeCFResponse[models.CFSubmission](body2)
		if err != nil {
			fmt.Println("⚠️ user.status decode failed:", err)
		} else {
			service.PopulateFromSubmissions(subs, &profile)
		}
	}

	// 3️⃣ user.rating
	resp3, err := http.Get("https://codeforces.com/api/user.rating?handle=" + username)
	if err != nil {
		fmt.Println("⚠️ Failed to fetch user.rating:", err)
	} else {
		defer resp3.Body.Close()
		body3, _ := io.ReadAll(resp3.Body)
		ratings, err := decodeCFResponse[models.CFContestRating](body3)
		if err != nil {
			fmt.Println("⚠️ user.rating decode failed:", err)
		} else {
			service.PopulateFromRatings(ratings, &profile)
		}
	}

	// ✅ Export final profile for UI
	api.UserProfile = profile
}
