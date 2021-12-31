package contributor

import (
	"fmt"
	"github.com/gitdog7/gitdog/src/repo"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"math"
	"sort"
)

// ContributeGraphOption graph options.
type ContributeGraphOption struct {
	// TopK only draw top k contributors, 0 means all
	TopK int

	// Type generate graph type, "circular", or "force"
	Type string
}

// GenerateGraph generate a graph, given github repo and option
func GenerateGraph(repo *repo.GitHubRepo, option ContributeGraphOption) *charts.Graph {
	return genGraph(repo, option)
}

////////////////////////////////////////////
// begin privates
////////////////////////////////////////////

// genNodes for each contributor in the repo, generate a node to represent.
// the more contributions of this contributor, the larger node symbolSize
func genNodes(repo *repo.GitHubRepo, option ContributeGraphOption) []opts.GraphNode {
	nodes := make([]opts.GraphNode, 0)
	for _, c := range repo.Contributors {
		_, isMember := repo.Members[*c.Login]
		category := ContributorID
		if isMember {
			category = MemberID
		}
		nodes = append(nodes, opts.GraphNode{
			Name:       fmt.Sprintf("%v(%v)", *c.Login, *c.Contributions),
			SymbolSize: getNodeSymbolSizeFactor(len(repo.Contributors))*math.Log10(float64(*c.Contributions)) + 10,
			Value:      float32(*c.Contributions)/100.0 + 10,
			Category:   category,
		})
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Value > nodes[j].Value
	})

	// keep only topk contributors
	if option.TopK > 0 && len(nodes) > option.TopK {
		nodes = nodes[0:option.TopK]
	}

	return nodes
}

func getNodeSymbolSizeFactor(num int) float64 {
	if num >= 100 {
		return 10.0
	} else {
		return 20.0 - float64(num)/100.0*10.0
	}
}

// genLinks generate follower edge for two contributors
// generate a direct link from A->B if A follows B in GitHub
func genLinks(repo *repo.GitHubRepo, option ContributeGraphOption) []opts.GraphLink {
	links := make([]opts.GraphLink, 0)
	for user1, followers := range repo.Followers {
		for _, user2 := range followers {
			if c1, f1 := repo.Contributors[user1]; f1 {
				if c2, f2 := repo.Contributors[user2]; f2 {
					links = append(links, opts.GraphLink{
						// c2 follows c1
						Source: fmt.Sprintf("%v(%v)", *c2.Login, *c2.Contributions),
						Target: fmt.Sprintf("%v(%v)", *c1.Login, *c1.Contributions),
						Value:  float32(*c1.Contributions),
						Label: &opts.EdgeLabel{
							Show: false,
						},
					})
				}
			}

		}
	}
	return links
}

// generateCategories categories is member or contributor
func generateCategories() []*opts.GraphCategory {
	return []*opts.GraphCategory{
		{
			Name: "Member",
		},
		{
			Name: "Contributor",
		},
	}
}

func getForceOption(option ContributeGraphOption) *opts.GraphForce {
	if option.Type == "circular" {
		return &opts.GraphForce{
			InitLayout: "circular",
			Repulsion:  0,
			Gravity:    0,
			EdgeLength: 0,
		}
	} else {
		return &opts.GraphForce{
			InitLayout: "force",
			Repulsion:  100,
			Gravity:    0.1,
			EdgeLength: 150,
		}
	}
}

func getEdgeLineOption(option ContributeGraphOption) opts.LineStyle {
	if option.Type == "circular" {
		return opts.LineStyle{
			Color:     "target",
			Curveness: 0.3,
			Width:     1.1,
			Opacity:   1.0,
		}
	} else {
		return opts.LineStyle{
			Color:     "target",
			Curveness: 0.3,
			Width:     1.1,
			Opacity:   1.0,
		}
	}
}

const MemberID = 0
const ContributorID = 1

// genCircleGraph draw circular graph
func genGraph(repo *repo.GitHubRepo, option ContributeGraphOption) *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: repo.Owner + "/" + repo.Repository + " Contributors Graph(" + option.Type + ")",
			Subtitle: fmt.Sprintf("Top %v Contributors and Following Relationships\n", option.TopK) +
				"Node name: contributor(#contributions)\n" +
				"Edge A->B: A follows B in GitHub",
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
		genNodes(repo, option),
		genLinks(repo, option)).
		SetSeriesOptions(
			charts.WithGraphChartOpts(
				opts.GraphChart{
					Force:              getForceOption(option),
					FocusNodeAdjacency: true,
					Layout:             option.Type,
					Roam:               true,
					Draggable:          true,
					Categories:         generateCategories(),
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
			charts.WithLineStyleOpts(getEdgeLineOption(option)),
		)

	return graph
}
