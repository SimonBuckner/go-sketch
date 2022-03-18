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

func NewColor(r, g, b int) Color {
	return Color{R: r, G: g, B: b}
}

func (c Color) ToString() string {
	return fmt.Sprintf("0x%02x%02x%02x", c.R, c.G, c.B)
}

type Canvas2d struct {
	doc     *Document
	canvas  js.Value
	context js.Value
}

func NewCanvas2d(canvasId string) *Canvas2d {
	doc := GetDocument()
	canvas := doc.GetElementById(canvasId)
	context := canvas.Call("getContext", "2d")

	if doc.Truthy() && canvas.Truthy() && context.Truthy() {
		return &Canvas2d{
			doc:     doc,
			canvas:  canvas,
			context: context,
		}
	}
	fmt.Println("error getting canvas")
	return nil
}

func (c2d *Canvas2d) Width() float64 {
	return c2d.canvas.Get("width").Float()
}

func (c2d *Canvas2d) Height() float64 {
	return c2d.canvas.Get("height").Float()
}

func (c2d *Canvas2d) SetWidth(width float64) {
	c2d.canvas.Set("width", width)
}

func (c2d *Canvas2d) SetHeight(height float64) {
	c2d.canvas.Set("height", height)
}

func (c2d *Canvas2d) FillStyle() Color {
	fs := c2d.context.Get("fillStyle").String()

	r, _ := strconv.ParseInt(fs[1:3], 16, 8)
	g, _ := strconv.ParseInt(fs[3:5], 16, 8)
	b, _ := strconv.ParseInt(fs[5:7], 16, 8)

	return Color{R: int(r), G: int(g), B: int(b)}
}

func (c2d *Canvas2d) SetFillStyle(color Color) {
	fill := fmt.Sprintf("rgb(%d, %d, %d)", color.R, color.G, color.B)
	c2d.context.Set("fillStyle", fill)
}

func (c2d *Canvas2d) StrokeStyle() Color {
	fs := c2d.context.Get("strokeStyle").String()

	r, _ := strconv.ParseInt(fs[1:3], 16, 8)
	g, _ := strconv.ParseInt(fs[3:5], 16, 8)
	b, _ := strconv.ParseInt(fs[5:7], 16, 8)

	return Color{R: int(r), G: int(g), B: int(b)}
}

func (c2d *Canvas2d) SetStrokeStyle(color Color) {
	stroke := fmt.Sprintf("rgb(%d, %d, %d)", color.R, color.G, color.B)
	c2d.context.Set("strokeStyle", stroke)
}

func (c2d *Canvas2d) Arc(x, y, radius, startAngle, endAngle float64, counterclockwise bool) {
	c2d.context.Call("arc", x, y, radius, startAngle, endAngle, counterclockwise)
}

func (c2d *Canvas2d) ArcTo(x1, y1, x2, y2, radius float64) {
	c2d.context.Call("arcTo", x1, y1, x2, y2, radius)
}

func (c2d *Canvas2d) BeginPath() {
	c2d.context.Call("beginPath")
}

func (c2d *Canvas2d) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	c2d.context.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

func (c2d *Canvas2d) ClearRect() {
	c2d.context.Call("clearRect")
}

func (c2d *Canvas2d) Clip() {
	c2d.context.Call("fill")
}

func (c2d *Canvas2d) ClipNonZero() {
	c2d.context.Call("fill", "nonzero")
}

func (c2d *Canvas2d) ClipEvenOdd() {
	c2d.context.Call("fill", "evenodd")
}

func (c2d *Canvas2d) ClosePath() {
	c2d.context.Call("closePath")
}

// func (c2d *Canvas2d) createConicGradient() {
// 	c2d.context.Call("createConicGradient")
// }

// func (c2d *Canvas2d) createImageData() {
// 	c2d.context.Call("createImageData")
// }

// func (c2d *Canvas2d) createLinearGradient() {
// 	c2d.context.Call("createLinearGradient")
// }

// func (c2d *Canvas2d) createPattern() {
// 	c2d.context.Call("createPattern")
// }

// func (c2d *Canvas2d) createRadialGradient() {
// 	c2d.context.Call("createRadialGradient")
// }

// func (c2d *Canvas2d) drawFocusIfNeeded() {
// 	c2d.context.Call("drawFocusIfNeeded")
// }

// func (c2d *Canvas2d) drawImage() {
// 	c2d.context.Call("drawImage")
// }

// func (c2d *Canvas2d) drawWidgetAsOnScreen() {
// 	c2d.context.Call("drawWidgetAsOnScreen")
// }

// func (c2d *Canvas2d) drawWindow() {
// 	c2d.context.Call("drawWindow")
// }

// func (c2d *Canvas2d) ellipse() {
// 	c2d.context.Call("ellipse")
// }

func (c2d *Canvas2d) Fill() {
	c2d.context.Call("fill")
}

func (c2d *Canvas2d) FillNonZero() {
	c2d.context.Call("fill", "nonzero")
}

func (c2d *Canvas2d) FillEvenOdd() {
	c2d.context.Call("fill", "evenodd")
}

func (c2d *Canvas2d) FillRect(x, y, width, height float64) {
	c2d.context.Call("fillRect", x, y, width, height)
}

// func (c2d *Canvas2d) fillText() {
// 	c2d.context.Call("fillText")
// }

// func (c2d *Canvas2d) getContextAttributes() {
// 	c2d.context.Call("getContextAttributes")
// }

