package models

type Book struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorId    string `json:"author_id"`
	AuthorData  User   `json:"author_data"`
}
