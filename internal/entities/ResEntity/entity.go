package ResEntity

import (
	"test1/internal/entities"
	"time"
)

type ResponseMsg struct {
	Message string             `json:"message"`
	Status  int                `json:"status"`
	Time    string             `json:"time"`
	Error   string             `json:"error"`
	Data    *entities.VideoRes `json:"data"`
}

func CustomErrorResponse(err string, code int) *ResponseMsg {
	return &ResponseMsg{
		Message: "FAILED",
		Status:  code,
		Time:    time.Now().Format(time.RFC3339),
		Error:   err,
		Data:    nil,
	}
}

func CustomSuccessResponse(code int, data *entities.VideoRes) *ResponseMsg {
	return &ResponseMsg{
		Message: "SUCCESS",
		Status:  code,
		Time:    time.Now().Format(time.RFC3339),
		Error:   "",
		Data:    data,
	}
}
