package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	err := Solve(NewState(10))
	fmt.Printf("Finished in %v\n", time.Since(start))
	if err != nil {
		fmt.Println(err)
	}

}
