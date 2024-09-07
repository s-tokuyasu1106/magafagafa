package route

import (
	"fmt"
	"net/http"

	"github.com/magafagafa/go-app/controller"
	"github.com/magafagafa/go-app/controller/user"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		fmt.Println(err)                                    // 標準出力へ
		c.JSON(http.StatusInternalServerError, err.Error()) // ブラウザ画面へ
	}

	e.GET("/test", controller.Test)
	e.GET("/", controller.Index)
	e.GET("/userRegist", usercontroller.UsuerRegist)
	// e.POST("/regist", controller.RegistUser)
	// e.POST("/registReult", controller.RegistReult)
	e.GET("/user/:id", controller.GetUser)
	e.GET("/show", controller.Show)
	e.POST("/save", controller.Save)

	return e
}
