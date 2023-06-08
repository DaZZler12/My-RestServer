package handler

import (
	"net/http"

	models "github.com/Dazzler/My-RestServer/pkg/models"
	"github.com/Dazzler/My-RestServer/pkg/services/itemservice"
	"github.com/gin-gonic/gin"
)

// the controller will also be of type struct
// and the controller will interect with the service layer interface
type ItemController struct {
	ItemService itemservice.ItemService
}

// defining the constructor i.e. the controller will
// need to have a way by which they can initialize the services..
// thus we are defining the constrcutor below.

func NewItemHandler(itemservice itemservice.ItemService) ItemController {
	return ItemController{
		ItemService: itemservice,
	}
}

// now from the handler we will call the routes..
// and the specific router will call the specific method..

func (itemcontroller *ItemController) CreateItem(ctx *gin.Context) {
	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := itemcontroller.ItemService.CreateItem(&item)
	if err != nil {
		// BADGATEWAY BECAUSE THE ERROR OCCURED IN SAVING INTO MONGODB
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully Created"})
}

func (itemcontroller *ItemController) GetItem(ctx *gin.Context) {
	itemname := ctx.Param("name")
	item, err := itemcontroller.ItemService.GetItem(&itemname)
	if err != nil {
		// BADGATEWAY BECAUSE THE ERROR OCCURED IN SAVING TO MONGODB
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, item)

}

func (itemcontroller *ItemController) GetAllItem(ctx *gin.Context) {
	allitem, err := itemcontroller.ItemService.GetAllItem()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, allitem)

}

func (itemcontroller *ItemController) UpdateItem(ctx *gin.Context) {
	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := itemcontroller.ItemService.UpdateItem(&item)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully Updated"})

}
func (itemcontroller *ItemController) UpdateWholeItem(ctx *gin.Context) {
	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := itemcontroller.ItemService.UpdateWholeItem(&item)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully Updated the Whole item"})
}

func (itemcontroller *ItemController) DeleteItem(ctx *gin.Context) {
	var itemname string = ctx.Param("name")
	err := itemcontroller.ItemService.DeleteItem(&itemname)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully Deleted"})
}

// now we will define the routes. for that we will create the receiv er method..
func (itemcontroller *ItemController) RegisterItemRoues(rg *gin.RouterGroup) {

	// here we are grouping all the item-routes
	// under one name called as itemr
	itemroute := rg.Group("/item")
	itemroute.POST("", itemcontroller.CreateItem)
	itemroute.GET("/:name", itemcontroller.GetItem)
	itemroute.GET("", itemcontroller.GetAllItem)
	itemroute.PATCH("/", itemcontroller.UpdateItem)
	itemroute.PUT("/", itemcontroller.UpdateWholeItem)
	itemroute.DELETE("/:name", itemcontroller.DeleteItem)
}
