package repository

import "UnitTest/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
