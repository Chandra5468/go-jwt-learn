package controllers

import (
	"context"
	"net/http"
	"time"

	mango "github.com/Chandra5468/go-jwt-learn/database/Mango"
	"github.com/Chandra5468/go-jwt-learn/helpers"
	"github.com/Chandra5468/go-jwt-learn/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var userCollection *mongo.Collection = mango.OpenCollection(mango.MongoCon, "user")

func Signup() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func Login() {

}

func HashPassword() {
	// bcrypt.
}

func VerifyPassword() {

}

func GetUsers() {

}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("userId")

		if err := helpers.MatchUserTypetoUID(ctx, userId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx2, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User

		res := userCollection.FindOne(ctx2, bson.M{"user_id": userId})

		err := res.Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to decode/deserialzie req.body to user"})
			return
		}

		ctx.JSON(http.StatusOK, &user)
	}
}
