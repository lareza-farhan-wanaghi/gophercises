package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/lareza-farhan-wanaghi/gophercises/img"
)

// main provides the entry point of the app
func main() {
	inpath := os.Args[1]
	f, err := os.Open(inpath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	chartItems := []*img.ChartItem{}
	for _, d := range data {
		itemValue, err := strconv.Atoi(d[1])
		if err != nil {
			panic(err)
		}

		item := &img.ChartItem{
			Key:   d[0],
			Value: itemValue,
		}
		chartItems = append(chartItems, item)
	}

	chart := img.NewChart(chartItems, 50, 10, 25, 25, 5)

	outpath := os.Args[2]
	if filepath.Ext(outpath) == ".png" {
		err := chart.DrawPNG(outpath)
		if err != nil {
			fmt.Println(err)
		}
	} else if filepath.Ext(outpath) == ".svg" {
		err := chart.DrawSvg(outpath)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		panic("Unsupported file extension. Currently available extension are .png and .svg")
	}
}
