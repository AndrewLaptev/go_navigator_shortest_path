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
)

func main() {
	fileSymb, err := os.Open(nameSymbFile)
	if err != nil {
		panic(err)
	}
	defer fileSymb.Close()

	// fmt.Println("Shortest distance ", best.Distance, " following path ", best.Path)
	fmt.Println(createMapVertsSymbNum(fileSymb))
}

func shortestPath(nameGraphFile string, src, dst int) (int64, []int) {
	graph, _ := dijkstra.Import(nameGraphFile)
	best, err := graph.Shortest(src, dst)
	if err != nil {
		log.Fatal(err)
	}
	return best.Distance, best.Path
}

// Create map of graph vertexs from symbolic file to numeric list
// должна вернуть все таки две мапы: мапа вершин назв-номер и мапу ребер вершин номер верш-строковый массив ребер
func createMapVertsSymbNum(symbFileName *os.File) map[string]int {
	var mapVertsSymbNum = make(map[string]int)

	scanner := bufio.NewScanner(symbFileName)
	for i := 0; scanner.Scan(); i++ {
		vertSymb := strings.Split(scanner.Text(), " ")[0]
		mapVertsSymbNum[vertSymb] += i
	}

	return mapVertsSymbNum
}

func createMapDirGraph(mapVerts map[string]int, mapVertsArcs map[int][]string) map[int][]int {

}

// Create numeric map vertexs and arcs of graph from symbolic file.
// For undir graph arcs are duplicated
func createMapNumUndirGraph(symbFileName *os.File, mapVertsSymbNum map[string]int) map[int]int {
	var mapNumUndirGraph = make(map[int]int)
	return mapNumUndirGraph
}

func createUndirGraphNumFile(mapVertsSymbNum map[string]int, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	mapVertsArcs := make(map[int][]string)
	for i, v := range mapVertsSymbNum {
		mapVertsArcs[v] =
	}
}
