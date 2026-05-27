package controllers

import (
	"fmt"
	"go-todo/config"
	"go-todo/models"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lakicsdomi/argus"
)

var (
	id        int
	item      string
	completed bool
	view      = template.Must(template.ParseFiles("views/index.html"))
	database  = config.Database()  // Initialize the database connection
	logger, _ = argus.Init("logs") // Initialize the logger using Argus
)

func Show(w http.ResponseWriter, r *http.Request) {
	statement, err := database.Query("SELECT * FROM gotodo.todos")
	if err != nil {
		logger.Error.LogErr("CONTROLLER", "Database query failed in Show", err)
	}
	defer statement.Close() // Ensure the statement is closed after we're done with it

	// Iterate over the rows returned by the query and populate the todos slice
	var todos []models.Todo
	for statement.Next() {
		err = statement.Scan(&id, &item, &completed)
		if err != nil {
			logger.Error.LogErr("CONTROLLER", "Row scanning failed", err)
		}

		todo := models.Todo{
			Id:        id,
			Item:      item,
			Completed: completed,
		}

		todos = append(todos, todo)
	}

	data := models.View{
		Todos: todos,
	}

	// Execute the template with the data and write the output to the response writer
	_ = view.Execute(w, data)
	logger.Verbose.Log("CONTROLLER", fmt.Sprintf("Displaying %d todo items", len(todos)))
}

// Add inserts a new to-do item into the database
func Add(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")

	_, err := database.Exec("INSERT INTO todos (item) VALUE (?)", item)
	if err != nil {
		logger.Error.LogErr("CONTROLLER", "Failed to insert todo", err)
	}

	// Redirect the user back to the home page after adding the new item
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	logger.Verbose.Log("CONTROLLER", fmt.Sprintf("Added new todo item: %s", item))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Extract the url variables
	id := vars["id"]

	_, err := database.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		logger.Error.LogErr("CONTROLLER", "Failed to delete todo", err)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	logger.Verbose.Log("CONTROLLER", fmt.Sprintf("Deleted todo item with id: %s", id))
}

// Complete marks a to-do item as completed in the database
func Complete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Extract the URL variables using mux.Vars
	id := vars["id"]

	_, err := database.Exec("UPDATE todos SET completed = true WHERE ID = ?", id)
	if err != nil {
		logger.Error.LogErr("CONTROLLER", "Failed to mark todo as completed", err)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	logger.Verbose.Log("CONTROLLER", fmt.Sprintf("Marked todo item with id %s as completed", id))
}
