package service

import (
	"errors"
	"golang-rest-api/entity"
	"golang-rest-api/repository"
	"math/rand"
)

// Methods
type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

// This struct that is going to implement this PostService interface
type service struct{}

// Fue traido de routes.go NO SE SABE SI COPIAR O CORTAR
var (
	repo repository.PostRepository = repository.NewFirestoreRepository()
)

func NewPostService() PostService {
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("the post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("the post title is empty")
		return err
	}
	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)

}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()

}
