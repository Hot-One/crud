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

var cars []models.Car = []models.Car{
	{
		Id:            "123456789",
		Model:         "Onix",
		Year:          2023,
		Price:         22000,
		CountryNumber: "C 919 MA",
		UserId:        "5fde231f-29eb-469c-93a1-ca7f4c16fcfe",
		UserData: models.User{
			Id:       "5fde231f-29eb-469c-93a1-ca7f4c16fcfe",
			FullName: "Azizbek",
			Login:    "aziz",
			Password: "1234",
		},
	},
	{
		Id:            "987654321",
		Model:         "Tahoe",
		Year:          2021,
		Price:         65000,
		CountryNumber: "C 777 MA",
		UserId:        "5fde231f-29eb-469c-93a1-ca7f4c16fcfe",
		UserData: models.User{
			Id:       "5fde231f-29eb-469c-93a1-ca7f4c16fcfe",
			FullName: "Azizbek",
			Login:    "aziz",
			Password: "1234",
		},
	},
}

func (c *Controller) CarCRUD(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		c.GetCAR(w, r)
	} else if r.Method == "PUT" {
		c.PutCAR(w, r)
	} else if r.Method == "DELETE" {
		c.CarDELETE(w, r)
	}

}

func (c *Controller) GetCAR(w http.ResponseWriter, r *http.Request) {
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
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				w.Write(data)
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		w.Write([]byte("Not found car with this ID"))
		w.WriteHeader(http.StatusNotFound)

	} else {
		data, err := json.Marshal(cars)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write(data)
		w.WriteHeader(http.StatusOK)
	}
}

func (c *Controller) PutCAR(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	arr := strings.Split(url, "/")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Err while reading file:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var car models.Car
	err = json.Unmarshal(body, &car)
	if err != nil {
		log.Println("Err while unmarshaling car")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if arr[len(arr)-1] != "" {
		id := arr[len(arr)-1]
		for ind, val := range cars {
			if val.Id == id {
				cars[ind].Model = car.Model
				cars[ind].Price = car.Price
				cars[ind].CountryNumber = car.CountryNumber
				cars[ind].Year = car.Year
			}
		}
		return
	}

	w.Write(body)
	w.WriteHeader(http.StatusAccepted)
	return
}

func (c *Controller) CarDELETE(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	arr := strings.Split(url, "/")

	if arr[len(arr)-1] != "" {
		id := arr[len(arr)-1]

		for ind, val := range cars {
			if val.Id == id {
				cars = append(cars[:ind], cars[ind+1:]...)
				return
			}
		}
		w.Write([]byte("Not found car with this ID"))
		w.WriteHeader(http.StatusNotFound)
	}
}
