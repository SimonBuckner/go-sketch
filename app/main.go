package main

import (
	"fmt"
)

var c chan bool

const (
	width  = 1080
	height = 1080
)

// func add(this js.Value, i []js.Value) interface{} {
// 	js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
// 	fmt.Println(js.ValueOf(i[0].Int() + i[1].Int()))
// 	return nil
// }

// func registerCallbacks() {
// 	js.Global().Set("add", js.FuncOf(add))
// }

func main() {

	fmt.Println("Go WebAssembly Initialized")

	ctx := NewCanvas2d("sketch")
	if ctx == nil {
		fmt.Println("unable to get c2d")
		return
	}

	ctx.SetWidth(width)
	ctx.SetHeight(height)

	ctx.SetFillStyle(Color{R: 11, G: 5, B: 38})

	ctx.BeginPath()
	ctx.Rect(10, 10, 100, 100)
	ctx.Stroke()

	<-c
}
