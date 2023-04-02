package main

import (
	"os"

	"github.com/lareza-farhan-wanaghi/gophercises/img"
)

func main() {
	pngFile, err := os.OpenFile("demo.png", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer pngFile.Close()

	svgFile, err := os.OpenFile("demo.svg", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer svgFile.Close()

	chartItems := []*img.ChartItem{
		{Key: "Jan", Value: 50},
		{Key: "Feb", Value: 30},
		{Key: "Mar", Value: 40},
		{Key: "Apr", Value: 25},
		{Key: "May", Value: 30},
		{Key: "Jun", Value: 35},
		{Key: "Jul", Value: 15},
		{Key: "Aug", Value: 5},
		{Key: "Sep", Value: 40},
		{Key: "Oct", Value: 35},
		{Key: "Nov", Value: 30},
		{Key: "Des", Value: 40},
	}

	chart := img.NewChart(chartItems, 50, 10, 25, 25, 5)

	chart.DrawSvg(svgFile)
	chart.DrawPNG(pngFile)
}
