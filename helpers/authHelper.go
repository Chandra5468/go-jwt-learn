package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// func CheckUserType(ctx *gin.Context, role string) (err error) {
// 	userType := ctx.GetString("user_type")
// 	err = nil
// 	if userType != role{
// 		err = errors.New("Unauthorized ")
// 	}
// }

func MatchUserTypetoUID(ctx *gin.Context, userId string) (err error) {
	userType := ctx.GetString("user_type")
	uid := ctx.GetString("uid")
	err = nil
	if userType == "USER" && uid != userId {
		err = errors.New("unauthorized to access the resources")
	}
	// err = CheckUserType(ctx, userType)
	return err
}
