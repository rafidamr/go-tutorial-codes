package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Job struct {
	weight      int
	length      int
	score_diff  int
	score_ratio float64
}

func Schedule() {
	jobs := buildJobSequence("course3/week1_greedy_schedule_mst/jobs.txt")
	sorted_ratio := SortByRatio(jobs)
	sorted_diff := SortByDiff(jobs)
	tDiff := SumWeightedCompletion(sorted_diff)
	tRatio := SumWeightedCompletion(sorted_ratio)
	fmt.Println(tDiff, tRatio)
}

// sort in decreasing order
func SortByDiff(jobs *[]Job) *[]Job {
	sorted := slices.Clone(*jobs)
	slices.SortFunc(sorted, func(j1, j2 Job) int {
		if j1.score_diff == j2.score_diff {
			return cmp.Compare(j2.weight, j1.weight)
		}
		return cmp.Compare(j2.score_diff, j1.score_diff)
	})
	return &sorted
}

// sort in decreasing order
func SortByRatio(jobs *[]Job) *[]Job {
	sorted := slices.Clone(*jobs)
	slices.SortFunc(sorted, func(j1, j2 Job) int {
		return cmp.Compare(j2.score_ratio, j1.score_ratio)
	})
	return &sorted
}

func SumWeightedCompletion(jobs *[]Job) int {
	sum := 0
	totalLen := 0
	for _, j := range *jobs {
		totalLen += j.length
		sum += j.weight * totalLen
	}
	return sum
}

func buildJobSequence(filepath string) *[]Job {
	var jobs []Job
	ptr, _ := os.Open(filepath)
	defer ptr.Close()

	scanner := bufio.NewScanner(ptr)
	for scanner.Scan() {
		scanner.Text() // skip the first line
		break
	}
	for scanner.Scan() {
		line := scanner.Text()
		tuple := strings.Split(line, " ")
		w, _ := strconv.Atoi(strings.TrimSpace(tuple[0]))
		l, _ := strconv.Atoi(strings.TrimSpace(tuple[1]))
		j := Job{
			weight:      w,
			length:      l,
			score_diff:  w - l,
			score_ratio: float64(w) / float64(l),
		}
		jobs = append(jobs, j)
	}
	return &jobs
}
