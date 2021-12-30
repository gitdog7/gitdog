package painter

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/render"
	"io"
	"os"
)

// PaintGraph given a graph, paint with html format and output file to outputPath
func PaintGraph(graph *charts.Graph, outputPath string) {
	page := newPage(graph.Title.Title)
	page.AddCharts(
		graph,
	)

	f, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}

	if err := page.Render(io.MultiWriter(f)); err != nil {
		panic(err)
	}
}

// newPage create a new page with given title and size.
func newPage(title string) *components.Page {
	page := &components.Page{}
	page.Height = "720px"
	page.Width = "1280px"
	page.AssetsHost = "./"
	page.PageTitle = title
	page.Assets.InitAssets()
	page.Renderer = render.NewPageRender(page, page.Validate)
	page.Layout = components.PageCenterLayout
	return page
}
