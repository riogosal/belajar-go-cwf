package model

type Movie struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Year        int    `json:"year"`
}

<<<<<<< HEAD
var Film []Movie
=======
// export const movies = []
var Movies = []Movie{}

// func (m Movie) dataMovies() Movie {
// 	return Movie{m.Id, m.Title, m.Description, m.Year}
// }
>>>>>>> f233b3b3fa0a6dcfb26f02c1390a0fbb624732b9
