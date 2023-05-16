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

func (u *User) ToModel() models.User {
	return models.User{
		ID:       u.ID.Hex(),
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}
}

func (u *User) FromModel(user models.User) error {
	var err error
	u.ID, err = parseOrCreateId(user.ID)
	if err != nil {
		return err
	}
	u.Username = user.Username
	u.Password = user.Password
	u.Email = user.Email

	return nil
}

type UserShort struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"`
	Email    string             `bson:"email,omitempty"`
}

func (u *UserShort) ToModel() models.UserShort {
	return models.UserShort{
		ID:       u.ID.Hex(),
		Username: u.Username,
		Email:    u.Email,
	}
}
