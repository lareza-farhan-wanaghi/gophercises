package img

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"strconv"

	svg "github.com/ajstarks/svgo"
)

type ChartItem struct {
	Key   string
	Value int
}

type chart struct {
	items         []*ChartItem
	barWidth      int
	barHeight     int
	offsiteWidth  int
	offsiteHeight int
	valueInterval int
	width         int
	height        int
	maxItemValue  int
}

// init does the setup of the object
func (c *chart) init() {
	maxVal := math.MinInt
	for _, item := range c.items {
		if item.Value > maxVal {
			maxVal = item.Value
		}
	}
	c.maxItemValue = maxVal

	c.width = c.calculateWidth()
	c.height = c.calculateHeight()
}

// calculateWidth returns width calculated from the width-related variables
func (c *chart) calculateWidth() int {
	return c.barWidth*len(c.items) + 2*c.offsiteWidth + (len(c.items)-1)*c.offsiteWidth
}

// calculateWidth returns height calculated from the height-related variables
func (c *chart) calculateHeight() int {

	return c.maxItemValue*c.barHeight + c.offsiteHeight
}

// DrawPNG draws the chart int the png format
func (c chart) DrawPNG(filename string) error {
	if filepath.Ext(filename) != ".png" {
		return fmt.Errorf("filename must be a png format file path")
	}

	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	upLeft := image.Point{0, 0}
	lowRight := image.Point{c.width, c.height}

	result := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	cyan := color.RGBA{100, 200, 200, 0xff}
	for x := 0; x < c.width; x++ {
		for y := 0; y < c.height; y++ {
			clampedX := x % (c.offsiteWidth + c.barWidth)
			sequence := x / (c.offsiteWidth + c.barWidth)
			maxY := c.barHeight * (c.items[min(len(c.items)-1, sequence)].Value)
			switch {
			case clampedX > c.offsiteWidth && c.height-y < maxY:
				result.Set(x, y, cyan)
			default:
				result.Set(x, y, color.White)
			}
		}
	}

	png.Encode(w, result)
	return nil
}

// DrawSvg draws the chart in the svg format
func (c chart) DrawSvg(filename string) error {
	if filepath.Ext(filename) != ".svg" {
		return fmt.Errorf("filename must be an svg format file path")
	}

	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	padding := 75
	canvas := svg.New(w)
	cyan := canvas.RGB(100, 200, 200)
	paddedHeight := c.height + padding
	bottom := paddedHeight - padding/2 - c.offsiteHeight
	left := padding
	paddedWidth := left + c.width

	canvas.Start(paddedWidth, paddedHeight)
	canvas.Rect(0, 0, paddedWidth, paddedHeight, "fill:white")
	canvas.Gstyle("font-size:14pt;fill:rgb(150, 150, 150);text-anchor:middle")
	for i, item := range c.items {
		x := left + c.offsiteWidth*(i+1) + c.barWidth*i
		y := bottom - c.barHeight*item.Value
		canvas.Rect(x, y, c.barWidth, c.barHeight*item.Value, cyan)

		tx := x + c.barWidth/2
		ty := paddedHeight - padding/3
		canvas.Text(tx, ty, item.Key)
	}

	for i := 0; i < c.maxItemValue/c.valueInterval; i++ {
		tx := left / 2
		ty := bottom - (i+1)*c.valueInterval*c.barHeight
		canvas.Text(tx, ty+5, strconv.Itoa((i+1)*c.valueInterval))

		lx := left - 12
		canvas.Line(lx, ty, left, ty, "stroke: rgb(150, 150, 150); stroke-width:2")
	}
	canvas.Gend()
	canvas.Line(left, bottom, left, 0, "stroke: rgb(150, 150, 150); stroke-width:2")
	canvas.Line(left, bottom, paddedWidth, bottom, "stroke: rgb(150, 150, 150); stroke-width:2")
	canvas.End()

	return nil
}

// min returns the min value from the slice
func min(values ...int) int {
	minVal := math.MaxInt
	for _, val := range values {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

// NewChart returns a new chart
func NewChart(items []*ChartItem, barWidth, barHeight, offsiteWidth, offsiteHeight, valueInterval int) *chart {
	result := &chart{
		items:         items,
		barWidth:      barWidth,
		barHeight:     barHeight,
		offsiteWidth:  offsiteWidth,
		offsiteHeight: offsiteHeight,
		valueInterval: valueInterval}

	result.init()
	return result
}
