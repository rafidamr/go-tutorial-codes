package main

import "fmt"

func main() {
	jobs := buildJobSequence("course3/week1_greedy_schedule_mst/jobs.txt")
	sorted_ratio := SortByRatio(jobs)
	sorted_diff := SortByDiff(jobs)
	tDiff := SumWeightedCompletion(sorted_diff)
	tRatio := SumWeightedCompletion(sorted_ratio)
	fmt.Println(tDiff, tRatio)
}
