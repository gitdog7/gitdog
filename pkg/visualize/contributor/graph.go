package contributor

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/live77/gitdog/pkg/repo"
	"math"
	"sort"
)

// ContributeGraphViz is a visualizer, given a GitHubRepo data.
// it will output a circular graph, the nodes in the graph represent the contributors
// the edge represent the follow relationship in those contributors
type ContributeGraphViz struct {
	Repo *repo.GitHubRepo
}

func (v *ContributeGraphViz) GenerateGraph() *charts.Graph {
	return v.genCircleGraph(v.Repo)
}

////////////////////////////////////////////
// begin privates
////////////////////////////////////////////

// genNodes for each contributor in the repo, generate a node to represent.
// the more contributions of this contributor, the larger node symbolSize
func (v *ContributeGraphViz) genNodes() []opts.GraphNode {
	nodes := make([]opts.GraphNode, 0)
	for _, c := range v.Repo.Contributors {
		_, isMember := v.Repo.Members[*c.Login]
		category := ContributorID
		if isMember {
			category = MemberID
		}
		nodes = append(nodes, opts.GraphNode{
			Name:       fmt.Sprintf("%v(%v)", *c.Login, *c.Contributions),
			SymbolSize: 15*math.Log10(float64(*c.Contributions)) + 10,
			Value:      float32(*c.Contributions)/100.0 + 10,
			Category:   category,
		})
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Value > nodes[j].Value
	})
	return nodes
}

// genLinks generate follower edge for two contributors
// generate a direct link from A->B if A follows B in GitHub
func (v *ContributeGraphViz) genLinks() []opts.GraphLink {
	links := make([]opts.GraphLink, 0)
	for user1, followers := range v.Repo.Followers {
		for _, user2 := range followers {
			if c1, f1 := v.Repo.Contributors[user1]; f1 {
				if c2, f2 := v.Repo.Contributors[user2]; f2 {
					links = append(links, opts.GraphLink{
						Source: fmt.Sprintf("%v(%v)", *c2.Login, *c2.Contributions),
						Target: fmt.Sprintf("%v(%v)", *c1.Login, *c1.Contributions),
					})
				}
			}

		}
	}
	return links
}

// generateCategories categories is member or contributor
func (v *ContributeGraphViz) generateCategories() []*opts.GraphCategory {
	return []*opts.GraphCategory{
		{
			Name: "Member",
		},
		{
			Name: "Contributor",
		},
	}
}

const MemberID = 0
const ContributorID = 1

// genCircleGraph draw circular graph
func (v *ContributeGraphViz) genCircleGraph(repo *repo.GitHubRepo) *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Contributors and Following Relationships",
			Subtitle: repo.Owner + "/" + repo.Repository + "\n" +
				"node name format: contributor(#contributions)\n" +
				"edge A->B means A follows B in GitHub",
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
			Data: []string{"Member", "Contributor"},
			Left: "left",
			Top:  "10%",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Height: "720px",
			Width:  "1280px",
			Theme:  "white",
		}),
	)

	graph.AddSeries("contributors",
		v.genNodes(),
		v.genLinks()).
		SetSeriesOptions(
			charts.WithGraphChartOpts(
				opts.GraphChart{
					Force: &opts.GraphForce{
						InitLayout: "circular",
						Repulsion:  0,
						Gravity:    0,
						EdgeLength: 0,
					},
					FocusNodeAdjacency: true,
					Layout:             "circular",
					Roam:               true,
					Draggable:          true,
					Categories:         v.generateCategories(),
					EdgeSymbol:         []string{"circle", "arrow"},
					EdgeSymbolSize:     []int{0, 10},
				}),
			charts.WithLabelOpts(
				opts.Label{
					Show:      true,
					Position:  "right",
					Formatter: "{b}",
				}),
			charts.WithEmphasisOpts(opts.Emphasis{
				Label: &opts.Label{
					Show:      true,
					Position:  "right",
					Formatter: "{b}",
				},
				ItemStyle: &opts.ItemStyle{
					BorderWidth: 1.0,
					BorderColor: "red",
					Opacity:     1.0,
				},
			}),
			charts.WithCircularStyleOpts(opts.CircularStyle{
				RotateLabel: true,
			}),
			charts.WithLineStyleOpts(opts.LineStyle{
				Color:     "target",
				Curveness: 0.3,
				Width:     1.1,
				Opacity:   1.0,
			}),
		)

	return graph
}