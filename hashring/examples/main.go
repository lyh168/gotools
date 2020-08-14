package main

import (
	"log"

	"github.com/yacen/gotools/hashring"
	"github.com/yacen/gotools/randomUtils"
)

func main() {
	weights := make(map[string]int)
	weights["192.168.0.246:11212"] = 1
	weights["192.168.0.247:11212"] = 2
	weights["192.168.0.249:11212"] = 1

	ring := hashring.NewWithWeights(weights)
	for i := 0; i < 100; i++ {
		server, ok := ring.GetNode(randomUtils.NumericSequence(10))
		log.Println(server, ok)
	}
}
