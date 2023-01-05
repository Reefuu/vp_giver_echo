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

func AddCoin(nama string, intKoin int64) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "UPDATE users SET koin=? WHERE nama=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(intKoin, nama)

	if err != nil {
		return res, err
	}

	rowAffectedID, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"row_affected_id": rowAffectedID,
	}

	return res, nil
}
