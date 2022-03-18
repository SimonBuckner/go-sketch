package main

type SketchVector struct {
	Sketch
	Settings *SettingPane
}

func NewSketchVector(ctx *Canvas2d) *SketchVector {
	sp := NewSettingPane("settings", "Vector Settings")

	sp.AddInputControl("title", "text", "Sketch Title", "go-sketch")
	sp.AddInputControl("crossSize", "number", "Cross Size", int64(100))
	sp.AddInputControl("arrowSize", "number", "Arrow Size", int64(200))
	sp.AddInputControl("speed", "number", "Rotation Speed", float64(0.1))

	sketch := &SketchVector{}
	sketch.ctx = ctx
	sketch.Settings = sp

	return sketch
}

func (sketch *SketchVector) Activate() {
	sketch.Settings.Activate()
}

func (sketch *SketchVector) UpdateSetting(id string, value interface{}) {
	sketch.Settings.SetValue(id, value)
}

func (sketch *SketchVector) Render() {

	sketch.Settings.RefreshValues()

	speed := sketch.Settings.GetValueAsFloat("speed", -99)
	sketch.angle += speed

	title := sketch.Settings.GetValueAsString("title", "go-sketch")
	h1 := GetDocument().GetElementById("sketchTitle")
	h1.Set("textContent", title)

	sketch.ctx.SetFillStyle(Color{R: 11, G: 5, B: 38})
	sketch.ctx.FillRect(0, 0, width, height)

	sketch.ctx.SetStrokeStyle(NewColor(255, 255, 255))

	crossSize := sketch.Settings.GetValueAsFloat("crossSize", 50.0)
	sketch.ctx.Save()
	sketch.ctx.Translate(width*0.5, height*0.5)
	sketch.ctx.SetStrokeStyle(NewColor(255, 255, 255))
	cross.Width = crossSize
	cross.Height = crossSize
	cross.Angle = sketch.angle
	cross.Stroke(sketch.ctx)
	sketch.ctx.Restore()

	arrowSize := sketch.Settings.GetValueAsFloat("arrowSize", 50.0)
	sketch.ctx.Save()
	sketch.ctx.Translate(width*0.5, height*0.5)
	sketch.ctx.SetStrokeStyle(NewColor(50, 255, 500))
	arrow.Width = arrowSize
	arrow.Height = arrowSize
	arrow.Angle = 360 - sketch.angle
	arrow.Stroke(sketch.ctx)
	sketch.ctx.Restore()

	sketch.Settings.SetValue("speed", speed+0.001)
}
