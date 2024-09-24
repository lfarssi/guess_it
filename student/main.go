package main

import (
	"bufio"
	"fmt"
	"guess-it/student/functions"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var result []float64
	for scanner.Scan() {
		data := scanner.Text()
		nb, err := strconv.ParseFloat(data, 64)
		if err != nil {
			fmt.Println("Invalid input!!!!")
			continue
		}
		result = append(result, nb)
		if len(result) > 1 {
			nb1, nb2 := functions.Interval(result)
			fmt.Printf("%d %d\n", nb1, nb2)
		}
	}
}
