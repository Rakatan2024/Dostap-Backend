package service

import (
	"hellowWorldDeploy/pkg/entity"
	"hellowWorldDeploy/pkg/repo"
	"log"
)

type SvcInterface interface {
	SignUp(user *entity.User) error
}

type Service struct {
	log  *log.Logger
	repo repo.RepInterface
}

func CreateService(repo repo.RepInterface, l *log.Logger) SvcInterface {
	return Service{repo: repo, log: l}
}