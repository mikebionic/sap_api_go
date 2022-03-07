package helper

import (
	"strings"
)

//Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
	Total   int         `json:"total`
}

//EmptyObj object is used when data doesnt want to be null on json
type EmptyObj struct{}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	// fmt.Print(reflect.TypeOf(data).Kind())
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
		Total:   1,
	}
	return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
		Total:   0,
	}
	return res
}
