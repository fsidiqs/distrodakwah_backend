package httpHelper

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	BadRequestMessage          = "bad_request"
	InternalServerErrorMessage = "internal_server_error"
	StatusOKMessage            = "success"
)
