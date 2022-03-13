package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

type Color struct {
	R int
	G int
	B int
	A int
}

func (c Color) ToString() string {
	return fmt.Sprintf("0x%02x%02x%02x", c.R, c.G, c.B)
	// return fmt.Sprintf("0x%02x%02x%02x%02x", c.R, c.G, c.B, c.A)
}

type Canvas2d struct {
	doc     js.Value
	canvas  js.Value
	context js.Value
}

func NewCanvas2d(canvasId string) *Canvas2d {
	doc := js.Global().Get("document")
	canvas := doc.Call("getElementById", canvasId)
	context := canvas.Call("getContext", "2d")

	if doc.Truthy() && canvas.Truthy() && context.Truthy() {
		return &Canvas2d{
			doc:     doc,
			canvas:  canvas,
			context: context,
		}
	}
	return nil
}

func (c *Canvas2d) Width() float64 {
	return c.canvas.Get("width").Float()
}

func (c *Canvas2d) Height() float64 {
	return c.canvas.Get("height").Float()
}

func (c *Canvas2d) SetWidth(width float64) {
	c.canvas.Set("width", width)
}

func (c *Canvas2d) SetHeight(height float64) {
	c.canvas.Set("height", height)
}

func (c *Canvas2d) FillStyle() Color {
	fs := c.context.Get("fillStyle").String()

	r, _ := strconv.ParseInt(fs[1:3], 16, 8)
	g, _ := strconv.ParseInt(fs[3:5], 16, 8)
	b, _ := strconv.ParseInt(fs[5:7], 16, 8)

	return Color{R: int(r), G: int(g), B: int(b)}
}

func (c *Canvas2d) SetFillStyle(color Color) {
	fill := fmt.Sprintf("rgb(%d, %d, %d)", color.R, color.G, color.B)
	c.context.Set("fillStyle", fill)
}

func (c *Canvas2d) StrokeStyle() Color {
	fs := c.context.Get("strokeStyle").String()

	r, _ := strconv.ParseInt(fs[1:3], 16, 8)
	g, _ := strconv.ParseInt(fs[3:5], 16, 8)
	b, _ := strconv.ParseInt(fs[5:7], 16, 8)

	return Color{R: int(r), G: int(g), B: int(b)}
}

func (c *Canvas2d) SetStrokeStyle(color Color) {
	stroke := fmt.Sprintf("rgb(%d, %d, %d)", color.R, color.G, color.B)
	c.context.Set("strokeStyle", stroke)
}

// arc()
func (c *Canvas2d) Arc() {
	c.context.Call("arc")
}

func (c *Canvas2d) ArcTo() {
	c.context.Call("arcTo")
}

func (c *Canvas2d) BeginPath() {
	c.context.Call("beginPath")
}

func (c *Canvas2d) bezierCurveTo() {
	c.context.Call("bezierCurveTo")
}

func (c *Canvas2d) ClearRect() {
	c.context.Call("clearRect")
}

func (c *Canvas2d) clip() {
	c.context.Call("clip")
}

func (c *Canvas2d) ClosePath() {
	c.context.Call("closePath")
}

func (c *Canvas2d) createConicGradient() {
	c.context.Call("createConicGradient")
}

func (c *Canvas2d) createImageData() {
	c.context.Call("createImageData")
}

func (c *Canvas2d) createLinearGradient() {
	c.context.Call("createLinearGradient")
}

func (c *Canvas2d) createPattern() {
	c.context.Call("createPattern")
}

func (c *Canvas2d) createRadialGradient() {
	c.context.Call("createRadialGradient")
}

func (c *Canvas2d) drawFocusIfNeeded() {
	c.context.Call("drawFocusIfNeeded")
}

func (c *Canvas2d) drawImage() {
	c.context.Call("drawImage")
}

func (c *Canvas2d) drawWidgetAsOnScreen() {
	c.context.Call("drawWidgetAsOnScreen")
}

func (c *Canvas2d) drawWindow() {
	c.context.Call("drawWindow")
}

func (c *Canvas2d) ellipse() {
	c.context.Call("ellipse")
}

func (c *Canvas2d) fill() {
	c.context.Call("fill")
}

func (c *Canvas2d) fillRect() {
	c.context.Call("fillRect")
}

func (c *Canvas2d) fillText() {
	c.context.Call("fillText")
}

func (c *Canvas2d) getContextAttributes() {
	c.context.Call("getContextAttributes")
}

func (c *Canvas2d) getImageData() {
	c.context.Call("getImageData")
}

func (c *Canvas2d) getLineDash() {
	c.context.Call("getLineDash")
}

func (c *Canvas2d) getTransform() {
	c.context.Call("getTransform")
}

func (c *Canvas2d) isPointInPath() {
	c.context.Call("isPointInPath")
}

func (c *Canvas2d) isPointInStroke() {
	c.context.Call("isPointInStroke")
}

func (c *Canvas2d) lineTo() {
	c.context.Call("lineTo")
}

func (c *Canvas2d) measureText() {
	c.context.Call("measureText")
}

func (c *Canvas2d) moveTo() {
	c.context.Call("moveTo")
}

func (c *Canvas2d) putImageData() {
	c.context.Call("putImageData")
}

func (c *Canvas2d) quadraticCurveTo() {
	c.context.Call("quadraticCurveTo")
}

func (c *Canvas2d) Rect(x, y, w, h float64) {
	c.context.Call("rect", x, y, w, h)
}

func (c *Canvas2d) resetTransform() {
	c.context.Call("resetTransform")
}

func (c *Canvas2d) Restore() {
	c.context.Call("restore")
}

func (c *Canvas2d) rotate() {
	c.context.Call("rotate")
}

func (c *Canvas2d) Save() {
	c.context.Call("save")
}

func (c *Canvas2d) scale() {
	c.context.Call("scale")
}

func (c *Canvas2d) scrollPathIntoView() {
	c.context.Call("scrollPathIntoView")
}

func (c *Canvas2d) setLineDash() {
	c.context.Call("setLineDash")
}

func (c *Canvas2d) setTransform() {
	c.context.Call("setTransform")
}

func (c *Canvas2d) Stroke() {
	c.context.Call("stroke")
}

func (c *Canvas2d) strokeRect() {
	c.context.Call("strokeRect")
}

func (c *Canvas2d) strokeText() {
	c.context.Call("strokeText")
}

func (c *Canvas2d) transform() {
	c.context.Call("transform")
}

func (c *Canvas2d) translate() {
	c.context.Call("translate")
}
