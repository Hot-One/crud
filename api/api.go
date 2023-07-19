package api

import (
	"net/http"

	"user/api/handler"
	"user/config"
	"user/storage"
)

func NewApi(cfg *config.Config, strg storage.StorageI) {
	handler := handler.NewHandler(cfg, strg)

	http.HandleFunc("/user", handler.User)
}
