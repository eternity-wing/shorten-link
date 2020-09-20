package model

type Link struct {
	ID        int    `json:”id,omitempty” bson:”id,omitempty”`
	URL       string `json:”url,omitempty” bson:”url,omitempty”`
	UserID    int    `json:”user_id,omitempty” bson:”user_id,omitempty”`
	ShortLink string
}
