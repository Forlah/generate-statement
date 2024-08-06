package dto

type IbanInfoData struct {
	IBan    string `json:"iban"`
	BIC     string `json:"bic"`
	Message string `json:"message"`
}

type ProductData struct {
	Name           string `json:"name"`
	Currency       string `json:"currency"`
	OpeningBalance string `json:"openingBalance"`
	MoneyOut       string `json:"moneyOut"`
	MoneyIn        string `json:"moneyIn"`
	ClosingBalance string `json:"closingBalance"`
}

type Transaction struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	MoneyOut    string `json:"moneyOut"`
	MoneyIn     string `json:"moneyIn"`
	Balance     string `json:"balance"`
	Currency    string `json:"currency"`
}

type AccountStatement struct {
	FullName     string         `json:"fullName"`
	Address      string         `json:"address"`
	FooterNote   string         `json:"footerNote"`
	IBANInfo     []IbanInfoData `json:"ibanInfo"`
	Product      ProductData    `json:"product"`
	StartDate    string         `json:"startDate"`
	EndDate      string         `json:"endDate"`
	Date         string         `json:"date"`
	Transactions []Transaction  `json:"transactions"`
}
