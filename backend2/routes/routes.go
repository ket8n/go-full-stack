package routes

import (
	"net/http"
	"backend2/controllers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(router *mux.Router, client *mongo.Client) {
	studentController := controllers.NewStudentController(client)

	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/students", studentController.CreateStudent).Methods("POST")
	router.HandleFunc("/students", studentController.GetAllStudents).Methods("GET")
	router.HandleFunc("/students/{id}", studentController.GetStudentByID).Methods("GET")
	router.HandleFunc("/students/{id}", studentController.UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", studentController.DeleteStudent).Methods("DELETE")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}