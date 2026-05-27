package models

// Wraps a slice of Todos to be passed into the HTML template
type View struct {
	Todos []Todo
}
