package preorderControllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/tyler-mcmullin/go-backend/data"
	"github.com/tyler-mcmullin/go-backend/db"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetLatest(c *gin.Context) {
	opts := options.FindOne().SetSort(bson.D{{Key: "_id", Value: -1}})

	collection := db.MongoClient.Database("posts").Collection("posts")

	var item data.Item
	err := collection.FindOne(context.TODO(), bson.D{}, opts).Decode(&item)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func PreorderTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
