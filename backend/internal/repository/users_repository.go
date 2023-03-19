package repository

import (
	"shodo/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepository struct {
	Client *mongo.Client
}

func NewUsersRepository(client *mongo.Client) *UsersRepository {
	return &UsersRepository{Client: client}
}

func (this *UsersRepository) CreateUser(user models.User) (primitive.ObjectID, error) {
	result, err := this.Client.Database("shodo").Collection("users").InsertOne(nil, user)
	return result.InsertedID.(primitive.ObjectID), err
}

// TODO: checking is not repository responsibility
func (this *UsersRepository) CheckUserExists(email string) (bool, error) {
	count, err := this.Client.Database("shodo").Collection("users").CountDocuments(nil, models.User{Email: email})
	return count > 0, err
}

func (this *UsersRepository) DeleteUser(id primitive.ObjectID) error {
	_, err := this.Client.Database("shodo").Collection("users").DeleteOne(nil, models.User{ID: id})
	return err
}
