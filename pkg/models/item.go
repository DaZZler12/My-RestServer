package models

type Item struct {
	Brand     string  `json:"brand" bson:"brand"`
	Model     string  `json:"model" bson:"model"`
	Item_Name string  `json:"item_name" bson:"item_name"`
	Year      int64   `json:"year" bson:"year"`
	Price     float64 `json:"price" bson:"price"`
}
