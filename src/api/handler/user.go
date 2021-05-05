package handler

import (
	"go-clean-arch/src/domain"
	helpers "go-clean-arch/src/helper"
	middlewares "go-clean-arch/src/middleware"
	Error "go-clean-arch/src/pkg/error"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (a *AppHandler) Login(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	var body domain.Login
	err := c.ShouldBindJSON(&body)
	if err != nil {
		statusCode = 406
		Error.Error(err)
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	data, err := a.Entity.Login(c, body)
	if err != nil {
		statusCode = 400
		Error.Error(err)
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	// Compare Password
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(body.Password))
	if err != nil {
		statusCode = 400
		Error.Error(err)
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	token, err := middlewares.CreateToken(data)
	if err != nil {
		statusCode = 400
		Error.Error(err)
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	data.Password = ""

	params := map[string]interface{}{
		"meta": map[string]interface{}{
			"token": token,
		},
		"payload": data,
	}
	c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(params))
}

func (a *AppHandler) GetUserByUuid(c *gin.Context) {
	errorParams := map[string]interface{}{}
	User := middlewares.GetUserCustom(c)

	res, err := a.Entity.GetUserByUuid(c, User["user_uuid"].(string))
	if err != nil {
		Error.Error(err)
		statusCode := http.StatusBadRequest
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helpers.OutputAPIResponseWithPayload(errorParams))
		return
	}

	params := map[string]interface{}{
		"meta":    "success",
		"payload": res,
	}
	c.JSON(http.StatusOK, helpers.OutputAPIResponseWithPayload(params))
}
