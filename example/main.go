package main

import (
	"fmt"
	"math/rand"

	"github.com/gravitamp/kmeans"
)

var d kmeans.Observations
var count []int

func main() {
	//setup data
	setupData("Mall_Customers.csv")
	// Partition the data points into 5 clusters

	// km := kmeans.New()
	// clusters, _ := km.Partition(d, 5, 5)
	rand.Seed(5)
	km, _ := kmeans.NewWithOptions(0.01, kmeans.SimplePlotter{})
	clusters, _ := km.Partition(d, 5, 5)
	for i, c := range clusters {
		count = append(count, len(c.Observations))
		fmt.Println(len(c.Observations))
		fmt.Printf("Centered at x: %.2f y: %.2f\n", c.Center[0], c.Center[1])
		fmt.Printf("Matching data points: %+v\n", c.Observations)
		fmt.Printf("total %d: %d\n", i, len(c.Observations))
	}
}
