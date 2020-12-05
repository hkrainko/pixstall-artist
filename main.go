package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer cancel()
	defer func() {
		if err = dbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	r := gin.Default()

	//authGroup := r.Group("/artist")
	//{
	//	ctr := InitAuthController(conn, dbClient.Database("pixstall-user"))
	//	authGroup.POST("/getAuthUrl", ctr.GetAuthURL)
	//}

	err = r.Run(":9002")
	print(err)
}
