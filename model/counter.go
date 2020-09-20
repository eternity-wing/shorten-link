package model

type Counter struct {
	ID    string `json:”id,omitempty” bson:”id,omitempty”`
	Value int    `json:”value,omitempty” bson:”value,omitempty”`
}
