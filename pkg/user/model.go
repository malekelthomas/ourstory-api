package user

import (
	"errors"

	"github.com/malekelthomas/ourstory-api/pkg/permissions"
)

type User struct {
	ID             string `json:"id"`
	UserName       string `json:"username"`
	securePassword string
	Role           UserRole `json:"role"`
	Permissions    []string `json:"permissions"`
}

type UserRole int64

const (
	Viewer UserRole = iota
	Submitter
	Approver
	RoleLimit
)

func (u UserRole) String() string {
	switch u {
	case Viewer:
		return "viewer"
	case Submitter:
		return "submitter"
	case Approver:
		return "approver"
	default:
		return "err"
	}
}

func (u UserRole) getPermissions() []string {
	switch u {
	case Viewer:
		return permissions.PermissionsViewer.ToArray()
	case Submitter:
		return permissions.PermissionsSubmitter.ToArray()
	case Approver:
		return permissions.PermissionsApprover.ToArray()
	default:
		return []string{"err"}
	}
}

func NewUser(id, username, password string, role UserRole) (*User, error) {

	if role.String() == "err" {
		return nil, errors.New("invalid role provided")
	}
	u := &User{
		ID:             id,
		UserName:       username,
		securePassword: password,
		Role:           role,
		Permissions:    role.getPermissions(),
	}
	return u, nil
}
