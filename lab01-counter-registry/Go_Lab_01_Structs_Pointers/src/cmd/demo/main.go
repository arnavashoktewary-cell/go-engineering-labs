package main

import (
	"fmt"

	"example.com/lab01"
)

func main() {
	r := lab01.NewRegistry()
	if err := r.Register("requests", 0); err != nil {
		fmt.Println("register error:", err)
		return
	}
	if err := r.Increment("requests"); err != nil {
		fmt.Println("increment error:", err)
		return
	}
	counter, ok := r.Get("requests")
	if !ok {
		fmt.Println("registered counter missing")
		return
	}
	fmt.Printf("%s: %d\n", counter.Name, counter.Value)
}
