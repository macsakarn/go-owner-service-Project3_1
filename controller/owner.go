package controller

import (
	"accounts/config"
	"accounts/model"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var accountsCollection *mongo.Collection = config.GetMongoDbCollection(config.Client, "Owner")

func GetInfo() func(c *gin.Context) {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second) //time Out
		defer cancel()
		var result = &model.QueryOwner{}
		err := accountsCollection.FindOne(ctx, bson.M{}).Decode(result)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, result)

	}
}

func UpdateInfo() func(c *gin.Context) {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second) //time Out
		defer cancel()

		var data model.UpdateOwner
		err := c.ShouldBindJSON(&data)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		editInfo := bson.M{"$set": data}
		result, err := accountsCollection.UpdateOne(
			ctx,
			bson.M{},
			editInfo,
		)

		_ = result

		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "msg"})
			return
		}

		c.JSON(http.StatusOK, true)
	}
}

func GetAccounts() func(c *gin.Context) {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second) //time Out
		defer cancel()

		var result = &model.QueryAccount{}
		err := accountsCollection.FindOne(ctx, bson.M{}).Decode(result)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, result)

	}
}

func InsertAccounts() func(c *gin.Context) {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second) //time Out
		defer cancel()

		var data model.InsertAccount
		err := c.ShouldBindJSON(&data)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		account := bson.M{"$push": bson.M{"account": data}}
		result, err := accountsCollection.UpdateOne(
			ctx,
			bson.M{},
			account,
		)
		_ = result

		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "msg"})
			return
		}

		c.JSON(http.StatusOK, true)
	}
}

func UpdateAccount() func(c *gin.Context) {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second) //time Out
		defer cancel()

		var data model.InsertAccount
		err := c.ShouldBindJSON(&data)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		queryAccounts := bson.M{"account.id": data.ID}
		editInfo := bson.M{"$set": bson.M{"account.$.name": data.Name, "account.$.number": data.Number}}

		result, err := accountsCollection.UpdateOne(ctx, queryAccounts, editInfo)

		_ = result

		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "msg"})
			return
		}

		c.JSON(http.StatusOK, true)
	}
}

func DeleteAccounts() func(c *gin.Context) {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second) //time Out
		defer cancel()

		var data model.DeleteAccounts
		err := c.ShouldBindJSON(&data)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		account := bson.M{"$pull": bson.M{"account": data}}
		result, err := accountsCollection.UpdateOne(
			ctx,
			bson.M{},
			account,
		)

		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "msg"})
			return
		}

		if result.ModifiedCount == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "non update", "staus": false})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Update done", "staus": true})
	}
}
