package walker

import (
	"math/rand"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

const (
	tol = 10e-4
	num = 10e6
)

func TestGenerate(t *testing.T) {
	tc := [][]float64{
		{0.5, 0.5},
		{0.2, 0.2, 0.2, 0.2, 0.2},
		{0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1},
		{0.15, 0.3, 0.05, 0.4, 0.1},
	}

	for i, c := range tc {
		// intentionally set seeds
		rand.Seed(time.Now().UnixNano())
		eRatio := make(map[int]float64, len(c))

		for j := range c {
			eRatio[j] = c[j]
		}

		s := GetSampler(c)
		ratio := make(map[int]float64, len(c))
		for j := 0; j < num; j++ {
			x := s.Generate()
			ratio[x]++
		}

		for j := range ratio {
			ratio[j] /= num
		}

		for j := range ratio {
			diff := ratio[j] - eRatio[j]
			if diff < 0 {
				diff *= -1
			}
			if diff > tol {
				t.Fatalf("%v-th test failed with \n ratio:%v \nexpected: %v", i, ratio, eRatio)
			}
		}
	}
}

// check if GetSampler doesn't cause panic
func TestGetSampler(t *testing.T) {
	tc := [][]float64{
		{0.5, 0.5},
		{0.2, 0.2, 0.2, 0.2, 0.2},
		{0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1},
		{0.1, 0.3, 0.2, 0.4, 0.1},
	}

	for _, c := range tc {
		GetSampler(c)
	}
}

func TestPop(t *testing.T) {
	tc := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 10},
		{1, 1, 1, 0},
		{1},
	}

	for _, c := range tc {
		as, p := pop(c)
		assert.Equal(t, len(c)-1, len(as))
		assert.Equal(t, c[len(c)-1], p)
	}
}
