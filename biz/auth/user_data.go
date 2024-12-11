package auth

import (
	"context"
	"time"

	"github.com/dgdts/UniversalServer/internal/utils"
	"github.com/dgdts/UniversalServer/pkg/mongo"
)

const (
	UserCollection    = "users"
	UserEmailField    = "email"
	UserPasswordField = "password"
)

type User struct {
	ID        string         `bson:"_id"`
	Username  string         `bson:"username"`
	Email     string         `bson:"email"`
	Password  HashedPassword `bson:"password"`
	Avatar    string         `bson:"avatar"`
	CreatedAt time.Time      `bson:"created_at"`
	UpdatedAt time.Time      `bson:"updated_at"`
}

func GetUserByEmail(ctx context.Context, email string) (*User, error) {
	r := mongo.Finder(utils.GlobalCollection(UserCollection)).FindOne(ctx, UserEmailField, email)
	if r.Error() != nil {
		return nil, r.Error()
	}

	var user User
	err := r.Read(&user)
	return &user, err
}

func CountUserByEmail(ctx context.Context, email string) (int64, error) {
	r, err := mongo.Finder(utils.GlobalCollection(UserCollection)).CountDocuments(ctx, UserEmailField, email)
	return r, err
}

func InsertUser(ctx context.Context, user *User) error {
	r := mongo.Inserter(utils.GlobalCollection(UserCollection)).Insert(ctx, user)
	return r.Error()
}

func UpdateUser(ctx context.Context, user *User) error {
	r := mongo.Updater(utils.GlobalCollection(UserCollection)).WithEqFilter(UserEmailField, user.Email).ReplaceOne(ctx, user)
	return r.Error()
}
