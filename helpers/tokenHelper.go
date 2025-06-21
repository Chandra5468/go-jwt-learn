package helpers

import (
	"log"
	"os"
	"time"

	mango "github.com/Chandra5468/go-jwt-learn/database/Mango"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// type SignedDetails struct {
// 	Email     string
// 	FirstName string
// 	LastName  string
// 	UserType  string
// 	jwt.Claims
// }

var userCollection *mongo.Collection = mango.OpenCollection(mango.MongoCon, "user")

var SECRET string = os.Getenv("SECRET")

func GenerateAllTokens(Email, FirstName, LastName, UserType *string) (signedToken, refreshToken string, err error) {
	// cms := &SignedDetails{
	// 	Email:      *Email,
	// 	FirstName:  *FirstName,
	// 	LastName:   *LastName,
	// 	UserType:   *UserType,
	// 	jwt.Claims: jwt.Claims{},
	// }

	claims := jwt.MapClaims{
		"Email":     *Email,
		"FirstName": *FirstName,
		"LastName":  *LastName,
		"UserType":  *UserType,
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(time.Hour * 3).Unix(), // expiration time after 3 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("SECRET")
	signedToken, err = token.SignedString(secret)

	if err != nil {
		log.Printf("Error in generating signed token for loggedin/registered users %s", err.Error())
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		""
	}

}
