package pdfgenerator

import (
	"fmt"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	assets "github.com/platnova/assests"
	"github.com/platnova/dto"
	"github.com/platnova/pdfclient"
)

var _ pdfclient.PDFGeneratorClient = (*PdfGeneratorClient)(nil)

const (
	fontPath = "static/fonts"
)

type PdfGeneratorClient struct {
	pdf      pdf.Maroto
	fontPath string
}

func New() *PdfGeneratorClient {
	return &PdfGeneratorClient{fontPath: fontPath}
}

func (client *PdfGeneratorClient) makeFooter(footerContent string) pdf.Maroto {
	client.pdf.RegisterFooter(func() {
		client.pdf.Row(10, func() {
			client.pdf.Col(12, func() {
				client.pdf.Text(footerContent, props.Text{
					Top:   20,
					Size:  5,
					Align: consts.Left,
					Color: color.Color{
						Red:   88,
						Green: 87,
						Blue:  87,
					},
				})
			})
		})
	})
	return client.pdf
}

func (client *PdfGeneratorClient) makeHeader() pdf.Maroto {
	client.pdf.RegisterHeader(func() {
		client.pdf.Row(15, func() {
			_ = client.pdf.Base64Image(assets.Base64Logo, consts.Png, props.Rect{
				Percent: 100,
			})

		})
		client.pdf.Row(8, func() {
			// Do nothing .
		})
	})

	return client.pdf
}

func (client *PdfGeneratorClient) initStatementPage(statement dto.AccountStatement) {
	pdfClient := pdf.NewMaroto(consts.Portrait, consts.A4)

	client.pdf = pdfClient
	client.pdf.SetPageMargins(20, 15, 20)

	// client.pdf.SetFontLocation(client.fontPath)
	// client.pdf.AddUTF8Font("CustomArial", consts.Normal, "arial-unicode-ms.ttf")
	// client.pdf.AddUTF8Font("CustomArial", consts.Bold, "arial-unicode-ms.ttf")
	// client.pdf.AddUTF8Font("CustomArial", consts.BoldItalic, "arial-unicode-ms.ttf")
	// client.pdf.AddUTF8Font("CustomArial", consts.Italic, "arial-unicode-ms.ttf")

	client.makeHeader()
	client.makeFooter(statement.FooterNote)
}

func (client *PdfGeneratorClient) statementTitle(statement dto.AccountStatement) pdf.Maroto {
	client.pdf.Row(4, func() {
		client.pdf.Col(12, func() {
			client.pdf.Text("EUR Statement", props.Text{
				Size:        13,
				Align:       consts.Right,
				Extrapolate: false,
				Style:       consts.Bold,
			})

		})
	})

	client.pdf.Row(8, func() {
		client.pdf.Col(12, func() {
			client.pdf.Text(fmt.Sprintf("Generated on the %s", statement.Date), props.Text{
				Style:       consts.Normal,
				Size:        6,
				Extrapolate: false,
				Align:       consts.Right,
				Top:         2,
				Color: color.Color{
					Red:   128,
					Green: 128,
					Blue:  128,
				},
			})

			client.pdf.Text("Revolut Bank UAB", props.Text{
				Style:       consts.Normal,
				Size:        6,
				Extrapolate: false,
				Align:       consts.Right,
				Top:         5,
				Color: color.Color{
					Red:   128,
					Green: 128,
					Blue:  128,
				},
			})

		})
	})

	return client.pdf
}

func (client *PdfGeneratorClient) setCustomerData(statement dto.AccountStatement) pdf.Maroto {
	client.pdf.Row(20, func() {
		client.pdf.Col(3, func() {
			client.pdf.Text(statement.FullName, props.Text{
				Style:       consts.Bold,
				Size:        8,
				Extrapolate: false,
				Top:         5,
			})

			client.pdf.Text(statement.Address, props.Text{
				Size:            5.5,
				Extrapolate:     false,
				Top:             10,
				Style:           consts.Bold,
				VerticalPadding: 1,
			})

		})

	})

	return client.pdf
}

func (client *PdfGeneratorClient) setIBans(statement dto.AccountStatement) pdf.Maroto {
	for i := 0; i < len(statement.IBANInfo); i++ {
		client.pdf.Row(4, func() {
			client.pdf.ColSpace(6)
			client.pdf.Col(25, func() {
				client.pdf.Text("IBAN", props.Text{
					Style:       consts.Bold,
					Size:        6,
					Extrapolate: false,
				})

				client.pdf.ColSpace(1)

				client.pdf.Text(statement.IBANInfo[i].IBan, props.Text{
					Size: 6,
				})

			})

		})

		client.pdf.Row(7, func() {
			client.pdf.ColSpace(6)
			client.pdf.Col(10, func() {
				client.pdf.Text("BIC", props.Text{
					Style:       consts.Bold,
					Size:        5,
					Extrapolate: false,
				})

				client.pdf.ColSpace(1)

				client.pdf.Text(statement.IBANInfo[i].BIC, props.Text{
					Size: 5,
				})

				if statement.IBANInfo[i].Message != "" {
					client.pdf.Text(statement.IBANInfo[i].Message, props.Text{
						Size:        5,
						Top:         3.5,
						Extrapolate: true,
					})
				}
			})

		})
	}

	return client.pdf
}

