package models

import "time"

type UserProfile struct {
	Handle       string 
	Rank         string 
	Rating       int    
	MaxRating    int    
	Country      string 
	Organization string 

	ProblemCount int    
	ContestCount int    
	MaxStreak    string 

	RecentContest string 
	RecentRank    int   
	RecentDelta   int  
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

