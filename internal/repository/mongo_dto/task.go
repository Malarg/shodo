package mongodto

import (
	"shodo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title,omitempty"`
}

func (t *Task) ToModel() models.Task {
	return models.Task{
		ID:    t.ID.Hex(),
		Title: t.Title,
	}
}

func (t *Task) FromModel(task models.Task) error {
	var err error
	t.ID, err = parseOrCreateId(task.ID)
	if err != nil {
		return err
	}

	t.Title = task.Title
	return nil
}
