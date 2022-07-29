package controllers

import(
	"github.com/gin-gonic/gin"
)

var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")


func GetTables() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err = tableCollection.Find(context.TODO(), bson.M{})

		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error while listing order items"})
			return
		}

		var allTables []bson.Message
		if err = result.All(ctx, &alltables); err != nil{
			log.Fatal(err)
			result
		}
		c.JSON(http.StatusOK, allTables)

	}
}

func getTable() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		tableId := c.Param("table_id")
		var table models.table

		err := tableCollection.FindOne(ctx, bson.M{"table_id": tableId}).Decode(&table)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured at finding order item"})
			return
		}
		c.JSON(http.StatusOK, table)
	}
}

func CreateTable() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		var table models.Table

		if err := c.BindJSON(&table); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
			validationErr = validate.Struct(table)

			if validationErr != nil{
			if err := c.BindJSON(&table); err != nil{
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			table.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			table.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

			table.ID = primitive.NewObjectID()
			table.Order = table.ID.Hex()

			result, insertErr := tableCollection.InsertOne(ctx, table)

			if insertErr != nil{
				msg : fmt.Sprintf("Table item was not created")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}

			defer cancel()

			c.JSON(http.StatusOK, result)

	}
}

func UpdateTable() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		var table models.Table

		tableId := c.Param("table_id")

		if err := c.BindJSON(&table); err != nil{
			c.JSON(http.StatusBadRequest, gin.H("error": err.Error()))
			return
		}

		var updateObj primitive.D

		if table.Number_of_guests != nil{
			updateObj = append(updateObj, bson.E{"number_of_guest", table.Number_of_guest})
		}

		if table.Table_number != nil{
			updateObj = append(updateObj, bson.E{"table_number", table.Table_number})
		}

		table.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		filter, err := bson.M{"table_id": tableId}

		tableCol.Updateone(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)

		if err != nil{
			msg := fmt.Sprintf("table item update failed")
			c.JSN(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()

		c.JSON(http.StausOk, result)

	}
}

