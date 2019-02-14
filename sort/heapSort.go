package main

/*
 //小堆


**/

func heapDown(source []int, i int, comparator2 comparator) {
	var (
		min int
		max int
		reserve int
	)
	min = 2 * i + 1
	max = 2 * i + 2
	reserve = source[i]
	for i * 2 + 2 <= len(source) {
		if i * 2 + 2 < len(source) {
			if !comparator2(source[min], source[max]) {
				min = max
			}
		}
		if (comparator2(reserve, source[min])) {
			break
		}
		source[i] = source[min]
		i = min
	}
	source[i] = reserve
}

func heapUp(source []int, i int, comparator2 comparator) {
	
}


type comparator func (int, int) bool

func Bigger(lhs, rhs int) bool{
	if lhs > rhs {
		return true;
	} else {
		return false;
	}
}

func Lesser(lhs, rhs int) bool {
	if lhs < rhs {
		return true
	} else {
		return false
	}
}