func (client *PdfGeneratorClient) setBalanceSummary(statement dto.AccountStatement) pdf.Maroto {
	client.pdf.Row(20, func() {
		client.pdf.Col(2, func() {
			client.pdf.Text("Balance summary", props.Text{
				Style:       consts.Bold,
				Size:        9,
				Top:         12,
				Extrapolate: false,
			})

		})
	})

	tableHeaders := []string{"Product", "Opening balance", "Money out", "Money in", "Closing balance"}

	td := make([]dto.ProductData, 0)
	td = append(td, statement.Product)
	td = append(td, dto.ProductData{
		Name:           "Total",
		Currency:       "£",
		OpeningBalance: "2.52",
		MoneyOut:       "1,944.09",
		MoneyIn:        "1,978.00",
		ClosingBalance: "36.43",
	})

	tableContents := make([][]string, len(td))
	for i := range tableContents {
		tableContents[i] = make([]string, 5)
	}

	for row := 0; row < len(td); row++ {
		tableContents[row][0] = td[row].Name
		tableContents[row][1] = fmt.Sprintf("%s%s", td[row].Currency, td[row].OpeningBalance)
		tableContents[row][2] = fmt.Sprintf("%s%s", td[row].Currency, td[row].MoneyOut)
		tableContents[row][3] = fmt.Sprintf("%s%s", td[row].Currency, td[row].MoneyIn)
		tableContents[row][4] = fmt.Sprintf("%s%s", td[row].Currency, td[row].ClosingBalance)
	}

	client.pdf.Row(3, func() {
		// Do nothing.
	})
	client.pdf.TableList(tableHeaders, tableContents, props.TableList{
		HeaderProp: props.TableListContent{
			Style:     consts.Bold,
			Size:      6.5,
			Family:    consts.Arial,
			GridSizes: []uint{3, 2, 2, 2, 2},
		},
		ContentProp: props.TableListContent{
			Family:    consts.Arial,
			Size:      6.0,
			Style:     consts.Bold,
			GridSizes: []uint{3, 2, 2, 2, 2},
		},
		Line: true,
		LineProp: props.Line{
			Style: consts.Solid,
			Width: 0.4,
		},
		Align:                  consts.Left,
		VerticalContentPadding: 4,
	})

	return client.pdf
}

func (client *PdfGeneratorClient) setTransactions(statement dto.AccountStatement) pdf.Maroto {
	title := fmt.Sprintf("Account transactions from %s to %s", statement.StartDate, statement.EndDate)

	client.pdf.Row(10, func() {
		// Do nothing.
	})
	client.pdf.Row(5, func() {
		client.pdf.Col(10, func() {
			client.pdf.Text(title, props.Text{
				Style:       consts.Bold,
				Size:        8,
				Extrapolate: false,
			})

		})
	})

	tableHeaders := []string{"Date", "Description", "Money out", "Money in", "Balance"}

	tableContents := make([][]string, len(statement.Transactions))
	for i := range tableContents {
		tableContents[i] = make([]string, 5)
	}

	for row := 0; row < len(statement.Transactions); row++ {
		tableContents[row][0] = statement.Transactions[row].Date
		tableContents[row][1] = statement.Transactions[row].Description
		if statement.Transactions[row].MoneyOut != "" {
			tableContents[row][2] = fmt.Sprintf("%s%s", statement.Transactions[row].Currency, statement.Transactions[row].MoneyOut)
		}
		if statement.Transactions[row].MoneyIn != "" {
			tableContents[row][3] = fmt.Sprintf("%s%s", statement.Transactions[row].Currency, statement.Transactions[row].MoneyIn)
		}

		tableContents[row][4] = fmt.Sprintf("%s%s", statement.Transactions[row].Currency, statement.Transactions[row].Balance)
	}

	client.pdf.Row(3, func() {
		// Do nothing.
	})
	client.pdf.TableList(tableHeaders, tableContents, props.TableList{
		HeaderProp: props.TableListContent{
			Style:     consts.Bold,
			Size:      6.5,
			Family:    consts.Arial,
			GridSizes: []uint{2, 3, 2, 2, 2},
		},
		ContentProp: props.TableListContent{
			Family:    consts.Arial,
			Size:      6.0,
			Style:     consts.Bold,
			GridSizes: []uint{2, 3, 2, 2, 2},
		},
		Line: true,
		LineProp: props.Line{
			Style: consts.Solid,
			Width: 0.4,
		},
		Align:                  consts.Left,
		VerticalContentPadding: 4,
	})

	return client.pdf
}

func (client *PdfGeneratorClient) GenerateAccountStatement(statement dto.AccountStatement) error {
	client.initStatementPage(statement)
	client.statementTitle(statement)
	client.setCustomerData(statement)
	client.setIBans(statement)
	client.setBalanceSummary(statement)
	client.setTransactions(statement)

	return client.pdf.OutputFileAndClose("account_statement.pdf")
}
