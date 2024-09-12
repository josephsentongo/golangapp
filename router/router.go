package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/josephsentongo/golangapp/services/classes/classcontroller"
	"github.com/josephsentongo/golangapp/services/students/controller"
	"github.com/josephsentongo/golangapp/services/subjects/controllersubject"
)

func NewRoute(student *controller.StudentController, class *classcontroller.ClassesController, subject *controllersubject.SubjectesController) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/students", func(r chi.Router) {
		r.Post("/", student.Create)
		r.Get("/{id}", student.StudentByID)
		r.Patch("/{id}", student.UpdateStudent)
		r.Delete("/{id}", student.DeleteStudent)
		
		r.Get("/", student.GetStudents)
	})
	router.Route("/class", func(r chi.Router) {
		r.Post("/", class.Create)
		r.Get("/", class.AllClass)

	})
router.Route("/users", func(r chi.Router) {
		r.Post("/", student.CreateUser)
		r.Get("/", student.GetUsers)

	})

	router.Route("/subject", func(r chi.Router) {
		r.Post("/", subject.Create)

	})

	return router
}
