package models

import "time"

type UserProfile struct {
	Handle       string // Codeforces handle
	Rank         string // e.g. "Expert"
	Rating       int    // current rating
	MaxRating    int    // max rating achieved
	Country      string // default to "N/A" if nil
	Organization string // default to "N/A" if nil

	ProblemCount int    // unique problems solved (from user.status)
	ContestCount int    // number of rated contests (from user.rating)
	MaxStreak    string // optional/future; placeholder now

	RecentContest string // name of latest contest
	RecentRank    int    // rank in latest contest
	RecentDelta   int    // delta in latest contest (new - old)
}

type StockResult struct {
	Handle        string
	ProblemName   string
	ContestID     int
	Index         string
	Rating        *int
	Verdict       string
	TimeFormatted string 
}

type AnalysisReport struct {
	GeneratedAt         time.Time
	User                CFUser
	Contests            []CFContestRating
	Submissions         []CFSubmission

	// Derived stats
	ContestsCount       int
	BestRank            int
	WorstRank           int
	RecentRanks         []int
	RatingTrend         []int
	TotalSolved         int
	UniqueProblems      int
	AvgProblemRating    float64
	MaxProblemRating    int
	FirstTryAccuracy    float64
	VerdictCounts       map[string]int
	TopTags             map[string]int
	WeakTags            map[string]int
}

