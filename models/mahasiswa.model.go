package models

import (
	"net/http"
	"vp_week11_echo/db"
)

type Mahasiswa struct {
	Nim      string `json:"nim"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Fakultas string `json:"fakultas"`
	Prodi    string `json:"prodi"`
}

// Read All

func FetchAllMahasiswa() (Response, error) {
	var obj Mahasiswa
	var arrObj []Mahasiswa
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM mahasiswa"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Nim, &obj.Name, &obj.Gender, &obj.Fakultas, &obj.Prodi)
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

// Insert Data

func StoreMahasiswa(nim string, name string, gender string, fakultas string, prodi string) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO mahasiswa(nim, name, gender, fakultas, prodi) VALUES (?, ?, ?, ?, ?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nim, name, gender, fakultas, prodi)

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

// Update Data

func UpdateMahasiswa(nim string, name string, gender string, fakultas string, prodi string) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "UPDATE mahasiswa SET name=?,gender=?,fakultas=?,prodi=? WHERE nim=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, gender, fakultas, prodi, nim)

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

// Delete Data

func DeleteMahasiswa(nim string) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "DELETE FROM mahasiswa WHERE nim=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nim)

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
