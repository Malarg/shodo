package mongodto

import (
	"shodo/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title,omitempty"`
}

func (this *Task) ToModel() models.Task {
	return models.Task{
		ID:    this.ID.Hex(),
		Title: this.Title,
	}
}

func (this *Task) FromModel(task models.Task) error {
	var err error
	this.ID, err = parseOrCreateId(task.ID)
	if err != nil {
		return err
	}

	this.Title = task.Title
	return nil
}