// func (c2d *Canvas2d) getImageData() {
// 	c2d.context.Call("getImageData")
// }

// func (c2d *Canvas2d) getLineDash() {
// 	c2d.context.Call("getLineDash")
// }

// func (c2d *Canvas2d) getTransform() {
// 	c2d.context.Call("getTransform")
// }

// func (c2d *Canvas2d) isPointInPath() {
// 	c2d.context.Call("isPointInPath")
// }

// func (c2d *Canvas2d) isPointInStroke() {
// 	c2d.context.Call("isPointInStroke")
// }

func (c2d *Canvas2d) LineTo(x, y float64) {
	c2d.context.Call("lineTo", x, y)
}

// func (c2d *Canvas2d) measureText() {
// 	c2d.context.Call("measureText")
// }

func (c2d *Canvas2d) MoveTo(x, y float64) {
	c2d.context.Call("moveTo", x, y)
}

// func (c2d *Canvas2d) putImageData() {
// 	c2d.context.Call("putImageData")
// }

// func (c2d *Canvas2d) quadraticCurveTo() {
// 	c2d.context.Call("quadraticCurveTo")
// }

func (c2d *Canvas2d) Rect(x, y, w, h float64) {
	c2d.context.Call("rect", x, y, w, h)
}

// func (c2d *Canvas2d) resetTransform() {
// 	c2d.context.Call("resetTransform")
// }

func (c2d *Canvas2d) Restore() {
	c2d.context.Call("restore")
}

func (c2d *Canvas2d) Rotate(angle float64) {
	c2d.context.Call("rotate", angle)
}

func (c2d *Canvas2d) Save() {
	c2d.context.Call("save")
}

func (c2d *Canvas2d) Scale(x, y float64) {
	c2d.context.Call("scale", x, y)
}

// func (c2d *Canvas2d) scrollPathIntoView() {
// 	c2d.context.Call("scrollPathIntoView")
// }

// func (c2d *Canvas2d) setLineDash() {
// 	c2d.context.Call("setLineDash")
// }

// func (c2d *Canvas2d) setTransform() {
// 	c2d.context.Call("setTransform")
// }

func (c2d *Canvas2d) Stroke() {
	c2d.context.Call("stroke")
}

func (c2d *Canvas2d) StrokePath(path Path2d) {
	c2d.context.Call("stroke", path.Value)
}

func (c2d *Canvas2d) StrokeRect(x, y, width, height float64) {
	c2d.context.Call("strokeRect", x, y, width, height)
}

func (c2d *Canvas2d) StrokeText(text string, x, y float64) {
	c2d.context.Call("strokeText", text, x, y)
}

func (c2d *Canvas2d) StrokeTextWithMaxWidth(text string, x, y, maxWidth float64) {
	c2d.context.Call("strokeText", text, x, y, maxWidth)
}
func (c2d *Canvas2d) Transform(a, b, c, d, e, f float64) {
	c2d.context.Call("transform", a, b, c, d, e, f)
}

func (c2d *Canvas2d) Translate(x, y float64) {
	c2d.context.Call("translate", x, y)
}

type Path2d struct {
	Value js.Value
}

func NewPath2d() Path2d {
	return Path2d{Value: js.Global().Get("Path2D").New()}
}

func NewPath2dFromVertices(vertices [][]float64) Path2d {
	path := NewPath2d()

	path.MoveTo(vertices[0][0], vertices[0][1])
	for i := 1; i < len(vertices); i++ {
		path.LineTo(vertices[i][0], vertices[i][1])
	}
	path.ClosePath()
	return path
}

// func NewScaledPath2dFromVertices(w, h float64, vertices [][]float64) Path2d {
// 	path := NewPath2d()

// 	path.MoveTo(vertices[0][0]*w, vertices[0][1]*h)
// 	for i := 1; i < len(vertices); i++ {
// 		path.LineTo(vertices[i][0]*w, vertices[i][1]*h)
// 	}
// 	path.ClosePath()
// 	return path
// }

func (path2d *Path2d) AddPath(path Path2d, a, b, c, d, e, f float64) {
	matrix := js.Global().Get("DOMMatrix").New([]float64{a, b, c, d, e, f})
	path2d.Value.Call("addPath", path.Value, matrix)
}

func (path2d *Path2d) ClosePath() {
	path2d.Value.Call("closePath")
}

func (path2d *Path2d) MoveTo(x, y float64) {
	path2d.Value.Call("moveTo", x, y)
}

func (path2d *Path2d) LineTo(x, y float64) {
	path2d.Value.Call("lineTo", x, y)
}

func (path2d *Path2d) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	path2d.Value.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

// func (path2d *Path2d) quadraticCurveTo() {
// 	path2d.Call("quadraticCurveTo")
// }

func (path2d *Path2d) Arc(x, y, radius, startAngle, endAngle float64, counterclockwise bool) {
	path2d.Value.Call("arc", x, y, radius, startAngle, endAngle, counterclockwise)
}

func (path2d *Path2d) ArcTo(x1, y1, x2, y2, radius float64) {
	path2d.Value.Call("arcTo", x1, y1, x2, y2, radius)
}

// func (path2d *Path2d) ellipse() {
// 	path2d.Call("ellipse")
// }

func (path2d *Path2d) Rect(x, y, w, h float64) {
	path2d.Value.Call("rect", x, y, w, h)
}

func (path2d *Path2d) ApplyTransform(a, b, c, d, e, f float64) {
	path2d.Value.Call("rect", a, b, c, d, e, f)
}
