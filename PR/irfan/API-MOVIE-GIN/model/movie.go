package model

type Labels []string

type Tags struct {
	Label Labels
}

type Movie struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Tags        `json:"tags"`
}

var Film = []Movie{
	{
		Id:          1,
		Title:       "AADC",
		Description: "Film Percitaan 2 Remaja",
		Year:        2002,
		Tags:        Tags{Label: Labels{"Romantic", "Drama", "Youngers"}},
	},
	{
		Id:          2,
		Title:       "Dilan1991",
		Description: "Film Percitaan tempo dulu",
		Year:        2018,
		Tags:        Tags{Label: Labels{"Romantic", "Drama", "Youngers"}},
	},
}
