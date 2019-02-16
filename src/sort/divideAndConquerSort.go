package main

import "fmt"

func DivideAndConquerSort(source []int) {
	var (
		low  = 0
		high = len(source) - 1
	)
	divideAndConquerSort(source, low, high)
}

func divideAndConquerSort(source []int, low, high int) {
	if low == high {
		return
	}
	// divide
	mid := (high-low)/2 + low
	divideAndConquerSort(source, low, mid)
	divideAndConquerSort(source, mid+1, high)

	// conquer
	conquer(source, low, mid, high)
}

func conquer(source []int, low, mid, high int) {
	var size = high - low + 1
	var (
		index = 0
		lhs   = low
		rhs   = mid + 1
	)
	temp := make([]int, size)
	for lhs <= mid && rhs <= high {
		if source[lhs] <= source[rhs] {
			//temp[index] = source[lhs]
			//index++
			//lhs++
			_add(temp, &lhs, source, &index)
		} else {
			//temp[index] = source[rhs]
			//index++
			//rhs++
			_add(temp, &rhs, source, &index)
		}
	}
	for lhs <= mid {
		//temp[index] = source[lhs]
		//index++
		//lhs++
		_add(temp, &lhs, source, &index)
	}

	for rhs <= high {
		//temp[index] = source[rhs]
		//index++
		//rhs++
		_add(temp, &rhs, source, &index)
	}
	copy(source[low:high+1], temp)
}

func _add(sl []int, idx *int, source []int, index *int) {
	sl[*index] = source[*idx]
	*index++
	*idx++
}

func main() {
	var source = []int{1, 4, 3, 2}
	DivideAndConquerSort(source)
	fmt.Println(source)
}
