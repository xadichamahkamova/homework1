package storage

import (
	"context"
	"errors"
	"service/models"
	"service/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepo struct {
	MDB *pkg.Mongo
}

func NewTaskRepo(db *pkg.Mongo) *TaskRepo {
	return &TaskRepo{db}
}

var ctx = context.Background()

func(db *TaskRepo) CreateTask(req *models.Task) (*models.Result, error) {

	resp := models.Result{}

	result, err := db.MDB.Collection.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}
	
	objectId, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		resp.Id = objectId.Hex()
	}else {
		return nil, errors.New("unexpected type")
	}
	resp.Status = "Task created succesfully"
	
	return &resp, nil
}

func(db *TaskRepo) GetTaskById(id string) (*models.Task, error) {

	resp := models.Task{}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id":objectId}

	result := db.MDB.Collection.FindOne(ctx, filter)

	if err := result.Decode(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func(db *TaskRepo) ListOfTask() ([]models.Task, error) {

	resp := []models.Task{}
	cursor, err := db.MDB.Collection.Find(ctx, options.Find())
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		item := models.Task{}
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		resp = append(resp, item)
	}

	return resp, nil
}

func(db *TaskRepo) UpdateTask(req *models.Task) (*models.Result, error) {

	resp := models.Result{}

	objectId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id":objectId}
	update := bson.M{
		"$set" :bson.M{
			"title": req.Title,
			"description":req.Description,
			"status":req.Status,
		},
	}

	result := db.MDB.Collection.FindOneAndUpdate(ctx, filter, update)
	if err := result.Decode(&resp); err != nil {
		return nil, err
	}
	resp.Status = "Task updated succesfully"
	return &resp, nil
}

func(db *TaskRepo) DeleteTask(id string) (*models.Result, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id":objectId}
	result, err := db.MDB.Collection.DeleteOne(ctx, filter)
	if err != nil || result.DeletedCount == 0 {
		return nil, err
	}
	
	return &models.Result{Id: id, Status: "User deleted succesfully"}, nil
}