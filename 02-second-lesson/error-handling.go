package main

import (
	"errors"
	"fmt"
	"log"
	// "net/http"
)

func main() {

	/* Handling HTTP errors

	res, err := http.Get("http://google.com")

	if err != nil {
		// aborts the program
		// panic()

		log.Fatal("Communication error")
	}

	fmt.Println(res.Header)
	*/

	// blank identifiers "_" can replace err, but it won't describe the problem.
	res, err := sum(7, 1)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res)

}

func sum(first int, second int) (int, error) {

	res := first + second

	if res > 10 {
		return -1, errors.New("the result was greater than 10")
	}

	return res, nil
}
