package main

import (
	"fmt"
	"mycomparator"
)

/*
*
*  v0 升序
 */

func QuickSortAsc(source []int) {
	quickSort(source, Asc)
}

func QuickSortDesc(source []int) {
	quickSort(source, Desc)
}

func quickSort(source []int, comparator QuickComparator) {
	var low = 0
	var high = len(source) - 1
	quickSortimp(source, low, high, comparator)
}

func quickSortimp(source []int, low, high int, comparator QuickComparator) {
	if low >= high {
		return
	}
	var q = paritition(source, low, high, comparator)
	quickSortimp(source, low, q-1, comparator)
	quickSortimp(source, q+1, high, comparator)
}

func paritition(source []int, low, high int, comparator QuickComparator) int {
	if low >= high {
		return low
	}
	// desIndex标记最后的未知，每有一个大于加1
	//  其二标记第一个大于swapNumber的数， 然后交换到后面去
	var desIndex = low
	// swapNumber标记交换的
	var swapNumber = source[low]
	for i := low + 1; i <= high; i++ {
		if comparator(swapNumber, source[i]) {
			desIndex++
			swap(source, i, desIndex)
		}
	}

	swap(source, low, desIndex)
	return desIndex

}

func swap(source []int, lhs, rhs int) {
	if lhs != rhs {
		var temp int
		temp = source[lhs]
		source[lhs] = source[rhs]
		source[rhs] = temp
	}
}

type QuickComparator func(int, int) bool

//func bigger(lhs, rhs int) bool {
//	if lhs > rhs {
//		return true
//	} else {
//		return false
//	}
//}

func Asc(lhs, rhs int) bool {
	return mycomparator.Bigger(lhs, rhs)
}

func Desc(lhs, rhs int) bool {
	return mycomparator.Bigger(rhs, lhs)
}

func main() {
	source := []int{4, 1, 3, 5, 2, 7, 8, 9, 11, 10}
	QuickSortDesc(source)
	fmt.Println(source)
}
