package pdfclient

import "github.com/platnova/dto"

//go:generate mockgen -source=pdf_client.go -destination=../../mocks/pdf_client_mock.go -package=mocks
type PDFGeneratorClient interface {
	GenerateAccountStatement(statement dto.AccountStatement) error
}
