package main

import "fmt"

func main() {
	ages := map[string]int{
		"Baby":   0,
		"Justin": 45,
		"Donald": 70,
		"Canada": 150,
	}
	// START OMIT
	ages["Baby"] = 1
	ages["SecondChild"] = 0
	fmt.Println(ages)
	// END OMIT
}
