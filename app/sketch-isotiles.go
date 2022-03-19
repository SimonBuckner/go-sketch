package main

type SketchIsoTiles struct {
	Sketch
	Settings *SettingPane
}

// type Tile struct {
// 	Vectors
// }

func NewSketchIsoTiles(ctx *Canvas2d) *SketchVector {
	sp := NewSettingPane("settings", "ISO Tiles")

	sp.AddInputControl("title", "text", "Sketch Title", "go-sketch")
	sp.AddInputControl("size", "number", "Tile Size", "100")

	sketch := &SketchVector{}
	sketch.ctx = ctx
	sketch.Settings = sp

	return sketch
}

func (sketch *SketchIsoTiles) Activate() {
	sketch.Settings.Activate()
}

func (sketch *SketchIsoTiles) UpdateSetting(id string, value interface{}) {
	sketch.Settings.SetValue(id, value)
}

func (sketch *SketchIsoTiles) Render() {

	speed := sketch.Settings.GetValueAsFloat("speed", -99)
	sketch.angle += speed

	title := sketch.Settings.GetValueAsString("title", "go-sketch")
	h1 := GetDocument().GetElementById("sketchTitle")
	h1.Set("textContent", title)

	sketch.ctx.SetFillStyle(Color{R: 11, G: 5, B: 38})
	sketch.ctx.FillRect(0, 0, width, height)

}
