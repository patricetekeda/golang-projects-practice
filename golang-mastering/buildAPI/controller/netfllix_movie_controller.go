package controller

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const connectionString = "mongodb+srv://ptekeda34:<db_password>@cluster0.x94xvoj.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection



func init() {}




