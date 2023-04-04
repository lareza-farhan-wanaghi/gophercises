package pdf

import (
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type certPDF struct {
	pdf     *gofpdf.Fpdf
	w, h, m float64
}

// setCanvas builds the base canvas of the certificate pdf
func (c *certPDF) setCanvas() {
	c.pdf = gofpdf.New("L", "mm", "A4", "")
	c.pdf.AddPage()
	c.w, c.h = c.pdf.GetPageSize()
	c.m = 0.065 * c.w
	c.pdf.SetMargins(c.m, 0, c.m)
}

// setBorders builds the borders of the certificate pdf
func (c *certPDF) setBorders() {
	headerLeftPts := []gofpdf.PointType{
		{X: 0, Y: 0}, {X: 0, Y: c.h * 0.12}, {X: c.w, Y: 0},
	}
	c.pdf.SetFillColor(122, 85, 102)
	c.pdf.Polygon(headerLeftPts, "F")

	headerRightPts := []gofpdf.PointType{
		{X: c.w, Y: 0}, {X: c.w, Y: c.h * 0.12}, {X: c.w * 0.025, Y: c.h * 0},
	}
	c.pdf.SetFillColor(102, 61, 79)
	c.pdf.Polygon(headerRightPts, "F")

	footerRightPts := []gofpdf.PointType{
		{X: c.w, Y: c.h}, {X: c.w, Y: c.h - c.h*0.12}, {X: c.w * 0.025, Y: c.h},
	}
	c.pdf.SetFillColor(122, 85, 102)
	c.pdf.Polygon(footerRightPts, "F")

	footerLeftPts := []gofpdf.PointType{
		{X: 0, Y: c.h}, {X: 0, Y: c.h - c.h*0.12}, {X: c.w, Y: c.h},
	}
	c.pdf.SetFillColor(102, 61, 79)
	c.pdf.Polygon(footerLeftPts, "F")
}

// setBorders builds the body section of the certificate pdf
func (c *certPDF) setBody(name string, date time.Time) {
	c.pdf.SetY(c.h * 0.2)
	c.pdf.SetFont("Times", "B", 48)
	c.pdf.SetTextColor(0, 0, 0)
	c.pdf.CellFormat(0, 7, "Certificate of Cempletion", "", 0, "C", false, 0, "")

	c.pdf.Ln(30)
	c.pdf.SetFont("Helvetica", "", 26)
	c.pdf.CellFormat(0, 7, "This certificate is awarded to", "", 0, "C", false, 0, "")

	c.pdf.Ln(23)
	c.pdf.SetFont("Times", "", 38)
	c.pdf.CellFormat(0, 7, name, "", 0, "C", false, 0, "")

	stringSplit := c.pdf.SplitText("For successfully completing all twenty programming exercises in the Gophercises Go Programming course", c.w*1.25)
	for i, str := range stringSplit {
		if i == 0 {
			c.pdf.Ln(23)
		} else {
			c.pdf.Ln(12)
		}
		c.pdf.SetFont("Helvetica", "", 22)
		c.pdf.CellFormat(0, 7, str, "", 0, "C", false, 0, "")
	}

	var opt gofpdf.ImageOptions
	opt.ImageType = "png"
	imageWidth := float64(40)
	c.pdf.ImageOptions("logo.png", c.w*0.5-imageWidth/2, c.h*0.7, imageWidth, 0, false, opt, 0, "")

	cellWidth := c.w * 0.31
	xOffset := c.w * 0.075
	c.pdf.SetFont("Times", "", 24)
	c.pdf.SetTextColor(139, 139, 139)
	c.pdf.SetXY(xOffset, c.h*0.75)
	c.pdf.CellFormat(cellWidth, 13, fmt.Sprintf("%02d/%02d/%04d", date.Month(), date.Day(), date.Year()), "B", 0, "C", false, 0, "")
	c.pdf.Ln(11.5)
	c.pdf.SetFont("Helvetica", "", 12)
	c.pdf.SetX(xOffset)
	c.pdf.CellFormat(cellWidth, 12, "Date", "", 0, "C", false, 0, "")

	xOffset = c.w - xOffset - cellWidth
	c.pdf.SetFont("Times", "", 24)
	c.pdf.SetTextColor(139, 139, 139)
	c.pdf.SetXY(xOffset, c.h*0.75)
	c.pdf.CellFormat(cellWidth, 13, "Jonathan Calhoun", "B", 0, "C", false, 0, "")
	c.pdf.Ln(11.5)
	c.pdf.SetFont("Helvetica", "", 12)
	c.pdf.SetX(xOffset)
	c.pdf.CellFormat(cellWidth, 12, "Instructor", "", 0, "C", false, 0, "")

}

// Create creates the certificate pdf on the given filename
func (c *certPDF) Create(filename string, name string, date time.Time) error {
	c.setCanvas()
	c.setBorders()
	c.setBody(name, date)

	err := c.pdf.OutputFileAndClose(filename)
	return err
}

// NewCertPDF returns a new certPDF instance
func NewCertPDF() *certPDF {
	return &certPDF{}
}
