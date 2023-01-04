package models

import (
	"net/http"
	"vp_giver_echo/db"
	"vp_giver_echo/helpers"
)

func RegisterNewAcc(nama string, email string, password string) (Response, error) {

	var res Response
	var image string

	image = "kosong"

	con := db.CreateCon()
	sqlStatement := "INSERT INTO users(nama, email, password, image) VALUES (?, ?, ?, ?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	hash, _ := helpers.HashPassword(password)

	result, err := stmt.Exec(nama, email, hash, image)

	if err != nil {
		return res, err
	}

	lastInsertedID, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedID,
	}

	return res, nil
}
