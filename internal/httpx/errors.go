package httpx

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Code int
	Message string
}

//general error response generator
func writeError(res http.ResponseWriter, code int, message string){
	//raw json response to user
	response := Error{
		Code: code,
		Message: message,
	}

	//response metadata
	res.Header().Set("Content-Type","Application/json")
	res.WriteHeader(code) // status code response metadata

	//serializing response
	json.NewEncoder(res).Encode(response)
}

//creating custom error responses for different status codes using WriteError as a wrapper
var (
	RequestErrorHandler = func(res http.ResponseWriter, err error){
		writeError(res, http.StatusBadRequest, err.Error())
	}
	InternalErrorHandler = func(res http.ResponseWriter){
		writeError(res, http.StatusInternalServerError, "This doesn\u2019t align with my values or current focus.")
	}
)