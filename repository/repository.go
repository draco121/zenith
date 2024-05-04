package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/draco121/common/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IBotRepository interface {
	InsertOne(ctx context.Context, bot *models.Bot) (*models.Bot, error)
	UpdateOne(ctx context.Context, bot *models.Bot) (*models.Bot, error)
	FindOneById(ctx context.Context, id primitive.ObjectID) (*models.Bot, error)
	FindOneByName(ctx context.Context, name string) (*models.Bot, error)
	DeleteOneById(ctx context.Context, id primitive.ObjectID) (*models.Bot, error)
	FindManyByProjectId(ctx context.Context, projectId primitive.ObjectID) (*[]models.Bot, error)
}

type botRepository struct {
	IBotRepository
	db *mongo.Database
}

func NewBotRepository(database *mongo.Database) IBotRepository {
	return &botRepository{db: database}
}

func (ur *botRepository) InsertOne(ctx context.Context, bot *models.Bot) (*models.Bot, error) {
	ownerId := ctx.Value("UserId").(primitive.ObjectID)
	result, _ := ur.FindOneByName(ctx, bot.Name)
	if result != nil {
		return nil, fmt.Errorf("record exists")
	} else {
		bot.ID = primitive.NewObjectID()
		bot.Owner = ownerId
		_, err := ur.db.Collection("bots").InsertOne(ctx, bot)
		if err != nil {
			return nil, err
		}
	}
	return bot, nil
}

func (ur *botRepository) UpdateOne(ctx context.Context, bot *models.Bot) (*models.Bot, error) {
	userId := ctx.Value("UserId").(string)
	ownerId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: bot.ID}, {Key: "owner", Value: ownerId}}
	bot.Owner = ownerId
	update := bson.M{"$set": bot}
	result := models.Bot{}
	err = ur.db.Collection("bots").FindOneAndUpdate(ctx, filter, update).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else {
		return &result, nil
	}
}

func (ur *botRepository) FindOneById(ctx context.Context, id primitive.ObjectID) (*models.Bot, error) {
	userId := ctx.Value("UserId").(primitive.ObjectID)
	filter := bson.D{{Key: "_id", Value: id}, {Key: "owner", Value: userId}}
	result := models.Bot{}
	err := ur.db.Collection("bots").FindOne(ctx, filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else {
		return &result, nil
	}

}

func (ur *botRepository) FindOneByName(ctx context.Context, name string) (*models.Bot, error) {
	ownerId := ctx.Value("UserId").(primitive.ObjectID)
	filter := bson.D{{Key: "name", Value: name}, {Key: "owner", Value: ownerId}}
	result := models.Bot{}
	err := ur.db.Collection("bots").FindOne(ctx, filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else {
		return &result, nil
	}
}

func (ur *botRepository) DeleteOneById(ctx context.Context, id primitive.ObjectID) (*models.Bot, error) {
	userId := ctx.Value("UserId").(primitive.ObjectID)

	filter := bson.D{{Key: "_id", Value: id}, {Key: "owner", Value: userId}}
	result := models.Bot{}
	err := ur.db.Collection("bots").FindOneAndDelete(ctx, filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else {
		return &result, nil
	}
}

func (ur *botRepository) FindManyByProjectId(ctx context.Context, projectId primitive.ObjectID) (*[]models.Bot, error) {
	userId := ctx.Value("UserId").(primitive.ObjectID)
	var bots []models.Bot
	filter := bson.D{{Key: "projectId", Value: projectId}, {Key: "Owner", Value: userId}}
	cur, err := ur.db.Collection("bots").Find(ctx, filter)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	} else {
		if err = cur.All(context.TODO(), &bots); err != nil {
			return nil, err
		}
		return &bots, nil
	}
}
