package schooner

import "testing"

func TestScore(t *testing.T) {
	tests := []struct {
		category Category
		diceRoll []int
		expected int
	}{
		{ONES, []int{1, 2, 1, 4, 5}, 2},
		{TWOS, []int{2, 2, 1, 4, 5}, 4},
		{THREES, []int{3, 3, 3, 4, 5}, 9},
		{FOURS, []int{4, 4, 4, 4, 4}, 20},
		{FIVES, []int{5, 1, 2, 3, 6}, 5},
		{SIXES, []int{6, 6, 6, 6, 6}, 30},
		{SEVENS, []int{7, 7, 7, 7, 7}, 35},
		{EIGHTS, []int{8, 8, 8, 8, 8}, 40},
		{THREE_OF_A_KIND, []int{3, 3, 3, 4, 5}, 18},
		{FOUR_OF_A_KIND, []int{4, 4, 4, 4, 5}, 21},
		{FULL_HOUSE, []int{3, 3, 3, 5, 5}, 25},
		{SMALL_STRAIGHT, []int{1, 2, 3, 4, 6}, 30},
		{ALL_DIFFERENT, []int{1, 2, 3, 4, 5}, 35},
		{LARGE_STRAIGHT, []int{2, 3, 4, 5, 6}, 40},
		{SCHOONER, []int{5, 5, 5, 5, 5}, 50},
		{CHANCE, []int{2, 3, 4, 5, 6}, 20},
		//// Invalid case
		{THREE_OF_A_KIND, []int{1, 2, 3, 4, 5}, 0},
	}

	for _, test := range tests {
		t.Run(test.category.String(), func(t *testing.T) {
			result := score(test.category, test.diceRoll)
			if result != test.expected {
				t.Errorf("Expected %d, got %d for %v in %v", test.expected, result, test.category, test.diceRoll)
			}
		})
	}
}

func TestTopCategories(t *testing.T) {
	tests := []struct {
		diceRoll []int
		expected []Category
	}{
		{[]int{3, 3, 3, 6, 7}, []Category{THREE_OF_A_KIND, CHANCE}},
		{[]int{1, 2, 3, 4, 5}, []Category{LARGE_STRAIGHT, ALL_DIFFERENT, SMALL_STRAIGHT, CHANCE}},
		{[]int{8, 8, 8, 8, 8}, []Category{SCHOONER, CHANCE, THREE_OF_A_KIND, FOUR_OF_A_KIND}},
		{[]int{6, 6, 6, 8, 8}, []Category{THREE_OF_A_KIND, CHANCE, FULL_HOUSE}},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := TopCategories(test.diceRoll)
			if len(result) != len(test.expected) {
				t.Fatalf("Expected %v, got %v", test.expected, result)
			}
			for i, category := range result {
				if category != test.expected[i] {
					if !(sumCategory(category) && sumCategory(test.expected[i])) {
						t.Errorf("Mismatch at index %d: expected %v, got %v", i, test.expected[i], category)
					}
				}
			}
		})
	}
}

// CHANCE, THREE_OF_A_KIND, FOUR_OF_A_KIND have the same value - sum of all dices; but they can have randomly sorted
// in the same time it doesn't matter in the tests
// so we just ensure that these 3 will not fail it if test case has unexpected order
// I believe more edge cases can cause it. but for now I will skip it
func sumCategory(c Category) bool {
	m := map[Category]string{
		CHANCE:          "",
		THREE_OF_A_KIND: "",
		FOUR_OF_A_KIND:  "",
	}
	// check if category in the set
	_, inSet := m[c]
	return inSet
}
