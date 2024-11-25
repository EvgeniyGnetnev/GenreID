package service

import "github.com/EvgeniyGnetnev/GenreID/pkg/repository"

type Artist interface {
}

type Genre interface {
}

type Song interface {
}

type Service struct {
	Artist
	Genre
	Song
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
