package entity

import "fmt"

type Genre string

type Movies struct {
	UUID, Year           int
	Title, Desc, Country string
	Genre                []Genre
	IMDb                 float32
}

// data movies print
func (Movies1 *Movies) PrintFilmFavorite() {

	fmt.Println("Your Favorite Films...")
	fmt.Println("==========================================================>")
	fmt.Println("ID. Movies :", Movies1.UUID)
	fmt.Println("Title : ", Movies1.Title)
	fmt.Println("Description : ", Movies1.Desc)
	fmt.Println("Year : ", Movies1.Year)
	fmt.Println("Country : ", Movies1.Country)
	fmt.Println("Genre : ")
	for i, _ := range Movies1.Genre {
		fmt.Print(Movies1.Genre[i] + ", ")
	}
	fmt.Println("")
	fmt.Println("Rating : ", Movies1.IMDb)

}
