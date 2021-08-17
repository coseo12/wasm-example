package main

import (
	"fmt"
	"log"
	"strconv"
	"syscall/js"
	"time"
)

func add(this js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1+int2)
	log.Println("result:", int1+int2)

	return js.ValueOf(int1 + int2)
}


func calculateGo(this js.Value, i []js.Value) interface{} {
	start := time.Now()
	number := i[0].Int()
	result := []int{0, 1}
	if number == 0 {
		fmt.Printf("%d\n", 0)
	}
	if number == 1 {
		fmt.Printf("%d,%d\n", 0, 1)
	}
	for i := 2; i < number; i++ {
		result = append(result, result[i-2] + result[i-1])
	}
	fmt.Println("calculateGo:", result)
	elapsed := time.Since(start)
	log.Println("calculateGo:", elapsed)
	return js.ValueOf(0)
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("calculateGo", js.FuncOf(calculateGo))
}

func main() {
	c1 := make(chan struct{})
	log.Println("WASM Go Initialized")
	registerCallbacks()
	js.Global().Call("subtract", []interface{}{"a", "b", "c"})
	<-c1
}

