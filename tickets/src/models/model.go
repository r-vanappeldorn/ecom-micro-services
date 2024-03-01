package models

import "go.mongodb.org/mongo-driver/mongo"

type Model interface {
	GetCollectionName() string
	FindAll(db *mongo.Database) ([]*Tickets, error)
	FindByID(db *mongo.Database, id string) (*Tickets, error)
}