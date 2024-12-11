package auth

import (
	"context"
	"errors"
	"time"

	"github.com/dgdts/UniversalServer/biz/model/auth"
	"github.com/dgdts/UniversalServer/internal/middleware"
	"github.com/dgdts/UniversalServer/pkg/global_id"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	count, err := CountUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := GeneratePassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:        global_id.GenerateUniqueID(),
		Email:     req.Email,
		Username:  req.Username,
		Password:  *hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{}, nil
}

func Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	user, err := GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = VerifyPassword(&user.Password, req.Password)
	if err != nil {
		return nil, err
	}

	token, err := middleware.GenerateToken(jwt.MapClaims{
		middleware.UserNameKey: user.Username,
		middleware.UserIDKey:   user.ID,
	})

	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{
		User: &auth.User{
			Id:        user.ID,
			Email:     user.Email,
			Username:  user.Username,
			Avatar:    user.Avatar,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
		AccessToken: token,
	}, nil
}
