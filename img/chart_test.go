package img

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"testing"
)

// TestCalculatedWidth tests the calculated width of the chart struct
func TestCalculatedWidth(t *testing.T) {
	for k, v := range testTable.calculatedWidth {
		splits := strings.Split(k, ",")

		numOfItems, err := strconv.Atoi(splits[0])
		if err != nil {
			t.Fatal(err)
		}

		chartItems := []*ChartItem{}
		for i := 0; i < numOfItems; i++ {
			chartItems = append(chartItems, &ChartItem{})
		}

		barWidth, err := strconv.Atoi(splits[1])
		if err != nil {
			t.Fatal(err)
		}

		offsiteWidth, err := strconv.Atoi(splits[2])
		if err != nil {
			t.Fatal(err)
		}

		chart := NewChart(chartItems, barWidth, 0, offsiteWidth, 0, 0)

		expectedWidth, err := strconv.Atoi(v)
		if err != nil {
			t.Fatal(err)
		}

		if expectedWidth != chart.width {
			t.Fatalf("expected %d but got %d. k:%s v:%s", expectedWidth, chart.width, k, v)
		}
	}
}

// TestCalculatedHeight tests the calculated height of the chart struct
func TestCalculatedHeight(t *testing.T) {
	for k, v := range testTable.calculatedHeight {
		splits := strings.Split(k, ",")

		chartItems := []*ChartItem{}
		for i := 0; i < len(splits)-2; i++ {
			itemValue, err := strconv.Atoi(splits[i])
			if err != nil {
				t.Fatal(err)
			}

			chartItems = append(chartItems, &ChartItem{Value: itemValue})
		}

		barHeight, err := strconv.Atoi(splits[len(splits)-2])
		if err != nil {
			t.Fatal(err)
		}

		offsiteHeight, err := strconv.Atoi(splits[len(splits)-1])
		if err != nil {
			t.Fatal(err)
		}

		chart := NewChart(chartItems, 0, barHeight, 0, offsiteHeight, 0)

		expectedHeight, err := strconv.Atoi(v)
		if err != nil {
			t.Fatal(err)
		}

		if expectedHeight != chart.height {
			t.Fatalf("expected %d but got %d. k:%s v:%s", expectedHeight, chart.height, k, v)
		}
	}
}

// TestDrawSvg tests the DrawSvg function of the chart struct
func TestDrawSvg(t *testing.T) {
	for k, v := range testTable.drawSvg {
		mockChartItems := []*ChartItem{
			{Key: "1", Value: 1},
			{Key: "2", Value: 2},
			{Key: "3", Value: 3},
		}
		chart := NewChart(mockChartItems, 10, 10, 10, 10, 10)
		err := chart.DrawSvg(k)

		if v == "error" {
			if err == nil {
				t.Fatalf("expected an error but got nil. k:%s v:%s", k, v)
			}
			continue
		}

		if err != nil {
			t.Fatal(err)
		}

		if _, err := os.Stat(v); errors.Is(err, os.ErrNotExist) {
			t.Fatalf("expected a file with a path of %s. k:%s v:%s", v, k, v)
		}

		err = os.RemoveAll(v)
		if err != nil {
			t.Fatal(err)
		}
	}
}

// TestDrawPng tests the DrawPng function of the chart struct
func TestDrawPng(t *testing.T) {
	for k, v := range testTable.drawPng {
		mockChartItems := []*ChartItem{
			{Key: "1", Value: 1},
			{Key: "2", Value: 2},
			{Key: "3", Value: 3},
		}
		chart := NewChart(mockChartItems, 10, 10, 10, 10, 10)
		err := chart.DrawPNG(k)
		if v == "error" {
			if err == nil {
				t.Fatalf("expected an error but got nil. k:%s v:%s", k, v)
			}
			continue
		}

		if err != nil {
			t.Fatal(err)
		}

		if _, err := os.Stat(v); errors.Is(err, os.ErrNotExist) {
			t.Fatalf("expected a file with a path of %s. k:%s v:%s", v, k, v)
		}

		err = os.RemoveAll(v)
		if err != nil {
			t.Fatal(err)
		}
	}
}
