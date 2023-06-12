package itemservice

import (
	models "github.com/Dazzler/My-RestServer/pkg/models"
	"github.com/gin-gonic/gin"
)

// here the interface has been created that will help us to
// define the service-contracts/ API-contracts
type ItemService interface {
	CreateItem(*models.Item) error                             // used to create abn item
	GetItem(*string) (*models.Item, error)                     // used to get an item
	GetAllItem(*gin.Context, int, int) ([]*models.Item, error) // used to get all the item and []*models.Item   --> this is a slice that will contain all the item-objects
	UpdateItem(*models.Item) error                             // used to update an item using PETCH
	UpdateWholeItem(*models.Item) error                        // used to update an whole item using the PUT
	DeleteItem(*string) error                                  // used to delete an item
}
