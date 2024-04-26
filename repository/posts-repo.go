package repository

import (
	"golang-rest-api/entity"
)

// import (
// 	"../entity"
// )

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
