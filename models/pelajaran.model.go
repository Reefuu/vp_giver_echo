package models

import (
	"net/http"
	"vp_giver_echo/db"
)

type Pelajaran struct {
	Pelajaran_ID   int    `json:"pelajaran_id"`
	Nama_Pelajaran string `json:"nama_pelajaran"`
	Kelas          int    `json:"kelas"`
	Image_Cover    string `json:"image_cover"`
	Image_Banner   string `json:"image_banner"`
}

func FetchAllPelajaran() (Response, error) {
	var obj Pelajaran
	var arrObj []Pelajaran
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM pelajaran"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Pelajaran_ID, &obj.Nama_Pelajaran, &obj.Kelas, &obj.Image_Cover, &obj.Image_Banner)
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

func FetchPelajaranKls(intkelas int64) (Response, error) {
	var obj Pelajaran
	var arrObj []Pelajaran
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM pelajaran WHERE kelas = ?"

	rows, err := con.Query(sqlStatement, intkelas)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Pelajaran_ID, &obj.Nama_Pelajaran, &obj.Kelas, &obj.Image_Cover, &obj.Image_Banner)
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

func StorePelajaran(nama_pelajaran string, kelas int64, image string, imageBanner string) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO pelajaran(nama_pelajaran, kelas, image, imageBanner) VALUES (?, ?, ?, ?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama_pelajaran, kelas, image, imageBanner)

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
