package main

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
func doesIntervalOverlap(a []int,  b []int) bool{
	return min(a[1], b[1]) - max(a[0], b[0]) >= 0
}

func mergeIntervals(a []int, b []int) []int{
	return []int{min(a[0], b[0]), max(a[1], b[1])};
}

func insertInterval(intervals [][]int, newInterval []int) {
	var isIntervalInserted bool= false;
	for i := 0; i < len(intervals); i++ {
		if (newInterval[0] < intervals[i][0]) {
			// Found the position, insert the interval and break from the loop.

			intervals = append(intervals, 0)
			copy(intervals[i+1:], intervals[i:])
			intervals[i] = newInterval
			// intervals.insert(intervals.begin() + i, newInterval);
			isIntervalInserted = true;
			break;
		}
	}
	
	// If there is no interval with a greater value of start value,
	// then the interval must be inserted at the end of the list.
	if (!isIntervalInserted) {
		intervals = append(intervals, newInterval)
	}
}

func insert(intervals [][]int,  newInterval []int) []int{
	// Insert the interval first before merge processing.
	insertInterval(intervals, newInterval);
	
	var answer [][]int;
	for i := 0; i < len(intervals); i++ {
		 currInterval := []int{intervals[i][0], intervals[i][1]};
		// Merge until the list gets exhausted or no overlap is found.
		for i < len(intervals) && doesIntervalOverlap(currInterval, intervals[i]) {
			currInterval = mergeIntervals(currInterval, intervals[i]);
			i++;
		}
		// Decrement to ensure we don't skip the interval due to outer for-loop incrementing.
		i--;
		answer= append(answer, currInterval);
	}
	
	return answer;
}