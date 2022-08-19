package helper

import "strings"

//Response struct json
type Response struct {
	Meta  MetaResponse `json:"meta"`
	Error interface{}  `json:"error"`
	Data  interface{}  `json:"data"`
}

type MetaResponse struct {
	Message string `json:"message"`
}

type ErrorDetail struct {
	Message string `json:"msg"`
}

//Struct to give Empty Object
type EmptyObj struct {
}

//function Build Response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Meta: MetaResponse{
			Message: message,
		},
		Error: nil,
		Data:  data,
	}
	return res
}

//function if the response is error
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")

	res := Response{
		Meta: MetaResponse{
			Message: message,
		},
		Error: ErrorDetail{
			Message: splittedError[0],
		},
		Data: data,
	}
	return res
}
