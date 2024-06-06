package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/waanvieira/price_dolar_server/models"
)

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

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	http.HandleFunc("/cotacao", inputDolar)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func inputDolar(w http.ResponseWriter, r *http.Request) {
	dolar, err := SearchDolar()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(dolar)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err = InsertDolar(ctx, (*Dolar)(&dolar.USDBRL))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dolar.USDBRL)
}

func SearchDolar() (*models.ApiResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	data, error := io.ReadAll(res.Body)
	if error != nil {
		log.Fatal(error)
	}

	var apiRes models.ApiResponse
	error = json.Unmarshal([]byte(data), &apiRes)
	if error != nil {
		return nil, error
	}

	return &apiRes, nil
}

func InsertDolar(ctx context.Context, dol *Dolar) (string, error) {
	db, err := sql.Open("sqlite3", "./price_db")
	stmt, _ := db.Prepare("CREATE TABLE IF NOT EXISTS prices (id VARCHAR(36), code VARCHAR(255), codein VARCHAR(255), name VARCHAR(255), high VARCHAR(255), low VARCHAR(255), varBid	VARCHAR(255), pctChange	VARCHAR(255), bid VARCHAR(255), ask	VARCHAR(255), timestamp	VARCHAR(255), create_date	timestamp NULL)")

	// ( id INTEGER PRIMARY KEY, name VARCHAR(50), description TEXT, price REAL, amount INT )")
	stmt.Exec()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	select {
	case <-ctx.Done():
		panic(ctx.Err())
	default:
		uuid := uuid.New().String()
		_, err = db.Exec(fmt.Sprintf("INSERT INTO prices VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', date('now'))", uuid, dol.Code, dol.Codein, dol.Name, dol.High, dol.Low, dol.VarBid, dol.PctChange, dol.Bid, dol.Ask, dol.Timestamp))
		if err != nil {
			panic(err)
		}
	}

	return "cadastrado", nil
}
