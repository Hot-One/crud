package models

type User struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type UserCreate struct {
	FullName string `json:"full_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserUpdate struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type UserGetListResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"users"`
}
