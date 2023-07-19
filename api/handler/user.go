package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"user/models"
)

func (h *handler) User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.CreateUser(w, r)
	case "GET":
		var (
			values = r.URL.Query()
			method = values.Get("method")
		)

		if method == "GET_LIST" {
			h.UserGetList(w, r)
		} else if method == "GET" {
			h.UserGetById(w, r)
		}
	case "PUT":
		h.UserUpdate(w, r)
	case "DELETE":
		h.DeleteUser(w, r)
	}
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var CreateUser models.UserCreate

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while creating user"))
	}

	err = json.Unmarshal(body, &CreateUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while unmarshal"))
	}

	userId, err := h.strg.User().Create(&CreateUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while unmarshal user"))
	}

	user, err := h.strg.User().GetById(&models.UserPrimaryKey{
		Id: userId.Id,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while getbyid user"))
	}

	resp, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while marshal user"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h *handler) UserGetById(w http.ResponseWriter, r *http.Request) {
	var id string = r.URL.Query().Get("id")

	resp, err := h.strg.User().GetById(&models.UserPrimaryKey{Id: id})
	if err != nil {
		log.Println("Error while user get by id: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while user get by id"))
		return
	}

	body, err := json.Marshal(resp)
	if err != nil {
		log.Println("Error while user marshal" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while user marshal"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}

func (h *handler) UserGetList(w http.ResponseWriter, r *http.Request) {
	var (
		offsetStr = r.URL.Query().Get("offset")
		limitStr  = r.URL.Query().Get("limit")
		search    = r.URL.Query().Get("search")
	)

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		log.Println("error while offset: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while offset"))
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Println("error while limit: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while limit"))
		return
	}

	resp, err := h.strg.User().GetList(&models.UserGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		log.Println("error while storage user get list: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while storage user get list:"))
		return
	}

	body, err := json.Marshal(resp)
	if err != nil {
		log.Println("Error while user marshal" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while user marshal"))
		return
	}

	log.Println("Success")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *handler) UserUpdate(w http.ResponseWriter, r *http.Request) {
	var Updateuser models.UserUpdate

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while user ioutil read" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while user ioutil read"))
		return
	}

	err = json.Unmarshal(body, &Updateuser)
	if err != nil {
		log.Println("Error while user unmarshal" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while user iunmarshal"))
		return
	}

	_, err = h.strg.User().Update(&Updateuser)
	if err != nil {
		log.Println("Error while user update" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while user ioutil read"))
		return
	}

	user, err := h.strg.User().GetById(&models.UserPrimaryKey{Id: Updateuser.Id})
	if err != nil {
		log.Println("Error while user get by id" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while user get by id"))
		return
	}

	resp, err := json.Marshal(user)
	if err != nil {
		log.Println("Error while user marshal" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while user marshal"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var id string = r.URL.Query().Get("id")
	log.Println(id)

	err := h.strg.User().Delete(&models.UserPrimaryKey{Id: id})
	if err != nil {
		log.Println("Error while user delete" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while user delete"))
		return
	}
	w.WriteHeader(http.StatusOK)
}
