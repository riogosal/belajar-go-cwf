package model

type Hobby []string

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Hobby   Hobby  `json:"hobby"`
}
