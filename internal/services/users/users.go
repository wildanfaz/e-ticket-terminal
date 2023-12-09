package users

import (
	"github.com/sirupsen/logrus"
	"github.com/wildanfaz/e-ticket-terminal/internal/repositories"
)

type Service struct {
	usersRepo repositories.UsersRepository
	log       *logrus.Logger
}

func NewService(
	usersRepo repositories.UsersRepository,
	log *logrus.Logger,
) *Service {
	return &Service{
		usersRepo: usersRepo,
		log:       log,
	}
}
