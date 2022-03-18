package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

type SettingPane struct {
	doc       *Document
	Container js.Value
	Title     string
	Controls  map[string]*SettingControl
}

type SettingControl struct {
	Id        string
	InputType string
	Label     string
	Value     js.Value
	Container *SettingPane
}

func JsEvent(this js.Value, event js.Value) interface{} {
	return false
}

func NewSettingPane(containerId string, title string) *SettingPane {

	doc := GetDocument()
	panel := doc.GetElementById(containerId)
	div := doc.CreateElement("div")

	sp := SettingPane{
		doc:       doc,
		Container: div,
		Title:     title,
		Controls:  make(map[string]*SettingControl, 0),
	}

	h1 := doc.CreateElement("h1")
	h1.Set("textContent", title)
	panel.Call("appendChild", h1)
	panel.Call("appendChild", div)

	div.Set("id", "sp")
	div.Set("style", "display: grid; grid-template-columns: 1fr 1fr; grid-template-rows: 1fr 1fr 1fr; gap: 0px 0px;")

	return &sp
}

func (sp *SettingPane) AddInputControl(id string, inputType string, label string, defaultValue interface{}) {
	control := &SettingControl{
		Id:        id,
		InputType: inputType,
		Label:     label,
		Value:     js.ValueOf(defaultValue),
	}
	control.Container = sp
	sp.Controls[control.Id] = control
}

func (sp *SettingPane) Render() {
	if len(sp.Controls) > 0 {
		for _, ic := range sp.Controls {
			label, input := ic.Render()
			sp.Container.Call("appendChild", label)
			sp.Container.Call("appendChild", input)
			fmt.Printf("%v : %v\n", ic.Id, ic.Value.String())
		}
	}
}

func (sp *SettingPane) AppendChild(element js.Value) {
	sp.Container.Call("appendChild", element)
}

func (sp *SettingPane) GetValue(id string) js.Value {
	if ic, found := sp.Controls[id]; found {
		return ic.Value
	}
	return js.Value{}
}

func (sp *SettingPane) GetValueAsFloat(id string, defaultValue float64) float64 {
	if ic, found := sp.Controls[id]; found {

		val, err := strconv.ParseFloat(ic.Value.String(), 64)
		if err == nil {
			return val
		}
		return defaultValue
	}
	return defaultValue
}

func (sp *SettingPane) GetValueAsInt(id string, defaultValue int64) int64 {
	if ic, found := sp.Controls[id]; found {

		val, err := strconv.ParseInt(ic.Value.String(), 10, 64)
		if err == nil {
			return val
		}
		return defaultValue
	}
	return defaultValue
}

func (sp *SettingPane) GetValueAsString(id string, defaultValue string) string {
	if ic, found := sp.Controls[id]; found {
		return ic.Value.String()
	}
	return defaultValue
}

// func (sp *SettingPane) SetValue(id string, value interface{}) {
// 	if ic, found := sp.Controls[id]; found {
// 		ic.Value = js.ValueOf(value)
// 	}
// }

// func (sp *SettingPane) SetValueAsString(id string, value string) {
// 	if ic, found := sp.Controls[id]; found {
// 		ic.Value = js.ValueOf(value)
// 	}
// }

// func (sp *SettingPane) SetValueAsFloat(id string, value float64) {
// 	if ic, found := sp.Controls[id]; found {
// 		ic.Value = js.ValueOf(value)
// 		fmt.Printf("%v - %v\n", value, ic.Value.String())
// 	}
// }

// func (sp *SettingPane) SetValueAsInt(id string, value int64) {
// 	if ic, found := sp.Controls[id]; found {
// 		ic.Value = js.ValueOf(value)
// 	}
// }

func (control *SettingControl) Render() (label, input js.Value) {

	doc := control.Container.doc

	label = doc.CreateElement("label")
	input = doc.CreateElement("input")

	label.Set("htmlFor", control.Id+"_IC")
	label.Set("textContent", control.Label)

	input.Set("type", control.InputType)
	input.Set("id", "setting_"+control.Id)
	input.Set("value", control.Value)
	input.Call("addEventListener", "input", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			return js.Value{}
		}
		event := args[0]
		return control.OnInput(this, event)
	}))

	return
}

func (control *SettingControl) OnInput(this js.Value, event js.Value) interface{} {
	control.Value = this.Get("value")
	return false
}
