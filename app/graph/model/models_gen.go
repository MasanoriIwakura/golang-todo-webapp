// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Task        string `json:"task"`
	Description string `json:"description"`
}

type Todo struct {
	ID          int    `json:"id"`
	Task        string `json:"task"`
	Description string `json:"description"`
}
