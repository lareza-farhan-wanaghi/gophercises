package main

import (
	"fmt"

	"github.com/lareza-farhan-wanaghi/gophercises/img"
)

// main provides the entry point of the app
func main() {

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

	err := chart.DrawSvg("demo-svg.png")
	if err != nil {
		fmt.Println(err)
	}
	err = chart.DrawPNG("demo-png.svg")
	if err != nil {
		fmt.Println(err)
	}
}
