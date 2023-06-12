package itemservice

/*
description: The service layer will interaect with the
database layer thus it has the database object
*/
import (
	"context"
	"errors"
	"strconv"
	"time"

	models "github.com/Dazzler/My-RestServer/pkg/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// it will have struct

type ItemServiceMethod struct {
	itemcollection *mongo.Collection // pointer becuase we will get the address of the object and point to that by the help of itemcollection
	ctx            context.Context
}

// now we need to initialize the ItemServiceMethod struct
// the below method will actas a constructor
func NewItemService(itemcollection *mongo.Collection, ctx context.Context) ItemService {
	return &ItemServiceMethod{
		itemcollection: itemcollection,
		ctx:            ctx,
	}
}

func (u *ItemServiceMethod) CreateItem(item *models.Item) error {
	// adding item to mongoDB
	item.CreatedAt = time.Now()
	_, err := u.itemcollection.InsertOne(u.ctx, item)
	return err
}

func (u *ItemServiceMethod) GetItem(item_name *string) (*models.Item, error) {
	var item *models.Item
	query := bson.D{bson.E{Key: "item_name", Value: item_name}}
	//alternate query below
	// query := bson.M{"item_name": item_name}
	err := u.itemcollection.FindOne(u.ctx, query).Decode(&item)
	return item, err
}

// func (u *ItemServiceMethod) GetAllItem() ([]*models.Item, error) {
// 	// here we need to deal with cursor because we will need to
// 	// get the user one by one from the DB
// 	var itemslice []*models.Item // creating a slice
// 	cursor, err := u.itemcollection.Find(u.ctx, bson.D{{}})
// 	if err != nil {
// 		return nil, err
// 	}
// 	for cursor.Next(u.ctx) {
// 		var item models.Item
// 		err := cursor.Decode(&item)
// 		if err != nil {
// 			return nil, err
// 		}
// 		itemslice = append(itemslice, &item)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}
// 	cursor.Close(u.ctx)

// 	if len(itemslice) == 0 {
// 		return nil, errors.New("documents not found")
// 	}

// 	// Sort itemslice based on creation time in descending order
// 	sort.Slice(itemslice, func(i, j int) bool {
// 		return itemslice[i].CreatedAt.After(itemslice[j].CreatedAt)
// 	})

// 	fmt.Println(itemslice)
// 	return itemslice, nil
// }

func (u *ItemServiceMethod) GetAllItem(ctx *gin.Context, start, end int) ([]*models.Item, error) {
	// Validate start and end values
	if start < 0 || end < 0 || start > end {
		return nil, errors.New("invalid start and end values")
	}

	// Calculate the limit and skip values for pagination
	limit := end - start + 1
	skip := start

	// Query the database with limit and skip parameters
	options := options.Find().SetLimit(int64(limit)).SetSkip(int64(skip)).SetSort(bson.M{"created_at": -1})
	cursor, err := u.itemcollection.Find(u.ctx, bson.D{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(u.ctx)
	var itemslice []*models.Item // creating a slice
	for cursor.Next(u.ctx) {
		var item models.Item
		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}
		itemslice = append(itemslice, &item)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if len(itemslice) == 0 {
		return nil, errors.New("documents not found")
	}
	// Get the total count of items
	totalCount, err := u.itemcollection.CountDocuments(u.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	// Set the X-Total-Count header in the response
	ctx.Header("X-Total-Count", strconv.Itoa(int(totalCount)))

	return itemslice, nil
}

func (u *ItemServiceMethod) UpdateItem(item *models.Item) error {
	filter := bson.D{primitive.E{Key: "item_name", Value: item.Item_Name}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "item_name", Value: item.Item_Name}, primitive.E{Key: "brand", Value: item.Brand}, primitive.E{Key: "model", Value: item.Model}, primitive.E{Key: "year", Value: item.Year}, primitive.E{Key: "price", Value: item.Price}}}}
	result, _ := u.itemcollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *ItemServiceMethod) UpdateWholeItem(item *models.Item) error {
	filter := bson.D{primitive.E{Key: "item_name", Value: item.Item_Name}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "brand", Value: item.Brand},
			primitive.E{Key: "model", Value: item.Model},
			primitive.E{Key: "year", Value: item.Year},
			primitive.E{Key: "price", Value: item.Price},
		}},
	}

	result, err := u.itemcollection.UpdateOne(u.ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}

	return nil
}
func (u *ItemServiceMethod) DeleteItem(item_name *string) error {
	filter := bson.D{primitive.E{Key: "item_name", Value: item_name}}
	result, _ := u.itemcollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
