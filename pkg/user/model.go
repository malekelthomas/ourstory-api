package user

import (
	"github.com/malekelthomas/ourstory-api/pkg/permissions"
)

type User struct {
	ID          string   `json:"id"`
	UserName    string   `json:"username"`
	Role        UserRole `json:"role"`
	Permissions []string `json:"permissions"`
	Archived    bool     `json:"archived"`
}

type UserDTO struct {
	ID             string
	UserName       string
	SecurePassword string `bson:"secure_password"`
	Role           UserRole
	Permissions    []string
	Salt           string
	Archived       bool
}

func (u UserDTO) ToUser() *User {
	return &User{
		ID:          u.ID,
		UserName:    u.UserName,
		Role:        u.Role,
		Permissions: u.Permissions,
		Archived:    u.Archived,
	}
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

type UpdateUserOpts struct {
	UserName    *string   `json:"username"`
	Role        *UserRole `json:"role"`
	Password    *string   `json:"password"`
	OldPassword *string   `json:"old_password"`
}
