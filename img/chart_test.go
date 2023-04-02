package img

import (
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
