## go-walker
![CircleCI](https://circleci.com/gh/mathetake/go-walker.svg?style=shield&circle-token=9099b55c8773ac52035f39cec9e88e42930945c4)  [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

go-walker is a tiny package which implements Walker's alias sampling method (https://en.wikipedia.org/wiki/Alias_method) for sampling from an arbitrary discrete probability distribution with complexity O(1).

## usage

```golang
import (
	"github.com/mathetake/go-walker"
	"fmt"
)

func main() {
	ws := []float64{0.3, 0.7}
	s := walker.GetSampler(ws)

	for i := 0; i < 10; i++ {
		fmt.Println(s.Generate())
	}
}
```


## references

[1] http://www.keithschwarz.com/darts-dice-coins

[2] https://en.wikipedia.org/wiki/Alias_method

## license

MIT
