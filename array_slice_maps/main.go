package arrayslicemaps

// Link: https://www.hackerrank.com/challenges/array-left-rotation/problem
func rotateLeft(d int32, arr []int32) []int32 {
	for i := 1; int32(i) <= d; i++ {
		slice := arr[1:]
		arr = append(slice, arr[0])
	}

	return arr
}

// This is the first version
// Link: https://www.hackerrank.com/challenges/arrays-ds/problem
func reverseArray1(a []int32) []int32 {
	reversedSlice := make([]int32, 0, cap(a))
	for i := len(a) - 1; i >= 0; i-- {
		reversedSlice = append(reversedSlice, a[i])
	}
	return reversedSlice
}

// This is the second version
// Link: https://www.hackerrank.com/challenges/arrays-ds/problem
func reverseArray(a []int32) []int32 {
	var limit int
	if len(a)%2 == 0 {
		limit = len(a) / 2
	} else {
		limit = len(a)/2 + 1
	}

	for i := 0; i < limit; i++ {
		if i != (len(a)-1)-i {
			a[i], a[(len(a)-1)-i] = a[(len(a)-1)-i], a[i]
		}
	}
	return a
}

// Link: https://www.hackerrank.com/challenges/find-the-median/problem
// This is not the best answer, If I find time, I will resolve this problem again
func FindMedian(arr []int32) int32 {
	SortSlice(arr)
	return arr[len(arr)/2]
}

func SortSlice(n []int32) {
	var i = 1
	for i < len(n) {
		var j = i
		for j >= 1 && n[j] < n[j-1] {
			n[j], n[j-1] = n[j-1], n[j]

			j--
		}

		i++
	}
}
