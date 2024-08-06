package pdfgenerator

import (
	"testing"

	"github.com/platnova/dto"
	"github.com/stretchr/testify/assert"
)

func TestPDFClient_GenerateAccountStatement(t *testing.T) {
	const (
		success = iota
		successWithProtection
		errorOccurred
	)

	tests := []struct {
		name     string
		testType int
	}{
		{
			name:     "Test success",
			testType: success,
		},
	}

	mockAccountStatement := dto.AccountStatement{
		FullName:   "SANDRA SAULGRIEZE",
		Address:    "14 The Dale Whitefield hall Bettystorm Meath A9N27C",
		FooterNote: "Revolut Ltd is registered in England and Wales (No. 08804411), 7 Westferry Circus, Canary Wharf, London, England, E14 4HD and is authorised by the Financial Conduct Authority under the Electronic Money Regulations 2011 (Firm Reference 900562). Revolut Travel Ltd is authorised by the Financial Conduct Authority to undertake insurance distribution activities (FCA No: 780586). Our insurance products are arranged by Revolut Travel Ltd and Revolut Ltd, which is an appointed representative of Revolut Travel Ltd. Revolut’s stock trading products are provided by Revolut Trading Ltd (No. 832790), an appointed representative of Resolution Compliance Ltd, which is authorised and regulated by the Financial Conduct Authority.",
		IBANInfo: []dto.IbanInfoData{
			{
				IBan:    "IE30REV099036022547749",
				BIC:     "REV0IE23",
				Message: "",
			},

			{
				IBan:    "LT093250041069208595",
				BIC:     "REVOLT21",
				Message: "(You cannot use this IBAN for bank transfers. Please use the IBAN found in the app)",
			},

			{
				IBan:    "LT087070024346246713",
				BIC:     "RETBLT21",
				Message: "(You cannot use this IBAN for bank transfers. Please use the IBAN found in the app)",
			},
		},
		Product: dto.ProductData{
			Name:           "Account (Current Account)",
			Currency:       "£",
			OpeningBalance: "2.52",
			MoneyOut:       "1,944.09",
			MoneyIn:        "1,978.00",
			ClosingBalance: "36.43",
		},
		StartDate: "1 February 2023",
		EndDate:   "29 March 2023",
		Date:      "20 May 2023",
		Transactions: []dto.Transaction{
			{
				Date:        "3 Feb 2023",
				Description: "Apple Pay Top-Up by *5453",
				MoneyOut:    "",
				MoneyIn:     "50.00",
				Balance:     "52.52",
				Currency:    "£",
			},

			{
				Date:        "3 Feb 2023",
				Description: "Apple Pay Top-Up by *5453",
				MoneyOut:    "",
				MoneyIn:     "100.00",
				Balance:     "152.52",
				Currency:    "£",
			},

			{
				Date:        "3 Feb 2023",
				Description: "To LINA MILLER SAULGREZE",
				MoneyOut:    "100",
				Balance:     "52.52",
				Currency:    "£",
			},

			{
				Date:        "7 Feb 2023",
				Description: "To LINA MILLER SAULGREZE",
				MoneyOut:    "10",
				Balance:     "42.52",
				Currency:    "£",
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			pdfclient := New()

			switch testCase.testType {
			case success:
				err := pdfclient.GenerateAccountStatement(mockAccountStatement)
				assert.NoError(t, err)

			}
		})
	}
}
