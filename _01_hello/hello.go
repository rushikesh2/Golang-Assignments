package main

import (
	"fmt"
)

const usixteenbitmax float64 = 65535
const kmhmul float64 = 1.603

type car struct {
	gas      uint16
	brak     uint16
	steering int16
	speedkmh float64
}

func (c car) kmph() float64 {
	return float64(c.gas) * (c.speedkmh / usixteenbitmax)
}
func main() {

	acar := car{gas: 22314,
		brak:     0,
		steering: 12573,
		speedkmh: 225}
	fmt.Println(acar.speedkmh)
	fmt.Println(acar.kmph())
	// http.HandleFunc("/", Myhandler1)
	// http.ListenAndServe(":8000", nil)
}

// func Myhandler1(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "welcome to go programming")
// }
