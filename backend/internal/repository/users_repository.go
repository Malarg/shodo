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

func (this *UsersRepository) CreateUser(user models.User) (string, error) {
	mongoUser := mongodto.User{}
	mongoUser.FromModel(user)

	result, err := this.getUsersCollection().InsertOne(context.TODO(), mongoUser)
	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

// TODO: checking is not repository responsibility
func (this *UsersRepository) CheckUserExists(email string) (bool, error) {
	count, err := this.getUsersCollection().CountDocuments(context.TODO(), mongodto.User{Email: email})
	return count > 0, err
}

func (this *UsersRepository) DeleteUser(id string) error {
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = this.getUsersCollection().DeleteOne(context.TODO(), mongodto.User{ID: mongoId})
	return err
}

func (this *UsersRepository) GetUserByEmail(email string) (models.User, error) {
	var user mongodto.User
	err := this.getUsersCollection().FindOne(context.TODO(), mongodto.User{Email: email}).Decode(&user)
	return user.ToModel(), err
}

func (this *UsersRepository) GetAllUsers(id string) ([]models.UserShort, error) {
	var users []models.UserShort

	cursor, err := this.getUsersCollection().Find(context.TODO(), mongodto.User{})
	if err != nil {
		return users, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
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

func (this *UsersRepository) GetUserById(id string) (models.User, error) {
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}

	var user mongodto.User
	err = this.getUsersCollection().FindOne(context.TODO(), mongodto.User{ID: mongoId}).Decode(&user)
	return user.ToModel(), err
}

func (this *UsersRepository) getUsersCollection() *mongo.Collection {
	return this.Client.Database(this.Config.DbName).Collection("users")
}
