package models

type Task struct {
	Id          string `json:"id" bson:"_id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type Result struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}
