import (
	"slices"
)

// returns time for eating all the bananas and the 'k' associated w/ it.
func getSum(piles []int, k int) int {
	total:=0

	for _,v := range piles {
		// mocking the ceil func
		total += ( (v + k - 1) / k)
	}

	return total
}

// now only to optimize to trim down the search space - ofc now's the time for Binary Seach
// to shine!
func minEatingSpeed(piles []int, h int) int {
	// alright, confession time. I have zero ideas as to how to solve this one.
	// just gonna go w/ the hints now.

	// first starting w/ the upper bound: forgetting about the `min` constraint.

	maxPile := slices.Max(piles)
	// so now it takes just one hour to eat from all the piles. right?
	// k becomes len(nums) -> i know, i know. can't be the answer since a minimum can exist.

	// so let's start from... 1? to the maxPile.
	for i:=1; i<=maxPile; i++ {
		timeToEatAllBananas := getSum(piles, i)

		if timeToEatAllBananas > h {
			continue
		}

		return i
	}

	return -1
}
