package schema

import "go.mongodb.org/mongo-driver/v2/bson"

type Mailbox struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	UserID       bson.ObjectID `bson:"user_id"`
	Name         string        `bson:"name"`
	UnreadCount  int64         `bson:"unread_count"`
	ApiKey       string        `bson:"api_key,omitempty"`
	SMTPUsername string        `bson:"smtp_username,omitempty"`
	SMTPPassword string        `bson:"smtp_password,omitempty"`
	DeletedAt    int64         `bson:"deleted_at,omitempty"`
	CreatedAt    int64         `bson:"created_at"`
	UpdatedAt    int64         `bson:"updated_at"`
}

func (m Mailbox) Collection() string {
	return "mailboxes"
}
