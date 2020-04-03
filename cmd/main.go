package main

import (
	"math/rand"
	"sandbox/genetic_algorithm"
	"time"
)

type PopulationInterface interface {
	CalculateFitness()
	DebugPrint()
	CreateNextGeneration()
	Mutate()
	IsFit() bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	population := genetic_algorithm.New(
		0.001,
		1000,
		[]rune("To be, or not to be: that is the question."),
		[]rune{
			'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
			' ', ',', ':', '.',
		})
	run(population)
}

func run(population PopulationInterface) {
	timeStart := time.Now()
	for {
		population.CalculateFitness()
		if population.IsFit() {
			population.DebugPrint()
			break
		}
		if time.Since(timeStart) > time.Second {
			population.DebugPrint()
			timeStart = time.Now()
		}
		population.CreateNextGeneration()
		population.Mutate()

	}
}
