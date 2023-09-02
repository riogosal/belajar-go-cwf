package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Species struct {
	Name           string `json:"name" bson:"name"`
	ScientificName string `json:"scientific_name" bson:"scientific_name"`
	Code           string `json:"code" bson:"code"`
	LoinsPerFish   int    `json:"loins_per_fish" bson:"loins_per_fish"`
}

type Type struct {
	Name     string `json:"name" bson:"name"`
	Code     string `json:"code" bson:"code"`
	Whole    bool   `json:"whole" bson:"whole"`
	Gradable bool   `json:"gradable" bson:"gradable"`
	Multiple bool   `json:"multiple" bson:"multiple"`
}

type Supplier struct {
	UUID      string `json:"uuid" bson:"uuid"`
	Name      string `json:"name" bson:"name"`
	Code      string `json:"code" bson:"code"`
	BatchCode string `json:"batch_code" bson:"batch_code"`
}

type Content struct {
	Number      int       `json:"number" bson:"number"`
	Count       int       `json:"count" bson:"count"`
	Weight      float64   `json:"weight" bson:"weight"`
	Grade       string    `json:"grade" bson:"grade"`
	ContainerID int       `json:"container_id" bson:"container_id"`
	Date        string    `json:"date" bson:"date"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

type GroupedContent struct {
	Weight float64 `json:"weight" bson:"weight"`
	Count  int     `json:"count" bson:"count"`
}

type RawMaterial struct {
	ID              primitive.ObjectID        `bson:"_id,omitempty" json:"_id"`
	Ilc             string                    `json:"ilc" bson:"ilc"`
	DeliveryNumber  string                    `json:"delivery_number" bson:"delivery_number"`
	Species         Species                   `json:"species" bson:"species"`
	Type            Type                      `json:"type" bson:"type"`
	Supplier        Supplier                  `json:"supplier" bson:"supplier"`
	TotalWeight     float64                   `json:"total_weight" bson:"total_weight"`
	TotalCount      int64                     `json:"total_count" bson:"total_count"`
	CreatedAt       time.Time                 `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time                 `json:"updated_at" bson:"updated_at"`
	Contents        []Content                 `json:"contents" bson:"contents"`
	GroupedContents map[string]GroupedContent `json:"grouped_contents" bson:"grouped_contents"`
	// GroupedContents []GroupedContent `json:"grouped_contents" bson:"grouped_contents"`
}

func salam() string {
	now := time.Now()
	hour := now.Hour()

	if hour >= 6 && hour <= 9 {
		return "Selamat pagi"
	} else if hour >= 10 && hour <= 13 {
		return "Selamat siang"
		// waktu makan yang berwarna, Selera terbuka, nikmati makanan dengan perasaan, Setelah santap siang, kembali bekerja dengan giat, Sebelum lanjut kerja, pinjam dulu seratus."
	} else if hour >= 14 && hour <= 18 {
		return "Sore yang cerah, semangat terakhir sebelum pulang."
	} else {
		return "Selamat malam, semoga Anda memiliki waktu istirahat yang baik"
	}
}

func waktu(r RawMaterial) string {
	t := time.UnixMilli(r.CreatedAt.UnixMilli())

	hari := t.Day()
	bulan := t.Month() + 1
	tahun := t.Year()

	return fmt.Sprintf("%d-%d-%d", hari, bulan, tahun)
}

func (r RawMaterial) CetakStrRawmaterialLot() string {
	result := ""
	for k, v := range r.GroupedContents {
		result += fmt.Sprintf("%s: %.2f %d\n", k, v.Weight, v.Count)
	}

	return fmt.Sprintf(
		"Salam :%s\nIlc :%s\nTanggal :%s\nSuplayer :%s\nDetail :\n%s", salam(), r.Ilc, waktu(r), r.Supplier.Name, result,
	)
}
