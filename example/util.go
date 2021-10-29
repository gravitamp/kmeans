package main

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"

	"github.com/gravitamp/kmeans"
)

func setupData(file string) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	csvReader := csv.NewReader(f)
	csvData, _ := csvReader.ReadAll()

	//read without header
	for i := 1; i < len(csvData); i++ {
		// val, _ := strconv.ParseFloat(csvData[i][1], 64)
		lat, _ := strconv.ParseFloat(csvData[i][2], 64)
		lng, _ := strconv.ParseFloat(csvData[i][3], 64)
		lon, _ := strconv.ParseFloat(csvData[i][4], 64)
		d = append(d, kmeans.Coordinates{
			// val,
			lng,
			lon,
			lat,
			//add random number, because why not?
			float64(rand.Intn(30-10+1) + 10),
		})

	}
}
