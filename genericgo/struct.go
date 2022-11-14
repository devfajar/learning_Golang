package main

import (
	"fmt"
)

type userModel[T int | float64] struct {
	Name string
	Scores []T
}

func (m *userModel[int]) SetScoresA(scores []int) {
	m.Scores = scores
}

func (m *userModel[float64]) SetScoresB(scores []float64) {
	m.Scores = scores
}

func main() {
	var m1 userModel[int]
	m1.Name = "Fajar"
	m1.Scores = []int{1, 2, 3}
	fmt.Println("Scores: ", m1.Scores)

	var m2 userModel[float64]
	m2.Name = "Fajar"
	m2.SetScoresB([]float64{10, 11})
	fmt.Println("Scores: ", m2.Scores)
}