package array

import "fmt"

func array() {
	var m [5]int
	m[4] = 10
	fmt.Println(m)

	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)

}
