package appmongo

import (
	"context"
	"log"

	"github.com/lea55/BACKENDCANDITICKET/src/global/core"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func Start() {
	cnn := core.GetAppConfig()
	opts := options.Client().ApplyURI(cnn.MongoCnn)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("CANDITICKETS: Error en conexión a la base de datos mongodb", err.Error())
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("CANDITICKETS: error en ping a la base de datos", err.Error())
	}

	db = client.Database(cnn.MongoDBName)
	log.Println("CANDITICKETS: Base de datos mongo db conectada y lista para usarse")
}

func GetCollection(name string) *mongo.Collection {
	if db == nil {
		Start()
	}

	return db.Collection(name)
}
