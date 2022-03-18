package main

import (
	"fmt"
	"syscall/js"
)

const (
	width  = 1024
	height = 768
)

type Sketcher interface {
	Activate()
	Render()
	UpdateSetting(id string, value interface{})
}

type Sketch struct {
	ctx   *Canvas2d
	angle float64
}

var sketches []Sketcher
var currentSketch Sketcher

func main() {

	fmt.Println("Go WebAssembly Initialized 3")

	ctx := NewCanvas2d("sketch")
	if ctx == nil {
		fmt.Println("unable to get c2d")
		return
	}

	ctx.SetWidth(width)
	ctx.SetHeight(height)

	sketches = make([]Sketcher, 2)
	sketches[0] = NewSketchVector(ctx)
	sketches[1] = NewSketchIsoTiles(ctx)
	currentSketch = sketches[0]
	currentSketch.Activate()
	currentSketch.UpdateSetting("title", "So how to we push updates?")

	sketchLoop := make(chan bool)
	js.Global().Call("setTimeout", js.FuncOf(RenderLoop))

	<-sketchLoop
}

func RenderLoop(this js.Value, args []js.Value) interface{} {
	currentSketch.Render()
	js.Global().Call("setTimeout", js.FuncOf(RenderLoop))
	return nil
}
