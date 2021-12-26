package main

import (
	"fmt"
	"math"
)

type InputData struct {
	data []float32
}

func (d *InputData) merge(p int, q int, r int) {
	var left_list = make([]float32, len(d.data[p:q]))
	var right_list = make([]float32, len(d.data[q:r]))
	var left_index int = 0
	var right_index int = 0
	// fmt.Println(p, q, r)
	// left_list = d.data[p:q]
	// right_list = d.data[q:r]\

	// 深拷贝，否则会使用同一个内存地址
	copy(left_list, d.data[p:q])
	copy(right_list, d.data[q:r])
	right_list = append(right_list, float32(math.Inf(1)))
	left_list = append(left_list, float32(math.Inf(1)))
	// fmt.Println(left_list, right_list, d.data)
	// fmt.Printf("%p\n", left_list)
	// fmt.Printf("%p\n", right_list)
	// fmt.Printf("%p\n", d.data)
	for i := p; i < r; i++ {
		if left_list[left_index] > right_list[right_index] {
			d.data[i] = right_list[right_index]
			right_index++
		} else {
			d.data[i] = left_list[left_index]
			left_index++
		}
	}
}

func (d *InputData) mergeSort(p int, r int) {
	var q int
	if p < r-1 {
		q = (p + r) / 2
		d.mergeSort(p, q)
		d.mergeSort(q, r)
		d.merge(p, q, r)
	}
}

func main() {
	// var Rander = rand.New(rand.NewSource(time.Now().UnixNano()))
	d := InputData{[]float32{65, 54, 75, 9, 57, 25, 91, 61, 11, 5, 37, 43, 81, 39, 86, 49}}
	// a := Rander.Perm(16)
	// fmt.Println(a)
	fmt.Println("input data is: ", d.data)
	d.mergeSort(0, len(d.data))
	fmt.Println("sorted data is: ", d.data)
}
