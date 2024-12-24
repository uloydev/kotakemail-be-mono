package repository

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	appcontext "kotakemail.id/pkg/context"
	"kotakemail.id/pkg/database"
	"kotakemail.id/shared/schema"
)

type MailboxRepo interface {
	// GetByID returns a mailbox by ID
	GetByID(appCtx *appcontext.AppContext, id string) (*schema.Mailbox, error)
	// Create creates a new mailbox
	Create(appCtx *appcontext.AppContext, mailbox *schema.Mailbox) error
	// Update updates a mailbox
	Update(appCtx *appcontext.AppContext, mailbox *schema.Mailbox) error
	// Delete deletes a mailbox
	Delete(appCtx *appcontext.AppContext, id string) error
	// List returns a list of mailboxs
	List(appCtx *appcontext.AppContext) ([]*schema.Mailbox, error)
}

type mailboxRepo struct {
	db   database.Database
	conn *mongo.Database
}

func NewMailboxRepo(db database.Database) MailboxRepo {
	return &mailboxRepo{
		db:   db,
		conn: db.GetConnection().(*mongo.Database),
	}
}

func (r *mailboxRepo) GetByID(appCtx *appcontext.AppContext, id string) (*schema.Mailbox, error) {
	mailbox := &schema.Mailbox{}
	objID, _ := bson.ObjectIDFromHex(id)
	err := r.conn.Collection(mailbox.Collection()).FindOne(appCtx.Context(), bson.M{"_id": objID}).Decode(mailbox)
	return mailbox, err
}

func (r *mailboxRepo) Create(appCtx *appcontext.AppContext, mailbox *schema.Mailbox) error {
	_, err := r.conn.Collection(mailbox.Collection()).InsertOne(appCtx.Context(), mailbox)
	if err != nil {
		return err
	}
	return nil
}

func (r *mailboxRepo) Update(appCtx *appcontext.AppContext, mailbox *schema.Mailbox) error {
	_, err := r.conn.Collection(mailbox.Collection()).UpdateOne(appCtx.Context(), bson.M{"_id": mailbox.ID}, bson.M{"$set": mailbox})
	if err != nil {
		return err
	}
	return nil
}

func (r *mailboxRepo) Delete(appCtx *appcontext.AppContext, id string) error {
	objID, _ := bson.ObjectIDFromHex(id)
	_, err := r.conn.Collection((&schema.Mailbox{}).Collection()).DeleteOne(appCtx.Context(), bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}

func (r *mailboxRepo) List(appCtx *appcontext.AppContext) ([]*schema.Mailbox, error) {
	mailboxs := []*schema.Mailbox{}
	cursor, err := r.conn.Collection((&schema.Mailbox{}).Collection()).Find(appCtx.Context(), bson.M{})
	if err != nil {
		return mailboxs, err
	}
	defer cursor.Close(appCtx.Context())
	for cursor.Next(appCtx.Context()) {
		mailbox := &schema.Mailbox{}
		if err := cursor.Decode(mailbox); err != nil {
			return mailboxs, err
		}
		mailboxs = append(mailboxs, mailbox)
	}
	return mailboxs, nil
}
