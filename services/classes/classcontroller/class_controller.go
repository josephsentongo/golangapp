package classcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/josephsentongo/golangapp/helper"
	"github.com/josephsentongo/golangapp/services/classes/classservices"
	"github.com/josephsentongo/golangapp/services/classes/modelclass"
	"github.com/josephsentongo/golangapp/services/students/response"
)

type ClassesController struct {
	service   classservices.ClassesServices
	validator *validator.Validate
}

func NewClassesController(service classservices.ClassesServices) *ClassesController {
	return &ClassesController{service, validator.New(validator.WithRequiredStructEnabled())}
}

func (class *ClassesController) Create(w http.ResponseWriter, r *http.Request) {

	var createClassRequest modelclass.Classes

	err := json.NewDecoder(r.Body).Decode(&createClassRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Validate the request
	err = class.validator.Struct(createClassRequest)
	if err != nil {
		helper.ValidationErrors(w, err)
		return
	}

	classModel := createClassRequest
	saveErr := class.service.SaveClass(classModel)
	if saveErr != nil {
		http.Error(w, saveErr.Error(), http.StatusInternalServerError)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(webResponse)

}

func (class *ClassesController) AllClass(w http.ResponseWriter, r *http.Request) {

	saveErr := class.service.AllClass()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   saveErr,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(webResponse)

}
