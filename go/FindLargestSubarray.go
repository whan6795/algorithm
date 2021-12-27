package main

import (
	"fmt"
	"math"
)

type InputData struct {
	data []float64
}

func (d *InputData) find_max_crossing_subarray(p int, q int, r int) (int, int, float64) {
	var left_index int
	var right_index int
	var left_sum float64 = math.Inf(-1)
	var right_sum float64 = math.Inf(-1)
	var sum_num float64 = 0
	for i := q - 1; i >= p; i-- {
		sum_num += d.data[i]
		if sum_num > left_sum {
			left_sum = sum_num
			left_index = i
		}
	}
	sum_num = 0
	for i := q; i < r; i++ {
		sum_num += d.data[i]
		if sum_num > right_sum {
			right_sum = sum_num
			right_index = i
		}
	}
	return left_index, right_index, left_sum + right_sum
}

func (d *InputData) find_max_subarray(p int, q int) (int, int, float64) {
	var mid int
	if p == q-1 {
		return p, q, d.data[p]
	}
	mid = (p + q) / 2
	lp, lq, ls := d.find_max_subarray(p, mid)
	rp, rq, rs := d.find_max_subarray(mid, q)
	mp, mq, ms := d.find_max_crossing_subarray(p, mid, q)
	if ls >= rs && ls >= ms {
		return lp, lq, ls
	} else if rs >= ls && rs >= ms {
		return rp, rq, rs
	} else {
		return mp, mq, ms
	}
}

func main() {
	d := InputData{[]float64{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}}
	fmt.Println("input data is: ", d.data)
	li, ri, sum_num := d.find_max_subarray(0, len(d.data))
	fmt.Println("left index: ", li, "\nright_index:", ri, "\nsum number: ", sum_num)
}
