package controllers

import(
	"github.com/gin-gonic/gin"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")


func GetOrders() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := orderCollectionFind(context.TODO(), bson.M{})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured white listing order items"})
		}
		var allOrders []bson.M
		if err = result.All(ctx, &allOrders); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrders[0])
	}
}

func getOrder() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderId := c.Param("order_id")
		var order models.order

		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error occured while fetching the order item"})
		}
		c.JSON(http.StatusOK, order)
	}
}

func CreateOrder() gin.HandlerFunc{
	return func(c *gin.Context){
		var table models.Table
		var order modesl.Order

		if err := c.BindJSON(&order); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validateErr := validate.Struct(order)

		if validateErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
			return
		}

		if order.Table_id != nil{
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil{
				msg := fmt.Sprintf("message: Table was not found")
				c.JSON(http.StaStatusBadRequest, gin.H{"error":msg})
				return
			}

			order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

			order.ID = primitive.NewObjectID()
			order.Order_id = order.ID.Hex()

			result, insertErr := orderCollection.InsertOne(ctx, order)
			if inserErr != nil{
				msg :=fmt.Sprintf("order item was not created")
				c.JSON(http.StatusInternalServerError, gin.J{"error":msg})
				return
			}
			defer cancel()
			c.JSON(http.StatusOK, result)
		}
	}
}

func UpdateOrder() gin.HandlerFunc{
	return func(c *gin.Context){
		var table models.Table
		var order models.Order

		var updatedObj primitive.D

		orderId := c.Param("order_id")

		if err := c.BindJSON(&food); err != nil{
			c.JSON(http.StatusBadRequest, gin.H("error": err.Error()))
			return
		}
		if order.Table_id != nil{
			err = orderCollection.FindOne(ctx, bson.M{"table_id": food.Table_id}).Decode(&table)
		defer cancel()
		if err != nil{
			msg := fmt.Sprintf("Message: table not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}		
		updateObj = append(updateObj, bson.E{"menu", order.Table_id})
		}

		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"update_at", order.Update_at})

		upsert : true

		filter := bson.M{"order_id":orderId}
		opt := options.UpdateOptions{
			Upsert:&upsert,
		}

		oorderCollection.UpdateOne{
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		}

		if er != nil{
			msg := fmt.Sprintf("order item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return 
		}

		defer cancel()
		c.JSON(http.StatusOK, result)

	}
}

func OrderItemsOrderCreator(ordr models.Order) string{

	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.ID = primitive.NewObjectID()
	order.Order_id = order.ID.Hex()

	orderCollection.InsertOne(ctx, order)
	defer cancel()

	return order.Order_id
}