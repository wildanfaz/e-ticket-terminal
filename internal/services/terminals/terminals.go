package terminals

import (
	"github.com/sirupsen/logrus"
	"github.com/wildanfaz/e-ticket-terminal/internal/repositories"
)

type Service struct {
	terminalsRepo repositories.TerminalsRepository
	usersRepo     repositories.UsersRepository
	log           *logrus.Logger
}

func NewService(
	terminalsRepo repositories.TerminalsRepository,
	usersRepo repositories.UsersRepository,
	log *logrus.Logger,
) *Service {
	return &Service{
		terminalsRepo: terminalsRepo,
		usersRepo:     usersRepo,
		log:           log,
	}
}
