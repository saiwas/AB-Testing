package main

import (
	"ab-testing/lib"
	"fmt"
)

func main() {
	test := lib.Bucket("Test", "test-api-key", "00112233")

	fmt.Println(test)
}
