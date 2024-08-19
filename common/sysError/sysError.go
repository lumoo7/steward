package sys_error

import "fmt"

type SysError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewSysError(code int, msg string) *SysError {
	return &SysError{
		Code:    code,
		Message: msg,
	}
}

func (s *SysError) Error() string {
	return fmt.Sprintf("you have an error. error code is: %d, error message is: %s", s.Code, s.Message)
}
