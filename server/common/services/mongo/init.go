package mongo

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"

	"stolencarsproject/server/config"
)

const dbHost = config.MongoHost
const dbPort = config.MongoPort
const dbName = config.MongoDbName

var client *mongo.Client
var ctx context.Context

func init() {
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
}
