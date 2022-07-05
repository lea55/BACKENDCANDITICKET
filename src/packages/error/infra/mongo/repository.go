package mdberrlog

import (
	"context"

	appmongo "github.com/lea55/BACKENDCANDITICKET/src/global/infrastructure/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	cll *mongo.Collection
}

func NewRepo() errlog.Repository {
	return &repo{
		cll: appmongo.GetCollection("err_log"),
	}
}

func (r repo) Save(data errlog.Entity) error {
	var mdl model
	mdl.fromEntity(data)

	_, err := r.cll.InsertOne(context.Background(), mdl)
	if err != nil {
		return err
	}
	return nil
}

func (r repo) FindPaginated(page uint32) ([]errlog.Entity, error) {
	return nil, nil
}
