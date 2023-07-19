package handler

import (
	"user/config"
	"user/storage"
)

type handler struct {
	cfg  *config.Config
	strg storage.StorageI
}

func NewHandler(cfg *config.Config, strg storage.StorageI) *handler {
	return &handler{
		cfg:  cfg,
		strg: strg,
	}
}
