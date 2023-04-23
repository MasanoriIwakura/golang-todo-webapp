package graph

import (
	"github.com/MasanoriIwakura/golang-todo-webapp/graph/model"
	"github.com/jinzhu/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
	DB    *gorm.DB
}
