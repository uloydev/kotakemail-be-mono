package repository

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	appcontext "kotakemail.id/pkg/context"
	"kotakemail.id/pkg/database"
	"kotakemail.id/shared/schema"
)

type UserRepo interface {
	// GetByID returns a user by ID
	GetByID(appCtx *appcontext.AppContext, id string) (*schema.User, error)
	// GetByEmail returns a user by email
	GetByEmail(appCtx *appcontext.AppContext, email string) (*schema.User, error)
	// Create creates a new user
	Create(appCtx *appcontext.AppContext, user *schema.User) error
	// Update updates a user
	Update(appCtx *appcontext.AppContext, user *schema.User) error
	// Delete deletes a user
	Delete(appCtx *appcontext.AppContext, id string) error
	// List returns a list of users
	List(appCtx *appcontext.AppContext) ([]*schema.User, error)
}

type userRepo struct {
	db   database.Database
	conn *mongo.Database
}

func NewUserRepo(db database.Database) UserRepo {
	return &userRepo{
		db:   db,
		conn: db.GetConnection().(*mongo.Database),
	}
}

func (r *userRepo) GetByID(appCtx *appcontext.AppContext, id string) (*schema.User, error) {
	user := &schema.User{}
	objID, _ := bson.ObjectIDFromHex(id)
	err := r.conn.Collection(user.Collection()).FindOne(appCtx.Context(), bson.M{"_id": objID}).Decode(user)
	return user, err
}

func (r *userRepo) GetByEmail(appCtx *appcontext.AppContext, email string) (*schema.User, error) {
	user := &schema.User{}
	err := r.conn.Collection(user.Collection()).FindOne(appCtx.Context(), bson.M{"email": email}).Decode(user)
	return user, err
}

func (r *userRepo) Create(appCtx *appcontext.AppContext, user *schema.User) error {
	_, err := r.conn.Collection(user.Collection()).InsertOne(appCtx.Context(), user)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) Update(appCtx *appcontext.AppContext, user *schema.User) error {
	_, err := r.conn.Collection(user.Collection()).UpdateOne(appCtx.Context(), bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) Delete(appCtx *appcontext.AppContext, id string) error {
	objID, _ := bson.ObjectIDFromHex(id)
	_, err := r.conn.Collection((&schema.User{}).Collection()).DeleteOne(appCtx.Context(), bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) List(appCtx *appcontext.AppContext) ([]*schema.User, error) {
	users := []*schema.User{}
	cursor, err := r.conn.Collection((&schema.User{}).Collection()).Find(appCtx.Context(), bson.M{})
	if err != nil {
		return users, err
	}
	defer cursor.Close(appCtx.Context())
	for cursor.Next(appCtx.Context()) {
		user := &schema.User{}
		if err := cursor.Decode(user); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
