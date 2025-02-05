package middleware

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dgdts/CloudNoteServer/biz/biz_context"
	"github.com/dgdts/CloudNoteServer/pkg/utils"
	"github.com/dgrijalva/jwt-go"
)

const (
	TokenExpireTime = 24 * time.Hour
	Secret          = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1Ni10"
)

const (
	UserNameKey = "user_name"
	UserIDKey   = "user_id"
)

var userTokenMap = utils.NewTTLMap[string, interface{}]()

func SetBlockedToken(token string, ttl time.Duration) {
	userTokenMap.Set(token, nil, ttl)
}

func IsBlockedToken(token string) bool {
	_, ok := userTokenMap.Get(token)
	return ok
}

func getBizContext(c *app.RequestContext) *biz_context.BizContext {
	if bizCtx, ok := c.Get("biz_context"); ok {
		ret, ok := bizCtx.(*biz_context.BizContext)
		if !ok {
			hlog.CtxErrorf(context.Background(), "biz context is not *biz_context.BizContext")
		}
		return ret
	}
	return nil
}

func getAuthFromJWTToken(token string) (map[string]interface{}, error) {
	if token == "" {
		return nil, errors.New("token is empty")
	}

	token = strings.TrimPrefix(token, "Bearer ")

	if !ValidateJWTToken(token) {
		return nil, errors.New("token is invalid")
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(tokenObj *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, err
	}

	userID, ok := claims[UserIDKey]
	if !ok || userID == "" {
		return nil, errors.New("user_id is empty")
	}

	ret := map[string]interface{}{
		UserNameKey: claims[UserNameKey],
		UserIDKey:   claims[UserIDKey],
	}

	return ret, nil
}

func ValidateJWTToken(token string) bool {
	if valid, _ := validateJWTSign(token); !valid {
		return false
	}
	if !validateJWTExpire(token) {
		return false
	}
	return true
}

func validateJWTSign(token string) (bool, error) {
	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Secret), nil
	})

	return tokenObj != nil && err == nil, err
}

func validateJWTExpire(token string) bool {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(tokenObj *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	return err == nil && claims.VerifyExpiresAt(time.Now().Unix(), true)
}

func GenerateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Claims.(jwt.MapClaims)["exp"] = time.Now().Add(TokenExpireTime).Unix()
	return token.SignedString([]byte(Secret))
}
