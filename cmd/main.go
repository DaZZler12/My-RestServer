package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Dazzler/My-RestServer/pkg/database"
	"github.com/Dazzler/My-RestServer/pkg/handler"
	"github.com/Dazzler/My-RestServer/pkg/services/itemservice"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// these are variables that we required thus declared
// and will be the initialized in the init method..
var (
	server         *gin.Engine // server for the gin-framework
	iservice       itemservice.ItemService
	ihandler       handler.ItemController
	ctx            context.Context
	itemcollection *mongo.Collection
	mongodb        *mongo.Database
	err            error
)

func init() {
	// context is used to deal with time-out as of now we are not dealing with time out..
	// thus we will use TODO()
	ctx = context.TODO() // thus todo will create context object with no-cancellation

	// below is the mongo-connection-logic
	mongodb, err = database.ConnectToMongoDB()
	if err != nil {
		// if there is any error we will close the application
		log.Fatal(err)
	}
	fmt.Print("database connected", mongodb)

	itemcollection = mongodb.Collection("item-db")             //create the table
	iservice = itemservice.NewItemService(itemcollection, ctx) // initialize the service
	ihandler = handler.NewItemHandler(iservice)                // initialize the handler/controller
	server = gin.Default()                                     // initialize the gin server
}

func main() {

	defer database.DisconnectMongoDB(ctx) // discount from mongo if application shutdown.
	basepath := server.Group("/api")      // thus path will be: ==>  /api/item/create
	ihandler.RegisterItemRoues(basepath)
	log.Fatal(server.Run(":8080"))
}
