package domain

type Film struct {
	ID    int    `json:"id"`
	Judul string `json:"judul"`
}

type FilmRepository interface {
	FindAll() ([]Film, error)
	Create(film Film) error
	Update(id int, film Film) error
}
