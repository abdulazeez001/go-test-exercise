package manager

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type BasicResponse struct {
	Success    bool                   `json:"success"`
	Name       string                 `json:"name"`
	StatusCode int                    `json:"status_code"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data"`
}

func GetSuccessResponse(data map[string]interface{}, message string, code int, context *gin.Context) {
	response := BasicResponse{
		Success:    true,
		Message:    message,
		Data:       data,
		StatusCode: code,
	}
	marshalled, _ := json.Marshal(&response)
	var res map[string]interface{}
	_ = json.Unmarshal(marshalled, &res)

	context.JSON(response.StatusCode, res)
}

func GetErrorResponse(data map[string]interface{}, name string, message string, code int, context *gin.Context) {
	response := BasicResponse{
		Success:    false,
		Name:       name,
		Message:    message,
		Data:       data,
		StatusCode: code,
	}
	marshalled, _ := json.Marshal(&response)
	var res map[string]interface{}
	_ = json.Unmarshal(marshalled, &res)

	context.JSON(response.StatusCode, res)
}
