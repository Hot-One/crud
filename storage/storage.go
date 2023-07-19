package storage

import "user/models"

type StorageI interface {
	Close()
	User() UserRepoI
}

type UserRepoI interface {
	Create(*models.UserCreate) (*models.UserPrimaryKey, error)
	GetById(req *models.UserPrimaryKey) (*models.User, error)
	GetList(req *models.UserGetListRequest) (*models.UserGetListResponse, error)
	Update(req *models.UserUpdate) (*models.UserPrimaryKey, error)
	Delete(req *models.UserPrimaryKey) error
}
