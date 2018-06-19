package walker

import (
	"math/rand"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

const (
	tol = 10e-3
	num = 10e5
)

func TestGenerate(t *testing.T) {
	tc := []struct {
		ws        []float64
		normalize bool
	}{
		{
			[]float64{0.5, 0.5},
			false,
		},
		{
			[]float64{0.2, 0.2, 0.2, 0.2, 0.2},
			false,
		},
		{
			[]float64{0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2},
			true,
		},
		{
			[]float64{0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.2, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5},
			true,
		},
		{
			[]float64{0.5, 1.5, 100.3},
			true,
		},
		{
			[]float64{23424124.23, 0.5},
			true,
		},
	}

	for i, c := range tc {
		// intentionally initialize seeds
		rand.Seed(time.Now().UnixNano())
		eRatio := make(map[int]float64, len(c.ws))

		var z float64 = 1
		if c.normalize {
			z = 0
			for j := range c.ws {
				z += c.ws[j]
			}
		}

		for j := range c.ws {
			eRatio[j] = c.ws[j] / z
		}

		s := GetSampler(c.ws, c.normalize)
		ratio := make(map[int]float64, len(c.ws))
		for j := 0; j < num; j++ {
			x := s.Generate()
			ratio[x]++
		}

		for j := range ratio {
			ratio[j] /= num
		}

		for j := range ratio {
			if ratio[j]-eRatio[j] > tol {
				t.Fatalf("%v-th test failed with \n ratio:%v \nexpected: %v", i, ratio, eRatio)
			}
		}
	}
}

// check if GetSampler doesn't cause panic
func TestGetSampler(t *testing.T) {
	tc := []struct {
		ws        []float64
		normalize bool
	}{
		{
			[]float64{0.5, 0.5},
			false,
		},
		{
			[]float64{0.2, 0.2, 0.2, 0.2, 0.2},
			false,
		},
		{
			[]float64{0.5, 1.5, 100.3},
			true,
		},
		{
			[]float64{23424124.23, 0.5},
			true,
		},
	}

	for _, c := range tc {
		GetSampler(c.ws, c.normalize)
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
