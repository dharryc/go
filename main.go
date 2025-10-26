// perl weekly challenge 122

package main

import (
	"fmt"
	m "math/rand"
)

func main() {
	// part 1 setup
	averagesInputs := []int{}
	for range 10 {
		averagesInputs = append(averagesInputs, m.Intn(101))
	}

	// part 2 setup
	basketBallPoints := []int{}
	for range 3 {
		basketBallPoints = append(basketBallPoints, m.Intn(8))
	}

	// part 1 running
	for i, v := range runningAverage(averagesInputs) {
		fmt.Printf("Running average after %d values: %f. Added number from inputs: %d\n", i+1, v, averagesInputs[i])
	}

	//part 2 running
	for _, score := range basketBallPoints {
		fmt.Printf("\nScore: %d\n", score)
		combinations := basketballPointCombinations(score)
		for _, combo := range combinations {
			for _, point := range combo {
				fmt.Printf("%d ", point)
			}
			fmt.Println()
		}
	}
}

/*
You are given a stream of numbers, @N.
Write a script to print the average of the stream at every point.
*/
func runningAverage(input []int) []float64 {
	averages := []float64{}
	average := 0.0
	sum := 0
	for i := range input {
		sum += input[i]
		average = float64(sum) / float64(i+1)
		averages = append(averages, float64(average))
	}
	return averages
}

/*
You are given a score $S.
You can win basketball points e.g. 1 point, 2 points and 3 points.
Write a script to find out the different ways you can score $S.
*/
func basketballPointCombinations(score int) [][]int {
	resultsChan := make(chan [][]int, 3)

	for _, startPoint := range []int{1, 2, 3} {
		go func(start int) {
			localResults := [][]int{}
			current := []int{start}
			findCombinations(score-start, current, &localResults)
			resultsChan <- localResults
		}(startPoint)
	}

	allResults := [][]int{}
	for i := 0; i < 3; i++ {
		results := <-resultsChan
		allResults = append(allResults, results...)
	}

	return allResults
}

func findCombinations(remaining int, current []int, results *[][]int) {
	if remaining == 0 {
		combination := make([]int, len(current))
		copy(combination, current)
		*results = append(*results, combination)
		return
	}

	if remaining < 0 {
		return
	}

	for _, points := range []int{1, 2, 3} {
		current = append(current, points)
		findCombinations(remaining-points, current, results)
		current = current[:len(current)-1]
	}
}
