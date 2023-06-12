package handler

import (
	"net/http"
	"strconv"

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

// CreateItem godoc
// @Summary      Create an item
// @Description  Create a new item
// @Tags         item
// @Accept       json
// @Produce      json
// @Param        item   body     models.Item  true  "Item object"
// @Success      200  {object}  gin.H        "message: Successfully Created"
// @Failure      400  {object}  gin.H        "message: Bad Request"
// @Failure      502  {object}  gin.H        "message: Bad Gateway"
// @Router       /api/item [post]
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

// GetItem godoc
// @Summary      Get an item
// @Description  Get an item by its name
// @Tags         item
// @Accept       json
// @Produce      json
// @Param        name   path     string  true  "Item name"
// @Success      200  {object}  models.Item
// @Failure      400  {object}  gin.H    "message: Bad Request"
// @Failure      502  {object}  gin.H    "message: Bad Gateway"
// @Router       /api/item/:name [get]
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

// GetAllItem godoc
// @Summary      Get all items with pagination
// @Description  Get all items with optional pagination parameters
// @Tags         item
// @Accept       json
// @Produce      json
// @Param        _start   query    integer  false  "Start index for pagination"
// @Param        _end     query    integer  false  "End index for pagination"
// @Success      200  {array}   models.Item
// @Failure      400  {object}  gin.H    "message: Bad Request"
// @Failure      502  {object}  gin.H    "message: Bad Gateway"
// @Router       /api/item [get]
func (itemcontroller *ItemController) GetAllItem(ctx *gin.Context) {
	// Extract the start and end query parameters
	start := ctx.DefaultQuery("_start", "0")
	end := ctx.DefaultQuery("_end", "4")

	// Convert the start and end values to integers
	startInt, _ := strconv.Atoi(start)
	endInt, _ := strconv.Atoi(end)

	// Call the GetAllItem method with the start and end values
	allitem, err := itemcontroller.ItemService.GetAllItem(ctx, startInt, endInt)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, allitem)

}

// UpdateItem godoc
// @Summary      Update an item
// @Description  Update an item with new data
// @Tags         item
// @Accept       json
// @Produce      json
// @Param        item     body     models.Item  true  "Item object to update"
// @Success      200  {object}  gin.H         "message: Successfully Updated"
// @Failure      400  {object}  gin.H         "message: Bad Request"
// @Failure      502  {object}  gin.H         "message: Bad Gateway"
// @Router       /api/item [patch]
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

// UpdateWholeItem godoc
// @Summary      Update the whole item
// @Description  Update the whole item with new data
// @Tags         item
// @Accept       json
// @Produce      json
// @Param        item     body     models.Item  true  "Item object to update"
// @Success      200  {object}  gin.H         "message: Successfully Updated the Whole item"
// @Failure      400  {object}  gin.H         "message: Bad Request"
// @Failure      502  {object}  gin.H         "message: Bad Gateway"
// @Router       /api/item [put]
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

// DeleteItem godoc
// @Summary      Delete an item
// @Description  Delete an item by name
// @Tags         item
// @Accept       json
// @Produce      json
// @Param        name     path     string  true  "Item name to delete"
// @Success      200  {object}  gin.H   "message: Successfully Deleted"
// @Failure      400  {object}  gin.H   "message: Bad Gateway"
// @Failure      404  {object}  gin.H   "message: Not Found"
// @Router       /items/{name} [delete]
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
func (itemcontroller *ItemController) RegisterItemRoutes(rg *gin.RouterGroup) {
	itemroute := rg.Group("/item")
	itemroute.POST("", itemcontroller.CreateItem)
	itemroute.GET("/:name", itemcontroller.GetItem)
	itemroute.GET("", itemcontroller.GetAllItem)
	itemroute.PATCH("", itemcontroller.UpdateItem)
	itemroute.PUT("", itemcontroller.UpdateWholeItem)
	itemroute.DELETE("/:name", itemcontroller.DeleteItem)
}
