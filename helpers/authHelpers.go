package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserYpeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	err = CheckUserType(c, userType)
	return err

}

func CheckUserType(c *gin.Context, role string) (err error) {

	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorize to accexx this reaource")
		return err

	}
	return err

}
