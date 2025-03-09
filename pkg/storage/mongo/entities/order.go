package orderEntity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderEntity struct {
	ObjectId  primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	IsActive  bool               `bson:"is_active" json:"is_active"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
}
