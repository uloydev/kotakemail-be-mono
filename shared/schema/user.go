package schema

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Email     string        `bson:"email"`
	Password  string        `bson:"password"`
	BlockedAt int64         `bson:"blocked_at,omitempty"`
}

func (u User) Collection() string {
	return "users"
}
