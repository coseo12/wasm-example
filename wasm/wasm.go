package main

import (
	"log"
	"strconv"
	"syscall/js"
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

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
}

func main() {
	c1 := make(chan struct{})
	log.Println("WASM Go Initialized")
	registerCallbacks()
	js.Global().Call("subtract", []interface{}{"a", "b", "c"})
	<-c1
}

