package usercontroller

import (
	"log"
	"time"
	"strconv"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/magafagafa/go-app/model"
)

var layout = "2006-01-02 03:04:05"

func UsuerRegist(c echo.Context) error {
	now := time.Now()
	// userテーブルを初期化
	mAccesUsers := model.NewMAccesUsers()
	mAccesUsers.MUserEmail = c.QueryParam("email")
	mAccesUsers.MUserAuthId = c.QueryParam("auth0Id")
	mAccesUsers.MUserRegistDate = timeToString(now)
	mAccesUsers.MUserUpdateDate = timeToString(now)
	mAccesUsers.UserInsert()
	return c.String(http.StatusOK, "team:")
}

func queryParamInt(c echo.Context, name string) int {
	param := c.QueryParam(name)
	result, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err.Error())
	}
	return result
}

func timeToString(t time.Time) string {
    str := t.Format(layout)
    return str
}

func stringToTime(str string) time.Time {
    t, _ := time.Parse(layout, str)
    return t
}