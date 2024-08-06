package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/platnova/dto"
	"github.com/platnova/pdfclient/pdfgenerator"
)

func main() {
	data, err := os.ReadFile("account_statement.json")
	if err != nil {
		log.Fatal(err)
	}

	statement := &dto.AccountStatement{}
	if err := json.Unmarshal(data, statement); err != nil {
		log.Fatal(err)
	}

	if err := pdfgenerator.New().GenerateAccountStatement(*statement); err != nil {
		log.Panic("Generate Account statement error", err)
	}

	log.Println("Generate statement")
}
