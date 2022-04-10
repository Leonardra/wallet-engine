package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"walletEngine/configs"
	"walletEngine/models"
	"walletEngine/util"
)

func CreateMongoRepository() *MongoRepository{
	mongoRepository := new(MongoRepository)
	mongoRepository.Collection =  configs.GetCollection(configs.DbClient, "wallets")
	return mongoRepository
}

type MongoRepository struct{
	Collection *mongo.Collection
}



func (repository *MongoRepository) CreateWallet(wallet models.Wallet) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err :=repository.Collection.InsertOne(ctx, wallet)

	if err != nil {
		util.ApplicationLog.Printf("Error Saving wallet %v\n", err)
		return  err
	}
	return nil
}

func (repository *MongoRepository) GetWallet(id primitive.ObjectID) (*models.Wallet, error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var foundWallet models.Wallet

	filter := bson.D{{"_id", id}}
	err := repository.Collection.FindOne(ctx, filter).Decode(&foundWallet)
	if err == mongo.ErrNoDocuments {
		return nil , err
	} else if err != nil {
		return nil, err
	}

	return &foundWallet, nil
}

func (repository *MongoRepository) UpdateWallet(id primitive.ObjectID, wallet *models.Wallet) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"_id", id}}
	singleResult :=repository.Collection.FindOneAndReplace(ctx, filter, wallet)
	err := singleResult.Err()
	if err == mongo.ErrNoDocuments {
		return err
	} else if err != nil {
		return err
	}
	return nil
}

