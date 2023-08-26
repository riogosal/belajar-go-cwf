package main

import (
	"go-mongo/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	routes.SetupRoutes()
}

// func mongoURI() string {
// 	return fmt.Sprintf(
// 		"mongodb://%s",
// 		os.Getenv("DB_URI"),
// 	)
// }

// func initializeMongo() (*mongo.Client, error) {
// 	ctx := context.Background()
// 	mongo_client, err := mongo.Connect(ctx, options.Client().ApplyURI(
// 		mongoURI(),
// 	).SetReadPreference(readpref.Primary()))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return mongo_client, nil
// }

// func main() {
// 	godotenv.Load()
// 	routes.SetupRoutes()
// godotenv.Load()
// ctx := context.Background()

// mongo_client, err := initializeMongo()
// if err != nil {
// 	panic(err)
// }

// defer func() {
// 	if err := mongo_client.Disconnect(ctx); err != nil {
// 		panic(err)
// 	}
// }()

// db := mongo_client.Database(os.Getenv("DB_NAME"))
// filmCollections := db.Collection(collections.Film)

// cursor, err := filmCollections.Find(ctx, bson.M{"id": 2})

// if err != nil {
// 	panic(err)
// }

// var results []model.Film
// if err := cursor.All(ctx, &results); err != nil {
// 	panic(err)
// }

// for _, film := range results {
// 	fmt.Printf("Data yang ditemukan:\nID: %d\nJudul: %s\nDeskripsi: %s\nStudio: %s\nRating: %.2f\nDurasi: %d\nGenres: %v\n",
// 		film.ID, film.Judul, film.Deskripsi, film.Studio, film.Rating, film.Durasi, film.Genres)
// }
// fmt.Println("Jumlah hasil:", len(results))
// }
