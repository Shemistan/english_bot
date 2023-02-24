package service

import "github.com/Shemistan/english_bot/internal/domain"

type service struct {
	storage domain.Storage
}

func New(storage domain.Storage) domain.Service {
	return &service{storage: storage}
}

func (s *service) Get() {

}
