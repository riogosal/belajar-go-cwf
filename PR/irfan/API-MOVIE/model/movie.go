package model

type Movie struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Year        int    `json:"year"`
}

var Film []Movie
