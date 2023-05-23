package middlewares

import (
	"electro-item-management/configs"
	"time"

	"github.com/golang-jwt/jwt/v4"
)
// (userID int, email string, isAdmin bool) => for admin role
func CreateToken(userID int, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["email"] = email
	// claims["isAdmin"] = isAdmin
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.AppConfig.JWTKey))
}

// func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		user := c.Get("user").(*jwt.Token)
// 		claims := user.Claims.(jwt.MapClaims)
// 		isAdmin := claims["isAdmin"].(bool)

// 		if !isAdmin {
// 			return echo.ErrUnauthorized
// 		}

// 		return next(c)
// 	}
// }

// func ExtractTokenUserID(e echo.Context) float64 {
// 	user := e.Get("user").(*jwt.Token)
// 	if user.Valid {
// 		claims := user.Claims.(jwt.MapClaims)
// 		userId := claims["userID"].(float64)
// 		return userId
// 	}
// 	return 0
// }

// func ExtractTokenIsAdmin(e echo.Context) bool {
// 	user := e.Get("user").(*jwt.Token)
// 	if user.Valid {
// 		claims := user.Claims.(jwt.MapClaims)
// 		isAdmin := claims["isAdmin"].(bool)
// 		return isAdmin
// 	}
// 	return false
// }
