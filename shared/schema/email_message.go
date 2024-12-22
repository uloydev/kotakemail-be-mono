package schema

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type EmailMessage struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	MailboxID   bson.ObjectID `bson:"mailbox_id"`
	From        string        `bson:"from"`
	To          string        `bson:"to"`
	Cc          string        `bson:"cc"`
	Subject     string        `bson:"subject"`
	Body        string        `bson:"body"`
	Attachments []Attachment  `bson:"attachments"`
	CreatedAt   int64         `bson:"created_at"`
	ReadAt      int64         `bson:"read_at,omitempty"`
	DeletedAt   int64         `bson:"deleted_at,omitempty"`
}

type Attachment struct {
	Filename string `bson:"filename"`
	Url      string `bson:"url"`
}

func (e EmailMessage) Collection() string {
	return "email_messages"
}
