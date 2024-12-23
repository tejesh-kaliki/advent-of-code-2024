package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type Edge struct {
	V1, V2 string
}

type Graph struct {
	Vertices []string
	Edges    [][]bool
}

func FindInterconnectedComputersOfSize3(graph Graph) [][3]string {
	edges := graph.Edges
	groups := make([][3]string, 0)
	for i := 0; i < len(graph.Vertices); i++ {
		for j := i + 1; j < len(graph.Vertices); j++ {
			for k := j + 1; k < len(graph.Vertices); k++ {
				if edges[i][j] && edges[j][k] && edges[k][i] {
					groups = append(groups, [3]string{graph.Vertices[i], graph.Vertices[j], graph.Vertices[k]})
				}
			}
		}
	}
	return groups
}

func FindInterconnectedComputersOfSize3Indices(graph Graph) [][]int {
	edges := graph.Edges
	groups := make([][]int, 0)
	for i := 0; i < len(graph.Vertices); i++ {
		for j := i + 1; j < len(graph.Vertices); j++ {
			for k := j + 1; k < len(graph.Vertices); k++ {
				if edges[i][j] && edges[j][k] && edges[k][i] {
					groups = append(groups, []int{i, j, k})
				}
			}
		}
	}
	return groups
}

func DoesTheComputerGroupContainLetter(vertices [3]string, letter string) bool {
	for _, vertex := range vertices {
		if strings.HasPrefix(string(vertex), letter) {
			return true
		}
	}
	return false
}

func FindVerticesConnectingAll(graph Graph, vertIndices []int) []int {
	commonVertices := make([]int, 0)
	for i := vertIndices[len(vertIndices)-1]; i < len(graph.Vertices); i++ {
		allFound := true
		for _, index := range vertIndices {
			if !graph.Edges[i][index] {
				allFound = false
				break
			}
		}

		if allFound {
			commonVertices = append(commonVertices, i)
		}
	}

	return commonVertices
}

func SolvePart1(graph Graph) int {
	connectedComputers := FindInterconnectedComputersOfSize3(graph)
	total := 0
	for _, group := range connectedComputers {
		if DoesTheComputerGroupContainLetter(group, "t") {
			total++
		}
	}
	return total
}

func GetTextFromIndices(graph Graph, indices []int) string {
	text := ""
	for i, vert := range indices {
		if i == 0 {
			text += graph.Vertices[vert]
		} else {
			text += "," + graph.Vertices[vert]
		}
	}
	return text
}

func SolvePart2(graph Graph) string {
	groupsOfNIndices := FindInterconnectedComputersOfSize3Indices(graph)

	for {
		nextGroupsOfIndices := make([][]int, 0)

		for _, groupIndices := range groupsOfNIndices {
			connected := FindVerticesConnectingAll(graph, groupIndices)
			nextGroupsOfIndices = slices.Grow(nextGroupsOfIndices, len(connected))
			for _, vert := range connected {
				newGroup := make([]int, len(groupIndices), len(groupIndices)+1)
				copy(newGroup, groupIndices)
				newGroup = append(newGroup, vert)

				nextGroupsOfIndices = append(nextGroupsOfIndices, newGroup)
			}
		}

		if len(nextGroupsOfIndices) == 0 {
			return GetTextFromIndices(graph, groupsOfNIndices[0])
		}

		groupsOfNIndices = nextGroupsOfIndices
	}
}

func ReadInput(input string) Graph {
	lines := strings.Split(input, "\n")
	edges := make([]Edge, 0, len(lines))
	vertices := make([]string, 0, len(lines)*2)
	for _, line := range lines {
		v1, v2, _ := strings.Cut(line, "-")
		if !slices.Contains(vertices, string(v1)) {
			vertices = append(vertices, string(v1))
		}
		if !slices.Contains(vertices, string(v2)) {
			vertices = append(vertices, string(v2))
		}

		edges = append(edges, Edge{v1, v2})
	}

	slices.Sort(vertices)
	graph := Graph{Vertices: vertices, Edges: make([][]bool, len(vertices))}
	for i := range graph.Edges {
		graph.Edges[i] = make([]bool, len(vertices))
	}

	for _, edge := range edges {
		v1Index := slices.Index(graph.Vertices, edge.V1)
		v2Index := slices.Index(graph.Vertices, edge.V2)
		graph.Edges[v1Index][v2Index] = true
		graph.Edges[v2Index][v1Index] = true
	}

	return graph
}

func (graph Graph) Display() string {
	text := ""
	for _, row := range graph.Edges {
		for _, val := range row {
			if val {
				text += "1"
			} else {
				text += " "
			}
		}
		text += "\n"
	}
	return text
}

func main() {
	graph := ReadInput(input)
	fmt.Println("Part 1 Solution:", SolvePart1(graph))
	fmt.Println("Part 2 Solution:", SolvePart2(graph))
}
