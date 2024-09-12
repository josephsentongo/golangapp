package controllersubject

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/josephsentongo/golangapp/helper"
	"github.com/josephsentongo/golangapp/services/students/response"
	"github.com/josephsentongo/golangapp/services/subjects/servicesubject"
	"github.com/josephsentongo/golangapp/services/subjects/subjectmodel"
)

type SubjectesController struct {
	service   servicesubject.IServiceSubject
	validator *validator.Validate
}

func NewSubjectesController(service servicesubject.IServiceSubject) *SubjectesController {
	return &SubjectesController{service, validator.New(validator.WithRequiredStructEnabled())}
}

func (Subject *SubjectesController) Create(w http.ResponseWriter, r *http.Request) {

	var createSubjectRequest subjectmodel.Subjects

	err := json.NewDecoder(r.Body).Decode(&createSubjectRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Validate the request
	err = Subject.validator.Struct(createSubjectRequest)
	if err != nil {
		helper.ValidationErrors(w, err)
		return
	}
	SubjectModel := createSubjectRequest
	saveErr := Subject.service.SaveSubject(SubjectModel)
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
