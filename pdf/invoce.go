package pdf

import (
	"fmt"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

const sGap = 6.25
const mGap = 7

type InvoiceItem struct {
	Name         string
	PricePerUnit int
	Quantity     int
}

type invoicePDF struct {
	pdf     *gofpdf.Fpdf
	w, h, m float64
}

// setCanvas builds the base canvas of the invoice pdf
func (inv *invoicePDF) setCanvas() {
	inv.pdf = gofpdf.New("P", "mm", "A4", "")
	inv.pdf.AddPage()
	inv.w, inv.h = inv.pdf.GetPageSize()
	inv.m = 0.065 * inv.w
	inv.pdf.SetMargins(inv.m, 0, inv.m)
}

// setHeader builds the header section of the invoice pdf
func (inv *invoicePDF) setHeader() {
	inv.pdf.SetFillColor(102, 61, 79)
	headerPts := []gofpdf.PointType{
		{X: 0, Y: 0}, {X: 0, Y: inv.h * 0.1}, {X: inv.w, Y: inv.h * 0.1}, {X: inv.w, Y: 0},
	}
	inv.pdf.Polygon(headerPts, "F")

	inv.pdf.SetFont("Arial", "B", 36)
	inv.pdf.SetTextColor(255, 255, 255)
	inv.pdf.SetXY(inv.m, inv.h*0.05)
	inv.pdf.Cell(0, 0, "INVOICE")

	var opt gofpdf.ImageOptions
	opt.ImageType = "png"
	inv.pdf.ImageOptions("logo.png", inv.w*0.375, inv.h*0.015, 0, inv.h*0.075, false, opt, 0, "")

	inv.pdf.SetFont("Arial", "", 11)
	inv.pdf.SetFillColor(0, 0, 0)
	inv.pdf.SetY(inv.h*0.05 - sGap)
	inv.pdf.CellFormat(inv.w*0.7, 0, "(000) 0000-000", "", 2, "R", false, 0, "")
	inv.pdf.SetY(inv.h * 0.05)
	inv.pdf.CellFormat(inv.w*0.7, 0, "exam@example.com", "", 2, "R", false, 0, "")
	inv.pdf.SetY(inv.h*0.05 + sGap)
	inv.pdf.CellFormat(inv.w*0.7, 0, "example.com", "", 2, "R", false, 0, "")

	inv.pdf.SetFillColor(0, 0, 0)
	inv.pdf.SetY(inv.h*0.05 - sGap)
	inv.pdf.CellFormat(0, 0, "123 Fake St0", "", 2, "R", false, 0, "")
	inv.pdf.SetY(inv.h * 0.05)
	inv.pdf.CellFormat(0, 0, "Some town, PA", "", 2, "R", false, 0, "")
	inv.pdf.SetY(inv.h*0.05 + sGap)
	inv.pdf.CellFormat(0, 0, "12345", "", 2, "R", false, 0, "")
}

// setFooter builds the footer section of the invoice pdf
func (inv *invoicePDF) setFooter() {
	footerPts := []gofpdf.PointType{
		{X: 0, Y: inv.h}, {X: 0, Y: inv.h * 0.95}, {X: inv.w, Y: inv.h * 0.975}, {X: inv.w, Y: inv.h},
	}
	inv.pdf.SetFillColor(102, 61, 79)
	inv.pdf.Polygon(footerPts, "F")

}

// setSummary builds the summary section of the invoice pdf
func (inv *invoicePDF) setSummary(subtotal int) {
	inv.pdf.SetFont("Times", "", 14)
	inv.pdf.SetTextColor(185, 185, 185)
	inv.pdf.SetXY(inv.m, inv.h*0.15)
	inv.pdf.Cell(0, 0, "Billed To")
	inv.pdf.SetTextColor(0, 0, 0)
	inv.pdf.SetXY(inv.m, inv.h*0.15+mGap)
	inv.pdf.Cell(0, 0, "Client Name")
	inv.pdf.SetXY(inv.m, inv.h*0.15+mGap+sGap)
	inv.pdf.Cell(0, 0, "1 CLient Address")
	inv.pdf.SetXY(inv.m, inv.h*0.15+mGap+sGap*2)
	inv.pdf.Cell(0, 0, "City, State, Country")
	inv.pdf.SetXY(inv.m, inv.h*0.15+mGap+sGap*3)
	inv.pdf.Cell(0, 0, "Postal Code")

	inv.pdf.SetTextColor(185, 185, 185)
	inv.pdf.SetXY(inv.w*0.35, inv.h*0.15)
	inv.pdf.Cell(0, 0, "Invoice Number")
	inv.pdf.SetTextColor(0, 0, 0)
	inv.pdf.SetXY(inv.w*0.35, inv.h*0.15+mGap)
	inv.pdf.Cell(0, 0, "00000000")
	inv.pdf.SetTextColor(185, 185, 185)
	inv.pdf.SetXY(inv.w*0.35, inv.h*0.15+mGap+sGap*1.875)
	inv.pdf.Cell(0, 0, "Date of Issue")
	inv.pdf.SetTextColor(0, 0, 0)
	inv.pdf.SetXY(inv.w*0.35, inv.h*0.15+mGap*2+sGap*1.875)
	inv.pdf.Cell(0, 0, "mm/dd/yyyy")

	inv.pdf.SetTextColor(185, 185, 185)
	inv.pdf.SetXY(0, inv.h*0.15)
	inv.pdf.CellFormat(0, 0, "Invoice Total", "", 2, "R", false, 0, "")
	inv.pdf.SetXY(0, inv.h*0.19)
	inv.pdf.SetFont("Times", "", 46)
	inv.pdf.SetTextColor(128, 98, 107)
	inv.pdf.CellFormat(0, 0, fmt.Sprintf("$%.02f", float64(subtotal)), "", 2, "R", false, 0, "")

	inv.pdf.SetFillColor(102, 61, 79)
	inv.pdf.SetXY(inv.m*0.5, inv.h*0.285)
	inv.pdf.CellFormat(inv.w-inv.m, 1, "", "", 2, "l", true, 0, "")
}

// setDetails builds the details section of the invoice pdf
func (inv *invoicePDF) setDetails(items []*InvoiceItem) int {
	margedWidth := inv.w - inv.m*2

	headers := []string{"Description", "Price Per Unit", "Quantity", "Amount"}
	cellWidth := []float64{margedWidth * 0.45, margedWidth * 0.2, margedWidth * 0.15, margedWidth * 0.2}

	inv.pdf.SetFont("Times", "", 14)
	inv.pdf.SetXY(inv.m, inv.h*0.3)
	for i, header := range headers {
		if i < len(headers)-1 {
			inv.pdf.SetTextColor(185, 185, 185)
		} else {
			inv.pdf.SetTextColor(0, 0, 0)
		}

		textAlign := "R"
		if i == 0 {
			textAlign = "L"
		}
		inv.pdf.CellFormat(cellWidth[i], 7, header, "", 0, textAlign, false, 0, "")
	}

	inv.pdf.SetTextColor(0, 0, 0)
	subtotal := 0
	for _, item := range items {
		amount := item.PricePerUnit * item.Quantity
		subtotal += amount

		nameSplit := inv.pdf.SplitText(item.Name, cellWidth[0])
		ppuSplit := inv.pdf.SplitText(fmt.Sprintf("$%.02f", float64(item.PricePerUnit)), cellWidth[1])
		qualtitySplit := inv.pdf.SplitText(strconv.Itoa(item.Quantity), cellWidth[2])
		amountSplit := inv.pdf.SplitText(fmt.Sprintf("$%.02f", float64(amount)), cellWidth[3])

		cellHeight := float64(13)
		maxLength := max(len(nameSplit), len(ppuSplit), len(qualtitySplit), len(amountSplit))
		for i := 0; i < maxLength; i++ {
			if i != 0 {
				inv.pdf.Ln(7.5)
			} else {
				inv.pdf.Ln(-1)
			}
			inv.pdf.SetX(inv.m)

			name := ""
			if i < len(nameSplit) {
				name = nameSplit[i]
			}
			inv.pdf.CellFormat(cellWidth[0], cellHeight, name, "", 0, "", false, 0, "")

			ppu := ""
			if i < len(ppuSplit) {
				ppu = ppuSplit[i]
			}
			inv.pdf.CellFormat(cellWidth[1], cellHeight, ppu, "", 0, "R", false, 0, "")

			quantity := ""
			if i < len(qualtitySplit) {
				quantity = qualtitySplit[i]
			}
			inv.pdf.CellFormat(cellWidth[2], cellHeight, quantity, "", 0, "R", false, 0, "")

			amount := ""
			if i < len(amountSplit) {
				amount = amountSplit[i]
			}
			inv.pdf.CellFormat(cellWidth[3], cellHeight, amount, "", 0, "R", false, 0, "")
		}
		inv.pdf.Ln(-1)
		inv.pdf.SetX(inv.m * 0.75)
		inv.pdf.CellFormat(inv.w-inv.m*1.5, 0.35, "", "", 2, "l", true, 0, "")
	}
	inv.pdf.Ln(10)
	inv.pdf.SetTextColor(185, 185, 185)
	inv.pdf.CellFormat(margedWidth*0.7, 13, "Subtotal", "", 0, "R", false, 0, "")
	inv.pdf.SetTextColor(0, 0, 0)
	inv.pdf.CellFormat(margedWidth*0.3, 13, fmt.Sprintf("$%.02f", float64(subtotal)), "", 0, "R", false, 0, "")
	return subtotal
}

// Create creates the invoice pdf on the given filename
func (inv *invoicePDF) Create(filename string, items []*InvoiceItem) error {
	inv.setCanvas()
	inv.setHeader()
	inv.setFooter()

	subtotal := inv.setDetails(items)
	inv.setSummary(subtotal)

	err := inv.pdf.OutputFileAndClose(filename)
	return err
}

// NewinvoicePDF returns a new invoicePDF instance
func NewInvoicePDF() *invoicePDF {
	return &invoicePDF{}
}

// max returns the max value from the slice
func max(values ...int) int {
	maxVal := 0
	for _, val := range values {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
