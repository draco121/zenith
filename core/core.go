package core

import (
	"context"
	"github.com/draco121/horizon/models"
	"github.com/draco121/horizon/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"zenith/repository"
)

type IBotService interface {
	CreateBot(ctx context.Context, bot *models.Bot) (*models.Bot, error)
	UpdateBot(ctx context.Context, bot *models.Bot) (*models.Bot, error)
	DeleteBot(ctx context.Context, id primitive.ObjectID) (*models.Bot, error)
	GetBotByName(ctx context.Context, name string) (*models.Bot, error)
	GetBotById(ctx context.Context, id primitive.ObjectID) (*models.Bot, error)
	GetBotsByProjectId(ctx context.Context, projectId primitive.ObjectID) (*[]models.Bot, error)
}

type botService struct {
	IBotService
	repo   repository.IBotRepository
	client *mongo.Client
}

func NewBotService(client *mongo.Client, repository repository.IBotRepository) IBotService {
	return &botService{
		repo:   repository,
		client: client,
	}
}

func (s *botService) CreateBot(ctx context.Context, bot *models.Bot) (*models.Bot, error) {
	mongoSession, err := s.client.StartSession()
	if err != nil {
		utils.Logger.Error("failed to start mongo mongoSession", "error: ", err.Error())
		return nil, err
	}
	defer mongoSession.EndSession(ctx)
	err = mongoSession.StartTransaction()
	if err != nil {
		utils.Logger.Error("failed to start mongo transaction", "error: ", err.Error())
		return nil, err
	}
	bot, err = s.repo.InsertOne(ctx, bot)
	if err != nil {
		utils.Logger.Error("failed to insert new bot", "error: ", err.Error())
		return nil, err
	} else {
		_ = mongoSession.CommitTransaction(ctx)
		utils.Logger.Info("inserted new bot", "id", bot.ID)
		return bot, nil
	}
}

func (s *botService) UpdateBot(ctx context.Context, bot *models.Bot) (*models.Bot, error) {
	mongoSession, err := s.client.StartSession()
	if err != nil {
		utils.Logger.Error("failed to start mongo mongoSession", "error: ", err.Error())
		return nil, err
	}
	defer mongoSession.EndSession(ctx)
	err = mongoSession.StartTransaction()
	if err != nil {
		utils.Logger.Error("failed to start mongo transaction", "error: ", err.Error())
		return nil, err
	}
	bot, err = s.repo.UpdateOne(ctx, bot)
	if err != nil {
		utils.Logger.Error("failed to update bot", "error: ", err.Error())
		return nil, err
	} else {
		_ = mongoSession.CommitTransaction(ctx)
		utils.Logger.Info("updated bot", "id", bot.ID)
		return bot, nil
	}
}

func (s *botService) DeleteBot(ctx context.Context, id primitive.ObjectID) (*models.Bot, error) {
	mongoSession, err := s.client.StartSession()
	if err != nil {
		utils.Logger.Error("failed to start mongo mongoSession", "error: ", err.Error())
		return nil, err
	}
	defer mongoSession.EndSession(ctx)
	err = mongoSession.StartTransaction()
	if err != nil {
		utils.Logger.Error("failed to start mongo transaction", "error: ", err.Error())
		return nil, err
	}
	bot, err := s.repo.DeleteOneById(ctx, id)
	if err != nil {
		utils.Logger.Error("failed to delete bot", "error: ", err.Error())
		return nil, err
	} else {
		_ = mongoSession.CommitTransaction(ctx)
		utils.Logger.Info("deleted bot", "id", id)
		return bot, nil
	}
}

func (s *botService) GetBotByName(ctx context.Context, name string) (*models.Bot, error) {
	mongoSession, err := s.client.StartSession()
	if err != nil {
		utils.Logger.Error("failed to start mongo mongoSession", "error: ", err.Error())
		return nil, err
	}
	defer mongoSession.EndSession(ctx)
	err = mongoSession.StartTransaction()
	if err != nil {
		utils.Logger.Error("failed to start mongo transaction", "error: ", err.Error())
		return nil, err
	}
	bot, err := s.repo.FindOneByName(ctx, name)
	if err != nil {
		utils.Logger.Error("failed to find bot", "error: ", err.Error())
		return nil, err
	} else {
		_ = mongoSession.CommitTransaction(ctx)
		utils.Logger.Info("found bot", "id", bot.ID)
		return bot, nil
	}
}

func (s *botService) GetBotById(ctx context.Context, id primitive.ObjectID) (*models.Bot, error) {
	mongoSession, err := s.client.StartSession()
	if err != nil {
		utils.Logger.Error("failed to start mongo mongoSession", "error: ", err.Error())
		return nil, err
	}
	defer mongoSession.EndSession(ctx)
	err = mongoSession.StartTransaction()
	if err != nil {
		utils.Logger.Error("failed to start mongo transaction", "error: ", err.Error())
		return nil, err
	}
	bot, err := s.repo.FindOneById(ctx, id)
	if err != nil {
		utils.Logger.Error("failed to find bot", "error: ", err.Error())
		return nil, err
	} else {
		_ = mongoSession.CommitTransaction(ctx)
		utils.Logger.Info("found bot", "id", bot.ID)
		return bot, nil
	}
}

func (s *botService) GetBotsByProjectId(ctx context.Context, projectId primitive.ObjectID) (*[]models.Bot, error) {
	mongoSession, err := s.client.StartSession()
	if err != nil {
		utils.Logger.Error("failed to start mongo mongoSession", "error: ", err.Error())
		return nil, err
	}
	defer mongoSession.EndSession(ctx)
	err = mongoSession.StartTransaction()
	if err != nil {
		utils.Logger.Error("failed to start mongo transaction", "error: ", err.Error())
		return nil, err
	}
	bots, err := s.repo.FindManyByProjectId(ctx, projectId)
	if err != nil {
		utils.Logger.Error("failed to find bots", "error: ", err.Error())
		return nil, err
	} else {
		_ = mongoSession.CommitTransaction(ctx)
		utils.Logger.Info("fetched bots")
		return bots, nil
	}
}
