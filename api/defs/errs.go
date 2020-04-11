package defs

import "net/http"

type Error struct {
    Error     string `json:"error"`      // 错误
    ErrorCode int    `json:"error_code"` // 错误码，内部使用
}

type ErrorResponse struct {
    HttpStatusCode int
    Error          Error
}

var (
    ErrorRequestBodyParseFailed = ErrorResponse{
        HttpStatusCode: http.StatusBadRequest,
        Error: Error{
            Error:     "Request body is not correct",
            ErrorCode: 001,
        },
    }
    ErrorNotAuthUser = ErrorResponse{
        HttpStatusCode: http.StatusUnauthorized,
        Error: Error{
            Error:     "User authentication failed.",
            ErrorCode: 002,
        },
    }
)
