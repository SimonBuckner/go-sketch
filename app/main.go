package main

import (
	"fmt"
	"syscall/js"
)

const (
	width  = 1024
	height = 768
)

// func add(this js.Value, i []js.Value) interface{} {
// 	js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
// 	fmt.Println(js.ValueOf(i[0].Int() + i[1].Int()))
// 	return nil
// }

// func registerCallbacks() {
// 	js.Global().Set("add", js.FuncOf(add))
// }

type Sketcher interface {
	RenderLoop()
}

type Sketch struct {
	ctx   *Canvas2d
	angle float64
}

func main() {

	fmt.Println("Go WebAssembly Initialized")

	ctx := NewCanvas2d("sketch")
	if ctx == nil {
		fmt.Println("unable to get c2d")
		return
	}

	ctx.SetWidth(width)
	ctx.SetHeight(height)

	sketch := NewSketchVector(ctx)
	sketch.sp.Render()

	sketchLoop := make(chan bool)
	var renderer js.Func
	renderer = js.FuncOf(func(this js.Value, arps []js.Value) interface{} {
		sketch.RenderLoop()
		js.Global().Call("setTimeout", renderer)
		return nil
	})
	js.Global().Call("setTimeout", renderer)

	<-sketchLoop
}
