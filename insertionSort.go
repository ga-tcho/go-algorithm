package main

//挿入ソート
//速度 Ο(n2)
//小さな配列に対しては高速
func InsertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i-j-1] > a[i-j] {
				a[i-j-1], a[i-j] = a[i-j], a[i-j-1]
			} else {
				break
			}
		}
	}
	return a
}
