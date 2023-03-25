package mongodto

import (
	"shodo/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func dtoTasksToModelTasks(tasks []Task) []models.Task {
	modelTasks := make([]models.Task, len(tasks))
	for i, task := range tasks {
		modelTasks[i] = task.ToModel()
	}
	return modelTasks
}

func modelTasksToDtoTasks(tasks []models.Task) []Task {
	dtoTasks := make([]Task, len(tasks))
	for i, task := range tasks {
		dtoTasks[i].FromModel(task)
	}
	return dtoTasks
}

func parseOrCreateId(id string) (primitive.ObjectID, error) {
	var objectId primitive.ObjectID
	var err error
	if id == "" {
		objectId = primitive.NewObjectID()

	} else {
		objectId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return objectId, err
		}
	}
	return objectId, nil
}

func parseIdList(ids []primitive.ObjectID) []string {
	idList := make([]string, len(ids))
	for i, id := range ids {
		idList[i] = id.Hex()
	}
	return idList
}
