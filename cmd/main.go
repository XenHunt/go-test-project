package main

import "fmt"

func main() {
	var intArr [3]int32 = [3]int32{1, 2, 3}
	intArrOther := [...]int8{1, 2, 3}
	fmt.Println(intArr, intArrOther)

}
