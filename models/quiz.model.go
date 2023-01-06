package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"vp_giver_echo/db"
)

type Quiz struct {
	Quiz_ID    int    `json:"quiz_id"`
	Pertanyaan string `json:"pertanyaan"`
	Jawaban    string `json:"jawaban"`
	Subbab_ID  int    `json:"subbab_id"`
}

func FetchAllQuiz() (Response, error) {
	var obj Quiz
	var arrObj []Quiz
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM quiz"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Quiz_ID, &obj.Pertanyaan, &obj.Jawaban, &obj.Subbab_ID)
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

func FetchQuizSubbab(intsubbab int64) (Response, error) {
	var obj Quiz
	var arrObj []Quiz
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM quiz WHERE subbab_id = ?"

	rows, err := con.Query(sqlStatement, intsubbab)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Quiz_ID, &obj.Pertanyaan, &obj.Jawaban, &obj.Subbab_ID)
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

func StoreQuiz(pertanyaan string, jawaban string, subbab_id int64) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO quiz(pertanyaan, jawaban, subbab_id) VALUES (?, ?, ?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(pertanyaan, jawaban, subbab_id)

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

func CheckQuiz(pertanyaan string, jawaban string) (bool, error) {

	var obj Quiz
	var jwbn string
	con := db.CreateCon()

	sqlStatement := "SELECT * FROM quiz WHERE pertanyaan = ?"

	err := con.QueryRow(sqlStatement, pertanyaan).Scan(&obj.Quiz_ID, &obj.Pertanyaan, &jwbn, &obj.Subbab_ID)

	if err == sql.ErrNoRows {
		fmt.Print("Username not found!")
		return false, err
	}

	if err != nil {
		fmt.Print("Query error!")
		return false, err
	}

	match := jawaban == jwbn

	if !match {
		fmt.Print("Jawaban Salah!")
		return false, err
	}

	return true, nil

}
