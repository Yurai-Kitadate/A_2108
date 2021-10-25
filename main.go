package main

import (
	"fmt"

	"github.com/jphacks/A_2108/src/router"
)

func main() {
	r := router.Route()
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
}
