package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	document := js.Global().Get("document")
	msgArea := document.Call("getElementById", "greetingMessage")
	nameInput := document.Call("getElementById", "name")
	js.Global().Set("greet", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		mh := "Hello, "
		mb := "WebAssembry"
		if name := nameInput.Get("value"); !name.Equal(js.ValueOf("")) {
			mb = name.String()
		}
		msgArea.Set("innerHTML", mh+mb+"!!")
		return nil
	}))
	js.Global().Set("reset", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		msgArea.Set("innerHTML", nil)
		nameInput.Set("value", nil)
		return nil
	}))
	select {}
}
