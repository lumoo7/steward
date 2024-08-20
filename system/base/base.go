package base

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"steward/system/constant"
	"strings"
)

type Page struct {
	PageSize  int `json:"pageSize"`
	PageIndex int `json:"pageIndex"`
}

type Sort struct {
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}

type Base struct{}

type PageResult struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Total   int64  `json:"total"`
	Page
}

type Result struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (s *Base) PageSuccess(ctx *gin.Context, data any, total int64, page Page) {
	res := &PageResult{
		Status:  true,
		Code:    0,
		Message: constant.RequestSuccess,
		Data:    data,
		Total:   total,
		Page:    page,
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *Base) Failure(ctx *gin.Context, code int, msg string) {
	res := &Result{
		Status:  false,
		Code:    code,
		Message: msg,
	}
	ctx.JSON(http.StatusInternalServerError, res)
}

func (s *Base) Success(ctx *gin.Context, data any) {
	res := &Result{
		Status:  true,
		Code:    1,
		Message: constant.RequestSuccess,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, res)
}

type Error struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

func NewError(code int, msg string) *Error {
	return &Error{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("an errror occur in %s. error code is:%d, error message is:%s",
		functionName(), e.ErrorCode, e.ErrorMsg)
}

// functionName 获取当前函数名
func functionName() string {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		parts := strings.Split(file, "/")
		fileName := parts[len(parts)-1]
		return fmt.Sprintf("%s:%d", fileName, line)
	}
	return ""
}
