package serializer

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

const (
	CodeNeedLogin = iota
	CodeDBError
	CodeParamError
	CodeBotError
	CodeTokenError
	CodeNotFoundError
	CodePasswordError
	CodeVerifiedError
	CodeRecordCreatedError
	CodeInternalServerError
)

func NeedLogin(c *gin.Context) {
	c.JSON(401, Response{
		Code:    CodeNeedLogin,
		Message: "The user is not logged in, please log in and then perform another operation",
	})
}

func NeedAccess(c *gin.Context) {
	c.JSON(403, Response{
		Code:    CodeNeedLogin,
		Message: "The user is not permitted for this operation",
	})
}

func Err(code int, msg string, err error) Response {
	res := Response{
		Code:    code,
		Message: msg,
	}

	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

// DBError database error
func DBError(err error) Response {
	msg := "the database service failed"
	return Err(CodeDBError, msg, err)
}

// ParamError Params Error
func ParamError(msg string, err error) Response {
	return Err(CodeParamError, msg, err)
}

// BotError bot error
func BotError(msg string, err error) Response {
	return Err(CodeBotError, msg, err)
}

// TokenError token error
func TokenError(msg string, err error) Response {
	return Err(CodeTokenError, msg, err)
}

// UserNotFoundError login user not found
func UserNotFoundError(msg string, err error) Response {
	return Err(CodeNotFoundError, msg, err)
}

// RecordCreated recoed created
func RecordCreatedError(msg string, err error) Response {
	return Err(CodeRecordCreatedError, msg, err)
}

// InternalServerError internal server error
func InternalServerError(err error) Response {
	msg := "Internal server error"
	return Err(CodeInternalServerError, msg, err)
}

func CustomVerifiedError() Response {
	msg := "custom password verify failed"
	return Err(CodeVerifiedError, msg, errors.New("custom password verify failed"))
}
