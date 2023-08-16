package entity

import "fmt"

type Genre string
type Rating float32

type Movies struct {
	UUID, Year           int
	Title, Desc, Country string
	Genre                []Genre
	IMDb                 Rating
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
	fmt.Println("Genre : ", Movies1.Genre)
	for _, v := range Movies1.Genre {
		fmt.Print(v + ", ")
	}
	fmt.Println("")
	fmt.Println("Rating : ", Movies1.IMDb)
	fmt.Println("==========================================================>")

}
