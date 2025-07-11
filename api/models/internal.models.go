package models

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
