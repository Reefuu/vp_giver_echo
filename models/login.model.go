package models

import (
	"database/sql"
	"fmt"
	"vp_giver_echo/db"
	"vp_giver_echo/helpers"
)

type User struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	NoTelp   string `json:"nomor_telepon"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Koin     int    `json:"koin"`
	Image    string `json:"image"`
}

func CheckLogin(username, password string) (bool, error) {

	var obj User
	var pwd string
	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE nama = ?"

	err := con.QueryRow(sqlStatement, username).Scan(&obj.Id, &obj.Nama, &obj.Email, &obj.NoTelp, &pwd, &obj.Role, &obj.Koin, &obj.Image)

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
