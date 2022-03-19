package main

import (
	"strconv"
	"syscall/js"
)

type SettingPane struct {
	Container js.Value
	Title     string
	Controls  map[string]*InputControl
	Keys      []string
}

func NewSettingPane(containerId string, title string) *SettingPane {

	doc := GetDocument()

	sp := SettingPane{
		Title:     title,
		Container: doc.GetElementById(containerId),
		Controls:  make(map[string]*InputControl, 0),
	}

	return &sp
}

func (sp *SettingPane) AddInputControl(id string, inputType string, label string, defaultValue interface{}) {
	control := NewInputControl(id, inputType, label, defaultValue)
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
