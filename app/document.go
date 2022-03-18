package main

import (
	"fmt"
	"syscall/js"
)

type Document struct {
	doc js.Value
}

func GetDocument() *Document {
	doc := js.Global().Get("document")
	return &Document{
		doc: doc,
	}
}

func (d *Document) GetElementById(id string) js.Value {
	element := d.doc.Call("getElementById", id)
	if element.Truthy() {
		return element
	}
	fmt.Printf("unable to find element by id '%v'", id)
	return js.Value{}
}

func (d *Document) CreateElement(tagType string) js.Value {
	return d.doc.Call("createElement", tagType)
}

func (d *Document) Call(m string, args ...interface{}) js.Value {
	return d.doc.Call(m, args)
}

func (d *Document) Truthy() bool {
	return d.doc.Truthy()
}
