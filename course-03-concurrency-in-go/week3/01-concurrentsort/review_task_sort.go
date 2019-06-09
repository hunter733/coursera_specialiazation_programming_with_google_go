package main

import (
	"fmt"
	"sort"
	"sync"
)

func doSort(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(arr)
}

type Idx struct {
	begin int
	end   int
}

func prepareRanges(len, goCnt int) []Idx {
	idxs := make([]Idx, goCnt)
	for i := 0; i < goCnt; i++ {
		idxs[i] = Idx{i * len / goCnt, (i + 1) * len / goCnt}
		if i == goCnt-1 {
			idxs[i].end = len
		}
	}
	return idxs
}

func parallelSort(arr []int, ranges []Idx) {
	var wg sync.WaitGroup

	subarrays := make(chan []int)
	go func() {
		for i := range ranges {
			fmt.Printf("subarray %d to be sorted:\n", i+1)
			fmt.Println(<-subarrays)
		}
	}()

	for _, idx := range ranges {
		wg.Add(1)
		go func(subarray []int) {
			defer wg.Done()
			subarrays <- subarray
			sort.Ints(subarray)
		}(arr[idx.begin:idx.end])
	}
	wg.Wait()
}

func mergeRanges(arr []int, idxs []Idx) []int {
	newArr := make([]int, len(arr))
	for i := range newArr {
		minIdx := -1
		for j, idx := range idxs {
			if idx.begin == idx.end {
				continue
			}
			if minIdx < 0 || arr[idx.begin] < arr[idxs[minIdx].begin] {
				minIdx = j
			}
		}
		newArr[i] = arr[idxs[minIdx].begin]
		idxs[minIdx].begin++
	}
	return newArr
}

func readIntArr() []int {
	fmt.Println("Please, enter ints delimited by spaces. Please enter to finish")
	arr := make([]int, 0)
	for {
		var i int
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			fmt.Println(err)
			break
		}
		arr = append(arr, i)
	}
	return arr
}

func main() {
	//arr := []int{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	arr := readIntArr()
	const goCnt = 4
	idxs := prepareRanges(len(arr), goCnt)
	parallelSort(arr, idxs)
	arr = mergeRanges(arr, idxs)
	fmt.Println()
	fmt.Println(arr)
}
