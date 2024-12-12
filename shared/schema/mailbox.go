package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mailbox struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserID       primitive.ObjectID `bson:"user_id"`
	Name         string             `bson:"name"`
	UnreadCount  int                `bson:"unread_count"`
	ApiKey       string             `bson:"api_key,omitempty"`
	SMTPUsername string             `bson:"smtp_username,omitempty"`
	SMTPPassword string             `bson:"smtp_password,omitempty"`
	DeletedAt    int64              `bson:"deleted_at,omitempty"`
}

func (m Mailbox) Collection() string {
	return "mailboxes"
}
