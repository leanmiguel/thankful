package models

type Entry struct {
	UserId      string   `json:"user_id"`
	CreatedTime string   `json:"created_time"`
	Entries     []string `json:"entries"`
}

type Entries []Entry
