// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"math"
)

func main() {

	// printing string raw literal using back tick
	fmt.Println(`printing some raw  string literals which is an 
	concatination of severals characters `)

	// printing random integer b/w [1,n)

	fmt.Println("My favorite number is" , rand.Intn(10))

	// printing square root of number , Always start method/properties with capital letter from package otherwise it will not be an exported 

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(9))

	// printing value of pi , using capital Pi

	fmt.Println(math.Pi)
	
}