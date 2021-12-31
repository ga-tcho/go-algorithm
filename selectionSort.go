package main

func Min(a []int) (idx, n int) {
	n = a[0]
	idx = 0
	for i, v := range a {
		if n > v {
			n = v
			idx = i
		}
	}
	return
}

//選択ソート
func SelectionSort(a []int) []int {
	for i := range a {
		idx, _ := Min(a[i:])
		a[i], a[i+idx] = a[i+idx], a[i]
	}
	return a
}
