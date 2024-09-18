package models

import "time"

type Post struct {
	ID              int       `json:"id"`
	Caption         string    `json:"caption"`
	ImageURL        string    `json:"imageUrl"`
	PostedTimestamp time.Time `json:"postedTimestamp"` // Ensure this is of type time.Time
	UserID          int       `json:"userId"`
}

func (p *Post) GetParsedTimestamp() (time.Time, error) {
	return time.Parse(time.RFC3339, p.PostedTimestamp.Format(time.RFC3339))
}
