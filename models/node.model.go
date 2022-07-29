package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Node struct{
	ID			primitive.ObjectID		`bson:"_id"`
	Text		string					`json:"text"`
	Title		string					`json:"title"`
	Created_at	time.time				`json:"created_at"`
	Updated_at	time.Time				`json:"updated_at"`
	Node_id		string					`json:"node_id"`
}