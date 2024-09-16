package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"backend2/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentController struct {
	collection *mongo.Collection
}

func NewStudentController(client *mongo.Client) *StudentController {
	collection := client.Database("test2").Collection("students")
	return &StudentController{collection}
}

func (sc *StudentController) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if student.Name == "" || student.Age == "" || student.Grade == "" {
		http.Error(w, "Send all required fields.", http.StatusBadRequest)
		return
	}

	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()

	result, err := sc.collection.InsertOne(context.Background(), student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	student.ID = result.InsertedID.(primitive.ObjectID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func (sc *StudentController) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	var students []models.Student
	cursor, err := sc.collection.Find(context.Background(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var student models.Student
		cursor.Decode(&student)
		students = append(students, student)
	}

	response := map[string]interface{}{
		"count": len(students),
		"data":  students,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (sc *StudentController) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var student models.Student
	err := sc.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&student)
	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func (sc *StudentController) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if student.Name == "" || student.Age == "" || student.Grade == "" {
		http.Error(w, "Send all required fields.", http.StatusBadRequest)
		return
	}

	student.UpdatedAt = time.Now()

	update := bson.M{
		"$set": student,
	}

	result, err := sc.collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Student updated successfully"))
}

func (sc *StudentController) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	result, err := sc.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Student deleted successfully"))
}