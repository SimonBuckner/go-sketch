package main

import (
	"strconv"
	"syscall/js"
)

type SettingPane struct {
	doc           js.Value
	Container     js.Value
	Title         string
	InputControls map[string]*InputControl
}

func NewSettingPane(containerId string, title string) *SettingPane {

	doc := js.Global().Get("document")

	sp := SettingPane{
		doc:           doc,
		Container:     doc.Call("getElementById", containerId),
		Title:         title,
		InputControls: make(map[string]*InputControl, 0),
	}

	h1 := doc.Call("createElement", "h1")
	h1.Set("textContent", title)
	sp.Container.Call("appendChild", h1)

	return &sp
}

func (sp *SettingPane) AddInputControl(inputControl *InputControl) {
	inputControl.Container = sp
	sp.InputControls[inputControl.Id] = inputControl
}

func (sp *SettingPane) Render() {
	if len(sp.InputControls) > 0 {
		for _, ic := range sp.InputControls {
			ic.Render()
		}
	}
}

func (sp *SettingPane) AppendChild(element js.Value) {
	sp.Container.Call("appendChild", element)
}

func (sp *SettingPane) GetValue(id string) js.Value {
	if ic, found := sp.InputControls[id]; found {
		return ic.Value
	}
	return js.Value{}
}

func (sp *SettingPane) GetValueAsFloat(id string, defaultValue float64) float64 {
	if ic, found := sp.InputControls[id]; found {

		val, err := strconv.ParseFloat(ic.Value.String(), 64)
		if err == nil {
			return val
		}
		return defaultValue
	}
	return defaultValue
}

func (sp *SettingPane) GetValueAsInt(id string, defaultValue int64) int64 {
	if ic, found := sp.InputControls[id]; found {

		val, err := strconv.ParseInt(ic.Value.String(), 10, 64)
		if err == nil {
			return val
		}
		return defaultValue
	}
	return defaultValue
}

type InputControl struct {
	Id        string
	InputType string
	Label     string
	Value     js.Value
	Container *SettingPane
}

func NewInputControl(id string, inputType string, label string, defaultValue interface{}) *InputControl {
	ic := InputControl{
		Id:        id,
		InputType: inputType,
		Label:     label,
		Value:     js.ValueOf(defaultValue),
	}
	return &ic
}

func (ic *InputControl) Render() {

	// fmt.Printf("Adding Input Control %v\n", ic.Id)

	doc := js.Global().Get("document")
	label := doc.Call("createElement", "label")
	input := doc.Call("createElement", "input")

	label.Set("htmlFor", ic.Id+"_IC")
	label.Set("textContent", ic.Id+"_IC")

	input.Set("type", ic.InputType)
	input.Set("id", ic.Id+"_IC")
	input.Set("value", ic.Value)
	input.Call("addEventListener", "input", js.FuncOf(ic.OnInput))

	// if ic.Container == nil {
	// 	fmt.Println("Container is empty")
	// }
	ic.Container.AppendChild(label)
	ic.Container.AppendChild(input)

}

func (ic *InputControl) OnInput(this js.Value, args []js.Value) interface{} {
	if len(args) >= 1 {
		ic.Value = this.Get("value")
		// fmt.Printf("OnChange: %v\n", ic.Value.String())
	}
	return false
}
