package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	mango "github.com/Chandra5468/go-jwt-learn/database/Mango"
	"github.com/Chandra5468/go-jwt-learn/helpers"
	"github.com/Chandra5468/go-jwt-learn/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var userCollection *mongo.Collection = mango.OpenCollection(mango.MongoCon, "user")
var validate = validator.New()

func Signup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx2, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(&user)

		if validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		countEmail, err := userCollection.CountDocuments(ctx2, bson.M{"email": user.Email})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		countPhone, err := userCollection.CountDocuments(ctx2, bson.M{"phone": user.Phone})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if countEmail > 0 || countPhone > 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "email or phonenumber already exists"})
			return
		}

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		_, err = userCollection.InsertOne(ctx2, &user)
		if err != nil {
			msg := fmt.Sprintf("User item is not created %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		// uid := res.InsertedID.(primitive.ObjectID)
		// THis should be mongo inserted document
		// user.User_id = math.rand()  //uid.Hex()
		token, refreshToken, _ := helpers.GenerateAllTokens(&user.Email, &user.FirstName, &user.LastName, &user.UserType)

		user.Token = token
		user.RefreshToken = refreshToken

		ctx.JSON(http.StatusCreated, user.User_id)

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
