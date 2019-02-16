package main

import (
	"fmt"
	"mycomparator"
)

/*
 //小堆


**/

func HeapSort(source []int, comparator2 comparator) {
	heapSort(source, comparator2)
	var result []int
	for 1 < len(source) {
		result = append(result, source[0])
		source[0] = source[len(source)-1]
		source = source[:len(source)-1]
		heapDown(source, 0, comparator2)
	}
	result = append(result, source[0])
	fmt.Println(result)
}

func heapSort(source []int, comparator2 comparator) {
	for i := len(source)/2 - 1; i >= 0; i-- {
		heapDown(source, i, comparator2)
	}
}

func heapDown(source []int, i int, comparator2 comparator) {
	var (
		min     int
		max     int
		reserve int
	)

	reserve = source[i]
	for i*2+2 <= len(source) {
		min = 2*i + 1
		max = 2*i + 2
		if i*2+2 < len(source) {
			if comparator2(source[max], source[min]) {
				min = max
			}
		}
		if comparator2(reserve, source[min]) {
			break
		}
		source[i] = source[min]
		i = min
	}
	source[i] = reserve
}

func heapUp(source []int, i int, comparator2 comparator) {
	var (
	//value int
	)

}

type comparator func(int, int) bool

func main() {
	var source = []int{1, 4, 2, 3}
	HeapSort(source, mycomparator.Lesser)
	//fmt.Println(source)
}
