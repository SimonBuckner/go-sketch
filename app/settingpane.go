package main

import (
	"fmt"
	"syscall/js"
)

// type SettingPaner interface {

// }
type SettingPane struct {
	doc           js.Value
	panel         js.Value
	form          js.Value
	Title         string
	Label         string
	InputControls []InputControl
}

func NewSettingPane(containerId string, title string) SettingPane {

	doc := js.Global().Get("document")

	sp := SettingPane{
		doc:           doc,
		panel:         doc.Call("getElementById", containerId),
		form:          doc.Call("createElement", "form"),
		Title:         title,
		InputControls: make([]InputControl, 0),
	}

	h1 := doc.Call("createElement", "h1")
	h1.Set("textContent", title)
	sp.panel.Call("appendChild", h1)
	sp.panel.Call("appendChild", sp.form)

	return sp
}

func (sp *SettingPane) AddInputControl(inputControl InputControl) {
	inputControl.Pane = sp
	sp.InputControls = append(sp.InputControls, inputControl)
}

func (sp *SettingPane) Render() {

	if len(sp.InputControls) > 0 {
		for _, ic := range sp.InputControls {
			fmt.Printf("calling redner for %v\n", ic.Id)
			ic.Render()
		}
	}
}

type InputControl struct {
	Id        string
	Class     string
	InputType string
	Value     js.Value
	Pane      *SettingPane
	Input     js.Value
	Label     js.Value
}

func NewInputControl(id string, inputType string) InputControl {
	ic := InputControl{
		Id:        id + "InputControl",
		InputType: inputType,
		Value:     js.Value{},
	}

	return ic
}

func (ic *InputControl) Render() {

	fmt.Printf("Adding Input Control %v\n", ic.Id)

	label := ic.Pane.doc.Call("createElement", "label")
	input := ic.Pane.doc.Call("createElement", "input")

	label.Set("htmlFor", ic.Id)
	label.Set("textContent", ic.Id)

	input.Set("type", ic.InputType)
	input.Set("id", ic.Id)
	input.Call("addEventListener", "change", js.FuncOf(ic.OnChange))
	// input.Call("addEventListener", "keydown", js.FuncOf(ic.OnKeyDown))
	// input.Call("addEventListener", "keydown", js.FuncOf(ic.OnInput))

	ic.Pane.form.Call("appendChild", label)
	ic.Pane.form.Call("appendChild", input)

}

func (ic *InputControl) OnChange(this js.Value, args []js.Value) interface{} {
	if len(args) == 1 {
		args[0].Call("preventDefault")
		input := ic.Pane.doc.Call("getElementById", ic.Id)
		if !input.Truthy() {
			fmt.Println("input not found")
		}
		fmt.Printf("OnChange: %v\n", input.Get("value").String())
	}
	return false
}

func (ic *InputControl) OnKeyDown(this js.Value, args []js.Value) interface{} {
	if len(args) == 1 {
		args[0].Call("preventDefault")
		input := ic.Pane.doc.Call("getElementById", ic.Id)
		if !input.Truthy() {
			fmt.Println("input not found")
		}
		fmt.Printf("OnKeyDown: %v\n", input.Get("value").String())
	}
	return false
}

func (ic *InputControl) OnInput(this js.Value, args []js.Value) interface{} {
	if len(args) == 1 {
		args[0].Call("preventDefault")
		input := ic.Pane.doc.Call("getElementById", ic.Id)
		if !input.Truthy() {
			fmt.Println("input not found")
		}
		fmt.Printf("OnInput: %v\n", input.Get("value").String())
	}
	return false
}
