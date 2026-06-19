package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/tyler-mcmullin/go-backend/data"
	"github.com/tyler-mcmullin/go-backend/db"
)

func GetItem(c *gin.Context) {
	id, err := bson.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	collection := db.MongoClient.Database("posts").Collection("posts")

	var item data.Item
	err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&item)
	if err == mongo.ErrNoDocuments {
		c.JSON(404, gin.H{"error": "item not found"})
		return
	} else if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, item)
}
