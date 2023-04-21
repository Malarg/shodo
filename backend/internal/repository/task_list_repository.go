package repository

import (
	"context"
	"errors"
	"fmt"
	"shodo/internal/config"
	"shodo/models"

	mongodto "shodo/internal/repository/mongo_dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskListRepository struct {
	Mongo  *mongo.Client
	Config *config.Config
}

func (this *TaskListRepository) CreateTaskList(ctx context.Context, taskList *models.TaskList) error {
	taskListDto := mongodto.TaskList{}
	taskListDto.FromModel(*taskList)

	_, err := this.getTaskListCollection().InsertOne(ctx, taskListDto)
	if err != nil {
		return err
	}

	return nil
}

func (this *TaskListRepository) DeleteTaskList(ctx context.Context, id string) error {
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = this.getTaskListCollection().DeleteOne(ctx, mongodto.TaskList{ID: mongoId})
	if err != nil {
		return err
	}

	return nil
}

func (this *TaskListRepository) AddUserToList(ctx context.Context, listId string, email string) error {
	mongoListId, err := primitive.ObjectIDFromHex(listId)
	if err != nil {
		return err
	}

	var mongoUser mongodto.User
	err = this.getUsersCollection().FindOne(ctx, bson.M{"email": email}).Decode(&mongoUser)
	if err != nil {
		return err
	}

	var list *mongodto.TaskList
	this.getTaskListCollection().FindOne(ctx, bson.M{"_id": mongoListId}).Decode(&list)

	list.SharedWith = append(list.SharedWith, mongoUser.ID)

	_, err = this.getTaskListCollection().UpdateOne(ctx, bson.M{"_id": mongoListId}, bson.M{"$set": bson.M{"sharedWith": list.SharedWith}})
	if err != nil {
		return err
	}

	return nil

}

func (this *TaskListRepository) RemoveUserFromList(ctx context.Context, listId string, email string) error {
	mongoListId, err := primitive.ObjectIDFromHex(listId)
	if err != nil {
		return err
	}

	var mongoUser mongodto.User
	err = this.getUsersCollection().FindOne(ctx, bson.M{"email": email}).Decode(&mongoUser)
	if err != nil {
		return err
	}

	var list *mongodto.TaskList
	this.getTaskListCollection().FindOne(ctx, bson.M{"_id": mongoListId}).Decode(&list)
	for i, u := range list.SharedWith {
		if u == mongoUser.ID {
			list.SharedWith = append(list.SharedWith[:i], list.SharedWith[i+1:]...)
		}
	}

	_, err = this.getTaskListCollection().UpdateOne(ctx, bson.M{"_id": mongoListId}, bson.M{"$set": bson.M{"sharedWith": list.SharedWith}})
	if err != nil {
		return err
	}

	return nil
}

func (this *TaskListRepository) AddTaskToList(ctx context.Context, listId *string, task *models.Task) (*string, error) {
	if listId == nil {
		return nil, errors.New("listId can not be nil")
	}
	mongoListId, err := primitive.ObjectIDFromHex(*listId)
	if err != nil {
		return nil, fmt.Errorf("failed to convertListId to ObjectId: %w", err)
	}

	taskDto := mongodto.Task{}
	taskDto.FromModel(*task)

	var list *mongodto.TaskList
	err = this.getTaskListCollection().FindOne(ctx, bson.M{"_id": mongoListId}).Decode(&list)
	if err != nil {
		return nil, fmt.Errorf("failed to find task list: %w", err)
	}
	list.Tasks = append(list.Tasks, taskDto)
	result, err := this.getTaskListCollection().UpdateOne(ctx, bson.M{"_id": mongoListId}, bson.M{"$set": bson.M{"tasks": list.Tasks}})
	if err != nil {
		return nil, fmt.Errorf("failed to update task list: %w", err)
	}

	if result.ModifiedCount == 0 {
		return nil, errors.New("Failed to add task to database")
	}

	id := taskDto.ID.Hex()

	return &id, nil
}

func (this *TaskListRepository) RemoveTaskFromList(ctx context.Context, listId *string, taskId *string) error {
	mongoListId, err := primitive.ObjectIDFromHex(*listId)
	if err != nil {
		return err
	}

	mongoTaskId, err := primitive.ObjectIDFromHex(*taskId)
	if err != nil {
		return err
	}

	var list *mongodto.TaskList
	err = this.getTaskListCollection().FindOne(ctx, bson.M{"_id": mongoListId}).Decode(&list)
	if err != nil {
		return err
	}
	for i, t := range list.Tasks {
		if t.ID == mongoTaskId {
			list.Tasks = append(list.Tasks[:i], list.Tasks[i+1:]...)
		}
	}

	_, err = this.getTaskListCollection().UpdateOne(ctx, bson.M{"_id": mongoListId}, bson.M{"$set": bson.M{"tasks": list.Tasks}})
	if err != nil {
		return err
	}

	return nil
}

func (this *TaskListRepository) GetTaskList(ctx context.Context, id *string) (models.TaskList, error) {
	mongoId, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return models.TaskList{}, err
	}
	var list *mongodto.TaskList
	err = this.getTaskListCollection().FindOne(ctx, mongodto.TaskList{ID: mongoId}).Decode(&list)
	return list.ToModel(), err
}

func (this *TaskListRepository) GetTaskLists(ctx context.Context, userId string) ([]models.TaskListShort, error) {
	mongoUserId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	var lists []*mongodto.TaskList
	var result []models.TaskListShort
	cursor, err := this.getTaskListCollection().Find(ctx, bson.M{"owner": mongoUserId})
	defer cursor.Close(ctx)

	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &lists)
	if err != nil {
		return nil, err
	}

	for _, list := range lists {
		result = append(result, list.ToShortModel())
	}

	cursor, err = this.getTaskListCollection().Find(ctx, bson.M{"sharedWith": bson.M{"$in": []primitive.ObjectID{mongoUserId}}})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &lists)
	if err != nil {
		return nil, err
	}

	for _, list := range lists {
		result = append(result, list.ToShortModel())
	}

	return result, nil
}

func (this *TaskListRepository) CheckTaskListExists(ctx context.Context, id string) (bool, error) {
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	var list *mongodto.TaskList
	err = this.getTaskListCollection().FindOne(ctx, mongodto.TaskList{ID: mongoId}).Decode(&list)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (this *TaskListRepository) CheckTaskExists(ctx context.Context, listId string, taskId string) (bool, error) {
	mongoListId, err := primitive.ObjectIDFromHex(listId)
	if err != nil {
		return false, err
	}
	mongoTaskId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return false, err
	}

	var list *mongodto.TaskList
	err = this.getTaskListCollection().FindOne(ctx, mongodto.TaskList{ID: mongoListId}).Decode(&list)
	if err != nil {
		return false, err
	}

	for _, t := range list.Tasks {
		if t.ID == mongoTaskId {
			return true, nil
		}
	}

	return false, nil
}

func (this *TaskListRepository) getTaskListCollection() *mongo.Collection {
	return this.Mongo.Database(this.Config.DbName).Collection("task_lists")
}

func (this *TaskListRepository) getUsersCollection() *mongo.Collection {
	return this.Mongo.Database(this.Config.DbName).Collection("users")
}
