package models

import (
	"net/http"
	"vp_giver_echo/db"
)

type Subbab struct {
	Subbab_Id    	int `json:"subbab_id"`
	Materi			string `json:"materi"`
	Bab_Id  		int `json:"bab_id"`
}

func FetchAllSubbab() (Response, error) {
	var obj Subbab
	var arrObj []Subbab
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM subbab"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Subbab_Id, &obj.Materi, &obj.Bab_Id)
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

func FetchSubbabBab(intbab int64) (Response, error) {
	var obj Subbab
	var arrObj []Subbab
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM subbab WHERE bab_id = ?"

	rows, err := con.Query(sqlStatement, intbab)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Subbab_Id, &obj.Materi, &obj.Bab_Id)
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

func StoreSubbab(materi string, bab_id int64) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO subbab(materi, bab_id) VALUES (?, ?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(materi, bab_id)

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