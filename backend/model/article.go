package model

type Article struct {
	Model
	Uid     int    `json:"uid"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
