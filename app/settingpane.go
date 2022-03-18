package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

type SettingPane struct {
	Container js.Value
	Title     string
	Controls  map[string]*SettingControl
	Keys      []string
}

func NewSettingPane(containerId string, title string) *SettingPane {

	doc := GetDocument()

	sp := SettingPane{
		Title:     title,
		Container: doc.GetElementById(containerId),
		Controls:  make(map[string]*SettingControl, 0),
	}

	return &sp
}

func (sp *SettingPane) AddInputControl(id string, inputType string, label string, defaultValue interface{}) {
	control := NewSettingControl(id, inputType, label, defaultValue)
	sp.Keys = append(sp.Keys, id)
	sp.Controls[control.Id] = control
}

func (sp *SettingPane) Activate() {
	doc := GetDocument()
	sp.Container.Set("textContent", "")

	h1 := doc.CreateElement("h1")
	h1.Set("id", "settings_title")
	h1.Set("textContent", sp.Title)
	sp.Container.Call("appendChild", h1)

	div := doc.CreateElement("div")
	div.Set("id", "sp")
	div.Set("style", "display: grid; grid-template-columns: 1fr 1fr; grid-template-rows: 1fr 1fr 1fr; gap: 0px 0px;")

	if len(sp.Keys) > 0 {
		for _, key := range sp.Keys {
			label, input := sp.Controls[key].Activate()
			div.Call("appendChild", label)
			div.Call("appendChild", input)
		}
	}

	sp.Container.Call("appendChild", div)

}

func (sp *SettingPane) RefreshValues() {
	doc := GetDocument()
	h1 := doc.GetElementById("settings_title")
	h1.Set("textContent", sp.Title)

	for _, control := range sp.Controls {
		control.RefreshValue()
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
		if ic.Value.Type() == js.TypeNumber {
			return ic.Value.Float()
		}
		val, err := strconv.ParseFloat(ic.Value.String(), 64)
		if err == nil {
			return val
		}
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

func (sp *SettingPane) SetValue(id string, value interface{}) {
	if ic, found := sp.Controls[id]; found {
		ic.Value = js.ValueOf(value)
	}
}

type SettingControl struct {
	Id           string
	InputType    string
	Label        string
	Value        js.Value
	DefaultValue js.Value
}

func NewSettingControl(id string, inputType string, label string, defaultValue interface{}) *SettingControl {
	control := &SettingControl{
		Id:           id,
		InputType:    inputType,
		Label:        label,
		Value:        js.ValueOf(defaultValue),
		DefaultValue: js.ValueOf(defaultValue),
	}
	return control
}

func (control *SettingControl) Activate() (label, input js.Value) {

	doc := GetDocument()

	label = doc.CreateElement("label")
	input = doc.CreateElement("input")

	label.Set("htmlFor", "setting_"+control.Id)
	label.Set("textContent", control.Label)

	input.Set("type", control.InputType)
	input.Set("id", "setting_"+control.Id)
	input.Set("value", control.Value)
	input.Set("placeHolder", control.DefaultValue)

	input.Call("addEventListener", "input", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			return js.Value{}
		}
		event := args[0]
		return control.OnInput(this, event)
	}))

	return
}

func (control *SettingControl) RefreshValue() {
	doc := GetDocument()
	input := doc.GetElementById("setting_" + control.Id)
	if input.Truthy() {
		input.Set("value", control.Value)
	} else {
		fmt.Printf("unable to get value %v\n", control.Id)
	}
}

func JsEvent(this js.Value, event js.Value) interface{} {
	return false
}

func (control *SettingControl) OnInput(this js.Value, event js.Value) interface{} {
	control.Value = this.Get("value")
	return false
}
