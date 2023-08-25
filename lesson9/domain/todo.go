package domain

type Todo struct {
	ID     int    `json:"todo_id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var TodosInMemory = []Todo{
	{
		ID:     1,
		Title:  "Pay phone bill",
		Status: false,
	},
	{
		ID:     2,
		Title:  "Buy milk",
		Status: false,
	},
}
