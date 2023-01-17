package model

import (
	"context"

	conf "github.com/codestates/WBA-BC-Project-02/daemon/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	Client         *mongo.Client
	ColUser        *mongo.Collection
	ColDexContract *mongo.Collection
	ColContract    *mongo.Collection
}

func NewModel() (*Model, error) {
	cf := conf.GetConfig("./config/config.toml")

	r := &Model{}

	var err error
	if r.Client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(cf.DB.Host)); err != nil {
		return nil, err
	} else if err := r.Client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.Client.Database("wemixon")
		r.ColUser = db.Collection("user")
		r.ColContract = db.Collection("contract")
		r.ColDexContract = db.Collection("dex_contract")
	}

	return r, nil
}
