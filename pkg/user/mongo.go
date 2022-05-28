package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/malekelthomas/ourstory-api/pkg/credentials"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository struct {
	collection *mongo.Collection
}

func NewUserMongoRepository(cl *mongo.Collection) UserMongoRepository {
	return UserMongoRepository{collection: cl}
}

func (u UserMongoRepository) Get(id string) (*User, error) {
	userDTO, err := u.get("id", id)
	if err != nil {
		return nil, err
	}
	return userDTO.ToUser(), nil
}

func (u UserMongoRepository) GetByUserName(username string) (*User, error) {
	userDTO, err := u.get("username", username)
	if err != nil {
		return nil, err
	}
	return userDTO.ToUser(), nil
}

func (u UserMongoRepository) get(key string, val string) (*UserDTO, error) {
	var user UserDTO
	if err := u.collection.FindOne(context.TODO(), bson.D{{Key: key, Value: val}}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u UserMongoRepository) Create(user *UserDTO) (*User, error) {
	if _, err := u.collection.InsertOne(context.TODO(), user); err != nil {
		return nil, fmt.Errorf("could not create user, err: %v", err)
	}
	return user.ToUser(), nil
}

func (u UserMongoRepository) Update(id string, opts UpdateUserOpts) (*User, error) {

	var update bson.D
	var user *User

	if opts.Password != nil {
		if opts.OldPassword == nil {
			return nil, errors.New("need to provide previous password")
		}
		user, err := u.get("id", id)
		if err != nil {
			return nil, err
		}

		if match := credentials.Verify(user.Salt, *opts.OldPassword, user.SecurePassword); !match {
			return nil, errors.New("incorrect password supplied")
		} else {
			salt, err := credentials.GenerateToken(len(*opts.Password))
			if err != nil {
				return nil, err
			}
			securePW := credentials.GenerateSecurePassword(salt, *opts.Password)

			update = append(update, bson.E{Key: "salt", Value: salt})
			update = append(update, bson.E{Key: "secure_password", Value: securePW})
		}
	}
	if opts.Role != nil {
		update = append(update, bson.E{Key: "role", Value: *opts.Role})
	}
	if opts.UserName != nil {
		update = append(update, bson.E{Key: "username", Value: *opts.UserName})
	}

	if err := u.collection.FindOneAndUpdate(context.TODO(), bson.D{{Key: "id", Value: id}}, update).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserMongoRepository) Archive(id string) (*User, error) {
	var user *User
	if err := u.collection.FindOneAndUpdate(context.TODO(), bson.D{{Key: "id", Value: id}}, bson.D{{Key: "archived", Value: true}}).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}
