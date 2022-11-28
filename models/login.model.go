package models

import (
	"database/sql"
	"fmt"
	"vp_week11_echo/db"
	"vp_week11_echo/helpers"
)

type User struct{
	Id			int 	`json:"id"`
	Username 	string 	`json:"username"`
	Password 	string 	`json:"password"`
	Email 		string 	`json:"email"`
}

func CheckLogin(username, password string) (bool, error){

	var obj User
	var pwd string
	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(&obj.Id, &obj.Username, &pwd, &obj.Email)

	if err == sql.ErrNoRows {
		fmt.Print("Username not found!")
		return false, err
	}
	
	if err != nil {
		fmt.Print("Query error!")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)

	if !match {
		fmt.Print("Hash and password doesn't match!")
		return false, err
	}


	return true, nil

}