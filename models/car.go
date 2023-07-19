package models

type Car struct {
	Id            string  `json:"id"`
	Model         string  `json:"model"`
	Year          int     `json:"year"`
	Price         float32 `json:"price"`
	CountryNumber string  `json:"country_number"`
	UserId        string  `json:"user_id"`
	UserData      User    `json:"user_data"`
}
