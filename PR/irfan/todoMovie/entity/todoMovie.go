package entity

import "fmt"

type Movies struct {
	UUID    int64
	Title   string
	Desc    string
	Year    int16
	Country string
	Genre   []Genre
	IMDb    float32
}

type Genre struct {
	Jenis string
}

func DataMovies() *Movies {
	DataHarryPotter := Movies{
		UUID:    1,
		Title:   "HARRY POTTER AND THE SORCERER’S STONE",
		Desc:    "This is the tale of Harry Potter (Daniel Radcliffe), an ordinary eleven-year-old boy serving as a sort of slave for his aunt and uncle who learns that he is actually a wizard and has been invited to attend the Hogwarts School for Witchcraft and Wizardry.",
		Year:    2001,
		Country: "United States",
		Genre: []Genre{
			{
				Jenis: "Adventure",
			},
			{
				Jenis: "Family",
			},
			{
				Jenis: "Fantasy",
			},
		},
		IMDb: 7.6,
	}

	return &DataHarryPotter
}

// func (m *Movies) cetakFilmDigimon() {}

func CetakMovies() {
	Movies1 := DataMovies()

	fmt.Println("Nowing Playings...")
	fmt.Println("================================================================")
	fmt.Println("ID. Movies ", Movies1.UUID)
	fmt.Println("Title : ", Movies1.Title)
	fmt.Println("Description : ", Movies1.Desc)
	fmt.Println("Year : ", Movies1.Year)
	fmt.Println("Country : ", Movies1.Country)
	fmt.Println("Genre : ")
	for i, _ := range Movies1.Genre {
		fmt.Print(Movies1.Genre[i].Jenis + ", ")
	}
	fmt.Println("")
	fmt.Println("Rating : ", Movies1.IMDb)

}
