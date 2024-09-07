package controller

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/magafagafa/go-app/model"
)

func Test(c echo.Context) error {
	// name := c.QueryParam("name")
	// password := c.QueryParam("password")
	// test := model.Newtest1(name, password)
	test := new(model.Test1)
	test.MUserName = c.QueryParam("name")
	test.MPassword = c.QueryParam("password")
	test.TestInsert()
	// if err != nil {
	// 	test.Common.Status = true
	// }

	return c.JSON(http.StatusOK, test)
}

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Helloddss World")
}

func RegistUser(c echo.Context) error {
	return c.Render(http.StatusOK, "regist", "")
}

// func RegistReult(c echo.Context) error {
// 	userId := c.FormValue("regis_userid")
// 	email := c.FormValue("regis_email")
// 	name := c.FormValue("regis_name")
// 	age := c.FormValue("regis_age")
// 	birth := c.FormValue("regis_birth")
// 	gender := c.FormValue("regist_gender")
// 	stay := c.FormValue("regis_stay")
// 	password := c.FormValue("regist_password")
// 	model.NewMAccesUser(userId)
// 	return c.Render(http.StatusOK, "regist-result", name)
// }

func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func Show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team"+team+" + member"+member)
}

func Save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
