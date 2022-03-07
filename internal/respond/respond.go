package respond

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  int         `json:"status,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Paging  interface{} `json:"paging,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type ErrorResponse struct {
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

type ResponseOption func(r *Response)

func Success(ctx *fiber.Ctx, options ...ResponseOption) error {
	res := &Response{}
	for _, opt := range options {
		opt(res)
	}

	if res.Status == 0 {
		res.Status = http.StatusOK
	}

	if res.Message == "" {
		res.Message = "available data"
	}

	return ctx.JSON(res)
}

func Fail(ctx *fiber.Ctx, options ...ResponseOption) error {
	res := &Response{}
	for _, opt := range options {
		opt(res)
	}

	if res.Status == 0 {
		res.Status = http.StatusBadRequest
	}

	if res.Message == "" {
		res.Message = "data not available"
	}

	return ctx.JSON(res)
}

func WithStatus(status int) ResponseOption {
	return func(res *Response) {
		res.Status = status
	}
}

func WithData(data interface{}) ResponseOption {
	return func(res *Response) {
		res.Data = data
	}
}

func WithMessage(message string) ResponseOption {
	return func(res *Response) {
		res.Message = message
	}
}

func WithPaging(paging interface{}) ResponseOption {
	return func(res *Response) {
		res.Paging = paging
	}
}

func WithError(err interface{}) ResponseOption {
	er, ok := err.(error)
	if ok {
		return func(res *Response) {
			res.Error = ErrorResponse{
				Detail: er.Error(),
			}
		}
	}

	return func(res *Response) {
		res.Error = err
	}
}
