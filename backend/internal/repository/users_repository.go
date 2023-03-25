package repository

import (
	"shodo/internal/models"
	mongodto "shodo/internal/repository/mongo_dto"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepository struct {
	Client *mongo.Client
}

func (this *UsersRepository) CreateUser(user models.User) (string, error) {
	mongoUser := mongodto.User{}
	mongoUser.FromModel(user)

	result, err := this.getUsersCollection().InsertOne(nil, mongoUser)
	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

// TODO: checking is not repository responsibility
func (this *UsersRepository) CheckUserExists(email string) (bool, error) {
	count, err := this.getUsersCollection().CountDocuments(nil, mongodto.User{Email: email})
	return count > 0, err
}

func (this *UsersRepository) DeleteUser(id string) error {
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = this.getUsersCollection().DeleteOne(nil, mongodto.User{ID: mongoId})
	return err
}

func (this *UsersRepository) GetUserByEmail(email string) (models.User, error) {
	var user mongodto.User
	err := this.getUsersCollection().FindOne(nil, mongodto.User{Email: email}).Decode(&user)
	return user.ToModel(), err
}

func (this *UsersRepository) getUsersCollection() *mongo.Collection {
	return this.Client.Database("shodo").Collection("users")
}
