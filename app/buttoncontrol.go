package main

import (
	"syscall/js"
)

type ButtonControl struct {
	Id        string
	InputType string
	Label     string
	Events    map[string]JsEventer
}

func NewButtonControl(id string, inputType string, label string) *ButtonControl {
	control := &ButtonControl{
		Id:        id,
		InputType: inputType,
		Label:     label,
		Events:    make(map[string]JsEventer),
	}

	return control
}

func (control *ButtonControl) Activate() (button js.Value) {

	doc := GetDocument()

	button = doc.CreateElement("button")

	button.Set("type", control.InputType)
	button.Set("id", "button_"+control.Id)
	button.Set("textContent", control.Label)

	for event, handler := range control.Events {
		button.Call("addEventListener", event, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if len(args) < 1 {
				return js.Value{}
			}
			return handler(this, args[0])
		}))
	}

	return
}

func (control *ButtonControl) RefreshValue() {
	// doc := GetDocument()
	// input := doc.GetElementById("setting_" + control.Id)
	// if input.Truthy() {
	// 	input.Set("value", control.Value)
	// } else {
	// 	fmt.Printf("unable to get value %v\n", control.Id)
	// }
}

// type JsEventer func(this js.Value, event js.Value) interface{}

// func (control *ButtonControl) OnInput(this js.Value, event js.Value) interface{} {
// 	control.Value = this.Get("value")
// 	return false
// }

func (control *ButtonControl) AddEventHandler(event string, handler JsEventer) {
	control.Events[event] = handler
}

func (control *ButtonControl) RemoveEventHandler(event string, handler JsEventer) {
	delete(control.Events, event)
}
