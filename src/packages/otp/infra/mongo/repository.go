package mdbotp

import (
	"context"

	"github.com/lea55/BACKENDCANDITICKET/src/global/infrastructure/mongo"
	otp "github.com/lea55/BACKENDCANDITICKET/src/packages/otp/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	cll *mongo.Collection
}

func New() otp.Repository {
	return &repo{
		cll: appmongo.GetCollection("otp_list"),
	}
}

func (r repo) Save(data otp.Entity) error {
	var model Model
	model.fromEntity(data)

	_, err := r.cll.InsertOne(context.Background(), model)
	if err != nil {
		return err
	}

	return nil
}

func (r repo) FindByCode(code string) (otp.Entity, error) {
	var model Model

	filter := bson.M{"code": code}
	err := r.cll.FindOne(context.Background(), filter).Decode(&model)

	if err != nil {
		return otp.Entity{}, err
	}

	result := model.toEntity()

	return result, nil
}
