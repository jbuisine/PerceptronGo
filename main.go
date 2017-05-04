package main

import (
	"fmt"
	"os"
	"math/rand"
	"strconv"
	"bufio"
)

type point struct {
	x float64
	y float64
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	nb, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	var points []point


	for i := 0; i < nb; i++  {
		points = append(points, point{x: rand.Float64(), y: rand.Float64()})
	}
	check(err)

	var lines []string
	for i := 0; i < nb; i++ {
		res := points[i].x + points[i].y - 1

		var tag int

		if(res > 0){
			tag = 1
		}else{
			tag = -1
		}

		line := fmt.Sprintf("%f %f %d", points[i].x, points[i].y, tag)

		lines = append(lines, line)
	}

	writeLines(lines, "points.txt")
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}