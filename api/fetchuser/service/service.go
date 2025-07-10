package service

import (
	"fmt"

	"github.com/arnavmahajan630/cfcli/api/fetchuser/models"
)

func PopulateFromSubmissions(subs []models.CFSubmission, profile *models.UserProfile) {
	unique := make(map[string]bool)
	for _, s := range subs {
		if s.Verdict == "OK" {
			key := fmt.Sprintf("%d-%s", s.Problem.ContestID, s.Problem.Index)
			unique[key] = true
		}
	}
	profile.ProblemCount = len(unique)
}

func PopulateFromRatings(ratings []models.CFContestRating, profile *models.UserProfile) {
	profile.ContestCount = len(ratings)
	if len(ratings) > 0 {
		last := ratings[len(ratings)-1]
		profile.RecentContest = last.ContestName
		profile.RecentRank = last.Rank
		profile.RecentDelta = last.NewRating - last.OldRating
	}
}

func PopulateFromInfo(u *models.CFUser, profile *models.UserProfile) {
	profile.Handle = u.Handle
	profile.Rank = u.Rank
	profile.Rating = u.Rating
	profile.MaxRating = u.MaxRating
	profile.Country = fallback(u.Country)
	profile.Organization = fallback(u.Organization)
}

func fallback(s *string) string {
	if s == nil || *s == "" {
		return "N/A"
	}
	return *s
}


