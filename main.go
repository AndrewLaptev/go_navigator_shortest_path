package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/RyanCarrier/dijkstra"
)

func main() {
	// fmt.Println("Shortest distance ", best.Distance, " following path ", best.Path)
	fmt.Println(graphFromSymbToNum("graph_symbolic.txt"))
}

func shortestPath(nameGraphFile string, src, dst int) (int64, []int) {
	graph, _ := dijkstra.Import(nameGraphFile)
	best, err := graph.Shortest(src, dst)
	if err != nil {
		log.Fatal(err)
	}
	return best.Distance, best.Path
}

func graphFromSymbToNum(nameSymbFile string /*, nameNumFile string*/) map[string]int {
	var vertexSymbToNum = make(map[string]int)

	fileSymb, err := os.Open(nameSymbFile)
	if err != nil {
		panic(err)
	}
	defer fileSymb.Close()

	scanner := bufio.NewScanner(fileSymb)
	for i := 0; scanner.Scan(); i++ {
		vertexSymbToNum[strings.Split(scanner.Text(), " ")[0]] += i
	}

	return vertexSymbToNum
}
