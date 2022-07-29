package controllers

import(
	"github.com/gin-gonic/gin"
)

type InvoiceViewFormat struct{
	Invoice_id			string
	Payment_method		string
	Order_id			string
	Payment_status		*string
	Payment_due			interface{}
	Table_number		interface{}
	Payment_due_date	time.Time
	Order_details		interface
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoic3")


func GetInvoices() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := invoiceCollection.Find(context.TODO(), bson.M{})
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured"})
		}
		var allInvoices []bson.M
		if err = result.All(ctx, &allInvoices); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allInvoices)
	}
}

func getInvoice() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		invoiceId := c.Param("invoice_id")

		var invoice models.invoice

		err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&invoice)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while lisitng invoice item"})
		}

		var invoiceCollection InvoiceViewFormat

		allOrderItems, err := ItemsbyOrder(invoice.Order_id)
		iew.Order_id = invoice.Order-id
		invoiceView.Payment_due_date = invoice.Payment_due_date

		invoiceView.Payment_method = "null"
		if invoice.Payment_method != nil{
			invoiceView.Payment_method = *invoice.Payment_method
		}

		invoiceView.Invoice_id = invoice.invoice_id
		invoiceView.payment_status = *&invoice.Payment_status
		invoiceView.Payment_due = allOrderItems[0]["payment_due"]
		invoiceView.Table_number = allOrderItems[0]["table_number"]
		invoiceView.Order_details = allOrderItems[0]["order_items"]

		c.JSON(http.StatusOK, invoiceWiew)

	}
}

func CreateInvoices() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var invoice models.Invoice

		if err := c.BindJSON(&order); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var order models.Order

		err := orderCollection.FindOne(ctx, bson.M{"order_id": invoice.Order_id}).Decode(&order)
		defer cancel()
		if err != nil{
			msg := msg.Sprintf("message: Order was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		status := "PENDING"
		if invoice.Payment_status == nil{
			invoice.Payment_status = &status
		}

		invoice.Payment_due_dateParse(time.RFC3339, time.Now().Format(time.RFC3339))time.Parse(time.RFC3339, time.Now().AddDate(0,0,1).Format(time.RFC3339))
		invoice.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.ID = primitive.NewObjectID()
		invoice.Invoice_id = invoice.ID.New()

		validateErr := validate.Struct(invoice)
		if validateErr !=  nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		}

		result, insertErr := invoiceCollection.InsertOne(ctx, invoice)
		if insertErr != nil{
			msg := fmt.Sprintf("Invoice was not updated")
			c.JSON(http.StatusInternalServerError, gin.H{'error': msg})
			return
		}
		defer cancel()

		c.JSON(http.StausOK, result)
	}
}

func UpdateInvoices() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var invoice models.Invoice
		invoiceId := c.Params["invoice_id"]
	}
	if err := c.BindJSON(&invoice); err != nil{
		c.JSON(http.StatusBadRequest, gin.H("error": err.Error()))
		return
	}

	filter := bson.M{"invoice_id": invoiceId}

	var updatedObj primitive.D

	
	if invoice.Payment_method != nil{
		updateObj = append(updatobj, bson.E{"payment_method", invoice.Payment_method})
	}

	if invoice.Payment_stauts != nil{
		updateObj = append(updatobj, bson.E{"payment_status", invoice.Payment_status})
	}

	invoice.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updatedObj = append(updateObj, bson.E{"update_at", invoice.Update_at})

	upsert := true
	opt := options.UpdatedOptions{
		Upsert: &upsert,
	}

	status := 	"PENDING"
	if invoice.Payment_status == nil{
		invoice.Payent_status = &status
	}

	result, err := invoiceCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateObj}
		},
		&opt,
	)	

	if err != nil{
		msg := fmt.Sprintf("invoice item update failed")
		c.JSON(http.StatusInternalServerError, gin.H("error": msg))
		return
	}

}

