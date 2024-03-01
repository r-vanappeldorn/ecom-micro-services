package routes

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Routes struct {
	DB *mongo.Database
}

func New(db *mongo.Database) *Routes {
	return &Routes{db}
}