package genetic_algorithm

import (
	"fmt"
	"math"
	"math/rand"
)

func New(mutationRate float64, size int, target []rune, alphabet []rune) *Population {
	population := Population{
		alphabet:      alphabet,
		numGeneration: 1,
		mutationRate:  mutationRate,
		size:          size,
		target:        target,
	}
	for i := 0; i < population.size; i++ {
		creature := Creature{
			genes:   make([]rune, len(target)),
			fitness: 0,
		}
		for j := 0; j < len(target); j++ {
			creature.genes[j] = population.newCharacter()
		}
		population.creatures = append(population.creatures, creature)
	}
	return &population
}

type Population struct {
	numGeneration int
	creatures     Creatures
	mutationRate  float64
	size          int
	target        []rune
	alphabet      []rune
	maxFitness    float64
}
type Creature struct {
	genes   []rune
	fitness float64
}

type Creatures []Creature

func (p Population) Mutate() {
	for idx, creature := range p.creatures {
		for i := 0; i < len(creature.genes); i++ {
			r := rand.Float64()
			if r < p.mutationRate {
				creature.genes[i] = p.newCharacter()
			}
		}
		p.creatures[idx] = creature
	}
}

func (p Population) newCharacter() rune {
	return p.alphabet[rand.Intn(len(p.alphabet))]
}

func (p *Population) CreateNextGeneration() {
	creatures := make(Creatures, len(p.creatures))
	copy(creatures, p.creatures)
	for i := 0; i < len(creatures); i++ {
		midPoint := rand.Intn(len(p.target))
		parentA := creatures.SelectParent(p.maxFitness)
		parentB := creatures.SelectParent(p.maxFitness)
		dataA := append([]rune{}, parentA.genes[:midPoint]...)
		dataB := append([]rune{}, parentB.genes[midPoint:]...)
		childCreature := Creature{
			genes:   append(dataA, dataB...),
			fitness: 0,
		}
		p.creatures[i] = childCreature
	}
	p.numGeneration++
}

func (p *Population) CalculateFitness() {
	p.maxFitness = 0
	for idx, creature := range p.creatures {
		creature.fitness = 0
		for i, c := range p.target {
			if creature.genes[i] == c {
				creature.fitness++
			}
		}
		if creature.fitness > p.maxFitness {
			p.maxFitness = creature.fitness
		}
		p.creatures[idx] = creature
	}
}

func (p Population) IsFit() bool {
	for _, creature := range p.creatures {
		if creature.fitness == float64(len(p.target)) {
			return true
		}
	}
	return false
}

func (p Population) DebugPrint() {
	var avgFitness float64
	var totalFitness float64
	var maxFitness float64
	var minFitness = math.MaxFloat64
	var fittestCreature Creature
	var fittestCreatureIndex int

	for creatureIndex, creature := range p.creatures {
		totalFitness += creature.fitness
		if maxFitness < creature.fitness {
			maxFitness = creature.fitness
			fittestCreature = creature
			fittestCreatureIndex = creatureIndex
		}
		if minFitness > creature.fitness {
			minFitness = creature.fitness
		}
	}
	avgFitness = totalFitness / float64(len(p.creatures))
	fmt.Println()
	fmt.Printf("population      : %v\n", len(p.creatures))
	fmt.Printf("mutation        : %v%%\n", 100*p.mutationRate)
	fmt.Printf("fitness-goal    : %v\n", len(p.target))
	fmt.Printf("generation      : %v\n", p.numGeneration)
	fmt.Printf("max fitness     : %v\n", maxFitness)
	fmt.Printf("avg fitness     : %v\n", avgFitness)
	fmt.Printf("min fitness     : %v\n", minFitness)
	fmt.Printf("fittest creature: (idx: %v)\n", fittestCreatureIndex)
	fmt.Printf("\tgenes  : %v\n", string(fittestCreature.genes))
	fmt.Printf("\tfitness: %v of max %v\n", fittestCreature.fitness, len(p.target))
}

func (c Creatures) SelectParent(selectionThreshold float64) Creature {
	var iterations int
	for {
		iterations++
		parent := c[rand.Intn(len(c))]
		fitnessThreshold := rand.Float64() * selectionThreshold
		if parent.fitness >= fitnessThreshold {
			return parent
		}
		if iterations > 10000 {
			panic("out of bounds selecting a parent")
		}
	}
}
