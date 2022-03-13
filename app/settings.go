package main

import (
	"fmt"
	"syscall/js"
)

type SettingPane struct {
	panel js.Value
	Label string
}

func NewSettingPane(id string) SettingPane {

	doc := js.Global().Get("document")

	div := doc.Call("getElementById", id)

	fmt.Printf("div %+v\n", div)

	h1 := doc.Call("createElement", "h1")
	h1.Set("textContent", "Test")

	div.Call("appendChild", h1)

	sp := SettingPane{
		panel: div,
	}

	return sp
}
