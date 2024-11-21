package handlers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"

	"github.com/Rohit-Prasanna/todo/db"
	"github.com/Rohit-Prasanna/todo/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// GetTodos retrieves all todos for a specific user
func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get userId from URL parameters
	params := mux.Vars(r)
	userID := params["userId"]

	collection := db.GetCollection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Filter by userId
	filter := bson.M{"userId": userID}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		http.Error(w, "Failed to fetch todos: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	var todos []models.Todo
	if err := cursor.All(ctx, &todos); err != nil {
		http.Error(w, "Failed to decode todos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with todos
	json.NewEncoder(w).Encode(todos)
}

// CreateTodo creates a new todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Assign a new ObjectID and timestamp to the todo
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	collection := db.GetCollection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert the todo into the database
	_, err := collection.InsertOne(ctx, todo)
	if err != nil {
		http.Error(w, "Failed to create todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the created todo
	json.NewEncoder(w).Encode(todo)
}

// UpdateTodo updates an existing todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the ID from URL parameters (treat as plain string)
	params := mux.Vars(r)
	id := params["id"]

	// Decode the updated todo from the request body
	var updatedTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Ensure the ID in the decoded todo matches the route parameter
	updatedTodo.ID = id
	updatedTodo.UpdatedAt = time.Now() // Update timestamps

	collection := db.GetCollection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Filter and update using string ID
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedTodo}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Failed to update todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the updated todo
	json.NewEncoder(w).Encode(updatedTodo)
}

// DeleteTodo deletes a todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the ID from URL parameters (treat as plain string)
	params := mux.Vars(r)
	id := params["id"]

	collection := db.GetCollection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Filter and delete using string ID
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		http.Error(w, "Failed to delete todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if any document was deleted
	if result.DeletedCount == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Respond with a success message (No Content status)
	w.WriteHeader(http.StatusNoContent)
}

// Function to delete all todos by user ID (path parameter)
func DeleteAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract userId from the path parameters
	vars := mux.Vars(r)
	userID := vars["userId"]
	if userID == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}

	// Log the received userId for debugging
	log.Println("Received userId:", userID)

	// Get MongoDB collection
	collection := db.GetCollection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Delete all documents with the specified user ID
	filter := bson.M{"userId": userID}
	result, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Println("Error deleting todos:", err)
		http.Error(w, "Error deleting todos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Log the number of deleted documents
	log.Printf("Deleted %d todos for userId: %s\n", result.DeletedCount, userID)

	// Respond with the number of deleted documents
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Todos deleted successfully",
		"deleted": result.DeletedCount,
	})
}
func GetRemainingTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract userId from the URL path
	vars := mux.Vars(r)
	userID := vars["userId"]
	if userID == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}

	// Get MongoDB collection
	collection := db.GetCollection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Count todos where checked is false for the specified user
	filter := bson.M{"userId": userID, "checked": false}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		http.Error(w, "Error counting todos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the count of remaining todos
	json.NewEncoder(w).Encode(map[string]int{
		"remaining": int(count),
	})
}
