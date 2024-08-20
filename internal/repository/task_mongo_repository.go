package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/harlitad/task-management-app/internal/model"
	"github.com/harlitad/task-management-app/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTaskRepository struct {
	Client            *mongo.Client
	Logger            logger.ILogger
	DBName            string
	TaskCollectionStr string
	TaskCollection    *mongo.Collection
}

func NewMongoTaskRepository(logger logger.ILogger, client *mongo.Client) ITaskRepository {
	collection := client.Database("tma").Collection("tasks")
	return &MongoTaskRepository{
		Client:            client,
		DBName:            "tma",
		TaskCollectionStr: "tasks",
		TaskCollection:    collection,
	}
}

func (r *MongoTaskRepository) UpdateTaskById(c context.Context, id uuid.UUID, newTask model.Task) (model.Task, error) {
	filter := bson.M{"id": id}

	// Use the helper function to build the update document
	update := buildUpdateDocument(newTask)

	// Perform the update operation
	result := r.TaskCollection.FindOneAndUpdate(c, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))

	var updatedTask model.Task
	err := result.Decode(&updatedTask)
	if err == mongo.ErrNoDocuments {
		return model.Task{}, fmt.Errorf("task with id %s not found", id)
	}
	if err != nil {
		return model.Task{}, fmt.Errorf("failed to update task, %w", err)
	}

	return updatedTask, nil
}

func (r *MongoTaskRepository) GetTasks(c context.Context) ([]model.Task, error) {
	var tasks []model.Task
	cursor, err := r.TaskCollection.Find(c, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks, %w", err)
	}
	if err := cursor.All(c, &tasks); err != nil {
		return nil, fmt.Errorf("failed to parse tasks, %w", err)
	}
	return tasks, nil
}

func (r *MongoTaskRepository) GetTasksToday(c context.Context) ([]model.Task, error) {
	var tasks []model.Task
	now := time.Now().UTC()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay := startOfDay.Add(24 * time.Hour)

	filter := bson.M{
		"due_date": bson.M{
			"$gte": startOfDay,
			"$lt":  endOfDay,
		},
	}

	cursor, err := r.TaskCollection.Find(c, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks for today, %w", err)
	}

	if err := cursor.All(c, &tasks); err != nil {
		return nil, fmt.Errorf("failed to parse tasks, %w", err)
	}

	return tasks, nil
}

func (r *MongoTaskRepository) GetTaskById(c context.Context, id uuid.UUID) (model.Task, error) {
	var task model.Task

	filter := bson.M{"id": id}

	err := r.TaskCollection.FindOne(c, filter).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return task, model.ErrorRecordNotFound
	}
	if err != nil {
		return task, fmt.Errorf("failed to get task by id, %w", err)
	}
	return task, nil
}

func (r *MongoTaskRepository) Get(c context.Context, id uuid.UUID, userId uuid.UUID) (model.Task, error) {
	var task model.Task

	objectId, err := primitive.ObjectIDFromHex(id.String())
	if err != nil {
		return task, err
	}
	creatorId, err := primitive.ObjectIDFromHex(userId.String())
	if err != nil {
		return task, err
	}

	filter := bson.M{"id": objectId, "creator_id": creatorId}
	err = r.TaskCollection.FindOne(c, filter).Decode(&task)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *MongoTaskRepository) Delete(c context.Context, id uuid.UUID) error {
	filter := bson.M{"id": id}
	res, err := r.TaskCollection.DeleteOne(c, filter)
	if err != nil {
		return fmt.Errorf("failed to delete task, %w", err)
	}
	if res.DeletedCount == 0 {
		return model.ErrorRecordNotFound
	}
	return nil
}

func (r *MongoTaskRepository) Create(c context.Context, newTask model.Task) error {
	_, err := r.TaskCollection.InsertOne(c, newTask)
	if err != nil {
		return fmt.Errorf("failed to create task, %w", err)
	}
	return nil
}

func (r *MongoTaskRepository) ListByAssigneeId(c context.Context, assigneeId string) ([]model.Task, error) {
	var tasks []model.Task
	filter := bson.M{"assignee_id": assigneeId}
	cursor, err := r.TaskCollection.Find(c, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks by assignee id, %w", err)
	}
	if err := cursor.All(c, &tasks); err != nil {
		return nil, fmt.Errorf("failed to parse tasks, %w", err)
	}
	return tasks, nil
}

func (r *MongoTaskRepository) GetTaskByAssigneeId(c context.Context, id uuid.UUID, assigneeId uuid.UUID) (model.Task, error) {
	var task model.Task
	filter := bson.M{"id": id.String(), "assignee_id": assigneeId.String()}
	err := r.TaskCollection.FindOne(c, filter).Decode(&task)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *MongoTaskRepository) GetTaskByDueDate(c context.Context, date time.Time) ([]model.Task, error) {
	var tasks []model.Task
	filter := bson.M{"due_date": date.Format("2006-01-02")}
	cursor, err := r.TaskCollection.Find(c, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks by due date, %w", err)
	}
	if err := cursor.All(c, &tasks); err != nil {
		return nil, fmt.Errorf("failed to parse tasks, %w", err)
	}
	return tasks, nil
}

func (r *MongoTaskRepository) UpdateStatus(c context.Context, id uuid.UUID, status string) error {
	filter := bson.M{"id": id.String()}
	update := bson.M{"$set": bson.M{"status": status}}
	_, err := r.TaskCollection.UpdateOne(c, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update task status, %w", err)
	}
	return nil
}

func buildUpdateDocument(task model.Task) bson.M {
	updateFields := bson.M{}

	if task.Title != "" {
		updateFields["title"] = task.Title
	}
	if task.Description != "" {
		updateFields["description"] = task.Description
	}
	if task.Status != "" {
		updateFields["status"] = task.Status
	}
	if !task.DueDate.IsZero() {
		updateFields["due_date"] = task.DueDate
	}
	if !task.CreatedAt.IsZero() {
		updateFields["created_at"] = task.CreatedAt
	}
	if !task.UpdatedAt.IsZero() {
		updateFields["updated_at"] = task.UpdatedAt
	}
	if task.AssigneeId != uuid.Nil {
		updateFields["assignee_id"] = task.AssigneeId
	}
	if task.CreatorId != uuid.Nil {
		updateFields["creator_id"] = task.CreatorId
	}

	return bson.M{"$set": updateFields}
}
