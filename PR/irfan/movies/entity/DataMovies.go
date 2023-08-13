package entity

import "fmt"

type Movies struct {
	UUID, Year           int16
	Title, Desc, Country string
	Genre                []Genre
	IMDb                 float32
}

type Genre struct {
	Jenis string
}

// data movies print
func (Movies1 *Movies) CetakMovies() {

	fmt.Println("Your Favorite Films...")
	fmt.Println("==========================================================>")
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
