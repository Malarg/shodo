package mongodto

import (
	"shodo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskList struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Title      string             `bson:"title,omitempty"`
	Owner      primitive.ObjectID `bson:"owner,omitempty"`
	SharedWith []primitive.ObjectID
	Tasks      []Task `bson:"tasks,omitempty"`
}

func (l *TaskList) New() {
	l.ID = primitive.ObjectID{}
	l.Title = ""
	l.Owner = primitive.ObjectID{}
	l.SharedWith = []primitive.ObjectID{}
	l.Tasks = []Task{}
}

func (l *TaskList) ToModel() models.TaskList {
	return models.TaskList{
		ID:         l.ID.Hex(),
		Title:      l.Title,
		Owner:      l.Owner.Hex(),
		SharedWith: parseIdList(l.SharedWith),
		Tasks:      dtoTasksToModelTasks(l.Tasks),
	}
}

func (l *TaskList) ToShortModel() models.TaskListShort {
	return models.TaskListShort{
		ID:    l.ID.Hex(),
		Title: l.Title,
		Owner: l.Owner.Hex(),
	}
}

func (l *TaskList) FromModel(taskList models.TaskList) error {
	var err error
	l.ID, err = parseOrCreateId(taskList.ID)
	if err != nil {
		return err
	}

	ownerId, err := primitive.ObjectIDFromHex(taskList.Owner)
	if err != nil {
		return err
	}

	l.Title = taskList.Title
	l.Owner = ownerId
	l.Tasks = modelTasksToDtoTasks(taskList.Tasks)

	return nil
}
