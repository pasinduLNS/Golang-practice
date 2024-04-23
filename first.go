
// pracetise in : https://goplay.tools/

package main

// import functions
// all modules with go
import (
	"fmt"
	"math"
)

func main() {

	// define variables in golang
	// var a int = 10
	// var b string = "pasindu sankalpa"
	var c float64 = 34.56
	fmt.Printf("%.1f\n", c)

	// we can covert the values types with the go
	var d int64 = int64(c)
	fmt.Println("this is the value of converted c", d)

	// constants with go
	const p = "PASINDU"
	const s = "SANKALPA"
	fmt.Println(p)
	fmt.Println(s)

	// usage of math function
	fmt.Printf("%.2f\n", math.Floor(c))
	fmt.Printf("%.2f\n", math.Sqrt(c))
	fmt.Printf("%.2f\n", math.Pow(c, 2)) // powered by two

}
