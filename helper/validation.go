package helper

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/josephsentongo/golangapp/services/students/response"
)

func ValidationErrors(w http.ResponseWriter, error error) {

	errs := error.(validator.ValidationErrors)
	var errInput []map[string]string
	for _, e := range errs {
		jsonErrors := map[string]string{
			strings.ToLower(e.Field()): strings.ToLower(e.Field() + " is " + e.ActualTag()),
		}
		errInput = append(errInput, jsonErrors)
	}
	webResponse := response.ValidationErrors{
		Code:   http.StatusBadRequest,
		Status: "Bad Request",
		Errors: errInput,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(webResponse)
	return
}
