package repository

import (
	"context"
	"shodo/internal/models"

	mongodto "shodo/internal/repository/mongo_dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskListRepository struct {
	Mongo *mongo.Client
}

func (this *TaskListRepository) CreateTaskList(taskList *models.TaskList) error {
	taskListDto := mongodto.TaskList{}
	taskListDto.FromModel(*taskList)

	_, err := this.getCollection().InsertOne(nil, taskListDto)
	if err != nil {
		return err
	}

	return nil
}

func (this *TaskListRepository) DeleteTaskList(id string) error {
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = this.getCollection().DeleteOne(nil, mongodto.TaskList{ID: mongoId})
	if err != nil {
		return err
	}

	return nil
}

func (this *TaskListRepository) AddUserToList(listId string, userId string) error {
	mongoListId, err := primitive.ObjectIDFromHex(listId)
	if err != nil {
		return err
	}

	mongoUserId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	var list *mongodto.TaskList
	this.getCollection().FindOne(context.Background(), bson.M{"_id": mongoListId}).Decode(&list)
	list.SharedWith = append(list.SharedWith, mongoUserId)

	_, err = this.getCollection().UpdateOne(context.Background(), bson.M{"_id": mongoListId}, bson.M{"$set": bson.M{"sharedWith": list.SharedWith}})
	if err != nil {
		return err
	}

	return nil

}

func (this *TaskListRepository) RemoveUserFromList(listId string, userId string) error {
	mongoListId, err := primitive.ObjectIDFromHex(listId)
	if err != nil {
		return err
	}

	mongoUserId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	var list *mongodto.TaskList
	this.getCollection().FindOne(context.Background(), bson.M{"_id": mongoListId}).Decode(&list)
	for i, u := range list.SharedWith {
		if u == mongoUserId {
			list.SharedWith = append(list.SharedWith[:i], list.SharedWith[i+1:]...)
		}
	}

	_, err = this.getCollection().UpdateOne(context.Background(), bson.M{"_id": mongoListId}, bson.M{"$set": bson.M{"sharedWith": list.SharedWith}})
	if err != nil {
		return err
	}

	return nil
}

func (this *TaskListRepository) AddTaskToList(listId *string, task *models.Task) error {
	mongoListId, err := primitive.ObjectIDFromHex(*listId)
	if err != nil {
		return err
	}

	taskDto := mongodto.Task{}
	taskDto.FromModel(*task)

	var list *mongodto.TaskList
	err = this.getCollection().FindOne(context.Background(), bson.M{"_id": mongoListId}).Decode(&list)
	if err != nil {
		return err
	}
	list.Tasks = append(list.Tasks, taskDto)
	_, err = this.getCollection().UpdateOne(context.Background(), bson.M{"_id": mongoListId}, bson.M{"$set": bson.M{"tasks": list.Tasks}})
	if err != nil {
		return err
	}

	return nil
}

func (this *TaskListRepository) RemoveTaskFromList(listId *string, taskId *string) error {
	mongoListId, err := primitive.ObjectIDFromHex(*listId)
	if err != nil {
		return err
	}

	mongoTaskId, err := primitive.ObjectIDFromHex(*taskId)
	if err != nil {
		return err
	}

	var list *mongodto.TaskList
	err = this.getCollection().FindOne(context.Background(), bson.M{"_id": mongoListId}).Decode(&list)
	if err != nil {
		return err
	}
	for i, t := range list.Tasks {
		if t.ID == mongoTaskId {
			list.Tasks = append(list.Tasks[:i], list.Tasks[i+1:]...)
		}
	}

	_, err = this.getCollection().UpdateOne(context.Background(), bson.M{"_id": mongoListId}, bson.M{"$set": bson.M{"tasks": list.Tasks}})
	if err != nil {
		return err
	}

	return nil
}

func (this *TaskListRepository) GetTaskList(id *string) (models.TaskList, error) {
	mongoId, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return models.TaskList{}, err
	}
	var list *mongodto.TaskList
	err = this.getCollection().FindOne(nil, mongodto.TaskList{ID: mongoId}).Decode(&list)
	return list.ToModel(), err
}

func (this *TaskListRepository) getCollection() *mongo.Collection {
	return this.Mongo.Database("shodo").Collection("task_lists")
}

func (this *TaskListRepository) GetTaskListsForUser(userId *string) ([]models.TaskList, error) {
	mongoUserId, err := primitive.ObjectIDFromHex(*userId)
	if err != nil {
		return nil, err
	}

	var lists []*mongodto.TaskList
	var result []models.TaskList
	cursor, err := this.getCollection().Find(nil, bson.M{"owner": mongoUserId})
	defer cursor.Close(nil)

	if err != nil {
		return nil, err
	}
	err = cursor.All(nil, &lists)
	if err != nil {
		return nil, err
	}

	for _, list := range lists {
		result = append(result, list.ToModel())
	}

	cursor, err = this.getCollection().Find(nil, bson.M{"sharedWith": bson.M{"$in": []primitive.ObjectID{mongoUserId}}})
	if err != nil {
		return nil, err
	}
	err = cursor.All(nil, &lists)
	if err != nil {
		return nil, err
	}

	for _, list := range lists {
		result = append(result, list.ToModel())
	}

	return result, nil
}
