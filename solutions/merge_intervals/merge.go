package mergeintervals

import "sort"

// Interval definition for an interval.
type Interval struct {
	Start int
	End   int
}

// IntervalSlice slice of Interval
type IntervalSlice []Interval

// Len is the number of elements in the collection.
func (s IntervalSlice) Len() int {
	return len(s)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s IntervalSlice) Less(i, j int) bool {
	return s[i].Start < s[j].Start
}

// Swap swaps the elements with indexes i and j.
func (s IntervalSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{}
	}

	sort.Sort(IntervalSlice(intervals))

	// a another sort scheme, faster in benchmark
	//for keep := true; keep; {
	//	keep = false
	//	for i := 0; i < len(intervals)-1; i++ {
	//		if intervals[i].Start > intervals[i+1].Start {
	//			intervals[i], intervals[i+1] = intervals[i+1], intervals[i]
	//			keep = true
	//		}
	//	}
	//}

	res := make([]Interval, 0, len(intervals))
	current := intervals[0]
	for _, next := range intervals[1:] {
		if next.Start <= current.End {
			current.End = max(current.End, next.End)
		} else {
			res = append(res, current)
			current = next
		}
	}
	return append(res, current)
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
