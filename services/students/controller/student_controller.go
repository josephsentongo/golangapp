package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/josephsentongo/golangapp/helper"
	"github.com/josephsentongo/golangapp/services/students/model"
	"github.com/josephsentongo/golangapp/services/students/response"
	"github.com/josephsentongo/golangapp/services/students/service"
)

type StudentController struct {
	service   service.StudentServices
	validator *validator.Validate
}

func NewStudentController(service service.StudentServices) *StudentController {
	return &StudentController{service, validator.New(validator.WithRequiredStructEnabled())}
}

func (controller *StudentController) Create(w http.ResponseWriter, r *http.Request) {

	var createTagsRequest response.PostStudents

	err := json.NewDecoder(r.Body).Decode(&createTagsRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Validate the request
	err = controller.validator.Struct(createTagsRequest)
	if err != nil {
		helper.ValidationErrors(w, err)
		return
	}

	studentModel := createTagsRequest
	savedStudent, saveErr := controller.service.SaveStudents(studentModel.ToModel())
	if saveErr != nil {
		http.Error(w, saveErr.Error(), http.StatusInternalServerError)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   savedStudent,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(webResponse)

}




func (controller *StudentController) StudentByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	savedStudent, saveErr := controller.service.StudentByID(id)
	if saveErr != nil {
		http.Error(w, saveErr.Error(), http.StatusInternalServerError)
		return
	}

	webResponse := response.SingleResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   savedStudent,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(webResponse)

}

func (controller *StudentController) UpdateStudent(w http.ResponseWriter, r *http.Request) {

	var createTagsRequest response.PostStudents

	err := json.NewDecoder(r.Body).Decode(&createTagsRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	studentModel := createTagsRequest
	savedStudent, saveErr := controller.service.UpdateStudent(studentModel.ToModel(), id)
	if saveErr != nil {
		http.Error(w, saveErr.Error(), http.StatusInternalServerError)
		return
	}

	webResponse := response.SingleResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   savedStudent,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(webResponse)

}

func (controller *StudentController) DeleteStudent(w http.ResponseWriter, r *http.Request) {

	var createTagsRequest response.PostStudents

	err := json.NewDecoder(r.Body).Decode(&createTagsRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	saveErr := controller.service.DeleteStudentByID(id)
	if saveErr != nil {
		http.Error(w, saveErr.Error(), http.StatusInternalServerError)
		return
	}

	webResponse := response.SingleResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "deleted",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(webResponse)

}

func (controller *StudentController) GetStudents(w http.ResponseWriter, r *http.Request) {

	

	tags, totalRecords,pagination, err := controller.service.Students(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	webResponse := response.Response{
		Code:     http.StatusOK,
		Status:   "OK",
		Data:     tags,
		Total:    totalRecords,
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(webResponse)
}




func (controller *StudentController) GetUsers(w http.ResponseWriter, r *http.Request) {

	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSizeStr := r.URL.Query().Get("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	tags, totalRecords, err := controller.service.FindPosts(page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	webResponse := response.Response{
		Code:     http.StatusOK,
		Status:   "OK",
		Data:     tags,
		Total:    totalRecords,
		Page:     page,
		PageSize: pageSize,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(webResponse)
}


func (controller *StudentController) CreateUser(w http.ResponseWriter, r *http.Request) {

	var createTagsRequest model.Comments

	err := json.NewDecoder(r.Body).Decode(&createTagsRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Validate the request
	err = controller.validator.Struct(createTagsRequest)
	if err != nil {
		helper.ValidationErrors(w, err)
		return
	}

	studentModel := createTagsRequest
	 saveErr := controller.service.SaveComment(&studentModel)
	if saveErr != nil {
		http.Error(w, saveErr.Error(), http.StatusInternalServerError)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(webResponse)

}