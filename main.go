package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/RyanCarrier/dijkstra"
)

const (
	nameSymbFile = "graph_symbolic.txt"
	nameNumFile  = "graph_numeric.txt"
	// src          = "A274"
	// dst          = "A258"
)

func main() {
	mapVerts := createGraphNumFile(nameSymbFile, nameNumFile)

	var src, dst, ok string
	for {
		fmt.Println("Please enter source cabinet:")
		fmt.Scanln(&src)
		fmt.Println("Please enter destination cabinet:")
		fmt.Scanln(&dst)
		distance, path := shortestPath(nameNumFile, vertToNum(src, mapVerts), vertToNum(dst, mapVerts))
		fmt.Printf("Shortest distance between %s and %s is %d following path %v\n", src, dst, distance, pathToSymb(path, mapVerts))
		fmt.Println("Are you want to exit?:")
		fmt.Scanln(&ok)
		if ok == "yes" {
			break
		}
	}
}

// Find shortest path between two vertexs of graph from numeric file
func shortestPath(nameGraphFile string, src, dst int) (int64, []int) {
	graph, _ := dijkstra.Import(nameGraphFile)
	best, err := graph.Shortest(src, dst)
	if err != nil {
		log.Fatal(err)
	}
	return best.Distance, best.Path
}

// Translate vertex symbolic name to numeric
func vertToNum(name string, relation map[string]int) int {
	return relation[name]
}

// Translate numeric view of path to symbolic view
func pathToSymb(path []int, relation map[string]int) []string {
	var strPath = make([]string, 0)
	for _, v := range path {
		for i1, v1 := range relation {
			if v1 == v {
				strPath = append(strPath, i1)
			}
		}
	}
	return strPath
}

// Create two maps relation of graph vertexs and arcs from symbolic file
func createMapsRelVertsArcs(symbFileName *os.File) (map[string]int, map[int][]string) {
	var mapVerts = make(map[string]int)
	var mapVertsArcs = make(map[int][]string)

	scanner := bufio.NewScanner(symbFileName)
	for i := 0; scanner.Scan(); i++ {
		vertStr := strings.Split(scanner.Text(), " ")[0]
		mapVerts[vertStr] += i
		arcsStrArr := strings.Split(scanner.Text(), " ")[1:]
		mapVertsArcs[mapVerts[vertStr]] = arcsStrArr
	}
	return mapVerts, mapVertsArcs
}

// Create weighted map of direction graph from two relation maps
func createMapDirGraph(mapVerts map[string]int, mapVertsArcs map[int][]string) map[int][]int {
	var mapDirGraph = make(map[int][]int)

	for i, v := range mapVertsArcs {
		mapDirGraph[i] = func() []int {
			var numArr = make([]int, 0)
			for _, v1 := range v {
				strArr := strings.Split(v1, ",")
				numArr = append(numArr, mapVerts[strArr[0]])
				numArr = append(numArr, int(([]rune(strArr[1])[0] - '0')))
			}
			return numArr
		}()
	}
	return mapDirGraph
}

// Create map of undirection graph from direction graph map
func createMapUndirGraph(mapDirGraph map[int][]int) map[int][]int {
	var mapUndirGraph = mapDirGraph

	for i, v := range mapDirGraph {
		for i1 := 0; i1 < len(v); i1 += 2 {
			check := func() bool {
				for i2 := 0; i2 < len(mapUndirGraph[v[i1]]); i2 += 2 {
					if mapUndirGraph[v[i1]][i2] == i {
						return true
					}
				}
				return false
			}()
			if check {
				continue
			} else {
				mapUndirGraph[v[i1]] = append(mapDirGraph[v[i1]], i, v[i1+1])
			}
		}
	}
	return mapUndirGraph
}

// Create numeric file of undirect graph from symbolic graph file (direct).
//  Return map of relation symbolic name of vertexs and numeric name of vertexs
func createGraphNumFile(symbFileName, numFileName string) map[string]int {
	fileSymb, err := os.Open(nameSymbFile)
	if err != nil {
		panic(err)
	}
	defer fileSymb.Close()
	fileNum, err := os.Create(numFileName)
	if err != nil {
		panic(err)
	}
	defer fileNum.Close()

	mapVerts, mapVertsArcs := createMapsRelVertsArcs(fileSymb)
	mapUndirGraph := createMapUndirGraph(createMapDirGraph(mapVerts, mapVertsArcs))

	for i, v := range mapUndirGraph {
		str := fmt.Sprintf("%d ", i)
		for i1, v1 := range v {
			if i1%2 != 0 {
				str += fmt.Sprintf(",%d ", v1)
				continue
			}
			str += fmt.Sprintf("%d", v1)
		}
		fileNum.WriteString(str + "\n")
	}
	return mapVerts
}
