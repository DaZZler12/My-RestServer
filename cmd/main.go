package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Dazzler/My-RestServer/pkg/authroute"
	"github.com/Dazzler/My-RestServer/pkg/database"
	"github.com/Dazzler/My-RestServer/pkg/handler"
	"github.com/Dazzler/My-RestServer/pkg/middleware"
	"github.com/Dazzler/My-RestServer/pkg/services/itemservice"
	"github.com/Dazzler/My-RestServer/pkg/services/userservice"
	"github.com/Dazzler/My-RestServer/pkg/store"
	masteryaml "github.com/Dazzler/My-RestServer/pkg/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"go.elastic.co/apm/module/apmgin"
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
	userCollection *mongo.Collection
	mongodb        *mongo.Database
	err            error
	uservice       userservice.UserService
	uhandler       handler.UserController
	usertable      string
	itemtable      string
	serverport     string
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
	dbconfig, err := masteryaml.ExtractYamlForDB()
	if err != nil {
		log.Fatal(err)
	}
	usertable = dbconfig.Database.Usertable
	itemtable = dbconfig.Database.Itemtable
	serverconfig, err := masteryaml.ExtractYamlForServer()
	if err != nil {
		log.Fatal(err)
	}
	serverport = serverconfig.Serverconfig.Port
	serverport = ":" + serverport
	itemcollection, err = store.StoreCreateCollection(mongodb, itemtable) //create the table by calling method from Store
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("database connected", itemcollection)
	userCollection, err = store.StoreCreateCollection(mongodb, usertable) // create the table for user
	if err != nil {
		log.Fatal(err)
	}
	iservice = itemservice.NewItemService(itemcollection, ctx) // initialize the service
	ihandler = handler.NewItemHandler(iservice)

	uservice = userservice.NewUserService(userCollection, ctx)
	uhandler = handler.NewUserController(uservice)
	// initialize the handler/controller
	server = gin.Default() // initialize the gin server
	server.Use(apmgin.Middleware(server))

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	// Logs all panic to error log
	//   - stack means whether output the stack info.
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-BrowserFingerprint", "X-Workspace-ID"},
		ExposeHeaders:    []string{"Content-Length", "Content-Range", "X-Total-Count"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Register the user-related routes
	// Defining the authentication routes within the init function ensures that they are set up
	// and ready to handle requests when the server starts. It provides a clean and centralized
	// way to define the authentication endpoints before the main server starts running.
	// authGroup := r.Group("/auth")
	// {
	// 	authGroup.POST("/signup", uhandler.Signup)
	// 	authGroup.POST("/login", uhandler.Login)
	// }
}

func main() {
	defer database.DisconnectMongoDB(ctx) // disconnect from mongo if application shutdown.
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Initialize the authentication routes
	authroute.SetupAuthRoutes(server, uhandler)
	// Apply the JWT authentication middleware to private routes
	privateRoutes := server.Group("/api")
	privateRoutes.Use(middleware.JWTAuth)
	ihandler.RegisterItemRoutes(privateRoutes)
	log.Fatal(server.Run(serverport))
}
