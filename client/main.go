package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var db *sql.DB

type Dolar struct {
	Bid string `json:"bid"`
}

func main() {
	http.HandleFunc("/price", GetPriceDolar)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func GetPriceDolar(w http.ResponseWriter, r *http.Request) {
	data, err := GetCotacao()

	fmt.Println(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func GetCotacao() (*Dolar, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
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
		fmt.Println(error)
	}
	fmt.Println(string(data))
	var dol Dolar
	error = json.Unmarshal([]byte(data), &dol)
	if error != nil {
		return nil, error
	}
	jsonBytes, err := json.Marshal(dol.Bid)
	WriteDoc(string(jsonBytes))
	return &dol, nil
}

func WriteDoc(input string) {
	file, err := os.Create("./cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	nb, err := file.WriteString("Dólar: " + input + "\n")
	fmt.Printf("foram escritos %d bytes", nb)
}
