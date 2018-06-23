package main

import (
	"ab-testing/lib"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	config := make(map[string]uint64)
	config["A"] = 60
	config["B"] = 40

	result := make(map[string]int)
	result["A"] = 0
	result["B"] = 0
	result["C"] = 0
	result["D"] = 0

	count := 100000000

	fmt.Printf("the percentage of the versions : %v \n", config)
	fmt.Printf("%v Users are joining this test! \n", count)

	start := time.Now()

	chanVersion := make(chan string, 10)
	receiveCount := 0
	buckets := lib.CreateBucket("Test", "test-api-key", config)

	go func() {
		for i := 0; i < count; i++ {
			ID := RandStringRunes(16)
			test := lib.GetVersion(buckets, ID)
			// sent msg function
			chanVersion <- test
		}
	}()

	// After change to ListenAndServe, We can remove this
	for receiveCount != count {
		version := <-chanVersion
		result[version] += 1
		receiveCount += 1
	}
	close(chanVersion)
	// end
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("The results: %v \n", result)
	fmt.Printf("Time spend : %v \n", elapsed)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
