package model

import (
	"fmt"
	"log"
	"github.com/magafagafa/go-app/db"
)

type MAccesUsers struct {
	MAccId          int        `json:"m_acc_id"`
	MUserName       string     `json:"m_user_name"`
	MUserEmail      string     `json:"m_user_email"`
	MUserStatus     string     `json:"m_user_status"`
	MUserRegistDate string     `json:"m_user_regist_date"`
	MUserUpdateDate string     `json:"m_user_update_date"`
	MUserAuthId     string     `json:"m_user_auth_id"`
}

func NewMAccesUsers() *MAccesUsers {
	mAccesUsers := &MAccesUsers{}
	return mAccesUsers
}


/*
* userINsert
*
*/
func (mAccesUsers *MAccesUsers) UserInsert() {
	db, err := db.ConnectDb()
	if err != nil {
		log.Println(err.Error())
	}
	dbClose, err := db.DB()
	if err != nil {
		defer dbClose.Close()
	}

	fmt.Println(*mAccesUsers)

	result := db.Create(&mAccesUsers)
	fmt.Println(result)
}