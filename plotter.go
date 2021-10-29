package kmeans

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

// SimplePlotter is the default standard plotter for 2-dimensional data sets
type SimplePlotter struct {
}

// A monokai-ish color palette
var colors = []drawing.Color{
	drawing.ColorFromHex("f92672"),
	drawing.ColorFromHex("89bdff"),
	drawing.ColorFromHex("66d9ef"),
	drawing.ColorFromHex("67210c"),
	drawing.ColorFromHex("7acd10"),
	drawing.ColorFromHex("af619f"),
	drawing.ColorFromHex("fd971f"),
	drawing.ColorFromHex("dcc060"),
	drawing.ColorFromHex("545250"),
	drawing.ColorFromHex("4b7509"),
	drawing.ColorFromHex("ff00ff"),
	drawing.ColorFromHex("ff2354"),
	drawing.ColorFromHex("452c1a"),
	drawing.ColorFromHex("224f42"),
	drawing.ColorFromHex("7aab23"),
	drawing.ColorFromHex("af002f"),
	drawing.ColorFromHex("fd093c"),
	drawing.ColorFromHex("dc88f9"),
	drawing.ColorFromHex("54926e"),
	drawing.ColorFromHex("c3210a"),
}

// Plot draw a 2-dimensional data set into a PNG file named {k_iteration}.png
func (p SimplePlotter) Plot(cc Clusters, iteration int) error {
	var series []chart.Series

	// draw data points
	for i, c := range cc {
		series = append(series, chart.ContinuousSeries{
			Style: chart.Style{
				StrokeWidth: chart.Disabled,
				DotColor:    colors[i%len(colors)],
				DotWidth:    8,
			},
			XValues: c.PointsInDimension(0),
			YValues: c.PointsInDimension(1),
		})
	}

	// draw cluster center points
	series = append(series, chart.ContinuousSeries{
		Style: chart.Style{
			StrokeWidth: chart.Disabled,
			DotColor:    drawing.ColorBlack,
			DotWidth:    16,
		},
		XValues: cc.CentersInDimension(0),
		YValues: cc.CentersInDimension(1),
	})

	graph := chart.Chart{
		Height: 1024,
		Width:  1024,
		Series: series,
		XAxis: chart.XAxis{
			Style: chart.Style{},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fmt.Sprintf("%d_%d.png", len(cc), iteration), buffer.Bytes(), 0644)
}

// Plot draw a 2-dimensional data set into a PNG file named {k_iteration}.png
func (p SimplePlotter) Plot2(cc Clusters, iteration int) error {
	var series []chart.Series

	// draw data points
	for i, c := range cc {
		series = append(series, chart.ContinuousSeries{
			Style: chart.Style{
				StrokeWidth: chart.Disabled,
				DotColor:    colors[i%len(colors)],
				DotWidth:    8,
			},
			XValues: c.PointsInDimension(0),
			YValues: c.PointsInDimension(1),
		})
	}

	// draw cluster center points
	series = append(series, chart.ContinuousSeries{
		Style: chart.Style{
			StrokeWidth: chart.Disabled,
			DotColor:    drawing.ColorBlack,
			DotWidth:    16,
		},
		XValues: cc.CentersInDimension(0),
		YValues: cc.CentersInDimension(1),
	})

	graph := chart.Chart{
		Height: 1024,
		Width:  1024,
		Series: series,
		XAxis: chart.XAxis{
			Style: chart.Style{},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fmt.Sprintf("new_%d_%d.png", len(cc), iteration), buffer.Bytes(), 0644)
}
