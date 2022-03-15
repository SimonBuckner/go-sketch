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
	panel := doc.Call("getElementById", containerId)

	sp := SettingPane{
		doc:           doc,
		panel:         panel,
		Title:         title,
		InputControls: make([]InputControl, 0),
		form:          js.Value{},
	}

	return sp
}

func (sp *SettingPane) AddInputControl(inputControl InputControl) {
	inputControl.Pane = sp
	sp.InputControls = append(sp.InputControls, inputControl)
}

func (sp *SettingPane) Render() {

	h1 := sp.doc.Call("createElement", "h1")
	h1.Set("textContent", sp.Title)
	sp.panel.Call("appendChild", h1)

	sp.form = sp.doc.Call("createElement", "form")
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
}

func NewInputControl(id string, inputType string) InputControl {
	ic := InputControl{
		Id:        id,
		InputType: inputType,
		Value:     js.Value{},
		Class:     "",
	}

	return ic
}

func (ic *InputControl) Render() {

	fmt.Printf("Adding Input Control %v\n", ic.Id)
	// div := ic.Pane.doc.Call("createElement", "div")
	label := ic.Pane.doc.Call("createElement", "label")
	control := ic.Pane.doc.Call("createElement", "input")

	inputName := ic.Id + "InputControl"

	label.Set("for", "inputName")
	label.Set("class", "inputName")
	label.Set("textContent", ic.Id)
	label.Set("id", "BALS")

	control.Set("type", ic.InputType)
	control.Set("id", inputName)
	control.Set("name", inputName)

	// div.Call("appendChild", control)
	// div.Call("appendChild", label)
	ic.Pane.panel.Call("appendChild", label)
	ic.Pane.panel.Call("appendChild", control)
	// ic.Pane.form.Call("appendChild", div)

}
