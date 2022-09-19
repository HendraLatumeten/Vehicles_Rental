package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	var status string
	if statusCode == 422 {
		status = "Unprocessable Entity"
	} else if statusCode == 400 {
		status = "Bad Request"
	} else if statusCode == 200 {
		status = "OK"
	} else if statusCode == 201 {
		status = "Created"
	} else if statusCode == 202 {
		status = "Accepted"
	}
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     status,
		"statusCode": statusCode,
		"data":       &data,
	})
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
