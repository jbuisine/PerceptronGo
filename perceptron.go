package main

import (
	"math/rand"
	"os"
	"bufio"
	"log"
	"strings"
	"fmt"
	"strconv"
)


/**
	Neurone structure type
 */
type Neurone struct{
	biais float64
	output float64
	w[] float64
}

/**
	Example structure type
 */
type Example struct{
	x []float64
	etiquette int
}

//Walking step
const STEP_WALK float64 = 0.01

func main() {

	filename := os.Args[1]

	examples := load_examples(filename)

	//Initialization of neurone
	n := Neurone{biais: 0.5, output: 0.0, w: []float64{rand.Float64(), rand.Float64()}}

	for i := 0; i < 10; i++{

		nb_error := 0

		for j := 0; j < len(examples); j++ {
			output := compute_neurone(&n, examples[j])

			if output != examples[j].etiquette {
				nb_error++
				update(&n, examples[j], output)
			}
		}
		fmt.Printf("Nombre d'erreurs %d, itération n°%d\n", nb_error, i)

		if nb_error == 0 {
			break;
		}
	}
}


func compute_neurone(n *Neurone, e Example) int{

	sigma := 0.0

	for i := 0; i < len(n.w); i++ {
		sigma += n.w[i] + e.x[i]
	}

	sigma -= n.biais

	if sigma > 0{
		return 1
	}else{
		return -1
	}
}

func update(n *Neurone, e Example, output int){

	n.biais += STEP_WALK * float64(e.etiquette - output) * (-0.5)

	for i := 0; i < len(n.w); i++ {
		n.w[i] += STEP_WALK * float64(e.etiquette - output) * e.x[i]
 	}
}

func load_examples(filename string) []Example {

	var examples []Example

	// open a file
	if file, err := os.Open(filename); err == nil {

		// make sure it gets closed
		defer file.Close()

		// create a new scanner and read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			v := strings.Split(scanner.Text(), " ")

			var values []float64

			for k := 0; k < len(v)-1; k++ {
				val, err := strconv.ParseFloat(v[k], 64)

				if err != nil {
					fmt.Println(err)
				}
				values = append(values, val)
			}

			etiquette, err := strconv.Atoi(v[len(v)-1])

			if err != nil {
				fmt.Println(err)
			}

			examples = append(examples, Example{x:values, etiquette:etiquette})
		}

		// check for errors
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

	} else {
		log.Fatal(err)
	}

	return examples
}
