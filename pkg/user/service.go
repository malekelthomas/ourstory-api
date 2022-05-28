package user

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/malekelthomas/ourstory-api/pkg/credentials"
)

type UserService interface {
	Get(id string) (*User, error)
	Create(username, plainTextPassword string, role UserRole) (*User, error)
	UpdatePassword(id, oldPassword, newPassword string) (*User, error)
	Archive(id string) (*User, error)
}
type service struct {
	r UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &service{r: r}
}

func (s *service) Get(id string) (*User, error) {
	return s.r.Get(id)
}

func (s *service) Create(username, plainTextPassword string, role UserRole) (*User, error) {
	if us, _ := s.r.GetByUserName(username); us != nil {
		return nil, errors.New("username taken")
	}
	if role.String() == "err" {
		return nil, errors.New("invalid role provided")
	}

	salt, err := credentials.GenerateToken(len(plainTextPassword))
	if err != nil {
		return nil, fmt.Errorf("unable to create user: %v", err)
	}

	securePassword := credentials.GenerateSecurePassword(salt, plainTextPassword)
	u := &UserDTO{
		ID:             uuid.NewString(),
		UserName:       username,
		SecurePassword: securePassword,
		Role:           role,
		Permissions:    role.getPermissions(),
		Salt:           salt,
	}
	return s.r.Create(u)
}

func (s *service) UpdatePassword(id, oldPassword, newPassword string) (*User, error) {
	uo := UpdateUserOpts{
		OldPassword: &oldPassword,
		Password:    &newPassword,
	}

	return s.r.Update(id, uo)
}

func (s *service) Archive(id string) (*User, error) {
	return s.r.Archive(id)
}
