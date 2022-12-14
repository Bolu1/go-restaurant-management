package controllers

import(
	"github.com/gin-gonic/gin"
	"time"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"resturant-management/database"
	"resturant-management/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/bluesuncorp/validator.v5"
	"gopkg.in/mgo.v2/bson"

)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFood() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error occured while fetching the food item"})
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = cotext.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		var food models.Food

		if err := c.BindJSON(&food); err !=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		validationErr := validate.Struct(food)
		if validationErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":validatonErr.Error()})
		}

		err := menuCollection.FindOne(ctx, bson.H{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()
		if err != nil{
			msg:=fmt.Sprintf("Menu was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.newObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result, insertErr := foodCollection.InsertOne(ctx, food)
		if insertErr !=nil{
			msg := fmt.Sprintf("Food item not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func GetFoods() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		recordPerPage := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1{
			recordPerPage = 10
		}

		page, err := strcov.Atoi(c.Query("page"))
		if err != nil || page < 1{
			page = 1
		}

		startIndex := (page-1)* recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D({"$match", bson.D{{}}})
		groupStage	:= bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}}, {"total_count", bson.D{{"$sum, 1"}}}, {"data", bson.D{{"$push", "$$ROOT"}}} }}}
		projectStage := bson.D{
			{
				"$project", bson.D{
					{"_id", 0},
					{"total_count", 1},
					{"food_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
			}
		}
	}

	reslt, err := foodCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, groupStage, projectStage
	})
	defer cancel()

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H("error":"error occured while listing food items"))
	}

	var allFoods []bson.M
	if err = result.All(ctx, &allFoods); err != nil{
		log.Fatal(err)
	}
	}
	c.JSON(http.StatusOk, allFoods[0])
}

func round(num float64) int{
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64{
	output := mat.Pow(10, float64(precision))
	return float64(round(num*output)) /output
}

func UpdateFood() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		var food models.food
		
		foodId := c.Params("food_id")

		if err := c.BindJSON(&food); err != nil{
			c.JSON(http.StatusBadRequest, gin.H("error": err.Error()))
			return
		}

		var updateObj primitive.D

		if food.Name != nil{
			updateObj = append(updateObj, bson.E{"name", food.Name})
		}

		if food.Price != nil{
			updateObj = append(updateObj, bson.E{"price", food.Price})
		}

		if food.Food_image != nil{
			updateObj = append(updateObj, bson.E{"food_image", food.Food_image})
		}

		if food.Menu_id := nil{
		err = menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()
		if err != nil{
			msg := fmt.Sprintf("Message: Menu not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}		
		updateObj = append(updateObj, bson.E{"menu", food.Price})
	}

		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"update_at", food.Update_at})

		upsert := true
		filter := bson.M{"food_id": foodID}

		opt := options.UpdateOptions{
			Upsert : &upsert,
		}

		foodCollection.UpdateOne(
			ctx,
			filter,
			bson.D(
				{"$set", updateObj}
			),
			&opt
		)

		if err != nil{
			msg := fmt.Sprintf("food item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOk, result)

	}
}