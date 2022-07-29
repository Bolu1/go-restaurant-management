package controllers

import(
	"github.com/gin-gonic/gin"
)

type OrderItemPack struct{
	Table_id *string
	Order_items []models.OrderItemPack
}

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

func GetOrderItems() gin.handlerFunc{
	return func (c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err = orderItemCollection.Find(context.TODO(), bson.M{})

		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error while listing order items"})
			return
		}

		var allOrderItems []bson.Message
		if err = result.All(ctx, &allOrderItems); err != nil{
			log.Fatal(err)
			result
		}
		c.JSON(http.StatusOK, allOrderItems)

	}
}

func GetOrderItemsByOrder() gin.handlerFunc{
	return func (c *gin.Context){
		orderId := c.Params("order_id")

		allOrderItems, err := ItemsByOrder(orderId)

		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while listing order items by order ID"})
			return
		}
		c.JSON(http.StatusOK, allOrderItems)
	}
}

func ItemsByOrder(id string) (GetOrderItems []primitive.M, err error){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	
	matchStage := bson.D{{"$match", bson.D{{"order_id", id}}}}
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "food"}, {"localField", "food_id"}, {"foreignField", "food_id"}, {"as", "food"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$food"}, {"preserveNullAndEmptyArrays", true} }}}

	lookupOrderStage := bson.D{{"$lookup", bson.D{{"from", "order"}, {"localField", "order_id"}, {"foreignField", "order_id"}, {"as", "order"}}}
	unwindOrderStage := bson.D{{"$unwind", bson.D{{"path", "$order"}, {"preserveNullAndEmptyArrays", true}}}}

	lookupTableStage := bson.D{{"$lookup", bson.D{"from", "table"}, {"localField", "order.table_id"}, {"foreignField", "table_id"}, {"as", "table"}}}
	unWindTableStage := bson.D{{"$unwind", bson.D{{"path", "$table"}, {"preserveNullAndEmptyArrays", true}}}}

	projectStage := bson.D{
		{
			"$project", bson.D{
				{"id", 0},
				{"amount", "$food.price"},
				{"total_count", 1},
				{"food_name", "$food.name"},
				{"food_image", "$food.food_image"},
				{"table_number", "$table,table_number"},
				{"table_id", "$table.table_id"},
				{"order_id", "4order.order_id"},
				{"price", "$food.price"},
				{"quantity", 1},
			}
		}
	}

	groupStage :=  bson.D{{"$group", bson.D{{"_id", bson.D{{"order_id", "$order_id", {"table_id", "$table_id"}, {"table_number", "$table_number"}}}, {"payment_due", bson.D{{"$sum", "$amount"}}}, {"total_count", bson.D{{"$sum", 1}}}, {"order_items", bson}}}}}

	projectStage2 := bson.D{
		{"$project", vson.D{
			{"id", 0},
			{"payment_due", 1},
			{"total_count", 1},
			{"table_number", "$_id.table_number"},
			{"order_items", 1},
		}}

		result, err := orderItemCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage,
			lookupStage,
			unwindStage,
			lookupOrderStage,
			unwindOrderStage,
			lookupTableStage,
			unWindTableStage
			projectStage,
			groupStage,
			projectStage2
		})
	}

	if err != nil{
		panic(err)
	}

	result.All(ctx, &Orderitems) err != nil{
		panic(err)
	}

	defer cancel()

	return OrderItems, err
	
}

func GetOrderItem() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		orderItemId := c.Param("order_item_id")
		var orderItem models.OrderItem

		err := orderItemCollection.FindOne(ctx, bson.M{"orderitem_id": orderItemId}).Decode(&orderItem)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured at finding order item"})
			return
		}
		c.JSON(http.StatusOK, orderItem)
	}
}

func CreateOrderitem() gin.handlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var orderItemPack OrderItemPack
		var order models.Order

		if err := c.BindJSON(&orderItemPack); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		order.Order_Date, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		orderItemsToBeInserted := []interface{}{}
		order.Table_id = orderItemPack.Table_id
		order_id := OrderItemOrderCreator(order)

		range oderItem := range orderitemPack.Order_items{
			orderitem.Order_id = order_id

			validationErr := validate.Struct(orderItem)

			if validationErr != nil{
				c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
				return
			}
			orderItem.ID = primitive.NewObjectID()
			orderItem.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			orderItem.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			orderItem.Order_item_id = orderitem.ID.Hex()
			var num = toFixed(*orderItem.Unit_price, 2)
			orderItem.Unit_price = &num
			orderItemsToBeInserted = append(orderItemsToBeInserted, orderItem)

		}
		insertedOrderItems, err := orderItemCollection.InsertMany(ctx, orderItemsToBeInserted)

		if err != nil{
			log.Fatal(err)
		}
		defer cancel()

		c.JSON(http.StatusOK, orderItem)

	}
}

func UpdateOrderitems() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var orderItem models.OrderItemOrderCreator

		orderItemId := c.Param("order_item_id")

		filter := bson.M("order_item_id": orderItemId)

		var updateObj primitive.D

		if orderItem.Unit_price != {
			updateObj = append(updateObj, bson.E{"unit_price", *&orderItem.Unit_price})
		}

		if orderItem.Quantity != nil{
			updateObj = append(updateObj, bson.E{"quanitiy", *orderItem.Quantity})
		}

		if orderItem.Food_id != nil{
			updaObj = append(updateObj, bson.E{"food_id", *orderItem.Food_id})
		}

		orderItem.Updated_at, _ = time= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = aappend(updateObj, bson.E{"updated_at", orderItem.Updated_at})

		upsert := true

		opt := options.UpdateOptions{
			upsert: &upsert,
		}

		orderItemCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)

		if err != nil{
			msg := "order items update failed"
			c.JSON(http.StatusInternalServerError, bson.H{"error" msg})
			return
		}

		defer cancel()

		c.JSON(http.StatusOK, result)

	}
}