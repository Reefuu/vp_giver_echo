package models

import (
	"net/http"
	"vp_giver_echo/db"
)

type Users struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	NoTelp   string `json:"nomor_telepon"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Koin     int    `json:"koin"`
	Image    string `json:"image"`
}

func FetchAllUsers() (Response, error) {
	var obj Users
	var arrObj []Users
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Email, &obj.NoTelp, &obj.Password, &obj.Role, &obj.Koin, &obj.Image)
		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil

}
