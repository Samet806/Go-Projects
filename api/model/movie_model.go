package model

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     int    `json:"year"`
}
