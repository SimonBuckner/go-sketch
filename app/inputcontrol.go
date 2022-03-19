package main

import (
	"fmt"
	"syscall/js"
)

type InputControl struct {
	Id           string
	InputType    string
	Label        string
	Value        js.Value
	DefaultValue js.Value
	Events       map[string]JsEventer
}

func NewInputControl(id string, inputType string, label string, defaultValue interface{}) *InputControl {
	control := &InputControl{
		Id:           id,
		InputType:    inputType,
		Label:        label,
		Value:        js.ValueOf(defaultValue),
		DefaultValue: js.ValueOf(defaultValue),
		Events:       make(map[string]JsEventer),
	}

	if inputType != "button" {
		control.AddEventHandler("input", control.OnInput)
	}
	return control
}

func (control *InputControl) Activate() (label, input js.Value) {

	doc := GetDocument()

	label = doc.CreateElement("label")
	input = doc.CreateElement("input")

	label.Set("htmlFor", "setting_"+control.Id)
	label.Set("textContent", control.Label)

	input.Set("type", control.InputType)
	input.Set("id", "setting_"+control.Id)
	input.Set("value", control.Value)
	input.Set("placeHolder", control.DefaultValue)

	for event, handler := range control.Events {
		input.Call("addEventListener", event, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if len(args) < 1 {
				return js.Value{}
			}
			return handler(this, args[0])
		}))
	}

	return
}

func (control *InputControl) RefreshValue() {
	doc := GetDocument()
	input := doc.GetElementById("setting_" + control.Id)
	if input.Truthy() {
		input.Set("value", control.Value)
	} else {
		fmt.Printf("unable to get value %v\n", control.Id)
	}
}

type JsEventer func(this js.Value, event js.Value) interface{}

func (control *InputControl) OnInput(this js.Value, event js.Value) interface{} {
	control.Value = this.Get("value")
	return false
}

func (control *InputControl) DisableDefaultEvent(event string) {

}

func (control *InputControl) AddEventHandler(event string, handler JsEventer) {
	control.Events[event] = handler
}

func (control *InputControl) RemoveEventHandler(event string, handler JsEventer) {
	delete(control.Events, event)
}
