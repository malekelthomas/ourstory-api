package main

import (
	"github.com/malekelthomas/ourstory-api/mongo"
	"github.com/malekelthomas/ourstory-api/pkg/user"
	"github.com/malekelthomas/ourstory-api/server"
)

func main() {
	ourstoryServer := server.NewServer()
	mongoClient := mongo.NewMongoConn("mongodb://localhost:27017", "ourstory")
	umr := user.NewUserMongoRepository(mongoClient.DB.Collection("users"))
	us := user.NewUserService(umr)
	user.RegisterUserRoutes(ourstoryServer.Router, us)

	ourstoryServer.Run(":8080")
}
