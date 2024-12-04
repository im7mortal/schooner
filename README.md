# Schooner 

## Run

The fastest way to run it. Is use [Dockerfile](Dockerfile). 
```shell
 docker build -t schooner:latest . && \
 docker run schooner:latest
```

Or run [run.sh](run.sh)


## Code

I was trying to omit any Go-specific logic and make it almost like pseudocode.
Here are some comments about Go.

1. Map definition and population: The following code creates a map entry if it doesn't exist and immediately increments it.

``` go
	dCount := make(diceCount)
	for _, die := range diceRoll {
		dCount[die]++
	}
```

2. Golang-specific sorting construction: We need to define a type and implement the sort.Sorter methods. Once implemented, we can pass an array of this type to sort.Sort, and it will be recognized automatically.

```
type categorySort struct {
	key   Category
	value int
}

type categorySorter []categorySort

func (cs categorySorter) Len() int {
	return len(cs)
}

func (cs categorySorter) Less(i, j int) bool {
	return cs[i].value > cs[j].value
}

func (cs categorySorter) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}```