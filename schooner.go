package schooner

import (
	"sort"
)

type Category int

const (
	ONES            Category = 1
	TWOS            Category = 2
	THREES          Category = 3
	FOURS           Category = 4
	FIVES           Category = 5
	SIXES           Category = 6
	SEVENS          Category = 7
	EIGHTS          Category = 8
	THREE_OF_A_KIND Category = 9
	FOUR_OF_A_KIND  Category = 10
	CHANCE          Category = 11
	FULL_HOUSE      Category = 25
	SMALL_STRAIGHT  Category = 30
	ALL_DIFFERENT   Category = 35
	LARGE_STRAIGHT  Category = 40
	SCHOONER        Category = 50
)

var categories = map[Category]string{
	ONES:            "ONES",
	TWOS:            "TWOS",
	THREES:          "THREES",
	FOURS:           "FOURS",
	FIVES:           "FIVES",
	SIXES:           "SIXES",
	SEVENS:          "SEVENS",
	EIGHTS:          "EIGHTS",
	THREE_OF_A_KIND: "THREE_OF_A_KIND",
	FOUR_OF_A_KIND:  "FOUR_OF_A_KIND",
	CHANCE:          "CHANCE",
	FULL_HOUSE:      "FULL_HOUSE",
	SMALL_STRAIGHT:  "SMALL_STRAIGHT",
	ALL_DIFFERENT:   "ALL_DIFFERENT",
	LARGE_STRAIGHT:  "LARGE_STRAIGHT",
	SCHOONER:        "SCHOONER",
}

// as I got from description - List<Enum> topCategories(List<int> diceRoll) should ignore primitive categories
func isPrimitive(c Category) bool {
	m := map[Category]string{
		ONES:   "",
		TWOS:   "",
		THREES: "",
		FOURS:  "",
		FIVES:  "",
		SIXES:  "",
		SEVENS: "",
		EIGHTS: "",
	}
	// check if category in the set
	_, inSet := m[c]
	return inSet
}

func (c Category) String() string {
	// if category doesn't exist it will return ""
	return categories[c]
}

func rollSum(diceRoll []int) (sum int) {
	for i := range diceRoll {
		sum += diceRoll[i]
	}
	return sum
}

type diceCount map[int]int

func (d diceCount) kind(num int) bool {
	for _, count := range d {
		if count >= num {
			return true
		}
	}
	return false
}

func (d diceCount) isKind3() bool {
	return d.kind(3)
}

func (d diceCount) isKind4() bool {
	return d.kind(4)
}

func (d diceCount) isFullHouse() bool {
	hasThree := false
	hasTwo := false
	for _, count := range d {
		if count == 3 {
			hasThree = true
		}
		if count == 2 {
			hasTwo = true
		}
	}
	return hasThree && hasTwo
}

func (d diceCount) isAllDifferent() bool {
	return len(d) == 5
}

func (d diceCount) isSchooner() bool {
	return len(d) == 1
}

func (d diceCount) isStraight(length int) bool {

	// check that we have enough unique dices
	if !(len(d) >= length) {
		return false
	}

	// get unique values through map
	uniqueValues := make([]int, 0, len(d))
	for key := range d {
		uniqueValues = append(uniqueValues, key)
	}

	sort.Ints(uniqueValues)

	// Check for consecutive sequence of desired length
	consecutive := 1
	for i := 1; i < len(uniqueValues); i++ {
		if uniqueValues[i] == uniqueValues[i-1]+1 {
			consecutive++
			if consecutive >= length {
				return true
			}
		} else {
			consecutive = 1
		}
	}
	return false
}

func (d diceCount) isStraightSmall() bool {
	return d.isStraight(4)
}

func (d diceCount) isStraightLarge() bool {
	return d.isStraight(5)
}

func score(category Category, diceRoll []int) int {

	// create map with key as dice number and value as number of entrances
	dCount := make(diceCount)
	for _, die := range diceRoll {
		dCount[die]++
	}

	switch category {
	case ONES, TWOS, THREES, FOURS, FIVES, SIXES, SEVENS, EIGHTS:
		return dCount[int(category)] * int(category)
	case THREE_OF_A_KIND:
		if dCount.isKind3() {
			return rollSum(diceRoll)
		}
	case FOUR_OF_A_KIND:
		if dCount.isKind4() {
			return rollSum(diceRoll)
		}
	case FULL_HOUSE:
		if dCount.isFullHouse() {
			return int(FULL_HOUSE)
		}
	case SMALL_STRAIGHT:
		if dCount.isStraightSmall() {
			return int(SMALL_STRAIGHT)
		}
	case ALL_DIFFERENT:
		if dCount.isAllDifferent() {
			return int(ALL_DIFFERENT)
		}
	case LARGE_STRAIGHT:
		if dCount.isStraightLarge() {
			return int(LARGE_STRAIGHT)
		}
	case SCHOONER:
		if dCount.isSchooner() {
			return int(SCHOONER)
		}
	case CHANCE:
		return rollSum(diceRoll)
	}
	return 0
}

// Following section required in golang for sorting of non-primitive objects
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
}

// End of sort.Sorter interface definition

func TopCategories(diceRoll []int) []Category {

	var results []categorySort
	for category := range categories {
		// ignore primitive categories
		if isPrimitive(category) {
			continue
		}

		s := score(category, diceRoll)

		if s != 0 {
			results = append(results, categorySort{key: category, value: s})
		}
	}
	sort.Sort(categorySorter(results))

	bestCategories := make([]Category, len(results))
	for i := range results {
		bestCategories[i] = results[i].key
	}
	return bestCategories
}
