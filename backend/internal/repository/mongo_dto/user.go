package mongodto

import (
	"shodo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
	Email    string             `bson:"email,omitempty"`
}

func (this *User) ToModel() models.User {
	return models.User{
		ID:       this.ID.Hex(),
		Username: this.Username,
		Password: this.Password,
		Email:    this.Email,
	}
}

func (this *User) FromModel(user models.User) error {
	var err error
	this.ID, err = parseOrCreateId(user.ID)
	if err != nil {
		return err
	}
	this.Username = user.Username
	this.Password = user.Password
	this.Email = user.Email

	return nil
}
