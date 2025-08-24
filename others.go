package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// httpGet performs an HTTP GET request and returns the response body as a string
func httpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
func minimumBribes(q []int32) {
	// Write your code here

	// bribes := make(map[int]int)
	bribes := 0
	for i := len(q) - 1; i >= 0; i-- {
		if q[i] == int32(i+1) {
			continue
		} else {
			if i-1 >= 0 {
				fmt.Println("Testing waters...", q[i], int32(i-1), q[i]-int32(i-1))
				if int32(i-1)-q[i] < 0 {
					bribes++
				}
			}

			// fmt.Println("How many this guy bribe", q[i], "[", i, "]", " was : ", len(q)-i)
		}
	}
	fmt.Println("Too chaotic", bribes)
}

func minimumDistances(a []int32) int32 {
	type Dist struct {
		FirstIndex int32
		LastIndex  int32
		Dist       int32
	}
	if len(a) == 0 || len(a) == 1 {
		return -1
	}

	distMap := make(map[int32]Dist)
	min := int32(-1)
	for cI, val := range a {
		trueIndex := int32(cI) + 1
		if distMap[val].FirstIndex == 0 {
			distMap[val] = Dist{
				FirstIndex: trueIndex,
				LastIndex:  trueIndex,
				Dist:       0,
			}
		} else {
			if (trueIndex - 1) >= distMap[val].LastIndex {
				cp := distMap[val]
				cp.LastIndex = trueIndex
				cp.Dist = cp.LastIndex - cp.FirstIndex
				distMap[val] = cp

				if min == -1 || cp.Dist <= min {
					min = cp.Dist
				}
			}
		}
	}
	return min
}

func isSubsequence(s string, t string) bool {
	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		if string(t[j]) == string(s[i]) {
			i++
		}

		j++
	}
	return i == len(s)
}

func isHappy(n int) bool {
	if n <= 0 {
		return false
	}

	sn := strconv.Itoa(n)

	isRepeated := make(map[int]bool)

	sum := 0
	nextString := sn
	for {
		sum = 0
		for _, v := range nextString {
			nv, _ := strconv.Atoi(string(v))
			sum += nv * nv
		}
		nextSum := fmt.Sprintf("%d", sum)
		fmt.Println("nextString", nextString, "sum:", nextSum)
		if isRepeated[sum] || (n > 1 && sum == n) {
			return false
		}

		isRepeated[sum] = true
		if sum == 1 {
			return true
		}

		if nextSum == nextString {
			return false
		}
		nextString = nextSum
	}

}

func getValue[T any](data map[string]interface{}, key string) (T, bool) {
	val, ok := data[key]
	if !ok {
		var zero T
		return zero, false
	}
	converted, ok := val.(T)
	return converted, ok
}

func Example() {
	raw := []byte(`{"price": 10.99, "title": "Book"}`)

	var result map[string]interface{}
	json.Unmarshal(raw, &result)

	// Extracting float64
	price, ok := getValue[float64](result, "price")
	if !ok {
		fmt.Println("Could not get price")
		return
	}

	// Extracting string
	title, ok := getValue[string](result, "title")
	if !ok {
		fmt.Println("Could not get title")
		return
	}

	fmt.Printf("Title: %s - Price: %.2f\n", title, price)

}

func sendMessage1(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "Message from Goroutine 1"
}

func sendMessage2(ch chan<- string) {
	time.Sleep(1 * time.Second)
	ch <- "Message from Goroutine 2"
}

func startSelect() {
	// Create two channels for communication
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start two goroutines that will send messages
	go sendMessage1(ch1)
	go sendMessage2(ch2)

	// Use select to listen on multiple channels
	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	case <-time.After(3 * time.Second): // Timeout case to avoid waiting indefinitely
		fmt.Println("Timeout: No messages received")
	}
}
