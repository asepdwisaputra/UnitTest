package service

import (
	"UnitTest/entity"
	"UnitTest/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found") // nil bisa ganti category
	} else {
		return category, nil
	}
}
