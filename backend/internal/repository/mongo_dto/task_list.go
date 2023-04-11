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

func (this *TaskList) New() {
	this.ID = primitive.ObjectID{}
	this.Title = ""
	this.Owner = primitive.ObjectID{}
	this.SharedWith = []primitive.ObjectID{}
	this.Tasks = []Task{}
}

func (this *TaskList) ToModel() models.TaskList {
	return models.TaskList{
		ID:         this.ID.Hex(),
		Title:      this.Title,
		Owner:      this.Owner.Hex(),
		SharedWith: parseIdList(this.SharedWith),
		Tasks:      dtoTasksToModelTasks(this.Tasks),
	}
}

func (this *TaskList) ToShortModel() models.TaskListShort {
	return models.TaskListShort{
		ID:    this.ID.Hex(),
		Title: this.Title,
		Owner: this.Owner.Hex(),
	}
}

func (this *TaskList) FromModel(taskList models.TaskList) error {
	var err error
	this.ID, err = parseOrCreateId(taskList.ID)
	if err != nil {
		return err
	}

	ownerId, err := primitive.ObjectIDFromHex(taskList.Owner)
	if err != nil {
		return err
	}

	this.Title = taskList.Title
	this.Owner = ownerId
	this.Tasks = modelTasksToDtoTasks(taskList.Tasks)

	return nil
}
