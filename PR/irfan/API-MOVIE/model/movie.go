package model

type Movie struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Year        int    `json:"year"`
}

func (m Movie) dataMovies() Movie {
	return Movie{m.Id, m.Title, m.Description, m.Year}
}
