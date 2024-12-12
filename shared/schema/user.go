package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	BlockedAt int64              `bson:"blocked_at,omitempty"`
}

func (u User) Collection() string {
	return "users"
}
