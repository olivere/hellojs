package main

import (
	"fmt"
	"syscall/js"
)

var (
	beforeUnloadCh = make(chan struct{})
)

func main() {
	fmt.Println("Hello, world!")

	beforeUnloadCb := js.NewEventCallback(js.PreventDefault, func(js.Value) {
		beforeUnloadCh <- struct{}{}
	})
	defer beforeUnloadCb.Release()

	doc := js.Global().Get("document")
	win := js.Global().Get("window")

	fmt.Printf("document: %v\n", doc)
	fmt.Printf("window: %v\n", win)
	loc := win.Get("location")
	fmt.Printf("location: %v\n", loc)

	href := loc.Get("href")
	fmt.Printf("Location: %s\n", href.String())

	msg := doc.Call("getElementById", "message")
	fmt.Printf("Previous text: %s\n", msg.Get("innerHTML"))
	msg.Set("innerHTML", "WASM module has been loaded.")

	// Bind to button on HTML page
	btnClickCb := js.NewCallback(btnClick)
	defer btnClickCb.Release()
	doc.Call("getElementById", "btn").Call("addEventListener", "click", btnClickCb)
	doc.Call("getElementById", "btn").Set("disabled", false)

	<-beforeUnloadCh

	fmt.Println("Exiting WASM module.")
}

func btnClick(args []js.Value) {
	fmt.Println("Button clicked.")
}
