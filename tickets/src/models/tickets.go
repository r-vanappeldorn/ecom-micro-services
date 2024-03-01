package models

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Tickets struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Title  string             `json:"title,omitempty" bson:"title,omitempty"`
	Price  int                `json:"price,omitempty" bson:"price,omitempty"`
	UserID string             `json:"userId,omitempty" bson:"userId,omitempty"`
}

func (t *Tickets) FindAll(db *mongo.Database) ([]*Tickets, error) {
	collection := db.Collection(t.GetCollectionName())
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	tickets := []*Tickets{}
	if err = cursor.All(context.TODO(), &tickets); err != nil {
		return nil, err
	}

	return tickets, nil
}

func (t *Tickets) FindByID(db *mongo.Database, id string) (*Tickets, error) {
	collection := db.Collection(t.GetCollectionName())
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var ticket *Tickets
	if err = collection.FindOne(context.TODO(), bson.M{"id": objectID}).Decode(&ticket); err != nil {
		return nil, err
	}

	return ticket, nil
}

func (t *Tickets) Create(db *mongo.Database) error {
	collection := db.Collection(t.GetCollectionName())

	if t.Price == 0 {
		return errors.New("price is required")
	}

	if t.Title == "" {
		return errors.New("title is required")
	}

	if t.UserID == "" {
		return errors.New("userId is required")
	}

	_, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tickets) GetCollectionName() string {
	return "tickets"
}

func (t *Tickets) GetTicketsCollection() Model {
	return &Tickets{}
}

func NewTicket(title, userId string, price int) *Tickets {
	return &Tickets{
		Title:  title,
		UserID: userId,
		Price:  price,
	}
}
