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

var users []models.User = []models.User{
	{
		Id:       "5fde231f-29eb-469c-93a1-ca7f4c16fcfe",
		FullName: "Azizbek",
		Login:    "aziz",
		Password: "1234",
	},
	{
		Id:       "tge231f-29eb-469c-93a1-ca7f4c16fcfe",
		FullName: "Nurbek",
		Login:    "nur",
		Password: "5678",
	},
	{
		Id:       "par231f-29eb-469c-93a1-ca7f4c16fcfe",
		FullName: "Parviz",
		Login:    "par",
		Password: "5432",
	},
}

func (c *Controller) UserCRUD(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		c.GET_User(w, r)
	} else if r.Method == "PUT" {
		c.PUT_User(w, r)
	} else if r.Method == "DELETE" {
		c.DELETE(w, r)
	}
}

func (c *Controller) GET_User(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	arr := strings.Split(url, "/")
	fmt.Println(url)
	fmt.Println(arr)
	fmt.Println(len(arr))

	if arr[len(arr)-1] != "" {
		id := arr[len(arr)-1]
		for _, user := range users {
			if user.Id == id {
				data, err := json.Marshal(user)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				w.Write(data)
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		w.Write([]byte("Not found user with this ID"))
		w.WriteHeader(http.StatusNotFound)
	} else {
		data, err := json.Marshal(users)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write(data)
		w.WriteHeader(http.StatusOK)
	}
}

func (c *Controller) PUT_User(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	arr := strings.Split(url, "/")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error while ioutil:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("error while user update unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if arr[len(arr)-1] != "" {
		id := arr[len(arr)-1]
		for ind, val := range users {
			if val.Id == id {
				users[ind].FullName = user.FullName
				users[ind].Login = user.Login
				users[ind].Password = user.Password
			}
		}
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(body)
	return
}

func (c *Controller) DELETE(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	arr := strings.Split(url, "/")

	if arr[len(arr)-1] != "" {
		id := arr[len(arr)-1]
		for ind, user := range users {
			if user.Id == id {
				users = append(users[:ind], users[ind+1:]...)
				return
			}
		}
		w.Write([]byte("Not found user with this ID"))
		w.WriteHeader(http.StatusNotFound)
	}
}
