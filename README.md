## go-walker

go-walker is a tiny package which implements Walker's sampling method (https://en.wikipedia.org/wiki/Alias_method) for sampling from a discrete probability with complexity O(1).

## usage

```golang
import (
    "github.com/mathetake/go-walker"
	"fmt"
)

func main() {
	ws := []float64{0.3, 0.7}
	s := walker.GetSampler(ws, false)

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