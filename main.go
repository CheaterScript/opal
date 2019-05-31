package main

import (
	"fmt"
	"sort"
	"sync"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

type Window struct {
	start int
	end   int
}

func findMedianSortedArrays(nums1 []int, nums2 []int) {
	nums := append(nums1, nums2...)
	sort.Sort(sort.IntSlice(nums))
	fmt.Println("fuck you.")
	wg.Done()
}

func Max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	wg.Add(1)
	go findMedianSortedArrays([]int{1, 2}, []int{3, 5})
	wg.Wait()
}
