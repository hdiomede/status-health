package repository

import "github.com/hdiomede/status-health/domain"

type RequestRepository interface {
	Store(request *domain.Request) error
	//Find(id string) (*Request, error)
	//FindAll() []*Request
}