package middleware

import (
	"TheBoys/api/response"
	"time"

	"TheBoys/domain"
	"TheBoys/infrastructure/config"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func NewMiddleware(user domain.UserRepository) *Middleware {
	return &Middleware{user}
}

type Middleware struct {
	user domain.UserRepository
}

func (m Middleware) AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			response.UnauthorizedError(c, "Unauthorized")
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		userId, isWeb, err := ValidateToken(token)

		if err != nil {
			response.InternalServerError(c, err.Error())
			return
		}

		user, err := m.user.FindUserById(userId)

		if err != nil {
			response.InternalServerError(c, err.Error())
			return
		}

		if user == nil || (isWeb && (user.WebToken == nil || *user.WebToken != token)) || (!isWeb) {
			response.UnauthorizedError(c, "Invalid token")
			return
		}

		if !user.IsActive {
			response.UnauthorizedError(c, "User account is inactive, Kindly contact admin")
			return
		}

		c.Set("user", user)

		c.Next()

	}

}

func ValidateToken(tokenValue string) (uint, bool, error) {

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(config.Config.JwtSecretKey), nil
	})

	if err != nil {
		return 0, false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return 0, false, errors.New("Invalid token")
	}

	userId, ok := claims["user_id"].(float64)

	if !ok || userId == 0 {
		return 0, false, errors.New("Invalid token")
	}

	isWeb, ok := claims["is_web"].(bool)

	if !ok {
		return 0, false, errors.New("Invalid token")
	}

	return uint(userId), isWeb, nil
}

type LoginResponse struct {
	Id           uint      `json:"id" gorm:"column:id"`
	Name         string    `json:"name" gorm:"column:name"`
	Email        string    `json:"email" gorm:"column:email"`
	Mobile       string    `json:"mobile" gorm:"column:mobile"`
	RoleId       int       `json:"roleId" gorm:"column:roleId"`
	WebToken     *string   `json:"webToken" gorm:"column:webToken"`
	IsActive     bool      `json:"isActive" gorm:"column:isActive"`
	RoleName     string    `json:"roleName" gorm:"column:roleName"`
	OtpCreatedAt time.Time `json:"otpCreatedAt" gorm:"column:otpCreatedAt"`
}

func GetUserClaims(c *gin.Context) (*LoginResponse, error) {

	userClaims, isSuccess := c.Get("user")
	if !isSuccess {
		return nil, errors.New("User fetching error in middleware")
	}

	user, ok := userClaims.(*LoginResponse)
	if !ok {
		return nil, errors.New("User type conversion error")
	}

	return user, nil
}
