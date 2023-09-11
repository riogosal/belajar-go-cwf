package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RefinedSpecies struct {
	Name           string `json:"name" bson:"name"`
	ScientificName string `json:"scientific_name" bson:"scientific_name"`
	Code           string `json:"code" bson:"code"`
	LoinsPerFish   int    `json:"loins_per_fish" bson:"loins_per_fish"`
}

type RefinedType struct {
	Name     string `json:"name" bson:"name"`
	Code     string `json:"code" bson:"code"`
	Whole    bool   `json:"whole" bson:"whole"`
	Gradable bool   `json:"gradable" bson:"gradable"`
	Multiple bool   `json:"multiple" bson:"multiple"`
}

type RefinedSupplier struct {
	UUID      string `json:"uuid" bson:"uuid"`
	Name      string `json:"name" bson:"name"`
	Code      string `json:"code" bson:"code"`
	BatchCode string `json:"batch_code" bson:"batch_code"`
}

type RefinedContent struct {
	Number          int       `json:"number" bson:"number"`
	Count           int       `json:"count" bson:"count"`
	Weight          float64   `json:"weight" bson:"weight"`
	Grade           string    `json:"grade" bson:"grade"`
	ContainerNumber int       `json:"container_id" bson:"container_id"`
	Date            string    `json:"date" bson:"date"`
	CreatedAt       time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" bson:"updated_at"`
}

type RefinedGroupedContent struct {
	Weight float64 `json:"weight" bson:"weight"`
	Count  int     `json:"count" bson:"count"`
}

type RefinedMaterial struct {
	ID                     primitive.ObjectID               `json:"_id,omitempty" bson:"_id,omitempty"`
	Ilc                    string                           `json:"ilc" bson:"ilc"`
	RawIlc                 string                           `json:"raw_ilc" bson:"raw_ilc"`
	CustomerGroup          string                           `json:"customer_group" bson:"customer_group"`
	RefinedSpecies         RefinedSpecies                   `json:"species" bson:"species"`
	RefinedTypes           RefinedType                      `json:"type" bson:"type"`
	RefinedSuppliers       RefinedSupplier                  `json:"supplier" bson:"supplier"`
	RefinedContents        []RefinedContent                 `json:"contents" bson:"contents"`
	TotalWeight            float64                          `json:"total_weight" bson:"total_weight"`
	TotalCount             int64                            `json:"total_count" bson:"total_count"`
	CreatedAt              int64                            `json:"created_at" bson:"created_at"`
	UpdatedAt              int64                            `json:"updated_at" bson:"updated_at"`
	RefinedGroupedContents map[string]RefinedGroupedContent `json:"grouped_contents" bson:"grouped_contents"`
}

var RefinedResults []*RefinedMaterial

func (rf RefinedMaterial) PrintRefinedMaterialDataByDate() string {
	result := ""

	for k, v := range rf.RefinedGroupedContents {
		result += fmt.Sprintf(`
%*s %d  ekor | %*.2f Kg`, -3, k, v.Count, -8, v.Weight)
	}

	if rf.TotalCount != 0 {
		return fmt.Sprintf(`
-----------------------------
          %s        
-----------------------------
%s 
	`, rf.CustomerGroup, result)
	} else {
		return fmt.Sprintf(`
-----------------------------
       %s        
-----------------------------
Tidak ada Data
	`, rf.CustomerGroup)
	}
}
