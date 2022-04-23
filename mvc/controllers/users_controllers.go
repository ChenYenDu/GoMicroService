package controllers

import (
	"encoding/json"
	"mvc/services"
	"mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationEror{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if err != nil {
		// Handle the error and return to the client
		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(apiErr.Message))
		return
	}

	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}