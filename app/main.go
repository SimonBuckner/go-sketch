package main

import (
	"fmt"
	"syscall/js"
)

const (
	width  = 1080
	height = 1080
)

type Sketcher interface {
	Activate()
	Render()
	UpdateSetting(id string, value interface{})
}

type Sketch struct {
	ctx           *Canvas2d
	angle         float64
	sketches      []Sketcher
	currentSketch Sketcher
	buttons       []*ButtonControl
}

func main() {

	fmt.Println("Go WebAssembly Initialized 3")

	sketch := Sketch{}
	sketch.ctx = NewCanvas2d("sketch")
	if sketch.ctx == nil {
		fmt.Println("unable to get c2d")
		return
	}

	sketch.ctx.SetWidth(width)
	sketch.ctx.SetHeight(height)

	sketch.sketches = make([]Sketcher, 2)
	sketch.sketches[0] = NewSketchVector(sketch.ctx)
	sketch.sketches[1] = NewSketchIsoTiles(sketch.ctx)
	sketch.currentSketch = sketch.sketches[1]
	sketch.currentSketch.Activate()

	sketch.buttons = make([]*ButtonControl, 2)
	sketch.buttons[0] = NewButtonControl("vectors", "button", "Vectors")
	sketch.buttons[0].AddEventHandler("click", sketch.EnableVectors)

	sketch.buttons[1] = NewButtonControl("isoTiles", "button", "Iso Tiles")
	sketch.buttons[1].AddEventHandler("click", sketch.EnableIsoTiles)

	doc := GetDocument()
	nav := doc.GetElementById("sketch_nav")
	nav.Call("appendChild", sketch.buttons[0].Activate())
	nav.Call("appendChild", sketch.buttons[1].Activate())

	sketchLoop := make(chan bool)
	js.Global().Call("setTimeout", js.FuncOf(sketch.RenderLoop))

	<-sketchLoop
}

func (sketch *Sketch) RenderLoop(this js.Value, args []js.Value) interface{} {
	sketch.currentSketch.Render()
	js.Global().Call("setTimeout", js.FuncOf(sketch.RenderLoop))
	return nil
}

func (sketch *Sketch) EnableVectors(this js.Value, event js.Value) interface{} {
	sketch.currentSketch = sketch.sketches[0]
	sketch.currentSketch.Activate()
	return nil
}

func (sketch *Sketch) EnableIsoTiles(this js.Value, event js.Value) interface{} {
	sketch.currentSketch = sketch.sketches[1]
	sketch.currentSketch.Activate()
	return nil
}
