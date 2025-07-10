package models

type CFUser struct {
	Handle        string  `json:"handle"`
	FirstName     *string `json:"firstName"`
	LastName      *string `json:"lastName"`
	Country       *string `json:"country"`
	City          *string `json:"city"`
	Organization  *string `json:"organization"`
	Contribution  int     `json:"contribution"`
	Rank          string  `json:"rank"`
	Rating        int     `json:"rating"`
	MaxRank       string  `json:"maxRank"`
	MaxRating     int     `json:"maxRating"`
	LastOnline    int64   `json:"lastOnlineTimeSeconds"`
	RegisteredAt  int64   `json:"registrationTimeSeconds"`
	FriendOfCount int     `json:"friendOfCount"`
	Avatar        string  `json:"avatar"`
	TitlePhoto    string  `json:"titlePhoto"`
}



type CFSubmission struct {
	ID       int `json:"id"`
	Verdict  string `json:"verdict"`
	Problem  struct {
		ContestID int    `json:"contestId"`
		Index     string `json:"index"`
		Name      string `json:"name"`
	} `json:"problem"`
}



type CFContestRating struct {
	ContestID   int    `json:"contestId"`
	ContestName string `json:"contestName"`
	Handle      string `json:"handle"`
	Rank        int    `json:"rank"`
	RatingUpdateTimeSeconds int64 `json:"ratingUpdateTimeSeconds"`
	OldRating   int    `json:"oldRating"`
	NewRating   int    `json:"newRating"`
}
