package core

import core "github.com/Suinar/Bank-backend/bankBackend/internal/repository/postgres_db"

type UserService struct {
	Repository core.UserRepository
}
