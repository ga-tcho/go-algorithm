package main

import "math"

//マージソート
//データを分割し各々をソート行った後にマージ
//平均・最悪ともに計算が高速
//実行時間Θ(nlogn)
func MergeSort(a []int) {
	mergesort(a, 0, len(a)-1)
	return
}

func mergesort(a []int, p, r int) {
	if p < r {
		q := (p + r) / 2     // 分割
		mergesort(a, p, q)   // 前半部ソート
		mergesort(a, q+1, r) // 後半部ソート
		merge(a, p, q, r)    // マージ
	}
}

func merge(a []int, p, q, r int) {
	n1 := q - p + 1                         // p~qまでの範囲分の配列
	n2 := r - q                             // q+1~rまでの配列
	larr := make([]int, n1)                 // 前半部分配列
	rarr := make([]int, n2)                 // 後半部分配列を
	copy(larr, a[p:q+1])                    // 元配列の前半部分入れる
	larr = append(larr, int(math.MaxInt64)) // 門番を入れ実質終端とする
	copy(rarr, a[q+1:r+1])                  // 元配列の後半部分
	rarr = append(rarr, int(math.MaxInt64)) // 門番を代入
	i := 0                                  // larrのインデックス値
	j := 0                                  // rarrのインデックス値
	for k := p; r >= k; k++ {               // p~rの範囲内で部分配列を比較しながら統合
		// 昇順になっている部分配列同士を随時比較しながら統合する処理
		if larr[i] < rarr[j] { // 昇順ソート
			a[k] = larr[i]
			i++
		} else {
			a[k] = rarr[j]
			j++
		}
	}
}
