package models

import (
	"fmt"

	"github.com/waanvieira/price_dolar_server/db/connection"
)

type ApiResponse struct {
	USDBRL Dolar
}
type Dolar struct {
	Id         string `json:"id"`
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

// func (db *DB) Exec(query string, args ...any) (Result, error)

// func (d *Dolar) GetPrices() (sql.Result, error) {
// 	var db *sql.DB
// 	rows, err := db.Query("SELECT * FROM prices")
// 	if err != nil {
// 		log.Fatalf("Error: Unable to execute query: %v", err)
// 	}
// 	defer rows.Close()

// }

// albumsByArtist queries for albums that have the specified artist name.
func GetPrices(name string) ([]Dolar, error) {
	// An albums slice to hold data from returned rows.
	var dolars []Dolar
	fmt.Println(connection.DB)
	rows, err := connection.DB.Query("SELECT * FROM prices")

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var dlr Dolar
		if err := rows.Scan(&dlr.Id, &dlr.Code, &dlr.Codein, &dlr.PctChange); err != nil {
			return nil, fmt.Errorf("dlrumsByArtist %q: %v", name, err)
		}
		dolars = append(dolars, dlr)
	}
	fmt.Println(dolars)
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return dolars, nil
}

// func (d Dolar) {
// 	select * from person where id = ?;
// }
