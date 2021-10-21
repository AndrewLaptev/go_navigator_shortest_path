package main

import (
	"fmt"
	"log"

	"github.com/RyanCarrier/dijkstra"
)

func main() {
	graph, _ := dijkstra.Import("A.txt")
	best, err := graph.Shortest(0, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shortest distance ", best.Distance, " following path ", best.Path)

	best, err = graph.Longest(0, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Longest distance ", best.Distance, " following path ", best.Path)
}
