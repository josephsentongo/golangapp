package main

import (
	"log"
	"net/http"
	"time"

	"github.com/josephsentongo/golangapp/config"
	"github.com/josephsentongo/golangapp/router"
	"github.com/josephsentongo/golangapp/services/classes/classcontroller"
	"github.com/josephsentongo/golangapp/services/classes/classesrepository"
	"github.com/josephsentongo/golangapp/services/classes/classservices"
	"github.com/josephsentongo/golangapp/services/students/controller"
	"github.com/josephsentongo/golangapp/services/students/repository"
	"github.com/josephsentongo/golangapp/services/students/service"
	"github.com/josephsentongo/golangapp/services/subjects/controllersubject"
	"github.com/josephsentongo/golangapp/services/subjects/repositorysubject"
	"github.com/josephsentongo/golangapp/services/subjects/servicesubject"
)

func main() {
	db := config.DatabaseConnection()
	tagsRepository := repository.NewStudentRepositoryImp(db)

	tagsService := service.NewStudentRepositoryImp(tagsRepository)

	//router
	tagsController := controller.NewStudentController(tagsService)
	// classesController:= classes.NewClassesController(tagsService)

	classRepository := classesrepository.NewClassRepositoryImp(db)

	classService := classservices.NewClassServiceImp(classRepository)

	//router
	controllerClass2 := classcontroller.NewClassesController(classService)

	subjectRepository := repositorysubject.NewSubjectsRepository(db)

	subjectsService1 := servicesubject.NewServiceSubject(subjectRepository)

	//router
	controllerSubject := controllersubject.NewSubjectesController(subjectsService1)
	r := router.NewRoute(tagsController, controllerClass2, controllerSubject)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}

}
