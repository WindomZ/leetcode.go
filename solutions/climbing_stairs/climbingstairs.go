package climbingstairs

func climbStairs(n int) int {
	// can reach (i) step in one of the two ways:
	// 1. taking a single step from (i-1) step.
	// 2. taking a step of 2 from (i-2) step.
	// so
	// 'current' is (i-1) -> (i) step
	// 'next' is (i) -> (i+1) step
	// reach (i+1) step: (i) step + (i-1) step
	current, next := 1, 1
	for ; n > 0; n-- {
		current, next = next, current+next
	}
	return current
}
