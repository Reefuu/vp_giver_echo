package models

import (
	"net/http"
	"vp_giver_echo/db"
)

type Buku struct {
	Buku_Id    		int `json:"buku_id"`
	Nama_Buku		string `json:"nama_buku"`
	Pelajaran_Id  	int `json:"pelajaran_id"`
	Image_Cover     string `json:"image_cover"`
	Image_Banner    string `json:"image_banner"`
}

func FetchAllBuku() (Response, error) {
	var obj Buku
	var arrObj []Buku
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM buku"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Buku_Id, &obj.Nama_Buku, &obj.Pelajaran_Id, &obj.Image_Cover, &obj.Image_Banner)
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

func FetchBukuPljrn(intpelajaran int64) (Response, error) {
	var obj Buku
	var arrObj []Buku
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM buku WHERE pelajaran_id = ?"

	rows, err := con.Query(sqlStatement, intpelajaran)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Buku_Id, &obj.Nama_Buku, &obj.Pelajaran_Id, &obj.Image_Cover, &obj.Image_Banner)
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

func StoreBuku(nama_buku string, pelajaran_id int64, imageCover string, imageBanner string) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO buku(nama_buku, pelajaran_id, imageCover, imageBanner) VALUES (?, ?, ?, ?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama_buku, pelajaran_id, imageCover, imageBanner)

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