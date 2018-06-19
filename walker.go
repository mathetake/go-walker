package walker

import (
	"math/rand"
)

type Sampler struct {
	prob  map[int]float64
	alias map[int]int
	size  int
	seed  float64
}

func (s *Sampler) Generate() int {
	idx := rand.Intn(s.size)
	x := rand.Float64()
	if x < s.prob[idx] {
		return idx
	} else {
		return s.alias[idx]
	}
}

// GetSampler returns a sampler with Walker's sampling method.
// The following is implemented according to `A Practical Version of Vose's Algorithm` in
// http://www.keithschwarz.com/darts-dice-coins
func GetSampler(ws []float64) *Sampler {
	// initialize a sampler
	s := &Sampler{}
	s.size = len(ws)
	var size64 = float64(s.size)
	s.prob = make(map[int]float64, s.size)
	s.alias = make(map[int]int, s.size)

	// create two work lists
	var small, large []int

	// scale all of weights by `size`
	for i := range ws {
		ws[i] = ws[i] * size64
	}

	for i, p := range ws {
		if p < 1 {
			small = append(small, i)
		} else {
			large = append(large, i)
		}
	}

	var l, g int
	for len(small) > 0 {
		small, l = pop(small)
		large, g = pop(large)

		s.prob[l] = ws[l]
		s.alias[l] = g
		ws[g] = (ws[g] + ws[l]) - 1

		if ws[g] < 1 {
			small = append(small, g)
		} else {
			large = append(large, g)
		}
	}

	for len(large) > 0 {
		large, g = pop(large)
		s.prob[g] = 1
	}
	return s
}

func pop(s []int) ([]int, int) {
	return s[:len(s)-1], s[len(s)-1]
}
