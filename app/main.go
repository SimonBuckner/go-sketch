package main

import (
	"fmt"
	"syscall/js"
)

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

	sketch := Sketch{
		ctx:   ctx,
		angle: 0,
	}

	sketchLoop := make(chan bool)
	var renderer js.Func
	renderer = js.FuncOf(func(this js.Value, arps []js.Value) interface{} {
		sketch.renderLoop()
		js.Global().Call("setTimeout", renderer)
		return nil
	})
	js.Global().Call("setTimeout", renderer)
	<-sketchLoop
}

func (sketch *Sketch) renderLoop() {

	sketch.angle += 0.1

	sketch.ctx.SetFillStyle(Color{R: 11, G: 5, B: 38})
	sketch.ctx.FillRect(0, 0, width, height)

	sketch.ctx.SetStrokeStyle(NewColor(255, 255, 255))

	sketch.ctx.BeginPath()
	sketch.ctx.Rect(10, 10, 100, 100)
	sketch.ctx.Stroke()

	sketch.ctx.Save()
	sketch.ctx.Translate(width*0.5, height*0.5)
	sketch.ctx.SetStrokeStyle(NewColor(255, 255, 255))
	cross.Width = 50
	cross.Height = 50
	cross.Angle = sketch.angle
	cross.Stroke(sketch.ctx)
	sketch.ctx.Restore()

	sketch.ctx.Save()
	sketch.ctx.Translate(width*0.25, height*0.25)
	sketch.ctx.SetStrokeStyle(NewColor(50, 255, 500))
	arrow.Width = 200
	arrow.Height = 200
	arrow.Angle = sketch.angle
	arrow.Stroke(sketch.ctx)
	sketch.ctx.Restore()

}
