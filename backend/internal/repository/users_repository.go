package repository

import (
	"context"
	"shodo/internal/config"
	mongodto "shodo/internal/repository/mongo_dto"
	"shodo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepository struct {
	Client *mongo.Client
	Config *config.Config
}

func (this *UsersRepository) CreateUser(ctx context.Context, user models.User) (string, error) {
	mongoUser := mongodto.User{}
	mongoUser.FromModel(user)

	result, err := this.getUsersCollection().InsertOne(ctx, mongoUser)
	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

// TODO: checking is not repository responsibility
func (this *UsersRepository) CheckUserExists(ctx context.Context, email string) (bool, error) {
	count, err := this.getUsersCollection().CountDocuments(ctx, mongodto.User{Email: email})
	return count > 0, err
}

func (this *UsersRepository) DeleteUser(ctx context.Context, id string) error {
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = this.getUsersCollection().DeleteOne(ctx, mongodto.User{ID: mongoId})
	return err
}

func (this *UsersRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user mongodto.User
	err := this.getUsersCollection().FindOne(ctx, mongodto.User{Email: email}).Decode(&user)
	return user.ToModel(), err
}

func (this *UsersRepository) GetAllUsers(ctx context.Context, id string) ([]models.UserShort, error) {
	var users []models.UserShort

	cursor, err := this.getUsersCollection().Find(ctx, mongodto.User{})
	if err != nil {
		return users, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user mongodto.UserShort
		err = cursor.Decode(&user)
		if err != nil {
			return users, err
		}

		if user.ID.Hex() == id {
			continue
		}
		users = append(users, user.ToModel())
	}

	return users, nil
}

func (this *UsersRepository) GetUserById(ctx context.Context, id string) (models.User, error) {
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}

	var user mongodto.User
	err = this.getUsersCollection().FindOne(ctx, mongodto.User{ID: mongoId}).Decode(&user)
	return user.ToModel(), err
}

func (this *UsersRepository) getUsersCollection() *mongo.Collection {
	return this.Client.Database(this.Config.DbName).Collection("users")
}
