package main

import "fmt"

func main() {
	/* var d float64 = 6
	var g float64 = 6
	var total float64 = 0
	var discount float64 = 1
	for i := 0; i < 21; i++ {
		if total > 100 && total < 150 {
			discount = 0.8
		}
		if total > 150 {
			discount = 0.5
		}
		total += d * discount

		if total > 100 && total < 150 {
			discount = 0.8
		}
		if total > 150 {
			discount = 0.5
		}
		total += g * discount
		fmt.Println(total)
	} */
	// fmt.Println(total)
	// value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 9.824), 64)
	// fmt.Println(value) //9.82

	// value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", 9.826), 64)
	// fmt.Println(value) //9.83

	var a,b,c float64
   a=1.67*100
   b=1.7*10
   c=a*b/(100*10)
//正确结果2.873
   fmt.Println(c)
}
