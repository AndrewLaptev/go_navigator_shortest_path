package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/RyanCarrier/dijkstra"
)

const (
	nameSymbFile = "graph_symbolic.txt"
	nameNumFile  = "graph_numeric.txt"
)

func main() {

	// fmt.Println("Shortest distance ", best.Distance, " following path ", best.Path)
	fmt.Println(createMapSymbNum("graph_symbolic.txt"))
}

func shortestPath(nameGraphFile string, src, dst int) (int64, []int) {
	graph, _ := dijkstra.Import(nameGraphFile)
	best, err := graph.Shortest(src, dst)
	if err != nil {
		log.Fatal(err)
	}
	return best.Distance, best.Path
}

func createMapSymbNum(symbFile string) (map[string]int, map[int][]string) {
	var vertexSymbToNum = make(map[string]int)
	var arcsNumToSymb = make(map[int][]string)

	fileSymb, err := os.Open(symbFile)
	if err != nil {
		panic(err)
	}
	defer fileSymb.Close()

	scanner := bufio.NewScanner(fileSymb)
	for i := 0; scanner.Scan(); i++ {
		vertexSymbName := strings.Split(scanner.Text(), " ")[0]
		arcsNumName := strings.Split(scanner.Text(), " ")[1:]
		
		for _, s := range arcsNumName {
			s[0] = 
		}
		vertexSymbToNum[vertexSymbName] += i
		arcsNumToSymb[vertexSymbToNum[vertexSymbName]] = arcsSymbName
	}

	return vertexSymbToNum, arcsNumToSymb
}

func createUndirGraphNumFile(mapVertexSymbToNum map[string]int, mapArcsSymb map[int][]string, numFile string) {
	var str string

	fileNum, err := os.Create(numFile)
	if err != nil {
		panic(err)
	}
	defer fileNum.Close()

	for i, v := range mapArcsSymb {
		str = strconv.Itoa(i) + " "
	}
}
