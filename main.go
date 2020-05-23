package main

import (
	"fmt"
	"strconv"
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
	drawSvg(20, 30, 50, 50)
	select {}
}

func drawSvg(iniX, iniY, iniWidth, iniHeight int) {
	document := js.Global().Get("document")
	target := document.Call("getElementById", "target")
	chX := registerChangeEvent("input-x")
	chY := registerChangeEvent("input-y")
	chW := registerChangeEvent("input-width")
	chH := registerChangeEvent("input-height")

	oldX, oldY, oldW, oldH := iniX, iniY, iniWidth, iniHeight
	go func() {
		for {
			target.Call("setAttribute", "x", oldX)
			target.Call("setAttribute", "y", oldY)
			target.Call("setAttribute", "width", oldW)
			target.Call("setAttribute", "height", oldH)
			select {
			case oldX = <-chX:
			case oldY = <-chY:
			case oldW = <-chW:
			case oldH = <-chH:
			}
		}
	}()
}

func registerChangeEvent(elementID string) <-chan int {
	ch := make(chan int, 0)
	document := js.Global().Get("document")
	input := document.Call("getElementById", elementID)
	input.Call("addEventListener", "change", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		input := this.Get("value").String()
		fmt.Println(input)
		v, err := strconv.Atoi(input)
		if err == nil { //if NO error, send value.
			ch <- v
		}
		return nil
	}))
	return ch
}
