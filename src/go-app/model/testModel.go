package model

import (
	"fmt"
	"log"

	"github.com/magafagafa/go-app/db"
)

type Test1 struct {
	Common    Base
	MPassword string `json:"m_password"`
	MUserName string `json:"m_user_name"`
}

func (t *Test1) TestInsert() {
	db, err := db.ConnectDb()
	if err != nil {
		log.Println(err.Error())
	}
	dbClose, err := db.DB()
	if err != nil {
		defer dbClose.Close()
	}

	fmt.Println(*t)

	result := db.Create(&t)
	fmt.Println(result)
}
