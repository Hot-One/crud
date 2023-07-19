package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"user/models"
)

var books []models.Book = []models.Book{
	{
		Id:          "rede231f-29eb-469c-93a1-ca7f4c16fcfe",
		Title:       "Prince of Persie",
		Description: "Fantastic",
		AuthorId:    "abc",
		AuthorData: models.User{
			Id:       "5fde231f-29eb-469c-93a1-ca7f4c16fcfe",
			FullName: "Azizbek",
			Login:    "aziz",
			Password: "1234",
		},
	},
}

func (c *Controller) BookCRUD(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		c.GetBook(w, r)
	} else if r.Method == "PUT" {

	}
}

func (c *Controller) GetBook(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	arr := strings.Split(url, "/")
	fmt.Println(url)
	fmt.Println(arr)
	fmt.Println(len(arr))

	if arr[len(arr)-1] != "" {
		id := arr[len(arr)-1]
		for _, book := range books {
			if book.Id == id {
				data, err := json.Marshal(book)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(err.Error()))
					return
				}
				w.Write(data)
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		w.Write([]byte("Not found book with this ID"))
		w.WriteHeader(http.StatusNotFound)
	} else {
		data, err := json.Marshal(books)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write(data)
		w.WriteHeader(http.StatusOK)
	}
}

func (c *Controller) BookPUT(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	arr := strings.Split(url, "/")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("err while book unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var book models.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		log.Println("err while book unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if arr[len(arr)-1] != "" {
		id := arr[len(arr)-1]
		for ind, val := range books {
			if val.Id == id {
				books[ind].Title = book.Title
				books[ind].Description = book.Description
				books[ind].AuthorId = book.AuthorId
			}
		}
		return
	}
	w.Write(body)
	w.WriteHeader(http.StatusAccepted)
	return
}

func (c *Controller) BookDELETE(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	arr := strings.Split(url, "/")

	if arr[len(arr)-1] != "" {
		id := arr[len(arr)-1]

		for ind, book := range books {
			if book.Id == id {
				books = append(books[:ind], books[ind+1:]...)
				return
			}
		}
		w.Write([]byte("Not found book with this ID"))
		w.WriteHeader(http.StatusNotFound)
	}
}
