package user

type UserRepository interface {
	Get(id string) (*User, error)
	GetByUserName(username string) (*User, error)
	Create(user *UserDTO) (*User, error)
	Update(id string, opts UpdateUserOpts) (*User, error)
	Archive(id string) (*User, error)
}
