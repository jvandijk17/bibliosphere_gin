package middleware

import (
	"bibliosphere_gin/config"
	"bibliosphere_gin/domain"
	"bibliosphere_gin/service"
	"crypto/rand"
	"log"
	"regexp"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User = domain.User

var identityKey = config.AppConfig.IdentityKey
var jwtKey []byte
var authService service.AuthService

func init() {
	var err error
	jwtKey, err = generateSecureKey()
	if err != nil {
		log.Fatalf("Failed to generate JWT key: %v", err)
	}
}

func JWTMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       config.AppConfig.Realm,
		Key:         jwtKey,
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx *gin.Context) interface{} {
			claims := jwt.ExtractClaims(ctx)
			userID := uint(claims[identityKey].(float64))
			return &User{
				Model: gorm.Model{ID: userID},
			}
		},
		Authenticator: func(context *gin.Context) (interface{}, error) {
			var loginVals User
			if err := context.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			email := loginVals.Email
			password := loginVals.Password

			token, err := authService.AuthenticateUser(email, password, jwtKey)
			if err != nil {
				return nil, err
			}
			return token, nil
		},
		Authorizator: func(data interface{}, context *gin.Context) bool {
			claims := jwt.ExtractClaims(context)
			userRole := claims["roles"]

			requestPath := context.Request.URL.Path
			requestMethod := context.Request.Method

			var accessControlRules = []struct {
				Path    string
				Methods []string
				Roles   string
			}{
				{Path: "^/user", Methods: []string{"POST"}, Roles: config.AppConfig.PublicAccess},
				{Path: "^/library/preview_libraries", Methods: []string{"GET"}, Roles: config.AppConfig.PublicAccess},
				{Path: "^/", Methods: []string{"GET", "POST", "PUT", "DELETE"}, Roles: config.AppConfig.IsAuthenticatedFully},
			}

			for _, rule := range accessControlRules {
				matched, _ := regexp.MatchString(rule.Path, requestPath)
				methodAllowed := contains(rule.Methods, requestMethod)

				if matched && methodAllowed {
					if rule.Roles == config.AppConfig.PublicAccess {
						return true
					} else if rule.Roles == config.AppConfig.IsAuthenticatedFully && userRole != nil {
						return true
					}
				}
			}
			return false
		},
		Unauthorized: func(context *gin.Context, code int, message string) {
			context.JSON(code, gin.H{"code": code, "message": message})
		},
		TokenLookup:   config.AppConfig.IdentityKey,
		TokenHeadName: config.AppConfig.TokenHeadName,
		TimeFunc:      time.Now,
	})
}

func generateSecureKey() ([]byte, error) {
	key := make([]byte, 32) // 32 bytes for a 256-bit key
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
