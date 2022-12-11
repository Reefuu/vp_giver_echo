package models

import (
	"net/http"
	"vp_giver_echo/db"
)

type Bab struct {
	Bab_id    		int `json:"bab_id"`
	Bab_Nama     	string `json:"bab_nama"`
	Buku_Id  		int `json:"buku_id"`
	Image_Cover     string `json:"image_cover"`
	Image_Banner    string `json:"image_banner"`
}

func FetchAllBab() (Response, error) {
	var obj Bab
	var arrObj []Bab
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM bab"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Bab_id, &obj.Bab_Nama, &obj.Buku_Id, &obj.Image_Cover, &obj.Image_Banner)
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

func FetchBabBuku(intbuku int64) (Response, error) {
	var obj Bab
	var arrObj []Bab
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM bab WHERE buku_id = ?"

	rows, err := con.Query(sqlStatement, intbuku)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Bab_id, &obj.Bab_Nama, &obj.Buku_Id, &obj.Image_Cover, &obj.Image_Banner)
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

func StoreBab(bab_nama string, buku_id int64, imageCover string, imageBanner string) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO bab(bab_nama, buku_id, imageCover, imageBanner) VALUES (?, ?, ?, ?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(bab_nama, buku_id, imageCover, imageBanner)

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