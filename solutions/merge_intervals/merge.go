package mergeintervals

import "sort"

// Interval definition for an interval.
type Interval struct {
	Start int
	End   int
}

type IntervalSlice []Interval

// Len is the number of elements in the collection.
func (l IntervalSlice) Len() int {
	return len(l)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (l IntervalSlice) Less(i, j int) bool {
	return l[i].Start < l[j].Start
}

// Swap swaps the elements with indexes i and j.
func (l IntervalSlice) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{}
	}

	sort.Sort(IntervalSlice(intervals))

	res := make([]Interval, 1, len(intervals))
	res[0] = intervals[0]
	for i, j := 1, 0; i < len(intervals); i++ {
		if res[j].End < intervals[i].Start {
			res = append(res, intervals[i])
			j++
		} else {
			res[j].End = max(res[j].End, intervals[i].End)
		}
	}
	return res
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
