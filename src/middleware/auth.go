package middleware

import (
	"fmt"
	"net/http"
	"os"

	"go-clean-arch/src/domain"
	helpers "go-clean-arch/src/helper"
	Error "go-clean-arch/src/pkg/error"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Auth ...
func Auth(a domain.Entity) gin.HandlerFunc {
	return func(c *gin.Context) {
		errorParams := map[string]interface{}{}
		header := c.Request.Header.Get("Authorization")
		if header == "" {
			fmt.Println(11)
			errorParams["meta"] = map[string]interface{}{
				"status":  401,
				"message": "Unauthorized",
			}
			errorParams["code"] = 401
			c.JSON(http.StatusUnauthorized, helpers.OutputAPIResponseWithPayload(errorParams))
			c.Abort()
			return
		}
		runes := []rune(header)
		tokenString := string(runes[7:])
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			fmt.Println(22)
			errorParams["meta"] = map[string]interface{}{
				"status":  401,
				"message": "Unauthorized",
			}
			errorParams["code"] = 401
			c.JSON(http.StatusUnauthorized, helpers.OutputAPIResponseWithPayload(errorParams))
			c.Abort()
			return
		}
		if !token.Valid {
			fmt.Println(33)
			errorParams["meta"] = map[string]interface{}{
				"status":  401,
				"message": "Unauthorized",
			}
			errorParams["code"] = 401
			c.JSON(http.StatusUnauthorized, helpers.OutputAPIResponseWithPayload(errorParams))
			c.Abort()
			return
		}
		uid := claims["uid"].(string)
		res, err := a.GetUserByUuid(c, uid)
		if err != nil {
			fmt.Println(44)
			errorParams["meta"] = map[string]interface{}{
				"status":  401,
				"message": "Unauthorized",
			}
			errorParams["code"] = 401
			c.JSON(http.StatusUnauthorized, helpers.OutputAPIResponseWithPayload(errorParams))
			c.Abort()
			return
		}
		if res.UUID == "" {
			fmt.Println(55)
			errorParams["meta"] = map[string]interface{}{
				"status":  401,
				"message": "Unauthorized",
			}
			errorParams["code"] = 401
			c.JSON(http.StatusUnauthorized, helpers.OutputAPIResponseWithPayload(errorParams))
			c.Abort()
			return
		}

		claims["user_uuid"] = res.UUID
		c.Set("User", claims)
		c.Next()

	}
}

// GetUserCustom ...
func GetUserCustom(c *gin.Context) map[string]interface{} {
	User := c.MustGet("User").(jwt.MapClaims)
	return User
}

// CreateToken ...
func CreateToken(data domain.UserLogin) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["uid"] = data.UUID
	atClaims["data"] = data
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("SECRET_KEY")))

	// rtClaims := jwt.MapClaims{}
	// rtClaims["uid"] = data.ID
	// rtClaims["data"] = uuid.New()
	// rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	// rtoken, err := rt.SignedString([]byte(os.Getenv("SECRET_KEY")))

	// // Save to db
	// tokenModel := models.TokenModel{}
	// _, _ = tokenModel.Save(data.ID, rtoken)

	if err != nil {
		Error.Error(err)
		return "", err
	}
	return token, nil
}
